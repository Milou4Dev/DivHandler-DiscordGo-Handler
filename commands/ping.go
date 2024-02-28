package commands

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"time"
)

func init() {
	CommandHandlers["ping"] = handlePingCommand
}

func PingCommand() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:        "ping",
		Description: "Measures bot's response time.",
	}
}

func handlePingCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	startTime := time.Now()

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Calculating ping...",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
	if err != nil {
		fmt.Println("Error responding to ping command:", err)
		return
	}

	delay := time.Since(startTime).Milliseconds()
	apiLatency := s.HeartbeatLatency().Milliseconds()

	_, err = s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
		Content: &[]string{fmt.Sprintf("API Latency: %d ms\nLatency(Processing time): %d ms", apiLatency, delay)}[0],
	})
	if err != nil {
		fmt.Println("Error updating ping response:", err)
	}
}