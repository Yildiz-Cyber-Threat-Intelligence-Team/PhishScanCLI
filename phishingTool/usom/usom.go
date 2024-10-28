package usom

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Address struct {
	URL         string `json:"url"`
	Type        string `json:"type"`
	Desc        string `json:"desc"`
	Criticality int    `json:"criticality"`
}

type AddressResponse struct {
	TotalCount int       `json:"totalCount"`
	Models     []Address `json:"models"`
}

func NormalizeURL(url string) string {
	url = strings.TrimSpace(url)
	url = strings.TrimPrefix(url, "http://")
	url = strings.TrimPrefix(url, "https://")
	url = strings.TrimPrefix(url, "www.")
	return url
}

func CheckPhishing(url string) (bool, []Address) {
	normalizedURL := NormalizeURL(url)
	apiURL := fmt.Sprintf("https://www.usom.gov.tr/api/address/index?q=%s", normalizedURL)
	resp, err := http.Get(apiURL)
	if err != nil {
		fmt.Println("USOM API hatası:", err)
		return false, nil
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("USOM API yanıt okuma hatası:", err)
		return false, nil
	}
	var addressResp AddressResponse
	if err := json.Unmarshal(body, &addressResp); err != nil {
		fmt.Println("USOM JSON ayrıştırma hatası:", err)
		return false, nil
	}

	return addressResp.TotalCount > 0, addressResp.Models
}
