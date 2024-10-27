package main

import (
    "bufio"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
    "strings"
)

type Address struct {
    ID          int    json:"id"
    URL         string json:"url"
    Type        string json:"type"
    Desc        string json:"desc"
    Source      string json:"source"
    Date        string json:"date"
    Criticality int    json:"criticality_level"
}

type AddressResponse struct {
    TotalCount int       json:"totalCount"
    Count      int       json:"count"
    Models     []Address json:"models"
}

func checkPhishing(url string) (bool, []Address) {
    apiURL := fmt.Sprintf("https://www.usom.gov.tr/api/address/index?q=%s", url)

    req, err := http.NewRequest("GET", apiURL, nil)
    if err != nil {
        fmt.Println("Hata:", err)
        return false, nil
    }

    req.Header.Add("accept", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        fmt.Println("Hata:", err)
        return false, nil
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("Hata:", err)
        return false, nil
    }

    var addressResp AddressResponse
    err = json.Unmarshal(body, &addressResp)
    if err != nil {
        fmt.Println("Hata:", err)
        return false, nil
    }

    if addressResp.TotalCount > 0 {
        return true, addressResp.Models
    }

    return false, nil
}

func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Lütfen websitesini giriniz: ")
    inputURL, _ := reader.ReadString('\n')
    inputURL = strings.TrimSpace(inputURL)

    phishing, details := checkPhishing(inputURL)
    if !phishing && !strings.HasPrefix(inputURL, "www.") {
        phishing, details = checkPhishing("www." + inputURL)
    }

    if phishing {
        fmt.Printf("Verilen URL '%s' USOM listesinde oltalama/tehlikeli bağlantı olarak görünüyor..\n", inputURL)
        fmt.Println("Bulunan tehlikeli bağlantılar:")
        for _, model := range details {
            fmt.Printf("- %s (Tip: %s, Açıklama: %s, Kritiklik Seviyesi: %d)\n",
                model.URL, model.Type, model.Desc, model.Criticality)
        }
    } else {
        fmt.Printf("Verilen URL '%s' USOM listesinde bulunmamakta ve güvenli görünüyor\n", inputURL)
    }
}
