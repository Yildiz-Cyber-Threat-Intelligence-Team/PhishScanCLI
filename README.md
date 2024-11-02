# PhishScanCLI
PhishScanCLI is a CLI tool for identifying phishing URLs. It analyzes incoming URLs to detect if they belong to known phishing sites and identifies the company being impersonated. This tool provides real-time alerts to enhance online security and help users avoid fraudulent websites.

---

### EN

## Setting Up the .env File

After obtaining your API keys, follow these steps to set up your `.env` file for secure storage of sensitive information:

1. In the root directory of your project, create a new file named `.env`.
2. Open the `.env` file and insert the following keys, replacing the placeholders with your actual API keys:

   ```plaintext
   ABUSEIP_API_KEY=your_abuseip_api_key
   GOOGLESB_API_KEY=your_google_safe_browsing_api_key
   IPQS_API_KEY=your_ipqs_api_key
   VIRUSTOTAL_API_KEY=your_virustotal_api_key
   WHOIS_API_KEY=your_whois_api_key
   ```

3. Save the file. The `.env` file should be listed in `.gitignore` to prevent it from being uploaded to version control.

---

### TR

## .env Dosyasının Kurulumu

API anahtarlarınızı aldıktan sonra, hassas bilgileri güvenli bir şekilde saklamak için `.env` dosyanızı kurmak üzere aşağıdaki adımları izleyin:

1. Projenizin ana dizininde `.env` adlı yeni bir dosya oluşturun.
2. `.env` dosyasını açın ve aşağıdaki anahtarları gerçek API anahtarlarınızla değiştirerek ekleyin:

   ```plaintext
   ABUSEIP_API_KEY=abuseip_api_anahtariniz
   GOOGLESB_API_KEY=google_safe_browsing_api_anahtariniz
   IPQS_API_KEY=ipqs_api_anahtariniz
   VIRUSTOTAL_API_KEY=virustotal_api_anahtariniz
   WHOIS_API_KEY=whois_api_anahtariniz
   ```

3. Dosyayı kaydedin. `.env` dosyası `.gitignore` dosyasına eklenmelidir, böylece sürüm kontrolüne yüklenmez.

--- 