package main

import (
	"fmt"
	"strconv"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	tgBotMLForever *tgbotapi.BotAPI
	tgBotTest      *tgbotapi.BotAPI
	//chatId         = "-690605257"
	chatId = "-1001827918980"
)

func main() {
	tgBotMLForever = newMLBot() // 2036154314
	tgBotTest = newTestBot()

	fmt.Println("go go")
	go runTelegramResponse()
	fmt.Println("go go 2")

	time.Sleep(1 * time.Second)
	TeleMLSendMsg(chatId, "from the ml bot msg")
	fmt.Println("send msg")

	time.Sleep(120 * time.Second)
}

// newBot
func newMLBot() *tgbotapi.BotAPI {

	// ml forever bot
	token := "2036154314:AAFsrPCta3Sci20K6EAEh5C6xU8HTrijtUw"

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		bot.Debug = false
	}

	return bot

}

// newBot
func newTestBot() *tgbotapi.BotAPI {

	// test bot
	token := "5404679853:AAEH2EfY1h432ozPXz5MkuDvH3eCvn_0nHQ"

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		bot.Debug = false
	}

	return bot

}

func TeleMLSendMsg(chatId, s string) {

	chatIdInt64, _ := strconv.ParseInt(chatId, 10, 64)
	msg := tgbotapi.NewMessage(chatIdInt64, s)

	//msg := tgbotapi.NewMessageToChannel("bot_test_channel", s)

	_, err := tgBotMLForever.Send(msg)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TeleTestSendMsg(chatId, s string) {

	chatIdInt64, _ := strconv.ParseInt(chatId, 10, 64)
	msg := tgbotapi.NewMessage(chatIdInt64, s)

	_, err := tgBotTest.Send(msg)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func runTelegramResponse() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	fmt.Println("waiting .....")
	updates, err := tgBotTest.GetUpdatesChan(u)
	if err != nil {
		fmt.Println(err.Error())
	}

	for update := range updates {
		if update.ChannelPost != nil {
			fmt.Println(update.ChannelPost.Text)
		}
		if update.Message != nil {
			fmt.Println(update.Message.From.ID, update.Message.Text)
		}
	}
}
