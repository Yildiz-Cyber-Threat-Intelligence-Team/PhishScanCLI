package abuseIp

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"github.com/joho/godotenv"
)

var apiKey string

func init() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	apiKey = os.Getenv("ABUSEIP_API_KEY")
}

type AbuseIPDBResponse struct {
	Data struct {
		IsWhitelisted *bool `json:"isWhitelisted"`
	} `json:"data"`
}

func CheckURLInAbuseIPDB(rawURL string) (int, error) {

	apiURL := "https://api.abuseipdb.com/api/v2/check"

	queryParams := url.Values{}
	queryParams.Add("url", rawURL)
	queryParams.Add("maxAgeInDays", "90")

	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return 0, err
	}

	req.URL.RawQuery = queryParams.Encode()

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Key", apiKey) //API Key

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	var decodedResponse AbuseIPDBResponse
	err = json.Unmarshal(body, &decodedResponse)
	if err != nil {
		return 0, err
	}

	if decodedResponse.Data.IsWhitelisted == nil {
		return 0, nil
	} else if *decodedResponse.Data.IsWhitelisted {
		return -1, nil
	} else {
		return 1, nil
	}
}
