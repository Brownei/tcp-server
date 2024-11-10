package game

// SELF EXPLAINATRY YES?
type Room struct {
	Name    string
	Members map[string]*Client
}

func (r *Room) Broadcast(sender *Client, msg string) {
	for key, client := range r.Members {
		if key != sender.Conn.LocalAddr().String() {
			client.Msg(msg)
		}
	}
}
