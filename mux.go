package main

// This file adds the Disgord message route multiplexer, aka "command router".
// to the Disgord bot. This is an optional addition however it is included
// by default to demonstrate how to extend the Disgord bot.

import (
	"github.com/henesy/london/x/mux"
)

// Router is registered as a global variable to allow easy access to the
// multiplexer throughout the bot.
var Router = mux.New()

func init() {
	// Register the mux OnMessageCreate handler that listens for and processes
	// all messages received.
	Session.AddHandler(Router.OnMessageCreate)

	// Register the build-in help command.
	Router.Route("help", "Display this message.", Router.Help)

	Router.Route("about", "General information about the bot.", Router.About)

	Router.Route("dump", "Dump config to file.", Router.Dump)

	Router.Route("beer", ":beer:", Router.Beer)

	Router.Route("whiskey", ":tumbler_glass:", Router.Whiskey)

	Router.Route("wine", ":wine_glass:", Router.Wine)

	Router.Route("uptime", "Current bot uptime", Router.Uptime)
}
