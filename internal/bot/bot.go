package bot

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"mockBot/internal/config"
	"mockBot/internal/mock"
	"time"

	tb "gopkg.in/telebot.v3"
)

func mockHandler(c tb.Context) error {
	// The function mockHandler takes a query as an input and sends a message with the modified string.
	// The function is called when the user sends a message with the command /mock
	q := c.Query()
	if q == nil {
		return errors.New("query is nil")
	}

	if q.Text == "" {
		return errors.New("query text is empty")
	}
	// Log all requests for debugging crashes
	log.Printf("New query: '%s'", q.Text)
	mockText := mock.MockText(q.Text)

	results := make(tb.Results, 1)
	results[0] = &tb.ArticleResult{
		ResultBase: tb.ResultBase{},
		Title:      mockText,
		Text:       mockText,
		HideURL:    true,
	}

	return c.Bot().Answer(q, &tb.QueryResponse{
		Results:   results,
		CacheTime: 60,
	})
}

func StartBot(c config.Config) {
	// Random seed I typed by hand
	rand.Seed(4156453156)

	var poller tb.Poller
	if c.Webhook.Enabled {
		// Currently, we don't support all the fields of Webhook.
		// We could add stuff like TLS/Cert config later still
		poller = &tb.Webhook{
			Listen: c.Webhook.Listen,
			Endpoint: &tb.WebhookEndpoint{
				PublicURL: c.Webhook.Url,
			},
		}
		log.Println("Using webhook from config!")
	} else {
		poller = &tb.LongPoller{Timeout: 10 * time.Second}
		log.Println("Using long polling as fallback!")
	}

	b, err := tb.NewBot(tb.Settings{
		Token:  c.Token,
		Poller: poller,
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/start", func(c tb.Context) error {
		welcomeText := fmt.Sprintf("Hi! Thanks for using this bot. You can just type '@%s' in any chat and I'll provide you with a mocked version of your text.", b.Me.Username)
		return c.Send(welcomeText)
	})

	b.Handle(tb.OnQuery, mockHandler)
	log.Printf("Starting @%s\n", b.Me.Username)
	b.Start()
}
