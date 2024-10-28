package whois

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const apiKey = "at_kuDHGYSAMTMjkpB334ELf1OJt8hHw"

type WhoisRecord struct {
	CreatedDate string `json:"createdDate"`
	DomainName  string `json:"domainName"`
}

type WhoisResponse struct {
	WhoisRecord WhoisRecord `json:"WhoisRecord"`
}

// CheckDomainAge attempts to parse the creation date using multiple formats.
func checkDomainAge(createdDate string) int {
	// Updated formats to include cases with `T` separator and timezone offsets
	formats := []string{
		time.RFC3339,
		"2006-01-02T15:04:05-0700",  // With T separator and offset
		"2006-01-02T15:04:05",       // Without timezone
		"2006-01-02 15:04:05 -0700", // Space separator with offset
		"2006-01-02",                // Date only
	}

	var t time.Time
	var err error

	// Try parsing with each format until successful
	for _, format := range formats {
		t, err = time.Parse(format, createdDate)
		if err == nil {
			break
		}
	}
	if err != nil {
		fmt.Println("Date parsing failed:", err)
		return 0
	}

	now := time.Now()
	daysSinceRegistration := int(now.Sub(t).Hours() / 24)

	return daysSinceRegistration
}

// WHOIS API query function
func GetDomainAgeRiskPoint(domain string) (int, error) {
	apiURL := fmt.Sprintf("https://www.whoisxmlapi.com/whoisserver/WhoisService?apiKey=%s&domainName=%s&outputFormat=JSON", apiKey, domain)
	resp, err := http.Get(apiURL)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	// Parse the response
	var whoisResponse WhoisResponse
	if err := json.Unmarshal(body, &whoisResponse); err != nil {
		return 0, err
	}

	// Check domain age
	days := checkDomainAge(whoisResponse.WhoisRecord.CreatedDate)
	fmt.Printf("Domain '%s' registration date: %s (registered %d days ago).\n", whoisResponse.WhoisRecord.DomainName, whoisResponse.WhoisRecord.CreatedDate, days)

	if days < 30 {
		return 5, nil
	}
	return 0, nil
}
