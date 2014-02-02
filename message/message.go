package message

import (
	"encoding/json"
	"log"
	"os"
	"os/user"
)

// Message JSON
type MessageBody struct {
	Event    string `json:"entry"`
	User     string `json:"user"`
	Hostname string `json:"hostname"`
}

// Entry JSON
type Entry struct {
	MessageBody *MessageBody `json:"event"`
}

// Create JSON Byte data
func createJson(entry_data *Entry) []byte {
	json_data, err := json.Marshal(entry_data)
	if err != nil {
		log.Fatal(err)
	}

	return json_data
}

// Generate new event message
func NewMessage(messageStr string) []byte {
	// Current Username
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	// Hostname
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}

	// MessageBody with content
	message := &MessageBody{
		Event:    messageStr,
		User:     usr.Username,
		Hostname: hostname,
	}

	// Event Entry with Message content
	entry := &Entry{
		MessageBody: message,
	}

	return createJson(entry)
}
