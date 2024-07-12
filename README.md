# Github App Test

## Prerequisites

1. Ngrok
2. Asdf

## Overview

A demo app that prints your private repositories when installed in your repo:

```
/callback called with event  installation
installation created event, listing repos...

chrisjpalmer/private-repo1
chrisjpalmer/private-repo2
...
```

## Running

1. Setup an Ngrok account. Follow the instructions to get [setup for mac](https://dashboard.ngrok.com/get-started/setup/macos). Use the "Static Domain" option and copy the Static Domain.
1. Create a [new github app](https://github.com/settings/apps/new)
    - Homepage URL: `https://github.com/<your-user>`
    - Callback URL: `https://<static-domain>/callback`
    - Webhook URL: `https://<static-domain>/callback`
    - Where can this app be installed: **Only on this account**
2. Copy the *App ID* and paste it as the value `appID` in `main.go`.
3. Scroll down and click *Generate Private Key*, save it in the root of this project as `secret.pem`
4. Click *Permissions & events*. Grant:
    - `Repository > Metadata (read-only)`
5. Start ngrok locally on port 8080 `ngrok http --domain=<static-domain> 8080`
6. Start this project by running `go run ./cmd/github`
7. Back on github, click *Install App* and install the app for your user account.

The webhook should be received and your private repositories printed out in the terminal.

