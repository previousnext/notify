package main

import (
	"fmt"

	"gopkg.in/alecthomas/kingpin.v2"
	"github.com/nlopes/slack"
	"os"
)

var (
	cliToken     = kingpin.Flag("token", "Slack token for authentication").Required().Envar("SLACK_TOKEN").String()
	cliChannel   = kingpin.Flag("channel", "Slack channel for posting updates").Required().Envar("SLACK_CHANNEL").String()
	cliUsername  = kingpin.Flag("username", "The slack username").Default("M8s").String()
	cliIconEmoji = kingpin.Flag("icon-emoji", "The slack icon emoji").Default(":m8s:").String()
	cliMessage   = kingpin.Flag("message", "The slack message").Required().String()
)

func main() {
	kingpin.Parse()

	api := slack.New(*cliToken)

	msg := slack.PostMessageParameters{
		Username:  *cliUsername,
		IconEmoji: *cliIconEmoji,
	}

	_, _, err := api.PostMessage(*cliChannel, *cliMessage, msg)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
