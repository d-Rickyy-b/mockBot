# mockBot - A Telegram bot to mock your friends
[![build](https://github.com/d-Rickyy-b/mockBot/actions/workflows/release_build.yml/badge.svg)](https://github.com/d-Rickyy-b/mockBot/actions/workflows/release_build.yml)
[![test](https://github.com/d-Rickyy-b/mockBot/actions/workflows/test_push_pr.yml/badge.svg)](https://github.com/d-Rickyy-b/mockBot/actions/workflows/test_push_pr.yml)

Annoyed of someone in a Telegram chat? Want to mock them because of stupid statements they made?
This bot is for you! It gives you the ability to quickly send a mocked text.

Using "Oh come on, I wasn't that late" as a input, it will generate and send a mocked text like this: 
> oh COmE On, I wAsN'T That lATe!

You can find the hosted version at [@sp0ngeBot](https://t.me/sp0ngeBot) on Telegram.

## Configuration
This bot uses a configuration file. You can create your own using the sample file `config.template.json`.

```json
{
  "token": "<your_bot_token>",
  "webhook": {
    "enabled": false,
    "url": "example.com/test",
    "listen": "127.0.0.1:8123"
  }
}
```
Just fill in your bot token. If you want to use a webhook, use the webhook options.
The current config only allows for receiving webhooks from a reverse proxy such as nginx.

## Usage
Create your `config.json` file and put it into the same directory as the binary. 
You can then simply call the binary.

```
Usage of mockBot:
  -config string
        Path to config file (default "config.json")
```

You can also use the parameter `-config` to speficy a path to another config file.
