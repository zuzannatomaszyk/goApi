package models

type Message struct {
	User      string `json:"user"`
	Text      string `json:"text"`
	Timestamp string `json:"timestamp"`
}

type MessagesList struct {
	Messages []Message `json:"messages"`
}
