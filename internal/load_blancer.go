package internal

import (
	"io"
	"log"
	"net/http"
	"net/url"
)

func Forwarder(w http.ResponseWriter, r *http.Request) {
	client := &http.Client{}
	URL, err := url.Parse("https://api1.binance.com/api/v3/exchangeInfo")
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
