package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
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

func checkDomainAge(createdDate string) bool {
	// Tarih formatı: "2018-05-25T22:36:11Z"
	t, err := time.Parse(time.RFC3339, createdDate)
	if err != nil {
		fmt.Println("Tarih parse edilemedi:", err)
		return false
	}


	now := time.Now()

	return now.Sub(t).Hours() <= 30*24
}

func main() {
	urlFlag := flag.String("url", "", "WHOIS sorgulanacak alan adı")
	flag.Parse()

	if *urlFlag == "" {
		fmt.Println("Bir alan adı girmeniz gerekiyor.")
		os.Exit(1)
	}


	apiURL := fmt.Sprintf("https://www.whoisxmlapi.com/whoisserver/WhoisService?apiKey=%s&domainName=%s&outputFormat=JSON", apiKey, *urlFlag)


	resp, err := http.Get(apiURL)
	if err != nil {
		fmt.Println("API'ye istek gönderilirken hata:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()


	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Yanıt okunurken hata:", err)
		os.Exit(1)
	}


	var whoisResponse WhoisResponse
	if err := json.Unmarshal(body, &whoisResponse); err != nil {
		fmt.Println("Yanıt parse edilemedi:", err)
		os.Exit(1)
	}


	if checkDomainAge(whoisResponse.WhoisRecord.CreatedDate) {
		fmt.Printf("UYARI: Domain '%s' son 30 gün içinde kaydedildi.\n", whoisResponse.WhoisRecord.DomainName)
	} else {
		fmt.Printf("Domain '%s' kayıt tarihi: %s.\n", whoisResponse.WhoisRecord.DomainName, whoisResponse.WhoisRecord.CreatedDate)
	}
}
