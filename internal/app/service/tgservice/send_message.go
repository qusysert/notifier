package tgservice

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"notifier/internal/app/model"
)

func (s *Service) SendMessage(msg model.Message) error {
	log.Printf("Processing message request")
	bot, err := tgbotapi.NewBotAPI(msg.Token)
	if err != nil {
		return fmt.Errorf("error creating bot: %v", err)
	}

	for _, chatId := range msg.ChatIds {
		_, err = bot.Send(tgbotapi.NewMessage(chatId, msg.Message))
		if err != nil {
			return fmt.Errorf("error sending message: %v", err)
		}
	}
	return nil
}

//func (s *Service) SendMessage(msg model.Message) error {
//	log.Printf("Processing message request")
//	bot, err := tgbotapi.NewBotAPI(msg.Token)
//	if err != nil {
//		return fmt.Errorf("error creating bot: %v", err)
//	}
//	wg := &sync.WaitGroup{}
//	errChan := make(chan error)
//
//	for _, chatId := range msg.ChatIds {
//		wg.Add(1)
//		go func(chatId int64) {
//			defer wg.Done()
//			_, err = bot.Send(tgbotapi.NewMessage(chatId, msg.Message))
//			if err != nil {
//				errChan <- fmt.Errorf("error sending message: %v", err)
//			}
//		}(chatId)
//	}
//	wg.Wait()
//	for err := range errChan {
//		if err != nil {
//			return err
//		}
//	}
//	return nil
//}
