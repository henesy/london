package mux

import (
	"net"
	"log"
	"github.com/go-irc/irc"
)

type IrcConn struct {
	Client	*irc.Client
	Config	irc.ClientConfig
	
}

// Bridging facilities for IRC
func JoinIrc(srv	string, creds	Credential) (ic IrcConn) {
	conn, err := net.Dial("tcp", srv)
	if err != nil {
		log.Fatalln(err)
	}

	ic.Config = irc.ClientConfig{
		Nick: creds.User,
		Pass: creds.Pass,
		User: creds.User,
		Name: "London Bridge",
		Handler: irc.HandlerFunc(func(c *irc.Client, m *irc.Message) {
			if m.Command == "001" {
				// 001 is a welcome event, so we join channels there
				c.Write("JOIN " + Config.Db["debugChannel"])
			} else if m.Command == "PRIVMSG" && c.FromChannel(m) {
				// Create a handler on all messages.
				c.WriteMessage(&irc.Message{
					Command: "PRIVMSG",
					Params: []string{
						m.Params[0],
						m.Trailing(),
					},
				})
			}
		}),
	}

	// Create the client
	ic.Client = irc.NewClient(conn, ic.Config)
	err = ic.Client.Run()
	if err != nil {
		log.Fatalln(err)
	}

	return
}

