package bot

import (
    "os"
    "log"
    "time"

    tele "gopkg.in/telebot.v3"
)



// initialize a new bot struct
func NewBot() (interface{}, error) {
    token, err := os.ReadFile("telegram-token.txt")
    if err != nil {
        log.Println(err)
        os.Exit(1)
    }

    settings := tele.Settings{
        Token: string(token),
        Poller: &tele.LongPoller{
            Timeout:time.Second * 2},
        }

    return tele.NewBot{settings}, nil
}
