package telegram

import (
	"github.com/foundation-framework/notify"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type service struct {
	bot     *tgbotapi.BotAPI
	chatIds []int64
}

func NewService(token string, chatIds []int64) (notify.Service, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	return &service{
		bot:     bot,
		chatIds: chatIds,
	}, nil
}

func (s *service) Send(text string, attachments ...notify.Attachment) error {
	//defer func() {
	//	for _, attachment := range attachments {
	//		// Ignore any errors at this stage
	//		_ = attachment.Close()
	//	}
	//}()

	for _, chatId := range s.chatIds {
		if err := s.sendText(chatId, text); err != nil {
			return err
		}

		if err := s.sendAttachments(chatId, attachments); err != nil {
			return err
		}
	}

	return nil
}

func (s *service) sendText(chatId int64, text string) error {
	message := tgbotapi.NewMessage(chatId, text)
	message.ParseMode = "HTML"

	_, err := s.bot.Send(message)
	return err
}

func (s *service) sendAttachments(chatId int64, attachments []notify.Attachment) error {
	for _, attachment := range attachments {
		if err := attachment.Reset(); err != nil {
			return err
		}

		file := tgbotapi.FileReader{
			Name:   attachment.Name(),
			Reader: attachment.Reader(),
		}

		_, err := s.bot.SendMediaGroup(
			tgbotapi.NewMediaGroup(
				chatId,
				[]interface{}{
					tgbotapi.NewInputMediaDocument(file),
				},
			),
		)

		if err != nil {
			return err
		}
	}

	return nil
}
