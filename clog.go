package main

import (
  "encoding/json"
  "fmt"
  "flag"
  "io/ioutil"
  "os"
  "github.com/pwelch/clog/message"
  "github.com/pwelch/clog/transmit"
)

// Config File values
type Config struct {
  Server string `json:server`
  Api_key string `json:api_key`
}

func main() {
  // Parse command line flags
  messagePtr := flag.String("m", "", "Message to log")
  configPtr := flag.String("c", "/etc/clog/clog_config.json", "Path to configuration file")
  flag.Parse()

  if *messagePtr == "" {
    fmt.Println("Message can not be blank")
    os.Exit(1)
  }

  // Genereate new entry message
  entry_message := message.NewMessage(*messagePtr)

  // Read configuration file
  fileContents, file_error := ioutil.ReadFile(*configPtr)
  if file_error != nil {
    fmt.Println(file_error)
    os.Exit(1)
  }

  // Store config values
  var conf Config
  err := json.Unmarshal(fileContents, &conf)
  if err != nil {
    fmt.Println(err)
  }

  // HTTP Request to server
  response, err := transmit.NewRequest(conf.Server, conf.Api_key, entry_message)
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

  // Return Success or Failure
  if response.StatusCode == 201 {
    fmt.Println("OK")
  } else {
    fmt.Println("Failed")
    os.Exit(1)
  }
}
