package abuseIp

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

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
	req.Header.Add("Key", "c65c69ea981e2ac50aa4d6ee603d703d148b06ec50e4e1ecf4f701ed31d10cc6d85ef766023d5a65") //API Key

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

