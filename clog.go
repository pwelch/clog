package main

import (
  "fmt"
  "flag"
  "os"
  "github.com/pwelch/clog/message"
  "github.com/pwelch/clog/transmit"
)

func main() {
  // Parse command line flags
  messagePtr := flag.String("m", "", "Message to log")
  flag.Parse()

  if *messagePtr == "" {
    fmt.Println("Message can not be blank")
    os.Exit(1)
  }

  // Genereate new entry message
  entry_message := message.NewMessage(*messagePtr)

  api_url   := "http://localhost:3000/api"
  api_token := "d298a89be686d366a72a78d92e3e43e8"

  // HTTP Request to server
  response, _ := transmit.NewRequest(api_url, api_token, entry_message)
 fmt.Println(response)
 // print JSON for debuging
 // fmt.Println("JSON:", string(event_entry))
}
