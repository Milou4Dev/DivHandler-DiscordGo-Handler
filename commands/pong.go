package commands

import (
    "code/utils"
    "github.com/bwmarrin/discordgo"
)

func init() {
    CommandHandlers["pong"] = handlePongCommand
}

func PongCommand() *discordgo.ApplicationCommand {
    return &discordgo.ApplicationCommand{Name: "pong", Description: "Responds with a playful message."}
}

func handlePongCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
    utils.RespondWithMessage(s, i, "Wait, that's illegal.", discordgo.MessageFlagsEphemeral)
}