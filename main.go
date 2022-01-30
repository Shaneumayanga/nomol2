package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"sync"

	"github.com/Shaneumayanga/nomol/client"
	"github.com/Shaneumayanga/nomol/nomol"
	"github.com/Shaneumayanga/nomol/sites"
)

func init() {
	sites.Prepare()
}

func open(url string) error {
	var cmd string
	var args []string
	cmd = "cmd"
	args = []string{"/c", "start"}
	args = append(args, url)
	return exec.Command(cmd, args...).Start()
}

func main() {
	wg := &sync.WaitGroup{}
	fmt.Println("Developer :Shane umayanga")
	fmt.Println("me@shaneumayanga.com")

	wg.Add(2)
	go func() {
		defer wg.Done()
		handler := http.HandlerFunc(nomol.NoMol)
		log.Println("NoMol running on port :8080")
		log.Fatal(http.ListenAndServe(":8080", handler))
	}()

	go func() {
		defer wg.Done()
		svc := &http.Server{
			Addr:    ":3000",
			Handler: client.NewClient(),
		}
		log.Println("Client running on port :3000")
		log.Fatal(svc.ListenAndServe())
	}()

	open("http://localhost:3000")

	wg.Wait()
}
