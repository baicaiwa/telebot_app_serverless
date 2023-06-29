package text

import (
	"github.com/clin003/tgbot_app_dev/features"

	tele "gopkg.in/telebot.v3"
)

// Command: /start <PAYLOAD>
func OnText(c tele.Context) error {
	if c.Message().Private() ||
		c.Message().FromChannel() ||
		c.Message().IsReply() ||
		c.Message().IsForwarded() {
		return nil
	}

	if c.Message().SenderChat.Type != tele.ChatChannel || c.Message().FromGroup() {
		return nil
	}

	menu := &tele.ReplyMarkup{ResizeKeyboard: true}
	btn1 := menu.URL("薅羊毛📦", "https://t.me/haowu_push")
	btn2 := menu.URL("值得买🔥", "https://t.me/haowu_dw")
	menu.Reply(
		menu.Row(btn1),
		menu.Row(btn2),
	)

	return c.Reply("评论区请友好👬发言", menu)
}

func init() {
	features.RegisterFeature(tele.OnText, OnText)
}
