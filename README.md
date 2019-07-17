# website_check

website_check is a golang application for a simple check if a website has threats using the virustotal API.

## Notice

- The Public API is limited to 4 requests per minute.
- The Public API must **not be used in commercial** products or services.


## Usage

The Virus Total API Key would be put in config.json file with virus total urls

```json
{
	"apikey":	"your apikey",
	"url_scan":	"https://www.virustotal.com/vtapi/v2/url/scan",
	"url_report":	"https://www.virustotal.com/vtapi/v2/url/report"
}

```
After just compile
```bash
go build website_check.go website_vt.go website_structs.go
```
And run
```bash
./website_check -url=http://example.com
```
