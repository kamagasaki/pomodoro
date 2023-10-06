package main

import (
	"fmt"
	"testing"
)

func TestTimeStamp(t *testing.T) {
	// presensihariini := getPresensiTodayFromPhoneNumber(MongoConn, "6281312000300")
	// fmt.Println(presensihariini)
	url := "https://petapedia.github.io"
	test, msg := CheckURLStatus(url)
	fmt.Println(test, msg)
}
