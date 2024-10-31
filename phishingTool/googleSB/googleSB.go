package googleSB

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

const apiKey = "AIzaSyDjrV9r8kM-hr3MskbeqziqSXAa2PRrfas"

func CheckPhishingGoogleSB(url string) (int, error) {

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
	}`, url)

	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer([]byte(reqBody)))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return 0, nil
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request to API:", err)
		return 0, nil
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode == http.StatusOK {
		if bytes.Contains(body, []byte("matches")) {
			return 1, nil
		} else {
			return -1, nil // Not Phishing
		}
	}

	return 0, nil
}
