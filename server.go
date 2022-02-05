package httpdump

import (
	"fmt"
	"net/http"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type ServeOpts struct {
	Port int
	Ip   string
	Cert string
	Key  string
}

func Serve(opts ServeOpts) error {
	handler := NewHandler()
	addr := fmt.Sprintf("%s:%d", opts.Ip, opts.Port)

	if len(opts.Cert) > 0 {
		if len(opts.Key) == 0 {
			return fmt.Errorf("cert is specified, but key is not")
		}

		return serveH2(handler, addr, opts.Cert, opts.Key)
	} else {
		return serveH2c(handler, addr)
	}
}

func serveH2c(h http.Handler, addr string) error {
	h2s := &http2.Server{}
	server := &http.Server{
		Addr:    addr,
		Handler: h2c.NewHandler(h, h2s),
	}
	fmt.Printf("Listening http://%s...\n", addr)
	return server.ListenAndServe()
}

func serveH2(h http.Handler, addr string, cert string, key string) error {
	server := &http.Server{
		Addr:    addr,
		Handler: h,
	}

	fmt.Printf("Listening https://%s...\n", addr)
	return server.ListenAndServeTLS(cert, key)
}
