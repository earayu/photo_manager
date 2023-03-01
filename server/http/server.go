package main

import (
	"fmt"
	"github.com/spf13/pflag"
	"net/http"
	"os"
)

var (
	port int
)

func main() {

	fs := pflag.NewFlagSet("http", pflag.ExitOnError)
	fs.IntVar(&port, "port", 8080, "the port to listen on")
	fs.Parse(os.Args[1:])

	//open a http server and listen on port 8080 and handle '/hello'
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})
	//start the server
	fmt.Println("listen on port: ", port)
	http.ListenAndServe(":"+string(port), nil)

	select {}
}
