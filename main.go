package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"gopkg.in/alecthomas/kingpin.v1"
)

var (
	cliName    = kingpin.Flag("name", "Name your bot.").Default("Notify").OverrideDefaultFromEnvar("NOTIFY_NAME").String()
	cliEmoji   = kingpin.Flag("emoji", "Give your bot a custom image.").Default(":slack:").OverrideDefaultFromEnvar("NOTIFY_EMOJI").String()
	cliChannel = kingpin.Flag("channel", "Which channel do you wish to post in.").Default("#general").OverrideDefaultFromEnvar("NOTIFY_CHANNEL").String()
	cliMessage = kingpin.Arg("message", "The message you wish to send.").Required().String()
	cliUrl     = kingpin.Arg("url", "The url you wish to post to.").Required().String()
)

type Message struct {
	Username string `json:"username"`
	Emoji    string `json:"icon_emoji"`
	Channel  string `json:"channel"`
	Text     string `json:"text"`
}

func main() {
	kingpin.Parse()

	m := Message{
		Username: *cliName,
		Emoji:    *cliEmoji,
		Channel:  *cliChannel,
		Text:     *cliMessage,
	}

	// Convert the response object into a json string for posting.
	jsonStr, err := json.Marshal(m)
	Check(err)

	// Build a request that will be sent to Slack.
	client := &http.Client{}
	req, err := http.NewRequest("POST", *cliUrl, bytes.NewBuffer(jsonStr))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	_, err = client.Do(req)
	Check(err)
}

func Check(e error) {
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}
}
