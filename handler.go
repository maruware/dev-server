package httpdump

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

type RequestLog struct {
	Method string `json:"method"`
	Path   string `json:"path"`
	IsTLS  bool   `json:"is_tls"`
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	l := RequestLog{
		Method: r.Method,
		Path:   r.URL.Path,
		IsTLS:  r.TLS == nil,
	}

	b := bytes.NewBuffer(nil)
	enc := json.NewEncoder(b)

	if err := enc.Encode(l); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logger.Errorf(err.Error())
		fmt.Fprintf(w, "err.Error()")
		return
	}

	logger.Printf(b.String())

	fmt.Fprintf(w, "OK")
}
