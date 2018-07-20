# Send message via Slack incoming webhook

The very first practice of Golang.

## Installation

```bash
go get github.com/o-p/slack-webhook
cd ~/go/src/github.com/o-p/slack-webhook
```

## Set up

First of all, you have to add incoming webhook to your Slack channel. You'll get a URL that looks like `https://hooks.slack.com/services/..../..../......`.

Copy and replace the const `webhookURL` in `slack.go`

Now you can simply try it by

```bash
$ echo "Hello, world!" | go run slack.go
200 OK

$ ls -al | go run slack.go -u <your-id> -n "You know who" -emoji=ghost

# example to send a emoji from bot
$ echo ":hehopeless:" | go run slack.go \
  -image https://storage.googleapis.com/uploads-blog-icook/2017/12/28bbbc20-item_0000002430_04.jpg \
  -n 黑鮪魚 \
  -keep
200 OK
```

