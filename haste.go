package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

/*
Note:This isn't written that well but hey it works!
Copy pastes from...
post client
https://github.com/tenntenn/golang-samples/blob/master/http/post/client.go
json parsing
https://gobyexample.com/json
for bad cert hastebin has
https://stackoverflow.com/questions/12122159/golang-how-to-do-a-https-request-with-bad-certificate
And stdin
https://www.socketloop.com/tutorials/golang-check-if-os-stdin-input-data-is-piped-or-from-terminal
*/
type Haste struct {
	Key string `json:"key"`
}

func post(body string) {
	//Dealing with hastebin's bad ssl cert.
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	r, err := client.Post("https://hastebin.com/documents", "text", bytes.NewBuffer([]byte(body)))
	perror(err)
	response, err := ioutil.ReadAll(r.Body)
	perror(err)

	//parse the json
	var link map[string]interface{}
	json.Unmarshal(response, &link)
	//print as
	fmt.Printf("https://hastebin.com/%s\n", link["key"])
}

func main() {

	fi, _ := os.Stdin.Stat()
	if (fi.Mode() & os.ModeCharDevice) == 0 {

		// do things for data from pipe

		bytes, _ := ioutil.ReadAll(os.Stdin)
		str := string(bytes)
		if len(str) < 1 {
			fmt.Println("Not enough arguments")
			return
		}
		if len(str) > 0 {
			//Post the stdin
			post(string(str))
			return
		}

	}
	if len(os.Args) > 1 {
		post(os.Args[1])
		return
	}
	if len(os.Args) < 2 {
		fmt.Println("Not enough arguments")
		return
	}

}

//Error handling
func perror(err error) {
	if err != nil {
		panic(err)
	}
}