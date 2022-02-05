package httpdump

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
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
	Body   string `json:"body"`
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h.handle(w, r); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logger.Errorf(err.Error())
		fmt.Fprintf(w, err.Error())
	}
}

func (h *Handler) handle(w http.ResponseWriter, r *http.Request) error {
	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	l := RequestLog{
		Method: r.Method,
		Path:   r.URL.Path,
		IsTLS:  r.TLS != nil,
		Body:   string(body),
	}

	logBuf := bytes.NewBuffer(nil)
	enc := json.NewEncoder(logBuf)

	if err := enc.Encode(l); err != nil {
		return err
	}

	logger.Printf(logBuf.String())

	fmt.Fprintf(w, "OK")
	return nil
}
