package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/whatsauth/watoken"
)

func TestTimeStamp(t *testing.T) {
	// presensihariini := getPresensiTodayFromPhoneNumber(MongoConn, "6281312000300")
	// fmt.Println(presensihariini)
	url := "https://petapedia.github.io"
	test, msg := CheckURLStatus(url)
	fmt.Println(test, msg)
}

func TestInputURLGithub(t *testing.T) {
	// Test with valid URL
	url := "https://petapedia.github.io"
	urlvalid, msgerrurl := CheckURLStatus(url)

	if !urlvalid {
		t.Errorf("Expected URL %s to be valid, but got invalid with message: %s", url, msgerrurl)
	}

	// Test encoding functionality
	var alias = url
	var PrivateKey = "null"

	// Correct parameter order: url (id), alias, privateKey, hours
	hashurl, err := watoken.EncodeforHours(url, alias, PrivateKey, 3)

	if err != nil {
		t.Errorf("Failed to encode URL: %v", err)
	}

	// Print the generated token
	fmt.Println("Generated Token:", hashurl)

	if hashurl == "" {
		t.Error("Expected non-empty hash URL, but got empty string")
	}

	// Test with invalid URL
	invalidURL := "not-a-valid-url"
	urlvalid, msgerrurl = CheckURLStatus(invalidURL)

	if urlvalid {
		t.Errorf("Expected URL %s to be invalid, but got valid with message: %s", invalidURL, msgerrurl)
	}
}

func TestGTmetrixScraping(t *testing.T) {
	url := "https://gtmetrix.com/reports/go-colly.org/q2juSxHU/"
	
	// Tes scraper yang diperbaiki
	performanceData, err := ScrapeGTmetrixData(url)
	if err != nil {
		t.Errorf("Error scraping GTmetrix data: %v", err)
	}
	
	// Cetak hasil
	fmt.Println("Data Performa GTmetrix (Perbaikan):")
	for key, value := range performanceData {
		fmt.Printf("%s: %s\n", key, value)
	}
	
	// Format laporan
	formattedReport := FormatGTmetrixReport(performanceData, url)
	fmt.Println("\nLaporan Terformat:")
	fmt.Println(formattedReport)
	
	// Periksa jika kita mendapatkan setidaknya beberapa metrik
	requiredMetrics := []string{"Grade", "Performance", "Structure", "LCP", "TBT", "CLS"}
	metricsFound := 0
	
	for _, metric := range requiredMetrics {
		if _, ok := performanceData[metric]; ok {
			metricsFound++
		}
	}
	
	// Kita harus menemukan setidaknya 3 dari 6 metrik
	if metricsFound < 3 {
		t.Errorf("Gagal menemukan cukup metrik. Menemukan %d dari 6 metrik yang diperlukan", metricsFound)
	}
	
	// Tes contoh laporan WhatsApp
	sampleReport := CreateSampleReport(performanceData, url)
	fmt.Println("\nContoh Pesan WhatsApp:")
	fmt.Println(sampleReport)
}

func CreateSampleReport(performanceData map[string]string, url string) string {
	// Buat contoh pesan
	msg := "*Myika Pomodoro Report 1 cycle*" +
		"\nHostname : TestHost" +
		"\nIP : https://whatismyipaddress.com/ip/192.168.1.1" +
		"\nJumlah ScreenShoot : 4" +
		"\nYang Dikerjakan :\n|Test milestone" +
		"\n#" + url
	
	// Tambahkan data GTmetrix
	msg += FormatGTmetrixReport(performanceData, url)
	
	return msg
}

func TestScrapeGTmetrixWithChromedp(t *testing.T) {
	// URL GTmetrix untuk testing
	url := "https://gtmetrix.com/reports/go-colly.org/q2juSxHU/"
	
	// Catatan: Kita bisa mengatur timeout, tapi fungsi ScrapeGTmetrixWithChromedp
	// sudah memiliki timeout internal (60 detik)
	startTime := time.Now()
	
	fmt.Println("🧪 Memulai test scraping dengan Chromedp...")
	
	// Jalankan scraping dengan Chromedp
	performanceData, err := ScrapeGTmetrixWithChromedp(url)
	
	// Hitung durasi eksekusi
	duration := time.Since(startTime)
	fmt.Printf("⏱️ Waktu eksekusi: %.2f detik\n", duration.Seconds())
	
	// Cek error
	if err != nil {
		t.Errorf("Error saat scraping dengan Chromedp: %v", err)
		return
	}
	
	// Cek apakah data berhasil didapatkan
	if len(performanceData) == 0 {
		t.Error("Tidak ada data yang berhasil di-scrape")
		return
	}
	
	// Print hasil scraping
	fmt.Println("📊 Data yang berhasil di-scrape:")
	for key, value := range performanceData {
		fmt.Printf("  - %s: %s\n", key, value)
	}
	
	// Verifikasi data kunci yang diharapkan
	requiredMetrics := []string{"Grade", "Performance", "Structure"}
	for _, metric := range requiredMetrics {
		if value, exists := performanceData[metric]; !exists || value == "" {
			t.Errorf("Metric %s tidak ditemukan atau kosong", metric)
		}
	}
	
	// Jika semua data utama ada, test dianggap berhasil
	fmt.Println("✅ Test scraping dengan Chromedp berhasil!")
	
	// Test format output WhatsApp
	formattedReport := FormatGTmetrixReport(performanceData, url)
	fmt.Println("\n📱 Contoh output WhatsApp:")
	fmt.Println(formattedReport)
}
