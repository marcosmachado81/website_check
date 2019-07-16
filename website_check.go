package main

import (
  "fmt"
  "os"
  "flag"
  "github.com/gocolly/colly"
  "time"
  "net/url"

)

func CheckUrl(u string) {

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
    URLScan(&urlinfo)
    urlinfo.ScanInfoResult.Response = -2
    for urlinfo.ScanInfoResult.Response == -2 {
      time.Sleep(25 * time.Second)
      URLReport(&urlinfo)
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
    //var sites []UrlInfo
    CheckUrl(*urlPtr)

  return
}
