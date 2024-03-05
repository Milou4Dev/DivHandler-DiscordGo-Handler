package commands

import "github.com/bwmarrin/discordgo"

var CommandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){}

func GetApplicationCommands() []*discordgo.ApplicationCommand {
    return []*discordgo.ApplicationCommand{
        WhoIsMyDevCommand(),
        PongCommand(),
        PingCommand(),
    }
}