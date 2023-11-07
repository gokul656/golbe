package internal

import (
	"net/http"
	"net/url"
)

func cloneRequest(url *url.URL, request *http.Request) *http.Request {
	// We can user http.Request.Clone here, but we cannot modify RequestURI of the request,
	// so it'll panic.
	return &http.Request{
		URL:           url,
		Host:          url.Host,
		Method:        request.Method,
		Header:        request.Header,
		Body:          request.Body,
		ContentLength: request.ContentLength,
		Form:          request.Form,
		PostForm:      request.PostForm,
		MultipartForm: request.MultipartForm,
		TLS:           request.TLS,

		Proto:            request.Proto,
		ProtoMajor:       request.ProtoMajor,
		ProtoMinor:       request.ProtoMinor,
		TransferEncoding: request.TransferEncoding,
		GetBody:          request.GetBody,
		Close:            request.Close,
		Trailer:          request.Trailer,
		RemoteAddr:       request.RemoteAddr,
		Response:         request.Response,
	}
}
