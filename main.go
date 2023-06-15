package main

import (
    "log"
    "time"
    "os"

    telebot "gopkg.in/telebot.v3"
)

func main(){
    
    perf := telebot.Settings{
        Token: os.Getenv("TELEGRAM_TOKEN"),
        Poller: &telebot.LongPoller{
            Timeout: 3 * time.Second,
        },
    }

    bot, err := telebot.NewBot(perf)
    if err != nil {
        log.Println(err)
        os.Exit(1)
    }

    bot.Handle("/start", func(c telebot.Context)error{
        c.Send("Yo, how r u doin?")
        return nil
    })
    bot.Handle("/commands", func(c telebot.Context)error{
        c.Send(`Here are the list of commands
        /Greetings
        /time
        /date
        /news
        /others`)
        return nil
    })

    log.Println("It has begun")
    bot.Start()
}
