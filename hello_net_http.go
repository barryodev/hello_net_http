package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func root(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Ping!")
	fmt.Println("Received Request")
}

func listenForRequests() {
	http.HandleFunc("/", root)
	log.Fatal(http.ListenAndServe(":1337", nil))
}

func sendRequest() {
	resp, err := http.Get("http://localhost:1337")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(resp.Status)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	sb := string(body)
	fmt.Println(sb)
}

func main() {
	listenFlag := flag.Bool("listen", false, "listen on port 1337 for requests")

	sendFlag := flag.Bool("send", false, "send a request to http://localhost:1337")

	flag.Parse()

	if *listenFlag {
		fmt.Println("Listening...")
		listenForRequests()
	}

	if *sendFlag {
		fmt.Println("Sending...")
		sendRequest()
	}
}
