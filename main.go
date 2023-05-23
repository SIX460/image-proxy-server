package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
)

var (
	inMemory bool
	port     string
	imgCache map[string][]byte
	imgMutex sync.RWMutex
)

func main() {
	flag.BoolVar(&inMemory, "in-memory", false, "store images in memory")
	flag.StringVar(&port, "port", "3777", "the port number to listen on")
	flag.Parse()

	imgCache = make(map[string][]byte)
	if !inMemory {
		if _, err := os.Stat("./img"); os.IsNotExist(err) {
			os.Mkdir("./img", 0755)
		}
	}

	http.HandleFunc("/proxy", func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.Query().Get("url")
		hash := fmt.Sprintf("%x", md5.Sum([]byte(url)))

		var img []byte
		var err error

		if inMemory {
			imgMutex.RLock()
			img, _ = imgCache[hash]
			imgMutex.RUnlock()
		} else {
			img, err = ioutil.ReadFile(fmt.Sprintf("./img/%s", hash))
			if err != nil && !os.IsNotExist(err) {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

		if img == nil {
			client := &http.Client{}
			req, err := http.NewRequest("GET", url, nil)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
			req.Header.Set("Accept-Encoding", "gzip, deflate, br")
			req.Header.Set("Accept-Language", "en-GB,en;q=0.9")
			req.Header.Set("Sec-Ch-Ua", `"Google Chrome";v="113", "Chromium";v="113", "Not-A.Brand";v="24"`)
			req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
			req.Header.Set("Sec-Ch-Ua-Platform", "macOS")
			req.Header.Set("Sec-Fetch-Dest", "document")
			req.Header.Set("Sec-Fetch-Mode", "navigate")
			req.Header.Set("Sec-Fetch-Site", "none")
			req.Header.Set("Sec-Fetch-User", "?1")
			req.Header.Set("Upgrade-Insecure-Requests", "1")
			req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36")


			resp, err := client.Do(req)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			defer resp.Body.Close()

			img, err = ioutil.ReadAll(resp.Body)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			if inMemory {
				imgMutex.Lock()
				imgCache[hash] = img
				imgMutex.Unlock()
			} else {
				err = ioutil.WriteFile(fmt.Sprintf("./img/%s", hash), img, 0644)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
			}
		}

		w.Header().Set("Content-Type", "image/jpeg")
		w.Write(img)
	})

	fmt.Printf("Running on port %s\n", port)
	http.ListenAndServe(":"+port, nil)
}
