package main

type ScanResult struct {
  Response  int64   `json:"response_code"`
  ScanDate  string  `json:"scan_date"`
  Total     int64   `json:"total"`
  Positives int64   `json:"positives"`
  Permalink string  `json:"permalink"`
  Message   string  `json:"verbose_msg"`
  Sha256    string  `json:"sha256"`
  Md5       string  `json:"md5"`
}

type ScanRequest struct {
  Permalink   string  `json:"permalink"`
  Resource    string  `json:"resource"`
  ResposeCode int64   `json:"response_code"`
  ScanId      string  `json:"scan_id"`
  Message     string  `json:"verbose_msg"`
  Sha256      string  `json:"sha256"`
  ScanDate    string  `json:"scan_date"`
  UrlAddress  string  `json:"url"`
}

type Config struct {
  Apikey    string `json:"apikey"`
  UrlScan   string `json:"url_scan"`
  UrlReport string `json:"url_report"`
}

type UrlInfo struct {
  Address         string
  ScanInfoRequest ScanRequest
  ScanInfoResult  ScanResult
}
