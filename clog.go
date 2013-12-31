package main

import (
  "fmt"
  "flag"
  "net/http"
  "encoding/json"
  "bytes"
  "os"
  "os/user"
)

// Event JSON Message
//  {"event":{"entry":"foo"}}
type Message struct {
  // TO DO what the hell does ` do?
  Event string `json:"entry"`
  User string `json:"user"`
  Hostname string `json:"hostname"`
}

type Entry struct {
  Message *Message `json:"event"`
}

// Main
func main() {
  // Parse command line flags
  messagePtr := flag.String("m", "", "Message to log")
  flag.Parse()

  if *messagePtr == "" {
    fmt.Println("Message can not be blank")
    os.Exit(1)
  }

  // Get Current Username
  usr, err := user.Current()
  if err != nil {
    fmt.Printf("Error: %v\n", err)
  }

  // Get Hostname
  hostname, err := os.Hostname()
  if err != nil {
   fmt.Printf("Error: %v\n", err)
   return
  }

  message := &Message{
    Event: *messagePtr,
    User: usr.Username,
    Hostname: hostname,
  }

  entry := &Entry{
    Message: message,
  }

  // create JSON
  event_entry, err := json.Marshal(entry)
  if err != nil {
    fmt.Printf("Error: %v\n", err)
  }
 
  // print JSON for debuging
  fmt.Println("JSON:", string(event_entry))

 // HTTP request
 client := &http.Client{}
 req, err := http.NewRequest("POST", "http://localhost:3000/api", bytes.NewReader(event_entry))
 
 // Add HTTP Headers
 req.Header.Add("Authorization", "Token token=d298a89be686d366a72a78d92e3e43e8")
 req.Header.Add("Content-Type", "application/json")
 req.Header.Add("Accept", "application/json")

 resp, err := client.Do(req)
 defer resp.Body.Close()

 fmt.Println(resp)
}
