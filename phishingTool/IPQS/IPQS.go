package IPQS

import (
	"encoding/json"
	"fmt"
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

	apiKey = os.Getenv("IPQS_API_KEY")
}

type IPQS struct {
	Key string
}

func NewIPQSClient() *IPQS {
	return &IPQS{Key: apiKey}
}

func (ipqs *IPQS) MaliciousURLScannerAPI(inputURL string, params map[string]string) (map[string]interface{}, error) {
	encodedURL := url.QueryEscape(inputURL)

	apiURL := "https://www.ipqualityscore.com/api/json/url/" + ipqs.Key + "/" + encodedURL

	q := url.Values{}
	for key, value := range params {
		q.Add(key, value)
	}
	apiURL = apiURL + "?" + q.Encode()

	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return result, nil
}

func CheckPhishing(result map[string]interface{}) int {
	if phishingVal, ok := result["phishing"]; ok {
		if phishing, isBool := phishingVal.(bool); isBool {
			if phishing {
				return 1
			} else {
				return -1
			}
		}
	}
	return 0
}
