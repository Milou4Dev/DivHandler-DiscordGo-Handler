package modules

import (
    "code/commands"
    "github.com/bwmarrin/discordgo"
    "log"
)

func RegisterSlashCommands(session *discordgo.Session) {
    for _, cmd := range commands.GetApplicationCommands() {
        if _, err := session.ApplicationCommandCreate(session.State.User.ID, "", cmd); err != nil {
            log.Printf("Error creating global application command %s: %v", cmd.Name, err)
        } else {
            log.Printf("Registered global slash command: %s", cmd.Name)
        }
    }
}

func InteractionHandler(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
    if handler, exists := commands.CommandHandlers[interaction.ApplicationCommandData().Name]; exists {
        handler(session, interaction)
    } else {
        log.Printf("Unknown command: %s", interaction.ApplicationCommandData().Name)
    }
}