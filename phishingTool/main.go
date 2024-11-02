package main

import (
	"flag"
	"fmt"
	"net"
	"net/url"
	"os"
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

	"github.com/joho/godotenv"
)

func riskEvaluate(urlStr string) string {
	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		return "Invalid URL"
	}

	riskPoint := 0

	if parsedURL.Scheme != "https" {
		riskPoint += 2
		fmt.Println("Risk: HTTPS not used")
	}

	if len(parsedURL.String()) > 50 {
		riskPoint += 2
		fmt.Println("Risk: URL is too long")
	}

	shorteners := []string{"bit.ly", "tinyurl.com", "goo.gl", "t.co", "is.gd", "buff.ly",
		"ow.ly", "shorte.st", "adf.ly", "cli.re", "bl.ink", "v.gd", "qr.ae", "post.ly", "u.to",
		"short.ie", "wp.me", "snipr.com", "po.st", "fic.kr", "tweez.me", "lnkd.in", "v.gd"}
	for _, shortener := range shorteners {
		if strings.Contains(parsedURL.Host, shortener) {
			riskPoint += 3
			fmt.Println("Risk: Shortened URL detected")
			break
		}
	}

	phishingKeys := []string{"secure", "login", "account", "signin", "update", "verify", "password", "aws",
		"payment", "paypal", "confirm", "webscr", "restrict", "unusual", "activity", "suspend", "bank", "microsoft", "cloud"}
	for _, keyword := range phishingKeys {
		if strings.Contains(urlStr, keyword) {
			riskPoint += 2
			fmt.Println("Risk: Potential phishing keyword detected")
			break
		}
	}

	suspiciousExtensions := []string{".tk", ".ml", ".ga", ".cf", ".gq", ".xyz", ".pw", ".top", ".club",
		".info", ".cc", ".ws", ".buzz", ".space", ".review", ".biz", ".trade", ".bid", ".loan", ".date", ".faith",
		".racing", ".freenom", ".partners", ".ventures", ".cheap", ".guru", ".domains", ".plumbing"}
	for _, ext := range suspiciousExtensions {
		if strings.HasSuffix(parsedURL.Host, ext) {
			riskPoint += 3
			fmt.Println("Risk: Suspicious domain extension detected")
			break
		}
	}

	host := parsedURL.Hostname()
	if net.ParseIP(host) != nil {
		riskPoint += 3
		fmt.Println("Risk: IP address detected in URL")
	}

	subdomains := strings.Split(parsedURL.Host, ".")
	if len(subdomains) > 3 {
		riskPoint += 2
		fmt.Println("Risk: Excessive subdomains detected")
	}

	suspiciousCharPattern := `[@!%&\^\*\(\)\{\}\[\]\\:;\"\'<>,\?\/~]`
	if matched, _ := regexp.MatchString(suspiciousCharPattern, urlStr); matched {
		riskPoint += 2
		fmt.Println("Risk: Suspicious characters detected")
	}

	ageRisk, err := whois.GetDomainAgeRiskPoint(parsedURL.Hostname())
	if err != nil {
		fmt.Println("Error during domain age check", err)
	} else {
		riskPoint += ageRisk
	}

	if riskPoint >= 7 {
		return "Potential phishing site (Risk Point: " + fmt.Sprint(riskPoint) + ")"
	}
	return "Not a phishing site (Risk Point: " + fmt.Sprint(riskPoint) + ")"
}

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	ipqsApiKey := os.Getenv("IPQS_API_KEY")
	virustotalApiKey := os.Getenv("VIRUSTOTAL_API_KEY")

	yildizAnimation.PrintAnimation()
	fishAnimation.AnimateFish()
	urlPtr := flag.String("u", "", "URL to check")
	flag.Parse()

	if *urlPtr == "" {
		fmt.Println("\nPlease enter a valid URL")
		return
	}

	ipqs := IPQualityScore.IPQS{Key: ipqsApiKey}
	params := map[string]string{}
	ipqsResult, err := ipqs.MaliciousURLScannerAPI(*urlPtr, params)
	if err != nil {
		fmt.Println("Error during IPQualityScore check:", err)
		return
	}
	if ipqsCheck := IPQualityScore.CheckPhishing(ipqsResult); ipqsCheck == 1 {
		fmt.Println("URL found as phishing in IPQualityScore")
		return
	} else if ipqsCheck == -1 {
		fmt.Println("URL not found in IPQualityScore")
	}

	usomResult, usomDetails := usom.CheckPhishing(*urlPtr)
	if usomResult {
		fmt.Printf("URL found as phishing in USOM: %v\n", usomDetails)
		return
	} else {
		fmt.Println("URL not found in USOM")
	}

	googleSBResult, googleSBDetails := googleSB.CheckPhishingGoogleSB(*urlPtr)
	if googleSBResult == 1 {
		fmt.Printf("URL found as phishing in Google Safe Browsing: %v\n", googleSBDetails)
	} else if googleSBResult == -1 {
		fmt.Println("URL not found in Google Safe Browsing")
	} else if googleSBResult == 0 {
		fmt.Println("Google Safe Browsing result uncertain")
	}

	apiKey := virustotalApiKey
	vtResult := virustotal.CheckPhishingVirusTotal(apiKey, *urlPtr)
	if vtResult == 1 {
		fmt.Println("URL found as phishing in VirusTotal")
		return
	} else if vtResult == -1 {
		fmt.Println("URL not found in VirusTotal")
	}

	abuseIpResult, abuseIPDetails := abuseIp.CheckURLInAbuseIPDB(*urlPtr)
	if abuseIpResult == 1 {
		fmt.Printf("URL found as phishing in AbuseIP: %v\n", abuseIPDetails)
		return
	} else if abuseIpResult == -1 {
		fmt.Println("URL found safe in AbuseIP")
	} else if abuseIpResult == 0 {
		fmt.Println("URL not found in AbuseIP")
	}

	result := riskEvaluate(*urlPtr)
	fmt.Println(result)
}
