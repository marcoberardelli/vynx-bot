package main

import(
	"log"
	"github.com/nicklaw5/helix"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {

	bot, err := tgbotapi.NewBotAPI(TelegramAPIKey)
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true

	client, err := helix.NewClient(&helix.Options{
		ClientID: TwitchAPIKey,
	})
	if err != nil {
		log.Panic(err)
	}
	users := make([]string, 1)
	users[0] = UserLogin
	streamParams := &helix.StreamsParams{
		UserLogins: users,
	}

	live := false;
	for {
		res, _ := client.GetStreams(streamParams)
		if !(len(res.Data.Streams) == 0 ) {
			
			if !live {
				msg := tgbotapi.NewMessage(ChatID, MessageText)
				bot.Send(msg)
				live = true
			}

		} else {
			live = false
		}
	}	
}