package main

import (
	"flag"
	"fmt"
	"net"
	"net/url"
	IPQualityScore "phishingTool/IPQS"
	abuseIp "phishingTool/abuseIP"
	"phishingTool/fishAnimation"
	"phishingTool/googleSB"
	"phishingTool/usom"
	"phishingTool/virustotal"
	"phishingTool/whois"
	"phishingTool/yildizAnimation"
	"regexp"
	"strings"
)

func riskEvaluate(urlStr string) string {
	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		return "Geçersiz URL"
	}

	riskPoint := 0

	if parsedURL.Scheme != "https" {
		riskPoint += 2
		fmt.Println("Risk: HTTPS kullanılmıyor")
	}

	if len(parsedURL.String()) > 50 {
		riskPoint += 2
		fmt.Println("Risk: URL çok uzun")
	}

	shorteners := []string{"bit.ly", "tinyurl.com", "goo.gl", "t.co", "is.gd", "buff.ly",
		"ow.ly", "shorte.st", "adf.ly", "cli.re", "bl.ink", "v.gd", "qr.ae", "post.ly", "u.to",
		"short.ie", "wp.me", "snipr.com", "po.st", "fic.kr", "tweez.me", "lnkd.in", "v.gd"}
	for _, shortener := range shorteners {
		if strings.Contains(parsedURL.Host, shortener) {
			riskPoint += 3
			fmt.Println("Risk: Kısaltılmış URL kullanımı")
			break
		}
	}

	phishingKeys := []string{"secure", "login", "account", "signin", "update", "verify", "password", "aws",
		"payment", "paypal", "confirm", "webscr", "restrict", "unusual", "activity", "suspend", "bank", "microsoft", "cloud"}
	for _, keyword := range phishingKeys {
		if strings.Contains(urlStr, keyword) {
			riskPoint += 2
			fmt.Println("Risk: Phishing bağlantılı kelime kullanımı")
			break
		}
	}

	suspiciousExtensions := []string{".tk", ".ml", ".ga", ".cf", ".gq", ".xyz", ".pw", ".top", ".club",
		".info", ".cc", ".ws", ".buzz", ".space", ".review", ".biz", ".trade", ".bid", ".loan", ".date", ".faith",
		".racing", ".freenom", ".partners", ".ventures", ".cheap", ".guru", ".domains", ".plumbing"}
	for _, ext := range suspiciousExtensions {
		if strings.HasSuffix(parsedURL.Host, ext) {
			riskPoint += 3
			fmt.Println("Risk: Şüpheli domain uzantısı kullanımı")
			break
		}
	}

	host := parsedURL.Hostname()
	if net.ParseIP(host) != nil {
		riskPoint += 3
		fmt.Println("Risk: URL'de IP adresi kullanımı")
	}

	subdomains := strings.Split(parsedURL.Host, ".")
	if len(subdomains) > 3 {
		riskPoint += 2
		fmt.Println("Risk: Şüpheli subdomain kullanımı")
	}

	suspiciousCharPattern := `[@!%&\^\*\(\)\{\}\[\]\\:;\"\'<>,\?\/~]`
	if matched, _ := regexp.MatchString(suspiciousCharPattern, urlStr); matched {
		riskPoint += 2
		fmt.Println("Risk: Şüpheli karakterler içeriyor")
	}

	ageRisk, err := whois.GetDomainAgeRiskPoint(parsedURL.Hostname())
	if err != nil {
		fmt.Println("Domain yaşı kontrolü yapılırken hata:", err)
	} else {
		riskPoint += ageRisk
	}

	if riskPoint >= 7 {
		return "Phishing sitesi olabilir (Risk Puanı: " + fmt.Sprint(riskPoint) + ")"
	}
	return "Phishing sitesi değil (Risk Puanı: " + fmt.Sprint(riskPoint) + ")"
}

func main() {
	yildizAnimation.PrintAnimation()
	fishAnimation.AnimateFish()
	urlPtr := flag.String("u", "", "Kontrol edilecek URL")
	flag.Parse()

	if *urlPtr == "" {
		fmt.Println("\nLütfen geçerli bir URL giriniz.")
		return
	}

	ipqs := IPQualityScore.IPQS{Key: "bBCLuOX94Hag9c0DtlHpj5UZxYgyA9al"}
	params := map[string]string{}
	ipqsResult, err := ipqs.MaliciousURLScannerAPI(*urlPtr, params)
	if err != nil {
		fmt.Println("IPQualityScore kontrolü sırasında hata:", err)
		return
	}
	if ipqsCheck := IPQualityScore.CheckPhishing(ipqsResult); ipqsCheck == 1 {
		fmt.Println("URL IPQualityScore'da phishing olarak bulundu")
		return
	} else if ipqsCheck == -1 {
		fmt.Println("URL IPQualityScore içerisinde bulunamadı")
	}

	usomResult, usomDetails := usom.CheckPhishing(*urlPtr)
	if usomResult {
		fmt.Printf("URL USOM'da phishing olarak bulundu: %v\n", usomDetails)
		return
	} else {
		fmt.Println("URL USOM içerisinde bulunamadı")
	}

	googleSBResult, googleSBDetails := googleSB.CheckPhishingGoogleSB(*urlPtr)
	if googleSBResult == 1 {
		fmt.Printf("URL Google Safe Browsing'de phishing olarak bulundu: %v\n", googleSBDetails)
	} else if googleSBResult == -1 {
		fmt.Println("URL Google Safe Browsing içerisinde bulunamadı")
	} else if googleSBResult == 0 {
		fmt.Println("Google Safe Browsing sonucu belirsiz")
	}

	apiKey := "964c04c983e6f0f57f4d5a48e1c663abe9de95485119f376d870629f2e9c854d" // Replace with your actual API key
	vtResult := virustotal.CheckPhishingVirusTotal(apiKey, *urlPtr)
	if vtResult == 1 {
		fmt.Println("URL VirusTotal'de phishing olarak bulundu")
		return
	} else if vtResult == -1 {
		fmt.Println("URL VirusTotal içerisinde bulunamadı")
	}

	abuseIpResult, abuseIPDetails := abuseIp.CheckURLInAbuseIPDB(*urlPtr)
	if abuseIpResult == 1 {
		fmt.Printf("URL AbuseIP'de phishing olarak bulundu: %v\n", abuseIPDetails)
		return
	} else if abuseIpResult == -1 {
		fmt.Println("URL AbuseIP'de güvenli olarak bulundu")
	} else if abuseIpResult == 0 {
		fmt.Println("URL AbuseIP içerisinde bulunamadı")
	}

	result := riskEvaluate(*urlPtr)
	fmt.Println(result)
}
