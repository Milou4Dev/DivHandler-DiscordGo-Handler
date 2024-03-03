package modules

import (
    "code/commands"
    "github.com/bwmarrin/discordgo"
    "log"
)

func RegisterSlashCommands(session *discordgo.Session) {
    if session == nil || session.State == nil || session.State.User == nil {
        log.Println("Session or session state or session user is nil")
        return
    }
    appCommands := commands.GetApplicationCommands()
    for _, cmd := range appCommands {
        if _, err := session.ApplicationCommandCreate(session.State.User.ID, "", cmd); err != nil {
            log.Printf("Error creating global application command %s: %v", cmd.Name, err)
        } else {
            log.Printf("Registered global slash command: %s", cmd.Name)
        }
    }
}

func InteractionHandler(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
    if session == nil {
        log.Println("Session is nil")
        return
    }
    commandName := interaction.ApplicationCommandData().Name
    if handler, exists := commands.CommandHandlers[commandName]; exists {
        handler(session, interaction)
    } else {
        log.Printf("Unknown command: %s", commandName)
    }
}