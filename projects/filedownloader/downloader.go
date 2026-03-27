package filedownloader

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type Result struct {
	URL      string
	FileName string
	Size     float64
	Duration time.Duration
	Error    error
}

func downloadFile(url string, ch chan Result, wg *sync.WaitGroup) {
	defer wg.Done()
	fileName := ""
	kbWritten := 0.0
	var timeTaken time.Duration
	startTime := time.Now()

	fmt.Printf("Started Fetching URL %s response at %s\n", url, startTime)
	response, err := http.Get(url)
	if err != nil {
		ch <- Result{URL: url, Error: err}
		return
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		ch <- Result{
			URL:   url,
			Error: fmt.Errorf("Invalid Response from URL %d", response.StatusCode),
		}
		return
	} else {
		destDir := "./projects/filedownloader/Downloads"
		if err = os.MkdirAll(destDir, 0777); err != nil {
			ch <- Result{URL: url, Error: err}
			return
		}

		fileName = filepath.Base(url)
		filePath := filepath.Join(destDir, fileName)
		out, err := os.Create(filePath)
		if err != nil {
			ch <- Result{URL: url, Error: err}
			return
		}
		defer out.Close()

		bytesWritten, err := io.Copy(out, response.Body)
		if err != nil {
			ch <- Result{URL: url, Error: err}
			return
		}
		kbWritten = float64(bytesWritten) / 1024
		timeTaken = time.Since(startTime)
		fmt.Printf("Kilo Bytes written to file %s are %.2f kb\n", fileName, kbWritten)
		fmt.Printf("Completed fetching url %s in %s\n", url, time.Since(startTime))
	}
	ch <- Result{
		URL:      url,
		FileName: fileName,
		Size:     kbWritten,
		Duration: timeTaken,
		Error:    err,
	}
}

func concurrentDownloader(urls []string) {
	var wg sync.WaitGroup
	resultCh := make(chan Result)

	for _, url := range urls {
		wg.Add(1)
		go downloadFile(url, resultCh, &wg)
	}

	go func() {
		wg.Wait()
		close(resultCh)
	}()

	for res := range resultCh {
		fmt.Printf("-- Results --\nURL Processed: %s\nFilename: %s\nSize: %.2f\nTime Taken: %+v\nError: %+v\n",
			res.URL,
			res.FileName,
			res.Size,
			res.Duration,
			res.Error,
		)
	}
}

func TriggerDownload() {
	urls := []string{
		"https://go.dev/blog/go-brand/Go-Logo/PNG/Go-Logo_Blue.png",
		"https://wallpaperaccess.com/full/406168.jpg",
	}
	concurrentDownloader(urls)
}
