package main

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func SendMsg(token string, chatIDStr string, message string) {

	// 检查 Token 是否存在
	if token == "" {
		log.Fatal("错误: TELEGRAM_BOT_TOKEN 环境变量未设置。")
	}

	// 检查 Chat ID 是否存在
	if chatIDStr == "" {
		log.Fatal("错误: TELEGRAM_CHAT_ID 环境变量未设置。")
	}

	if message != "" {
		// 将 Chat ID 从字符串转换为 int64
		var chatID int64
		_, err := fmt.Sscanf(chatIDStr, "%d", &chatID)
		if err != nil {
			log.Fatalf("错误: 无法解析 TELEGRAM_CHAT_ID: %v", err)
		}

		// 创建新的 Bot 实例
		bot, err := tgbotapi.NewBotAPI(token)
		if err != nil {
			log.Fatalf("错误: 创建 Telegram Bot 实例失败: %v", err)
		}

		// 创建消息对象
		msg := tgbotapi.NewMessage(chatID, message)

		// 发送消息
		_, err = bot.Send(msg)
		if err != nil {
			log.Fatalf("错误: 发送消息失败: %v", err)
		}

		fmt.Println("消息发送成功！")
	}

}
