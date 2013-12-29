package main

import (
  "fmt"
  "flag"
  "net/http"
  "log"
  "encoding/json"
  "bytes"
)

// Event JSON Message
//  {"event":{"entry":"foo"}}
type Message struct {
  // TO DO what the hell does ` do?
  Event string `json:"entry"`
}

type Entry struct {
  Message *Message `json:"event"`
}

// Parsing CLI arguments. note, that variables are pointers
var strMessage = flag.String("message", "", "Message to log" )
func init() {
  // example with short version for long flag
  flag.StringVar(strMessage, "m", "", "Message to log")
}

// Main
func main() {
  message := &Message{Event: "clog-alpha"}

  entry := &Entry{
    Message: message,
  }

  // create JSON
  event_entry, err := json.Marshal(entry)
  if err != nil {
    log.Fatal(err)
  }
 
  // print JSON for debuging
  fmt.Println("JSON:", string(event_entry))
  flag.Parse()

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
