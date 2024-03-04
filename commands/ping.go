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
	return &discordgo.ApplicationCommand{Name: "ping", Description: "Measures bot's response time."}
}

func handlePingCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	startTime := time.Now()
	if err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{Content: "Calculating ping...", Flags: discordgo.MessageFlagsEphemeral},
	}); err != nil {
		fmt.Printf("Error sending interaction response: %v\n", err)
		return
	}
	delay := time.Since(startTime).Milliseconds()
	apiLatency := s.HeartbeatLatency().Milliseconds()
	content := fmt.Sprintf("API Latency: %d ms\nLatency (Processing time): %d ms", apiLatency, delay)
	if _, err := s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{Content: &content}); err != nil {
		fmt.Printf("Error editing interaction response: %v\n", err)
	}
}
