package main

import (
	"fmt"
	"testing"

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
