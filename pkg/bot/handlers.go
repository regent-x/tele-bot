package bot

import (
    "log"
    "fmt"

    tele "gopkg.in/telebot.v3"
)

func RegentBot(bot *tele.Bot)  {

    bot.Handle("/hello", func(c tele.Context) error {
        log.Println("/hello handle")
        c.Send("hey")
        return nil
     })

    bot.Handle(tele.OnText, func(c tele.Context) error {
        var user = c.Sender()
        var msg = c.Message()
        var reply = fmt.Sprintf("I don't understand (%s) means yet, mr %s", msg, user)

        log.Println("Unhandled")
        c.Send(reply)
        return nil
    })

}
