package main

import (
	"flag"
	"log"
	"mockBot/internal/bot"
	"mockBot/internal/config"
)

func main() {
	// configure logging
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	var configFile = flag.String("config", "config.json", "Path to config file")
	flag.Parse()

	c, _ := config.ReadConfig(*configFile)
	bot.StartBot(c)
}
