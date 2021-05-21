package main

import (
	"flag"
	"fmt"
	"github.com/bwmarrin/discordgo"
"github.com/henesy/london/x/mux"
	"log"
	"os"
	"os/signal"
	sc "strconv"
	"syscall"
	"time"
)

const Version = "v0.1.1"

// Session is declared in the global space so it can be easily used
// throughout this program.
// In this use case, there is no error that would be returned.
var Session, _ = discordgo.New()

// Read in all configuration options from both environment variables and
// command line arguments.
func init() {

	// Discord Authentication Token
	// Have to prefix "Bot [Token Here]" or 401 Forbidden
	Session.Token = os.Getenv("LONDON_TOKEN")
	if Session.Token == "" {
		flag.StringVar(&Session.Token, "t", "", "Discord Authentication Token")
	}
}

func main() {

	// Declare any variables needed later.
	var err error

	// Print out a fancy logo!
	fmt.Printf(` 
            __
           (  \
      __   \  '\
     (  "-_ \ .-'----._
     '-_  "v"         "-
	"Y'             ".
	 |                |
	 |        o     o |
	 |          .<>.  |
	  \         "Ll"  |
	   |             .'
	   |             |
	   (             /
	  /'\         . \
	  "--^--.__,\_)-'   %-16s\/`+"\n\n", Version)

	// Parse command line arguments
	flag.Parse()

	// Verify a Token was provided
	if Session.Token == "" {
		log.Println("You must provide a Discord authentication token.")
		return
	}

	// Verify the Token is valid and grab user information
	Session.State.User, err = Session.User("@me")
	if err != nil {
		log.Printf("error fetching user information, %s\n", err)
	}

	// Open a websocket connection to Discord
	err = Session.Open()
	if err != nil {
		log.Printf("error opening connection to Discord, %s\n", err)
		os.Exit(1)
	}

	// Init boot vars
	mux.StartTime = time.Now()

	// Init Mux daemons
	mux.Config.Init(Session)

	Session.UpdateGameStatus(0, "Bridge is Falling Down")

	// Wait for a CTRL-C
	log.Printf(`Now running on PID ` + sc.Itoa(os.Getpid()) + `. Press CTRL-C to exit.`)
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Clean up
	Session.Close()

	// Exit Normally.
}
