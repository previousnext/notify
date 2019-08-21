package main

import (
	"github.com/nlopes/slack"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	cliToken     = kingpin.Flag("token", "Slack token for authentication").Required().Envar("SLACK_TOKEN").String()
	cliChannel   = kingpin.Flag("channel", "Slack channel for posting updates").Required().Envar("SLACK_CHANNEL").String()
	cliUsername  = kingpin.Flag("username", "The slack username").Default("M8s").String()
	cliIconEmoji = kingpin.Flag("icon-emoji", "The slack icon emoji").Default(":m8s:").String()
	cliMessage   = kingpin.Flag("message", "The slack message").Required().String()
	cliColor     = kingpin.Flag("color", "The slack message color").Default("#32cd32").String()
)

func main() {
	kingpin.Parse()

	api := slack.New(*cliToken)
	
	attachment := slack.Attachment{
		Color: *cliColor,
		Fields: []slack.AttachmentField{
			{
				Value: *cliMessage,
			},
		},
	}

	msg := slack.PostMessageParameters{
		Username:  *cliUsername,
		IconEmoji: *cliIconEmoji,
	}

	_, _, err := api.PostMessage(*cliChannel, slack.MsgOptionPostMessageParameters(msg), slack.MsgOptionAttachments(attachment))

	if err != nil {
		panic(err)
	}
}
