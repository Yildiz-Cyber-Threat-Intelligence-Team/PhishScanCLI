package main

import (
	"bytes"
	"encoding/json"
	"flag"
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

func main() {
	apiURL := "https://www.virustotal.com/api/v3/urls"
	apiKey := "964c04c983e6f0f57f4d5a48e1c663abe9de95485119f376d870629f2e9c854d" // Geçerli API anahtarınızı ekleyin

	// URL'yi komut satırından al
	urlPtr := flag.String("url", "", "kontrol edilecek url")
	flag.Parse()

	if *urlPtr == "" {
		fmt.Println("Lütfen kontrol edilecek bir URL belirtin.")
		return
	}

	// URL'yi encode et
	encodedURL := url.QueryEscape(*urlPtr)
	payload := "url=" + encodedURL

	// URL'yi gönder
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer([]byte(payload)))
	if err != nil {
		fmt.Println("İstek oluşturma hatası", err)
		return
	}

	// Başlıkları ayarla
	req.Header.Add("accept", "application/json")
	req.Header.Add("x-apikey", apiKey) // API anahtarını buraya ekleyin
	req.Header.Add("content-type", "application/x-www-form-urlencoded")

	// İsteği gönder
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("API istek hatası:", err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Yanıt okuma hatası:", err)
		return
	}

	// İlk yanıtı yazdır
	fmt.Println("API Yanıtı:", string(body))

	// Eğer isteğin durumu başarısızsa, durumu kontrol et
	if res.StatusCode != http.StatusOK {
		var sendResponse VTSendResponse
		if err := json.Unmarshal(body, &sendResponse); err == nil {
			fmt.Printf("Hata: %d - %s\n", res.StatusCode, sendResponse.Error.Message)
		} else {
			fmt.Printf("Hata: %d - %s\n", res.StatusCode, http.StatusText(res.StatusCode))
		}
		return
	}

	// Eğer istek başarılıysa, ID'yi kullanarak analizi sorgula
	var sendResponse VTSendResponse
	if err := json.Unmarshal(body, &sendResponse); err != nil {
		fmt.Println("JSON unmarshal hatası:", err)
		return
	}

	// Analiz sonucunu al
	analysisID := sendResponse.Data.ID
	analysisURL := "https://www.virustotal.com/api/v3/analyses/" + analysisID

	analysisReq, err := http.NewRequest("GET", analysisURL, nil)
	if err != nil {
		fmt.Println("Analiz isteği oluşturma hatası:", err)
		return
	}

	// Başlıkları ayarla
	analysisReq.Header.Add("accept", "application/json")
	analysisReq.Header.Add("x-apikey", apiKey) // API anahtarını buraya ekleyin

	// Analiz isteğini gönder
	analysisRes, err := http.DefaultClient.Do(analysisReq)
	if err != nil {
		fmt.Println("Analiz API'si istek hatası:", err)
		return
	}
	defer analysisRes.Body.Close()

	analysisBody, err := io.ReadAll(analysisRes.Body)
	if err != nil {
		fmt.Println("Analiz yanıtı okuma hatası:", err)
		return
	}

	// Analiz sonucunu yazdır
	var analysisResponse VTAnalysisResponse
	if err := json.Unmarshal(analysisBody, &analysisResponse); err != nil {
		fmt.Println("Analiz JSON hatası:", err)
		return
	}

	fmt.Printf("Analiz ID: %s\n", analysisResponse.Data.ID)
	fmt.Printf("Kötü Amaçlı: %d\n", analysisResponse.Data.Attributes.Stats.Malicious)
	fmt.Printf("Belirsiz: %d\n", analysisResponse.Data.Attributes.Stats.Undetected)
	fmt.Printf("Spam: %d\n", analysisResponse.Data.Attributes.Stats.Spam)

	// Phishing durumu belirle
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

	fmt.Printf("Phishing Durumu: %d\n", phishingStatus)
}
