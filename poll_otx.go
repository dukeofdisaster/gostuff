package main
// Author: dukeofdisaster
// 2019
// Light weight otx pulse polling client. 
// Polls whatever the user is subscribed too based on recent modifications.
// My use case: pick some feeds and cron this little baby and send it to a file.
// TODO: add logging of response status codes and other minutiae

import (
  "fmt"
  "io"
  "io/ioutil"
  "net/http"
  "net/url"
  "os"
  "strings"
  "time"
)
  /*  
    url to hit: https://otx.alienvault.com/api/v1/pulses/subscribed?modified_since=2019-09-27T18%3A19%3A01-07%3A00&limit=20
  */

func main() {
  // make the time diff of whatever you want
  t := time.Now().Add(-10*time.Minute)
  // make it a string we can use in a header
  s := fmt.Sprintf(t.Format(time.RFC3339))

  // Create a client with a sensible timeout
  client := &http.Client{
    Timeout: time.Second * 10,
  }

  // build a url; tailor limit to your needs
  baseurl, err := url.Parse("https://otx.alienvault.com")
  baseurl.Path += "api/v1/pulses/subscribed"
  params := url.Values{}
  params.Add("modified_since", s)
  params.Add("limit", "30")
  baseurl.RawQuery = params.Encode()
  finalurl := baseurl.String()

  req, err :=  http.NewRequest("GET",finalurl, nil)
  req.Header.Set("X-OTX-API-KEY", "YOURAPIKEYHERE")

  // now that request is built, we can pass it to the client
  resp, err := client.Do(req)
  if err != nil {
    err_msg := fmt.Sprintf("Request failed: %s\n", err)
    io.WriteString(os.Stderr, err_msg)
  } else {
    // {"count": 0, "next": null, "results": [], "previous": null}
    data, _ := ioutil.ReadAll(resp.Body)
    data_string := string(data)
    if strings.Contains(data_string[:40], "\"results\": []") {
      io.WriteString(os.Stderr, "Empty poll\n")
    } else {
      io.WriteString(os.Stdout, string(data))
    }
  }
}

