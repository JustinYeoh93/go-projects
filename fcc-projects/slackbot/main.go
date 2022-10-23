package main

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/shomali11/slacker"
)

const (
	SLACK_APP_TOKEN = "xapp-1-A04831QGPRP-4247877634919-06e5804abade7fc082a2ad7a1e0cb940da049a99092dd75975f87d55029b6097"
	SLACK_BOT_TOKEN = "xoxb-4275040733233-4275068396241-e9dXaf1AXNh1bmO7OkJnZhWl"
)

func main() {
	bot := slacker.NewClient(SLACK_BOT_TOKEN, SLACK_APP_TOKEN)

	// logging purposes
	go printCommandEvents(bot.CommandEvents())

	// add bot command
	bot.Command("my yob is <year>", &slacker.CommandDefinition{
		Description: "year of birth calculator",
		Examples:    []string{"my yob is 1993"},
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				fmt.Println(err)
			}
			age := 2021 - yob
			r := fmt.Sprintf("age is %d", age)
			response.Reply(r)
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

func printCommandEvents(ch <-chan *slacker.CommandEvent) {
	for event := range ch {
		fmt.Println("==========")
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
	}
}
