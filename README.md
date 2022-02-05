# http-dump

A HTTP Dev Server to dump request

Support h2, h2c

## Installation

```
go install github.com/maruware/http-dump/cmd/http-dump
```

or Download from [Releases](releases)

## Usage

### Quick

```sh
$ http-dump
Listening http://0.0.0.0:8080...
2022/02/05 15:44:26 GET /sample HTTP/1.1 # curl http://localhost:8080/sample
2022/02/05 15:44:51 GET /sample HTTP/2.0 # curl --http2-prior-knowledge http://localhost:8080/sample
```

### Options

```sh
$ http-dump -h
NAME:
   http-dump - http dump dev server

USAGE:
   http-dump [global options] command [command options] [arguments...]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --port value, -p value    listen port (default: 8080)
   --bind value, -b value    listen ip (default: "0.0.0.0")
   --cert value              TLS cert file path
   --key value               TLS key file path
   --output value, -o value  output format (simple or json or simple_color) (default: "simple_color")
   --help, -h                show help (default: false)
```

### h2 (https)

Prepare cert

```sh
$ go run $(go env GOROOT)/src/crypto/tls/generate_cert.go -rsa-bits 2048 -host localhost
2022/02/05 15:50:18 wrote cert.pem
2022/02/05 15:50:18 wrote key.pem
```

```sh
$ http-dump --cert ./cert.pem --key ./key.pem
Listening https://0.0.0.0:8080...
2022/02/05 15:53:56 GET /sample HTTP/2.0 # curl --http2 --insecure https://localhost:8080/sample
```
