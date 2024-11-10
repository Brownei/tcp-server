package game

type commadnId int

const (
	CMD_BROWN commadnId = iota
	CMD_JOIN
	CMD_ROOMS
	CMD_MSG
	CMD_QUIT
)

// EVERY COMMAND HAS A CLIENT WHO HAS CALLED IT WITH ITS ARGUMENTS
type Command struct {
	Id     commadnId
	Client *Client
	Args   []string
}
