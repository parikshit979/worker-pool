package tasks

import (
	"context"
	"fmt"
	"time"
)

// WebCrawlerTask represents a task that crawls a web page.
// It implements the Task interface.
type WebCrawlerTask struct {
	URL string
}

// NewWebCrawlerTask creates a new WebCrawlerTask with the specified URL.
func NewWebCrawlerTask(url string) *WebCrawlerTask {
	return &WebCrawlerTask{
		URL: url,
	}
}

// Execute crawls the web page at the specified URL.
func (w *WebCrawlerTask) Execute(ctx context.Context) error {
	time.Sleep(5 * time.Second) // Simulate a delay for crawling the web page.
	// Simulate crawling a web page.
	fmt.Printf("Crawling web page at %s\n", w.URL)
	return nil
}
