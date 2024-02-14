package examples

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"

	rushgo "github.com/shelovesmox/rushgo/rushgo"
)

// SimpleGETRequest demonstrates a simple GET request using RushGo
func SimpleGETRequest() {
	client := rushgo.New(nil) // Using default configuration
	resp, err := client.Get("https://httpbin.org/get")
	if err != nil {
		fmt.Println("Error in GET request:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	fmt.Println("GET Response Body:", string(body))
}

// POSTRequestWithJSON demonstrates a POST request with a JSON body using RushGo
func POSTRequestWithJSON() {
	client := rushgo.New(nil) // Using default configuration

	jsonData := map[string]string{"key": "value"}
	jsonBody, _ := json.Marshal(jsonData)

	resp, err := client.Post("https://httpbin.org/post", jsonBody)
	if err != nil {
		fmt.Println("Error in POST request:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	fmt.Println("POST Response Body:", string(body))
}

// DownloadImageExample demonstrates downloading an image using RushGo
func DownloadImageExample() {
	client := rushgo.New(nil) // Using default configuration
	url := "https://httpbin.org/image/jpeg"
	savePath := "downloaded_image.jpg" // Path to save the image

	resp, err := client.DownloadImage(url, &savePath)
	if err != nil {
		fmt.Println("Error downloading image:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Image downloaded successfully. Status:", resp.Status)
}

// CustomHeadersAndCookiesExample demonstrates setting custom headers and cookies
func CustomHeadersAndCookiesExample() {
	client := rushgo.New(nil).
		WithHeaders(map[string]string{
			"Custom-Header": "HeaderValue",
		}).
		WithCookies(map[string]string{
			"session_token": "123456",
		})

	resp, err := client.Get("https://httpbin.org/headers")
	if err != nil {
		fmt.Println("Error in GET request with custom headers and cookies:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	fmt.Println("Response with custom headers and cookies:", string(body))
}

// You can call these functions from another part of your application.

func CustomUserAgentExample() {
    client := rushgo.New(nil).WithUserAgent("MyCustomUserAgent/1.0")

    resp, err := client.Get("https://httpbin.org/user-agent")
    if err != nil {
        fmt.Println("Error in GET request with custom User-Agent:", err)
        return
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("Error reading response body:", err)
        return
    }

    fmt.Println("Response with custom User-Agent:", string(body))
}

func CustomConfigExample() {
    cfg := rushgo.Config{
		EnableHTTP3: true, //change to false to use http.1.1
    }
    client := rushgo.New(&cfg)

    resp, err := client.Get("https://example.com")
    if err != nil {
        fmt.Println("Error in GET request with custom configuration:", err)
        return
    }
    defer resp.Body.Close()



    fmt.Println("Response with custom configuration:", string(resp.Proto))
}


func XmlParser () {
	cfg := rushgo.Config{
        EnableHTTP2: true, //change to false to use http.1.1
    }
    client := rushgo.New(&cfg)

    resp, err := client.Get("https://httpbin.org/xml")
	if err != nil {
		// Handle the error here.
	}

	defer resp.Body.Close()

	body := `
	<note>
	  <to>Tove</to>
	  <from>Jani</from>
	  <heading>Reminder</heading>
	  <body>Don't forget me this weekend!</body>
	</note>
	`

	xml, err := rushgo.ParseXML(strings.NewReader(body))
	if err != nil {
		fmt.Println("Error parsing XML:", err)
		return
	}


	fmt.Println(xml["from"])
}


func ExtractBetween () {

	cfg := rushgo.Config{
		EnableHTTP2: true, //change to false to use http.1.1
	}

	client := rushgo.New(&cfg)

	resp, err := client.Get("https://example.com")

	if err != nil {
		fmt.Println("Error in GET request:", err)
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	extractedTitle, err := rushgo.ExtractBetween(string(body), "This", "permission")

	if err != nil {
		fmt.Println("Error extracting title:", err)
		return
	}

	fmt.Println(extractedTitle)
}

func WebSocketExample () {

	rgClient := rushgo.New(nil)

    // Connect to the WebSocket server
    wsURL := "wss://socketsbay.com/wss/v2/1/demo/"
    conn, resp, err := rgClient.WebSocketConnect(wsURL)
    if err != nil {
        log.Fatalf("Error connecting to WebSocket: %v\n", err)
    }
    defer conn.Close()

    fmt.Printf("Connected to WebSocket server. HTTP status code: %d\n", resp.StatusCode)
}

func HTTP3() {
	cfg := rushgo.Config{
		EnableHTTP3: true, //change to false to use http.1.1
	}

	client := rushgo.New(&cfg)

	resp, err := client.Get("https://example.com")

	if err != nil {
		fmt.Println("Error in GET request:", err)
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	fmt.Println(string(body))

	if resp.ProtoMajor == 3 {
		fmt.Println("The request is using HTTP/3")
	} else {
		fmt.Println("The request is not using HTTP/3")
	}
}

func DeleteExample() {
	client := rushgo.New(nil)

	resp, err := client.Delete("https://httpbin.org/delete")
	if err != nil {
		fmt.Println("Error in DELETE request:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	fmt.Println("DELETE Response Body:", string(body))
}

func WithBasicAuthExample() {
	client := rushgo.New(nil).WithBasicAuth("username", "password")

	resp, err := client.Get("https://httpbin.org/basic-auth/username/password")
	if err != nil {
		fmt.Println("Error in GET request with Basic Auth:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	fmt.Println("Response with Basic Auth:", string(body))
}


func WithBearerTokenExample() {
	client := rushgo.New(nil).WithBearerToken("my-jwt-token")

	resp, err := client.Get("https://httpbin.org/bearer")
	if err != nil {
		fmt.Println("Error in GET request with Bearer Token:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	fmt.Println("Response with Bearer Token:", string(body))
}