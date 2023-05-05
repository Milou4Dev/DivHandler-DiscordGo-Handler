package main

import (
    "log"

    "Divbot/commands"
    "github.com/bwmarrin/discordgo"
)

const (
    grayColor  = "\033[90m"
    resetColor = "\033[0m"
)

func registerSlashCommands(s *discordgo.Session) {
    appCommands := []*discordgo.ApplicationCommand{
        commands.HelloCommand(),
        commands.PongCommand(),
    }

    for _, cmd := range appCommands {
        AddCommand(s, cmd)
    }

    s.AddHandler(interactionHandler)
}

func AddCommand(s *discordgo.Session, cmd *discordgo.ApplicationCommand) {
    _, err := s.ApplicationCommandCreate(s.State.User.ID, "", cmd)
    if err != nil {
        log.Printf("Error creating global application command %s: %v", cmd.Name, err)
    } else {
        log.Printf("%sRegistered global slash command: %s%s", grayColor, cmd.Name, resetColor)
    }
}

func interactionHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
    commandName := i.ApplicationCommandData().Name

    switch commandName {
    case "hello":
        commands.HandleHelloCommand(s, i)
    case "pong":
        commands.HandlePongCommand(s, i)
    default:
        log.Printf("Unknown command: %s", commandName)
    }
}
