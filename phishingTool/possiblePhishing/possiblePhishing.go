package possiblePhishing

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)

// Function to generate possible variations
func generateVariations(domain string) []string {
	parts := strings.Split(domain, ".")
	if len(parts) < 2 {
		log.Println("Invalid domain format.")
		return nil
	}

	domainName, tld := parts[0], parts[1]
	var variations []string
	variations = append(variations, swapSimilarCharacters(domainName+"."+tld)...)
	variations = append(variations, duplicateCharacters(domainName, tld)...)
	variations = append(variations, adjacentCharacterSwap(domainName, tld)...)
	variations = append(variations, missingCharacter(domainName, tld)...)
	variations = append(variations, reverseAdjacentCharacters(domainName, tld)...)
	variations = append(variations, addSubdomain(domainName, tld)...)
	variations = append(variations, changeTLD(domainName)...)
	variations = append(variations, addExtraCharacters(domainName, tld)...)
	variations = append(variations, homoglyphCharacters(domainName, tld)...)
	return variations
}

// Generating domain variations with similar characters
func swapSimilarCharacters(domain string) []string {
	similarChars := map[rune][]rune{
		'a': {'à', 'á', 'â', 'ä', 'ã', 'å', 'α', 'а'},
		'b': {'ɓ', 'β', 'в'},
		'c': {'ç', 'ć', 'č', 'ċ', 'с'},
		'd': {'đ', 'ɗ', 'д'},
		'e': {'è', 'é', 'ê', 'ë', 'ε', 'е'},
		'f': {'ƒ'},
		'g': {'ĝ', 'ğ', 'ǧ', 'ġ', 'г'},
		'h': {'ħ', 'н'},
		'i': {'ì', 'í', 'î', 'ï', 'ι', 'і'},
		'j': {'ј'},
		'k': {'ķ', 'к'},
		'l': {'ł', '1', 'I', 'ⅼ'},
		'm': {'м'},
		'n': {'ñ', 'ń', 'η', 'п'},
		'o': {'ò', 'ó', 'ô', 'ö', 'õ', 'ο', 'о'},
		'p': {'þ', 'ρ', 'р'},
		'q': {'ԛ'},
		'r': {'ř', 'г'},
		's': {'ş', 'ś', 'š', 'ѕ'},
		't': {'ť', 'т'},
		'u': {'ù', 'ú', 'û', 'ü', 'υ', 'у'},
		'v': {'ν'},
		'w': {'ŵ'},
		'x': {'х'},
		'y': {'ý', 'ÿ', 'у'},
		'z': {'ž', 'ź', 'ż', 'ʐ'},
		'0': {'ο', 'О', '0'},
		'1': {'ⅼ', 'І'},
		'2': {'２'},
		'3': {'з'},
		'4': {'４'},
		'5': {'５'},
		'6': {'６'},
		'7': {'７'},
		'8': {'８'},
		'9': {'９'},
	}

	var variations []string
	runes := []rune(domain)
	for i, char := range runes {
		if simChars, ok := similarChars[char]; ok {
			for _, simChar := range simChars {
				newRunes := make([]rune, len(runes))
				copy(newRunes, runes)
				newRunes[i] = simChar
				variations = append(variations, string(newRunes))
			}
		}
	}
	return variations
}

// Create variations with repeating characters
func duplicateCharacters(domainName, tld string) []string {
	var variations []string
	for i := 0; i < len(domainName); i++ {
		newDomain := domainName[:i] + string(domainName[i]) + domainName[i:] + "." + tld
		variations = append(variations, newDomain)
	}
	return variations
}

// Creating a variation by swapping two characters next to each other
func reverseAdjacentCharacters(domainName, tld string) []string {
	var variations []string
	for i := 0; i < len(domainName)-1; i++ {
		newDomain := domainName[:i] + string(domainName[i+1]) + string(domainName[i]) + domainName[i+2:] + "." + tld
		variations = append(variations, newDomain)
	}
	return variations
}

// Missing character variations
func missingCharacter(domainName, tld string) []string {
	var variations []string
	for i := 0; i < len(domainName); i++ {
		newDomain := domainName[:i] + domainName[i+1:] + "." + tld
		variations = append(variations, newDomain)
	}
	return variations
}

// Variations with extra characters
func addExtraCharacters(domainName, tld string) []string {
	extraChars := []string{"-", "1", "0"}
	var variations []string
	for _, char := range extraChars {
		variations = append(variations, domainName+char+"."+tld, char+domainName+"."+tld)
	}
	return variations
}

// Variations with homoglyph characters
func homoglyphCharacters(domainName, tld string) []string {
	homoglyphs := map[rune]rune{'o': 'ο', 'a': 'а', 'e': 'е'}
	var variations []string
	runes := []rune(domainName)
	for i, char := range runes {
		if glyph, ok := homoglyphs[char]; ok {
			newRunes := make([]rune, len(runes))
			copy(newRunes, runes)
			newRunes[i] = glyph
			variations = append(variations, string(newRunes)+"."+tld)
		}
	}
	return variations
}

// Variations by adding subdomains
func addSubdomain(domainName, tld string) []string {
	subdomains := []string{"www", "login", "secure", "account"}
	var variations []string
	for _, sub := range subdomains {
		variations = append(variations, sub+"."+domainName+"."+tld)
	}
	return variations
}

// Variations with different TLDs
func changeTLD(domainName string) []string {
	tlds := []string{
		"com", "net", "org", "info", "co",
		"biz", "xyz", "club", "online",
		"site", "shop", "top", "pro",
		"tech", "click",
	}
	var variations []string
	for _, tld := range tlds {
		variations = append(variations, domainName+"."+tld)
	}
	return variations
}

// Variations by swapping two neighboring characters
func adjacentCharacterSwap(domainName, tld string) []string {
	var variations []string
	for i := 0; i < len(domainName)-1; i++ {
		newDomain := domainName[:i] + string(domainName[i+1]) + string(domainName[i]) + domainName[i+2:]
		variations = append(variations, newDomain+"."+tld)
	}
	return variations
}

// Checking the domain by sending an HTTP request
func checkDomain(domain string) bool {
	return checkWithProtocol("https://"+domain) || checkWithProtocol("http://"+domain) || checkDNS(domain)
}

// helper function for HTTP request
func checkWithProtocol(url string) bool {
	client := &http.Client{
		Timeout: 5 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	resp, err := client.Get(url)
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	return resp.StatusCode == http.StatusOK
}

// Checking with DNS (this is also an additional verification method)
func checkDNS(domain string) bool {
	_, err := net.LookupHost(domain)
	return err == nil
}

func CheckPhishing(url string) {
	domain := url
	variations := generateVariations(domain)

	file, err := os.Create("suspectUrls.txt")
	if err != nil {
		fmt.Println("Dosya oluşturulamadı:", err)
		return
	}
	defer file.Close()

	for _, variation := range variations {
		if checkDomain(variation) {
			_, err := file.WriteString(variation + "\n")
			if err != nil {
				fmt.Println("Dosyaya yazılamadı:", err)
			}
		}
	}
}
