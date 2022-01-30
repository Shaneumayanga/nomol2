package nomol

import (
	"io"
	"log"
	"net"
	"net/http"
	"sync"

	"github.com/Shaneumayanga/nomol/sites"
)

func NoMol(w http.ResponseWriter, req *http.Request) {
	blockedSites := sites.GetSites()

	if req.Method != http.MethodConnect {
		log.Println(req.Method, "Not allowed")
		http.NotFound(w, req)
		return
	}
	log.Println(req.RequestURI)

	for _, site := range blockedSites {
		if req.RequestURI == site {
			http.Error(w, "Site blocked by nomol", http.StatusUnauthorized)
			log.Println("Site blocked")
			return
		}
	}

	dst, err := net.Dial("tcp", req.RequestURI)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer dst.Close()
	w.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))

	conn, bio, err := w.(http.Hijacker).Hijack()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		if n := bio.Reader.Buffered(); n > 0 {
			n64, err := io.CopyN(dst, bio, int64(n))
			if n64 != int64(n) || err != nil {
				log.Println("io.CopyN:", n64, err)
				return
			}
		}
		io.Copy(dst, conn)
	}()

	go func() {
		defer wg.Done()
		io.Copy(conn, dst)
	}()

	wg.Wait()
}
