package rushgo

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
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
func ParseXML(data io.Reader) (map[string][]string, error) {
    decoder := xml.NewDecoder(data)
    decoder.CharsetReader = charset.NewReaderLabel // Set the CharsetReader

    result := make(map[string][]string)
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
            currentText += string(tok)
        case xml.EndElement:
            if currentElement != "" {
                result[currentElement] = append(result[currentElement], currentText)
                currentElement = ""
            }
        }
    }

    return result, nil
}





// GetString retrieves a string value from a JSON object by key

// ResponseBodyContains checks if the response body contains a specific string
func ResponseBodyContains(responseBody []byte, searchStr string) bool {
    return bytes.Contains(bytes.ToLower(responseBody), []byte(strings.ToLower(searchStr)))
}
