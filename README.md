## Phishing Detection CLI Tool

This is a command-line tool designed to detect phishing URLs by analyzing them with multiple APIs and providing a risk score if no API flags it directly. The tool also has a domain-similarity feature to find potential spoofed domains. You can find the Turkish readme below.
### Features

- **URL Analysis**: Detects if a URL is a phishing site by using several APIs.
- Whois Query: It examines the whois query of the given URL and adds a risk score based on the domain acquisition date.
- **Risk Scoring**: If no API flags the URL, a risk score is calculated based on various heuristics.
- **Domain Similarity**: Lists potential spoofed domains using an approach similar to dnstwist.

### Installation

1. **Prerequisites**:
    - Ensure you have [Go](https://golang.org/dl/) installed.
2. **Clone the Repository**:
    
    ```bash
    git clone <repository-url>
    cd <repository-directory>
    ```
    
3. **Environment Variables**:

## **Setting Up the .env File**

After obtaining your API keys, follow these steps to set up your `.env` file for secure storage of sensitive information:

1. Install the `godotenv` Package

This package allows you to load environment variables from a `.env` file into your Go application.

**Install the Package**:

```bash

go get github.com/joho/godotenv

```

1. In the root directory of your project, create a new file named `.env`.
2. Open the `.env` file and insert the following keys, replacing the placeholders with your actual API keys:
    
    ```
    ABUSEIP_API_KEY=your_abuseip_api_key
    GOOGLESB_API_KEY=your_google_safe_browsing_api_key
    IPQS_API_KEY=your_ipqs_api_key
    VIRUSTOTAL_API_KEY=your_virustotal_api_key
    WHOIS_API_KEY=your_whois_api_key
    
    ```
    
3. Save the file. The `.env` file should be listed in `.gitignore` to prevent it from being uploaded to version control.

### Usage

- **URL Check**:
    
    ```bash
    go run main.go -u <URL>
    ```
    
    This command will check if the given URL is phishing or not based on API results and a calculated risk score.
    
- **Domain Similarity Check**:
    
    ```bash
    go run main.go -s <URL>
    ```
    
    This command lists possible spoofed domains similar to the provided URL.
    

### Example Output

- **Phishing Detected**:
    
    ```
    URL not found in IPQualityScore
    URL not found in USOM
    URL not found in Google Safe Browsing
    URL found as phishing in VirusTotal
    URL not found in Google Safe Browsing
    URL not found in AbuseIP
    ```
    
- **Not Phishing**:
    
    ```
    URL not found in IPQualityScore
    URL not found in USOM
    URL not found in Google Safe Browsing
    URL not found in VirusTotal
    URL not found in AbuseIP
    Risk: HTTPS not used
    Domain 'example.com' registration date:  (registered 101 days ago).
    
    ```
    
- **Risk Score**:
    
    ```
    Potential phishing site (Risk Point: 7)
    ```
    

---

## Phishing Tespit CLI Aracı

Bu komut satırı aracı, URL'leri çeşitli API'lerle analiz ederek phishing olup olmadığını tespit etmek ve doğrudan tespit edilemeyen URL'ler için risk puanı hesaplamak için tasarlanmıştır. Araç ayrıca potansiyel olarak taklit edilen alan adlarını bulmak için bir alan adı benzerlik özelliği sunar.

### Özellikler

- **URL Analizi**: Çeşitli API'leri kullanarak bir URL'nin phishing olup olmadığını tespit eder.
- **Whois Sorgusu**: Verilen URL’in whois sorgusunu inceleyerek, domain alınma tarihine göre risk puanı ekler.
- **Risk Puanı**: Hiçbir API URL'yi doğrudan işaretlemezse, çeşitli ölçütlere göre bir risk puanı hesaplanır.
- **Alan Adı Benzerliği**: dnstwist benzeri bir yöntem kullanarak potansiyel olarak taklit edilebilecek alan adlarını listeler.

### Kurulum

1. **Gereksinimler**:
    - [Go](https://golang.org/dl/) yüklü olduğundan emin olun.
2. **Depoyu Klonlayın**:
    
    ```bash
    git clone <repository-url>
    cd <repository-directory>
    
    ```
    
3. **Ortam Değişkenleri**:

1. `godotenv` Eklentisini Yükleyin

Bu paket, `.env` dosyasını okuyarak ortam değişkenlerini Go uygulamanıza yüklemenizi sağlar.

**Paketi Yükleyin**:

```bash
bash
Kodu kopyala
go get github.com/joho/godotenv

```

1. **.env Dosyasının Kurulumu**

API anahtarlarınızı aldıktan sonra, hassas bilgileri güvenli bir şekilde saklamak için `.env` dosyanızı kurmak üzere aşağıdaki adımları izleyin:

1. Projenizin ana dizininde `.env` adlı yeni bir dosya oluşturun.
2. `.env` dosyasını açın ve aşağıdaki anahtarları gerçek API anahtarlarınızla değiştirerek ekleyin:
    
    ```
    ABUSEIP_API_KEY=abuseip_api_anahtariniz
    GOOGLESB_API_KEY=google_safe_browsing_api_anahtariniz
    IPQS_API_KEY=ipqs_api_anahtariniz
    VIRUSTOTAL_API_KEY=virustotal_api_anahtariniz
    WHOIS_API_KEY=whois_api_anahtariniz
    
    ```
    
3. Dosyayı kaydedin. `.env` dosyası `.gitignore` dosyasına eklenmelidir, böylece sürüm kontrolüne yüklenmez.

### Kullanım

- **URL Kontrolü**:
    
    ```bash
    go run main.go -u <URL>
    
    ```
    
    Bu komut, verilen URL'nin API sonuçlarına ve hesaplanan bir risk puanına göre phishing olup olmadığını kontrol eder.
    
- **Alan Adı Benzerlik Kontrolü**:
    
    ```bash
    go run main.go -s <URL>
    
    ```
    
    Bu komut, sağlanan URL'ye benzer olabilecek potansiyel olarak taklit edilmiş alan adlarını listeler.
    

### Örnek Çıktılar

- **Phishing Tespit Edildi**:
    
    ```
    URL not found in IPQualityScore
    URL not found in USOM
    URL not found in Google Safe Browsing
    URL found as phishing in VirusTotal
    URL not found in Google Safe Browsing
    URL not found in AbuseIP
    ```
    
- **Phishing Değil**:
    
    ```
    URL not found in IPQualityScore
    URL not found in USOM
    URL not found in Google Safe Browsing
    URL not found in VirusTotal
    URL not found in AbuseIP
    Risk: HTTPS not used
    Domain 'example.com' registration date:  (registered 101 days ago).
    ```
    
- **Risk Puanı**:
    
    ```
    Potential phishing site (Risk Point: 7)
    ```
    

---
