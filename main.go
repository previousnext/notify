package main

import (
	"gopkg.in/alecthomas/kingpin.v2"
	"github.com/nlopes/slack"
	"github.com/pkg/errors"

	"fmt"
	"os"
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

	msg := slack.PostMessageParameters{
		Username:  *cliUsername,
		IconEmoji: *cliIconEmoji,
		Attachments: []slack.Attachment{
			{
				Color: *cliColor,
				Fields: []slack.AttachmentField{
					{
						Value: *cliMessage,
					},
				},
			},
		},
	}

	_, _, err := api.PostMessage(*cliChannel, "", msg)

	if err != nil {
		fmt.Println(errors.Wrap(err, "failed to send slack message"))
		os.Exit(1)
	}
}
