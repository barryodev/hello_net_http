package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"log"
	"net/http"
)

func root(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Ping!")
	fmt.Println("Received Request")

	pixel := Pixel{}

	data, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Println(err)
	}

	if err := proto.Unmarshal(data, &pixel); err != nil {
		fmt.Println(err)
	}

	println("{", pixel.GetX(), ",", pixel.GetY(), ":", pixel.GetColour(), "}")
}

func listenForRequests() {
	http.HandleFunc("/", root)
	log.Fatal(http.ListenAndServe(":1337", nil))
}

func sendRequest(pixel Pixel) {
	data, err := proto.Marshal(&pixel)
	if err != nil {
		fmt.Println(err)
		return
	}

	resp, err := http.Post("http://localhost:1337", "", bytes.NewBuffer(data))

	if err != nil {
		fmt.Println(err)
		return
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

	xFlag := flag.Int("x", 0, "X coordinate")

	yFlag := flag.Int("y", 0, "Y coordinate")

	colourFlag := flag.String("colour", "FFFFFF", "Colour hex code")


	flag.Parse()

	if *listenFlag {
		fmt.Println("Listening...")
		listenForRequests()
	}

	if *sendFlag {
		fmt.Println("Sending...")

		pixel := Pixel{X: int32(*xFlag), Y: int32(*yFlag), Colour: *colourFlag}
		sendRequest(pixel)
	}
}
