package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/brownei/game-server/cmd/game"
)

type Server struct {
	Rooms    map[string]*game.Room
	Commands chan game.Command
}

func NewServer() *Server {
	return &Server{
		Rooms:    make(map[string]*game.Room),
		Commands: make(chan game.Command),
	}
}

func (s *Server) Run() {
	for cmd := range s.Commands {
		switch cmd.Id {
		case game.CMD_MSG:
			s.messages(cmd.Client, cmd.Args)
		case game.CMD_BROWN:
			s.brown(cmd.Client, cmd.Args)
		case game.CMD_JOIN:
			s.joinRooms(cmd.Client, cmd.Args)
		case game.CMD_QUIT:
			s.quit(cmd.Client)
		case game.CMD_ROOMS:
			s.listRooms(cmd.Client)
		}
	}
}

func (s *Server) NewServerClient(conn net.Conn) {
	log.Printf("New connection from client: %s", conn.LocalAddr())
	client := game.NewClient(conn, s.Commands, "anonymous")

	client.ReadInput()
}

func (s *Server) messages(c *game.Client, args []string) {
	if c.Room == nil {
		c.Err(errors.New("You must join a room first"))
		return
	}

	c.Room.Broadcast(c, c.Brown+": "+strings.Join(args[1:], " "))
}

func (s *Server) brown(c *game.Client, args []string) {
	c.Brown = args[1]
	c.Msg(fmt.Sprintf("Hello, Your name is %s", c.Brown))
}

func (s *Server) joinRooms(c *game.Client, args []string) {
	roomName := args[1]
	r, ok := s.Rooms[roomName]
	if !ok {
		r = &game.Room{
			Name:    roomName,
			Members: make(map[string]*game.Client),
		}

		s.Rooms[roomName] = r
	}

	r.Members[c.Conn.RemoteAddr().String()] = c
	s.quitCurrentRoom(c)
	c.Room = r

	r.Broadcast(c, fmt.Sprintf("%s has joined room: %s", c.Brown, roomName))
	c.Msg(fmt.Sprintf("Welcome to %s, %s", roomName, c.Brown))
}

func (s *Server) quit(c *game.Client) {
	log.Println("Client has quit room")

	s.quitCurrentRoom(c)

	c.Msg("Good Riddance")
	c.Conn.Close()
}

func (s *Server) listRooms(c *game.Client) {
	var rooms []string

	for _, room := range s.Rooms {
		rooms = append(rooms, room.Name)
	}

	c.Msg(fmt.Sprintf("List of Rooms: %s", strings.Join(rooms, ", ")))
}

func (s *Server) quitCurrentRoom(c *game.Client) {
	if c.Room != nil {
		delete(c.Room.Members, c.Conn.LocalAddr().String())

		c.Room.Broadcast(c, fmt.Sprintf("%s has left the server", c.Brown))
	}
}
