package mux

import (
	"github.com/bwmarrin/discordgo"
)

// Display basic information about the bot
func (m *Mux) About(ds *discordgo.Session, dm *discordgo.Message, ctx *Context) {
	resp := ""

	resp += `Hi!
My name is London.
	
My source code is located at https://github.com/henesy/london. I am written in Go (https://tour.golang.org).
	`

	resp += "\n"

	ds.ChannelMessageSend(dm.ChannelID, resp)

	return
}

