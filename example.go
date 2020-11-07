package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/JohnBat26/go-httpclient/gohttp"
)

var githubHttpClient = getGithubHttpClient()

func getGithubHttpClient() gohttp.HttpClient {
	client := gohttp.New()

	client.DisableTimeouts(false)
	client.SetMaxIdleConnections(20)
	client.SetConnectTimeout(2 * time.Second)
	client.SetResponseTimeout(4 * time.Second)

	commonHeaders := make(http.Header)
	commonHeaders.Set("Authorization", "Bearer foo")
	client.SetHeaders(commonHeaders)

	return client
}

func main() {
	getUrls()
	getUrls()
	getUrls()
}

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func getUrls() {
	response, err := githubHttpClient.Get("https://api.github.com", nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(response.StatusCode)

	bytes, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(bytes))
}

func createUser(user User) {
	response, err := githubHttpClient.Post("https://api.github.com", nil, user)
	if err != nil {
		panic(err)
	}

	fmt.Println(response.StatusCode)

	bytes, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(bytes))
}
