package httpdump

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type ServeOpts struct {
	Port int
	Ip   string
}

func Serve(opts ServeOpts) error {
	handler := NewHandler()

	h2s := &http2.Server{}
	addr := fmt.Sprintf("%s:%d", opts.Ip, opts.Port)
	server := &http.Server{
		Addr:    addr,
		Handler: h2c.NewHandler(handler, h2s),
	}

	fmt.Printf("Listening %s...\n", addr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
	return nil
}
