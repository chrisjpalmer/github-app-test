package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/chrisjpalmer/github-app-test/internal/github"
	gogithub "github.com/google/go-github/github"
)

type Handler struct {
	AppID int64
}

func New(appID int64) *Handler {
	return &Handler{AppID: appID}
}

func (h *Handler) Callback(rw http.ResponseWriter, r *http.Request) {
	githubEvent := r.Header.Get("X-Github-Event")
	fmt.Println("/callback called with event ", githubEvent)

	if githubEvent != "installation" {
		fmt.Println("dropping event of type ", githubEvent)
		return
	}

	dec := json.NewDecoder(r.Body)
	var payload gogithub.InstallationEvent
	if err := dec.Decode(&payload); err != nil {
		http.Error(rw, "could not parse body", http.StatusBadRequest)
		return
	}

	if *payload.Action == "created" {
		fmt.Printf("installation created event, listing repos...\n\n")

		cl := github.NewClient(h.AppID, *payload.Installation.ID)
		repos, err := cl.ListRepos()
		if err != nil {
			http.Error(rw, "could not list repos", http.StatusInternalServerError)
			return
		}
		for _, r := range repos {
			fmt.Println(*r.FullName)
		}
	} else {
		fmt.Println("dropping event of action ", *payload.Action)
	}
}
