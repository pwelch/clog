package transmit

import (
  "log"
  "net/http"
  "bytes"
)

 // HTTP request
func NewRequest(api_url string, api_token string, message []byte) (*http.Response, error) {
  client := &http.Client{}
  request, err := http.NewRequest("POST", api_url, bytes.NewReader(message)) 
   if err != nil {
     log.Fatal(err)
   }
  // Add HTTP Headers
  request.Header.Add("Authorization", "Token token=" + api_token)
  request.Header.Add("Content-Type", "application/json")
  request.Header.Add("Accept", "application/json")

  // Send HTTP Request to API Server
  response, err := client.Do(request)
  defer response.Body.Close()

  return response, err
}
