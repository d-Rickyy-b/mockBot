package main1

import (
	"errors"
	"fmt"
	tb "gopkg.in/telebot.v3"
	"log"
	"math/rand"
	"strings"
	"time"
	"unicode/utf8"
)

func mock(input string) string {
	// The function mock takes a string as an input and randomly changes the case of the individual letters.
	// The function returns the modified string
	var output string
	for _, char := range input {
		if rand.Intn(2) == 0 {
			output += strings.ToUpper(string(char))
		} else {
			output += strings.ToLower(string(char))
		}
	}
	return output
}

// ellipsis takes a text and an integer to generate a new string that's <= 'max' characters (not bytes) long.
// Excess characters will be removed from the middle and replaced with "..."
func ellipsis(text string, max int) string {
	// subtract length of "..."
	max -= 5

	if max <= 5 {
		max = 5
	}

	if len(text) <= max {
		return text
	}

	charCount := 0

	endFirst := max / 2
	startLast := utf8.RuneCountInString(text) - (max / 2)
	result := ""

	for _, r := range text {
		charCount++

		if charCount > endFirst && charCount <= startLast {
			continue
		} else if charCount == endFirst {
			result += string(r)
			result += " ... "
			continue
		}

		result += string(r)
	}

	return result
}

func mockHandler(c tb.Context) error {
	// The function mockHandler takes a query as an input and sends a message with the modified string.
	// The function is called when the user sends a message with the command /mock
	q := c.Query()
	if q == nil {
		return errors.New("query is nil")
	}
	fmt.Println(q.Text)

	options := []string{}
	mockText := ""
	if q.Text != "" {
		mockText = mock(q.Text)
		options = append(options, mockText)
	}

	results := make(tb.Results, 1)
	results = append(results, &tb.ArticleResult{
		ResultBase: tb.ResultBase{},
		Title:      mockText,
		Text:       mockText,
		HideURL:    true,
	})

	return c.Bot().Answer(q, &tb.QueryResponse{
		Results:   results,
		CacheTime: 60,
	})
}

func main() {
	// Random seed I typed by hand
	rand.Seed(4156453156)

	pref := tb.Settings{
		Token:  "652299433:AAETY4DrM1BX3njHNrDfJsnND15F44aNZPY",
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tb.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/start", func(c tb.Context) error {
		welcomeText := fmt.Sprintf("Hi! Thanks for using this bot. You can just type '@%s' in any chat and I'll provide you with a mock text.", b.Me.Username)
		return c.Send(welcomeText)
	})

	b.Handle(tb.OnQuery, mockHandler)

	b.Start()
}
