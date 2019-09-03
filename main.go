package main

import (
	"fmt"
	"log"
	"os"

	"github.com/nlopes/slack"
)

func main() {
	fmt.Println("running...")

	api := slack.New(
		os.Getenv("SLACK_TOKEN"),
		slack.OptionDebug(false),
		slack.OptionLog(log.New(os.Stdout, "disglair: ", log.Lshortfile|log.LstdFlags)),
	)

	u, err := api.GetUsers()
	if err != nil {
		panic(err)
	}

	users := map[string]slack.User{}
	for _, us := range u {
		users[us.ID] = us
	}

	chann, err := api.GetChannels(true)
	if err != nil {
		panic(err)
	}

	channels := map[string]slack.Channel{}
	for _, c := range chann {
		channels[c.ID] = c
	}

	rtm := api.NewRTM()
	go rtm.ManageConnection()

	fmt.Println("started listener")
	for msg := range rtm.IncomingEvents {
		switch ev := msg.Data.(type) {
		case *slack.HelloEvent:
			// Ignore hello

		case *slack.ConnectedEvent:
			fmt.Println("Info:", ev.Info)
			fmt.Println("Connection counter:", ev.ConnectionCount)
			// Replace C2147483705 with your Channel ID
			//rtm.SendMessage(rtm.NewOutgoingMessage("Hello world", "C2147483705"))

		case *slack.MessageEvent:
			//fmt.Printf("Message: %+v\n", ev)
			fmt.Printf("%s %s | %s: %s\n", ev.Timestamp, channels[ev.Channel].Name, users[ev.User].Name, ev.Text)

		case *slack.PresenceChangeEvent:
			//fmt.Printf("Presence Change: %v\n", ev)

		case *slack.LatencyReport:
			//fmt.Printf("Current latency: %v\n", ev.Value)

		case *slack.RTMError:
			fmt.Printf("Error: %s\n", ev.Error())

		case *slack.InvalidAuthEvent:
			fmt.Printf("Invalid credentials")
			return

		default:
			// Ignore other events..
			//fmt.Printf("Unexpected: %v\n", msg.Data)
		}
	}
}
