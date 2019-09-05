package m3u8proxy

import (
	"log"
	"net/http"
)

type HTTPServer struct {
	RealHost string
	Proxy    Proxy
}

func (s *HTTPServer) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	url := s.RealHost + request.URL.RequestURI()
	log.Println("proxy ", url)
	b, err := s.Proxy.ReplaceURL(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/vnd.apple.mpegurl")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
