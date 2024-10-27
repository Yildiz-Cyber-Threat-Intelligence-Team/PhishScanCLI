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

func checkDomainAge(createdDate string) int {
	// Tarih formatı: "2018-05-25T22:36:11Z"
	t, err := time.Parse(time.RFC3339, createdDate)
	if err != nil {
		fmt.Println("Tarih parse edilemedi:", err)
		return 0
	}

	now := time.Now()
	daysSinceRegistration := int(now.Sub(t).Hours() / 24)

	return daysSinceRegistration
}

// WHOIS API sorgusu yapan fonksiyon
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

	// Yanıtı parse etme
	var whoisResponse WhoisResponse
	if err := json.Unmarshal(body, &whoisResponse); err != nil {
		return 0, err
	}

	// Domain yaşı kontrolü
	days := checkDomainAge(whoisResponse.WhoisRecord.CreatedDate)
	fmt.Printf("Domain '%s' kayıt tarihi: %s (%d gün önce kaydedildi).\n", whoisResponse.WhoisRecord.DomainName, whoisResponse.WhoisRecord.CreatedDate, days)

	if days < 30 {
		return 5, nil
	}
	return 0, nil
}

