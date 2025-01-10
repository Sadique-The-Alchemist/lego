package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	return url != "waat://furhurterwe.geds"
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://www.google.com",
		"http://www.gypsy.com",
		"waat://furhurterwe.geds",
	}

	want := map[string]bool{
		"http://www.google.com":   true,
		"http://www.gypsy.com":    true,
		"waat://furhurterwe.geds": true,
	}
	got := CheckWebsites(mockWebsiteChecker, websites)
	if reflect.DeepEqual(got, want) {
		t.Fatalf("got %v want %v ", got, want)
	}
}
func slowStubBenchmarkChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}
func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "aurl"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubBenchmarkChecker, urls)
	}
}
