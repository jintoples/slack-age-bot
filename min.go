package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/shomali11/slacker"
)

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-4223055417137-4210362238771-9VOmbUzuKW6SixBA7GqMXApu")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A0466A25SUB-4204274003814-adb0846a3f5fb5338553620c6dde06ee2b4336164d34d6e3551b559523b4de7c")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printComandEvents(bot.CommandEvents())

	bot.Command("my yob is <year>", &slacker.CommandDefinition{
		Description: "yob calculator",
		Examples:    []string{"my yob is 2020", "Atau contoh ini"},
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				println("error")
			}
			age := time.Now().Year() - yob
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

func printComandEvents(analyticsChanel <-chan *slacker.CommandEvent) {
	for event := range analyticsChanel {
		fmt.Println("Comand event")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}
