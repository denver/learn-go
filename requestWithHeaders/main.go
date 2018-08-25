package main

import (
	"fmt"
	"net/http"
)

func main() {
  url := "http://data.com"

  var bearer = "Bearer " + <ACCESS TOKEN HERE>
  req, err := http.NewRequest("GET", url, nil)
  req.Header.Add("authorization", bearer)

  client := urlfetch.Client(context)

  resp, err := client.Do(req)
  if err != nil {
  panic(err)
  }
  defer resp.Body.Close()

  body, _ := ioutil.ReadAll(resp.Body)
  writer.Write([]byte(body))

}
