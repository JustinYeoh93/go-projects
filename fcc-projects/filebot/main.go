package main

import (
	"fmt"

	"github.com/slack-go/slack"
)

const (
	SLACK_APP_TOKEN  = "xapp-1-A04831QGPRP-4247877634919-06e5804abade7fc082a2ad7a1e0cb940da049a99092dd75975f87d55029b6097"
	SLACK_BOT_TOKEN  = "xoxb-4275040733233-4275068396241-e9dXaf1AXNh1bmO7OkJnZhWl"
	ASOKO_CHANNEL_ID = "C047HPEM970"
)

func main() {
	api := slack.New(SLACK_BOT_TOKEN)
	channelArr := []string{ASOKO_CHANNEL_ID}
	fileArr := []string{"some.yaml"}

	for i := 0; i < len(fileArr); i++ {
		params := slack.FileUploadParameters{
			Channels: channelArr,
			File:     fileArr[i],
		}
		file, err := api.UploadFile(params)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Name: %s, URL: %s\n", file.Name, file.URLPrivate)
	}
}
