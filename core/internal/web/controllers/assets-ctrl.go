package controllers

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	"core/internal/plugins"
)

const polyfillBaseURL = "https://cdnjs.cloudflare.com/polyfill/v3/polyfill.min.js"

func NewAssetsCtrl(g *plugins.CoreGlobals) *AssetsCtrl {
	return &AssetsCtrl{g}
}

type AssetsCtrl struct {
	g *plugins.CoreGlobals
}

func (ctrl *AssetsCtrl) GetFavicon(w http.ResponseWriter, r *http.Request) {
	contents, err := os.ReadFile(ctrl.g.CoreAPI.Utl.Resource("assets/images/default-favicon-32x32.png"))
	if err != nil {
		ctrl.g.CoreAPI.HttpAPI.HttpResponse().Error(w, r, err, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "image/png")
	w.Write(contents)
}

func (ctrl *AssetsCtrl) Polyfill(w http.ResponseWriter, r *http.Request) {
	// Parse the polyfill.io base URL and append the query string from the client request.
	proxyURL, err := url.Parse(polyfillBaseURL)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Error parsing polyfill base URL:", err)
		return
	}

	// Append the client's query parameters to the proxy URL.
	if r.URL.RawQuery != "" {
		proxyURL.RawQuery = r.URL.RawQuery
	}

	// Create a new HTTP request for the polyfill server.
	proxyReq, err := http.NewRequest(r.Method, proxyURL.String(), r.Body)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Error creating request for polyfill server:", err)
		return
	}

	// Forward all headers from the original request.
	for key, values := range r.Header {
		for _, value := range values {
			proxyReq.Header.Add(key, value)
		}
	}

	// Make the HTTP request to the polyfill server.
	client := &http.Client{}
	resp, err := client.Do(proxyReq)
	if err != nil {
		http.Error(w, "Failed to fetch polyfill", http.StatusInternalServerError)
		log.Println("Error fetching polyfill:", err)
		return
	}
	defer resp.Body.Close()

	// Copy the response headers from the polyfill server to the client.
	for key, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}

	// Set the status code from the polyfill server response.
	w.WriteHeader(resp.StatusCode)

	// Stream the response body from the polyfill server to the client.
	_, err = io.Copy(w, resp.Body)
	if err != nil {
		log.Println("Error streaming response to client:", err)
	}
}
