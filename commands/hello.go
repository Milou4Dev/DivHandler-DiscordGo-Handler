package commands

import (
	"github.com/bwmarrin/discordgo"
)

func HelloCommand() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:        "hello",
		Description: "HELLO_WORLD 2.0",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "name",
				Description: "Name?",
				Required:    false,
			},
		},
	}
}

func HandleHelloCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	var name string
	if len(i.ApplicationCommandData().Options) > 0 {
		name = i.ApplicationCommandData().Options[0].StringValue()
	}
	response := "Hello World!"
	if name != "" {
		response = "Hello, " + name + "!"
	}
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: response,
		},
	})
	if err != nil {
		return
	}
}
