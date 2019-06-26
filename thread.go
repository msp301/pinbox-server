package pinbox

// A Thread represents a Message chain stored in the Mailbox.
type Thread struct {
	ID         string    `json:"id"`
	Type       string    `json:"type"`
	Subject    string    `json:"subject"`
	NewestDate int64     `json:"newestDate"`
	OldestDate int64     `json:"oldestDate"`
	Authors    []string  `json:"authors"`
	Messages   []Message `json:"messages"`
}
