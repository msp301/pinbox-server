package pinbox

// Sender represents a Message Transfer Agent (MTA)
type Sender interface {
	Send(msg MessageContent) error
}
