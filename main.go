package main

import(
	"log"
	"github.com/nicklaw5/helix"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"time"
)

func main() {

	// Creating the telegram bot
	bot, err := tgbotapi.NewBotAPI(TelegramAPIKey)
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true

	// Creating the twitch client
	client, err := helix.NewClient(&helix.Options{
		ClientID: TwitchAPIKey,
	})
	if err != nil {
		log.Panic(err)
	}
	// users contains the username used to login in twitch.
	// You can also use the user id, you just need to substitute the UserLogins field in the struct to UserIDs
	users := make([]string, 1)
	users[0] = UserLogin
	streamParams := &helix.StreamsParams{
		UserLogins: users,
	}

	// live is used to check when the telegram bot should send a message
	live := false;
	for {
		// Getting the stream info from twitch
		res, _ := client.GetStreams(streamParams)

		// Checking if the user is live
		if !(len(res.Data.Streams) == 0 ) {
			
			if !live {
				msg := tgbotapi.NewMessage(ChatID, MessageText)
				bot.Send(msg)
				// Setting the boolean value to true, so the bot send just one message: one each time the user goes live
				live = true
			}

		} else {
			live = false
		}

		time.Sleep(time.Minute * 5)
	}
}