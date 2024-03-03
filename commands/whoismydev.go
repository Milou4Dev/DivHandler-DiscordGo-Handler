package commands

import (
	"code/utils"
	"github.com/bwmarrin/discordgo"
)

func init() {
	CommandHandlers["whoismydev"] = handleWhoIsMyDevCommand
}

func WhoIsMyDevCommand() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{Name: "whoismydev", Description: "Reveals who the developer is."}
}

func handleWhoIsMyDevCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	utils.RespondWithMessage(s, i, "My dev is Milou4Dev.", discordgo.MessageFlagsEphemeral)
}
