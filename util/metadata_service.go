package util

import (
  "net/http"
)

func AWSMetadata(path string) (string, error) {
  value := ""

  url_prefix = "http://instance-data/latest/meta-data"
  url = url_prefix + "/" + path

  resp, err := http.Get(url)

  if err != nil {
  }

  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)

  return value, nil
}

