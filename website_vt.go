package main

import (
  "fmt"
  "net/http"
  "encoding/json"
  "net/url"
  "strings"
)

/*const APIKEY      = 
const url_scan    = "https://www.virustotal.com/vtapi/v2/url/scan"
const url_report  = "https://www.virustotal.com/vtapi/v2/url/report"*/

func URLScan(urlinfo *UrlInfo, config Config) {
  hc := http.Client{}

  form := url.Values{}
  form.Add("apikey", config.Apikey)
  form.Add("url", urlinfo.Address)
  req, err := http.NewRequest("POST", config.UrlScan, strings.NewReader(form.Encode()))
  //req.PostForm = form

  req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
  res, err := hc.Do(req)
	if err != nil {
    fmt.Println("Error sending request")
		return
	}

	// Check the response
  if res.StatusCode != http.StatusOK {
		err = fmt.Errorf("bad status: %s", res.Status)
    fmt.Printf("bad status: %s\n", res.Status)
	} else {
    decoder := json.NewDecoder(res.Body)
    err = decoder.Decode(&urlinfo.ScanInfoRequest)
  }
}

func URLReport(urlinfo *UrlInfo, config Config) {
  hc := http.Client{}
  req, err := http.NewRequest("GET", config.UrlReport, nil)

  form := req.URL.Query()
  form.Add("apikey", config.Apikey)
  form.Add("resource", urlinfo.ScanInfoRequest.Resource)
  req.URL.RawQuery = form.Encode()

  req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
  res, err := hc.Do(req)
	if err != nil {
    fmt.Println("Error sending result request")
		return
	}

	// Check the response
  if res.StatusCode != http.StatusOK {
		err = fmt.Errorf("bad status: %s", res.Status)
	} else {
    decoder := json.NewDecoder(res.Body)
    err = decoder.Decode(&urlinfo.ScanInfoResult)

  }


}
