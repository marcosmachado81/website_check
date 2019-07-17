package main

import (
  "fmt"
  "os"
  "flag"
  "github.com/gocolly/colly"
  "time"
  "net/url"
  "encoding/json"
  "io/ioutil"

)

func CheckUrl(u string, config Config) {

  domain, _ := url.Parse(u)
  fmt.Printf("Scanning address: %s\n", u)
  fmt.Println("The scan will be limited by domain ", domain.Hostname())
	c := colly.NewCollector(
		colly.AllowedDomains(domain.Hostname()),
		colly.MaxDepth(1),
	)

	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		c.Visit(e.Request.AbsoluteURL(link))
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
    fmt.Println("checking ", r.URL)
    var urlinfo UrlInfo

    urlinfo.Address = r.URL.String()
    URLScan(&urlinfo,config)
    urlinfo.ScanInfoResult.Response = -2
    for urlinfo.ScanInfoResult.Response == -2 {
      time.Sleep(25 * time.Second)
      URLReport(&urlinfo,config)
    }
    if urlinfo.ScanInfoResult.Response == 1 {
      fmt.Printf("\tTotal scans: %d, Treats: %d\n", urlinfo.ScanInfoResult.Total,urlinfo.ScanInfoResult.Positives)
      if urlinfo.ScanInfoResult.Positives > 0 {
        fmt.Printf("\tDetails: %s\n",urlinfo.ScanInfoResult.Permalink)
      }
    } else if urlinfo.ScanInfoResult.Response == 0 {
       fmt.Printf("The item was not present in VirusTotal's dataset")
    }
	})
	c.Visit(u)
}

func main() {

  urlPtr := flag.String("url", "", "the url with http://<domain>/")

  if len(os.Args) < 2 {
        fmt.Println("expected 'url'")
        fmt.Println("\tUsage")
        fmt.Printf("\t%s -url=http://<domain>/\n",os.Args[0])
        os.Exit(1)
  }
  flag.Parse()
    //load Config
    if _,err := os.Stat("config.json"); err == nil {
      cg,_:= ioutil.ReadFile("config.json")
      config := Config{}
      
      _= json.Unmarshal([]byte(cg),&config)
      
      CheckUrl(*urlPtr,config)
    } else if os.IsNotExist(err) {
      fmt.Println("You need create the file config.json with:")
      fmt.Printf("{\n\t\"apikey\":\t\"your apikey\",\n\t\"url_scan\":\t\"https://www.virustotal.com/vtapi/v2/url/scan\",\n\t\"url_report\":\t\"https://www.virustotal.com/vtapi/v2/url/report\"\n}\n")
      
    } else {
        fmt.Println("Something goes wrong")
        fmt.Println(err)
    }
    
    

  return
}
