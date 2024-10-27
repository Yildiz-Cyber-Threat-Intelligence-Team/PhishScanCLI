package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const apiKey = "AIzaSyDjrV9r8kM-hr3MskbeqziqSXAa2PRrfas"

func main() {
	urlFlag := flag.String("url", "", "Kontrol edilecek URL")
	flag.Parse()

	if *urlFlag == "" {
		fmt.Println("Bir URL girmeniz gerekiyor.")
		os.Exit(1)
	}

	//if !strings.HasPrefix(*urlFlag, "http://") && !strings.HasPrefix(*urlFlag, "https://") {
	//	*urlFlag = "https://" + *urlFlag
	//}

	apiURL := "https://safebrowsing.googleapis.com/v4/threatMatches:find?key=" + apiKey
	reqBody := fmt.Sprintf(`{
		"client": {
			"clientId": "Cuneyt",
			"clientVersion": "1.5.2"
		},
		"threatInfo": {
			"threatTypes": ["MALWARE", "SOCIAL_ENGINEERING"],
			"platformTypes": ["WINDOWS"],
			"threatEntryTypes": ["URL"],
			"threatEntries": [
				{"url": "%s"}
			]
		}
	}`, *urlFlag)

	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer([]byte(reqBody)))
	if err != nil {
		fmt.Println("Request oluşturulurken hata:", err)
		os.Exit(1)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("API'ye istek gönderilirken hata:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode == http.StatusOK {
		if bytes.Contains(body, []byte("matches")) {
			fmt.Println(1) // Phishing
		} else {
			fmt.Println(-1) // Phishing değil
		}
	} else {
		fmt.Println(0) // Bilinmiyor
	}
}
