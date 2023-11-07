package internal

import (
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/gokul656/golbe/config"
)

func SetupLb() {
	lb := config.LoadLBConfig().Config

	for key := range lb {
		createProxy(lb[key].Peers[0])
	}
}

func createProxy(targetUrl string) func(http.ResponseWriter, *http.Request) {
	URL, err := url.Parse(targetUrl)
	return func(w http.ResponseWriter, r *http.Request) {
		client := &http.Client{}
		if err != nil {
			log.Println("[ERROR]", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		newReq := cloneRequest(URL, r)

		response, err := client.Do(newReq)
		if err != nil {
			log.Println("[ERROR]", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		resp, err := io.ReadAll(response.Body)
		if err != nil {
			log.Println("[ERROR]", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		responseHeader := w.Header()
		for key, value := range response.Header {
			for _, val := range value {
				responseHeader.Add(key, val)
			}
		}

		w.Write(resp)
	}
}
