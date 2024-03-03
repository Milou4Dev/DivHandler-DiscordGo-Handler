package commands

import "github.com/bwmarrin/discordgo"

var CommandHandlers = map[string]func(*discordgo.Session, *discordgo.InteractionCreate){}

func GetApplicationCommands() []*discordgo.ApplicationCommand {
	return []*discordgo.ApplicationCommand{
		WhoIsMyDevCommand(),
		PongCommand(),
		PingCommand(),
	}
}
