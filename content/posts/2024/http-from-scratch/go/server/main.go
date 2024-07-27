package main

import (
	"bufio"
	"log"
	"net"
	"net/http"
	"net/url"
	"strings"
)

type Server struct {
	Addr    string
	Handler http.Handler
}

func (s *Server) ServeAndListen() error {
	if s.Handler == nil {
		panic("http server started without a handler")
	}
	l, err := net.Listen("tcp", s.Addr)

	if err != nil {
		return err
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go func(c net.Conn) {
			reader := bufio.NewReader(c)
			line, _, err := reader.ReadLine()
			if err != nil {
				return
			}

			fields := strings.Fields(string(line))
			if len(fields) < 2 {
				return
			}
			r := &http.Request{
				Method:     fields[0],
				URL:        &url.URL{Scheme: "http", Path: fields[1]},
				Proto:      "HTTP/0.9",
				ProtoMajor: 0,
				ProtoMinor: 9,
				RemoteAddr: c.RemoteAddr().String(),
			}

			s.Handler.ServeHTTP(newWriter(c), r)
			c.Close()
		}(conn)
	}
}

type responseBodyWriter struct {
	conn net.Conn
}

// Header implements http.ResponseWriter.
func (r *responseBodyWriter) Header() http.Header {
	panic("unsupported with HTTP/0.9")
}

// Write implements http.ResponseWriter.
// Subtle: this method shadows the method (Conn).Write of responseBodyWriter.Conn.
func (r *responseBodyWriter) Write(b []byte) (int, error) {
	return r.conn.Write(b)
}

// WriteHeader implements http.ResponseWriter.
func (r *responseBodyWriter) WriteHeader(statusCode int) {
	panic("unsupported with HTTP/0.9")
}

func newWriter(c net.Conn) http.ResponseWriter {
	return &responseBodyWriter{
		conn: c,
	}
}

func main() {
	addr := "127.0.0.1:9000"
	s := Server{
		Addr: addr,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello World!"))
		}),
	}
	log.Printf("Listing on %s", addr)
	if err := s.ServeAndListen(); err != nil {
		log.Fatal(err)
	}
}
