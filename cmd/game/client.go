package game

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

// EVERY CLIENT HAS A BROWN, CONN, IS IN A ROOM AND HAS COMMANDS TO GIVE TO THE TCP
type Client struct {
	Conn     net.Conn
	Brown    string
	Room     *Room
	Commands chan Command
}

func NewClient(conn net.Conn, serverCommands chan Command, name string) *Client {
	return &Client{
		Conn:     conn,
		Brown:    name,
		Commands: serverCommands,
	}
}

func (c *Client) ReadInput() {
	for {
		// read, _ := bufio.NewReader(c.Conn).ReadByte()
		msg, err := bufio.NewReader(c.Conn).ReadString('\n')
		if err != nil {
			return
		}

		fmt.Printf("Message: %s", msg)
		// fmt.Printf("READ: %v", read)
		msg = strings.Trim(msg, "\r\n")

		args := strings.Split(msg, " ")
		cmd := strings.TrimSpace(args[0])

		switch cmd {
		case "/brown":
			c.Commands <- Command{
				Id:     CMD_BROWN,
				Client: c,
				Args:   args,
			}
		case "/join":
			c.Commands <- Command{
				Id:     CMD_JOIN,
				Client: c,
				Args:   args,
			}
		case "/msg":
			c.Commands <- Command{
				Id:     CMD_MSG,
				Client: c,
				Args:   args,
			}
		case "/quit":
			c.Commands <- Command{
				Id:     CMD_QUIT,
				Client: c,
				Args:   args,
			}
		case "/rooms":
			c.Commands <- Command{
				Id:     CMD_ROOMS,
				Client: c,
				Args:   args,
			}
		default:
			c.Err(fmt.Errorf("Unknown command: %s", cmd))
		}
	}
}

func (c *Client) Err(e error) {
	c.Conn.Write([]byte("ERROR: " + e.Error() + "\n"))
}

func (c *Client) Msg(m string) {
	c.Conn.Write([]byte("MSG: " + m + "\n"))
}
