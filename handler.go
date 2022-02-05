package httpdump

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/fatih/color"
)

type Handler struct {
	format string
}

func NewHandler(format string) *Handler {
	return &Handler{format: format}
}

type RequestLog struct {
	Method   string      `json:"method"`
	Path     string      `json:"path"`
	IsTLS    bool        `json:"is_tls"`
	Body     string      `json:"body"`
	Header   http.Header `json:"header"`
	Protocol string      `json:"protocol"`
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
		Method:   r.Method,
		Path:     r.URL.Path,
		IsTLS:    r.TLS != nil,
		Body:     string(body),
		Header:   r.Header,
		Protocol: r.Proto,
	}

	switch h.format {
	case "json":
		s, err := l.formatJSON()
		if err != nil {
			return err
		}
		logger.Printf(s)
		break
	case "simple":
		s := l.formatSimple()
		logger.Printf(s)
		break
	default:
		s := l.formatSimpleColor()
		logger.Printf(s)
		break
	}

	fmt.Fprintf(w, "OK")
	return nil
}

func (rl RequestLog) formatSimple() string {
	return fmt.Sprintf("%s %s %s %s", rl.Method, rl.Path, rl.Protocol, rl.Body)
}

func (rl RequestLog) formatSimpleColor() string {
	return fmt.Sprintf("%s %s %s %s", color.GreenString(rl.Method), rl.Path, rl.Protocol, rl.Body)
}

func (rl RequestLog) formatJSON() (string, error) {
	logBuf := bytes.NewBuffer(nil)
	enc := json.NewEncoder(logBuf)

	if err := enc.Encode(rl); err != nil {
		return "", err
	}
	return logBuf.String(), nil
}
