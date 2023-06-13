package main

import (
    "log"

    "github.com/regent-x/tele-bot/pkg/bot"
)

func main(){
    var bot = NewBot()

    RegentBot(bot)


    log.Println("Regent is in the house")

    bot.Start()
}
