package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/line/line-bot-sdk-go/linebot"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Gzip())

	// CORS restricted
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods: []string{echo.GET},
	}))

	e.Static("/", "dist")

	lineBotAPI := e.Group("linebot/api/v1")
	{
		lineBotAPI.POST("/video", PutVideoAPIRequest())
	}

	if os.Getenv("GO_RUN_ENV") == "" {
		// Start server on production
		e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
	} else {
		// .env.development: For Development
		// .env.test: For test
		err := godotenv.Load(fmt.Sprintf(".env.%s", os.Getenv("GO_RUN_ENV")))
		if err != nil {
			log.Fatal("Error loading .env file")
		} else {
			// Start server on TEST or development
			e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
		}
	}
}

// PutVideoAPIRequest handles API call
func PutVideoAPIRequest() echo.HandlerFunc {
	return func(c echo.Context) error {
		bot, err := linebot.New(
			os.Getenv("CHANNEL_SECRET"),
			os.Getenv("CHANNEL_TOKEN"),
		)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("bot is ready")

		// Setup HTTP Server for receiving requests from LINE platform
		events, err := bot.ParseRequest(c.Request())
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				return c.NoContent(http.StatusBadRequest)
			}
			return c.NoContent(http.StatusInternalServerError)
		}
		for _, event := range events {
			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do(); err != nil {
						log.Print(err)
					}
				}
			}
		}
		return c.String(http.StatusOK, "Process is done.")
	}
}
