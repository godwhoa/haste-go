package main

import (
	"bytes"
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
for bad cert hastebin has update: now using http instead of https herp derp
https://stackoverflow.com/questions/12122159/golang-how-to-do-a-https-request-with-bad-certificate
And stdin
https://www.socketloop.com/tutorials/golang-check-if-os-stdin-input-data-is-piped-or-from-terminal
*/
type Haste struct {
	Key string `json:"key"`
}

func post(body string) {

	r, err := http.Post("http://hastebin.com/documents", "text", bytes.NewBuffer([]byte(body)))
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
		if Exist(os.Args[1]) {
			body, err := ioutil.ReadFile(os.Args[1])
			perror(err)
			post(string(body))
			return
		}
		fmt.Println("File doesn't exist.")
		return
	}
	if len(os.Args) < 2 {
		fmt.Println("Not enough arguments")
		return
	}

}

//File exists
func Exist(file string) bool {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return false
	}

	return true
}

//Error handling
func perror(err error) {
	if err != nil {
		panic(err)
	}
}
