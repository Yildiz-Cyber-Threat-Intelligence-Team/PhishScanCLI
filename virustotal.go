package virustotal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type VTSendResponse struct {
	Data struct {
		ID string `json:"id"`
	} `json:"data"`
	Error struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

type VTAnalysisResponse struct {
	Data struct {
		ID         string `json:"id"`
		Type       string `json:"type"`
		Attributes struct {
			Stats struct {
				Malicious  int `json:"malicious"`
				Undetected int `json:"undetected"`
				Spam       int `json:"spam"`
			} `json:"stats"`
		} `json:"attributes"`
	} `json:"data"`
}

func CheckPhishingVirusTotal(apiKey string, inputURL string) int {
	apiURL := "https://www.virustotal.com/api/v3/urls"
	encodedURL := url.QueryEscape(inputURL)
	payload := "url=" + encodedURL

	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer([]byte(payload)))
	if err != nil {
		fmt.Println("İstek oluşturma hatası", err)
		return 0
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("x-apikey", apiKey)
	req.Header.Add("content-type", "application/x-www-form-urlencoded")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("API istek hatası:", err)
		return 0
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Yanıt okuma hatası:", err)
		return 0
	}

	if res.StatusCode != http.StatusOK {
		var sendResponse VTSendResponse
		if err := json.Unmarshal(body, &sendResponse); err == nil {
			fmt.Printf("Hata: %d - %s\n", res.StatusCode, sendResponse.Error.Message)
		} else {
			fmt.Printf("Hata: %d - %s\n", res.StatusCode, http.StatusText(res.StatusCode))
		}
		return 0
	}

	var sendResponse VTSendResponse
	if err := json.Unmarshal(body, &sendResponse); err != nil {
		fmt.Println("JSON unmarshal hatası:", err)
		return 0
	}

	analysisID := sendResponse.Data.ID
	analysisURL := "https://www.virustotal.com/api/v3/analyses/" + analysisID

	analysisReq, err := http.NewRequest("GET", analysisURL, nil)
	if err != nil {
		fmt.Println("Analiz isteği oluşturma hatası:", err)
		return 0
	}

	analysisReq.Header.Add("accept", "application/json")
	analysisReq.Header.Add("x-apikey", apiKey)

	analysisRes, err := http.DefaultClient.Do(analysisReq)
	if err != nil {
		fmt.Println("Analiz API'si istek hatası:", err)
		return 0
	}
	defer analysisRes.Body.Close()

	analysisBody, err := io.ReadAll(analysisRes.Body)
	if err != nil {
		fmt.Println("Analiz yanıtı okuma hatası:", err)
		return 0
	}

	var analysisResponse VTAnalysisResponse
	if err := json.Unmarshal(analysisBody, &analysisResponse); err != nil {
		fmt.Println("Analiz JSON hatası:", err)
		return 0
	}

	var phishingStatus int
	if analysisResponse.Data.Attributes.Stats.Malicious > 0 {
		phishingStatus = 1 // Kötü amaçlı
	} else if analysisResponse.Data.Attributes.Stats.Spam > 0 {
		phishingStatus = 1 // Spam
	} else if analysisResponse.Data.Attributes.Stats.Undetected > 0 {
		phishingStatus = 0 // Belirsiz
	} else {
		phishingStatus = -1 // Kötü amaçlı değil
	}
	return phishingStatus
}
