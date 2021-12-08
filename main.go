package main

import (
    "fmt"
    "os"
    "github.com/bwmarrin/discordgo"

)
func main(){

    bot, err := discordgo.New("Bot " + os.Getenv("token"))

    if err != nil {
        panic(err)
    }


    bot.AddHandler(ready)
    bot.AddHandler(messageCreate)

    err = bot.Open()

	if err != nil {
		fmt.Println("Erro ao abrir sessão Discord: ", err)
	}

	fmt.Println("Bot está rodando.  Pressione CTRL-C para sair.")
	for {}

	bot.Close()
}

func ready(s *discordgo.Session, event *discordgo.Ready){
    s.UpdateStatus(0, "Com Go")
    fmt.Println("Usuário logado " +string(s.State.User.ID))
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate){

    if m.Content == "ping"{
        s.ChannelMessageSend(m.ChannelID, "pong")
    }
    
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate){

    if m.Content == "ola"{
        s.ChannelMessageSend(m.ChannelID, "Oie")
    }
    
}
