package rushgo

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"strings"

	"golang.org/x/net/html/charset"
)

// ParseJSON parses a byte slice of JSON data into a map
func ParseJSON(data []byte) (map[string]interface{}, error) {
    var result map[string]interface{}
    err := json.Unmarshal(data, &result)
    if err != nil {
        return nil, err
    }

    return result, nil
}

// ParseCookies extracts and parses cookies from an http.Response and returns them as a map
func ParseCookies(resp *http.Response) map[string]string {
    cookies := make(map[string]string)

    for _, cookie := range resp.Cookies() {
        cookies[cookie.Name] = cookie.Value
    }

    return cookies
}


// largely incomplete XML parser
func ParseXML(data io.Reader) (map[string]string, error) {
    decoder := xml.NewDecoder(data)
    decoder.CharsetReader = charset.NewReaderLabel // Set the CharsetReader

    result := make(map[string]string)
    var currentElement string
    var currentText string

    for {
        tok, err := decoder.Token()
        if err != nil {
            if err == io.EOF {
                break
            }
            return nil, err
        }

        switch tok := tok.(type) {
        case xml.StartElement:
            currentElement = tok.Name.Local
            currentText = ""
        case xml.CharData:
            currentText = strings.TrimSpace(currentText + string(tok))
        case xml.EndElement:
            if currentElement != "" {
                if existing, exists := result[currentElement]; exists {
                    result[currentElement] = existing + currentText
                } else {
                    result[currentElement] = currentText
                }
                currentElement = ""
            }
        }
    }

    return result, nil
}



// ResponseBodyContains checks if the response body contains a specific string
func ResponseBodyContains(responseBody []byte, searchStr string) bool {
    return bytes.Contains(bytes.ToLower(responseBody), []byte(strings.ToLower(searchStr)))
}


func ExtractBetween(body, left, right string) (string, error) {
    // Find the start index of the left delimiter
    start := strings.Index(body, left)
    if start == -1 {
        return "", fmt.Errorf("left string %s not found", left)
    }

    // Find the end index of the right delimiter starting from the end of the left delimiter
    end := strings.Index(body[start:], right)
    if end == -1 {
        return "", fmt.Errorf("right string %s not found", right)
    }
    // Adjust end to the actual index in `body`
    end += start + len(right)

    // Extract the target substring including the left and right delimiters
    return body[start:end], nil
}

