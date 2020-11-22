package hue

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/codescot/go-common/httputil"
)

// Hue Philips Hue interface
type Hue struct {
	AppName  string
	Address  string
	Username string
}

// AuthResult auth result
type AuthResult struct {
	Success struct {
		Username string `json:"username"`
	} `json:"success"`
}

func (h Hue) baseURL() string {
	return fmt.Sprintf("http://%s/api", h.Address)
}

func (h Hue) urlWithUsername() string {
	return fmt.Sprintf("http://%s/api/%s", h.Address, h.Username)
}

// Authenticate lets get started.
func (h Hue) Authenticate() {
	app := h.AppName
	name, _ := os.Hostname()

	req := newPOSTRequest(h.baseURL())
	req.Body = fmt.Sprintf(`{ "devicetype": "%s#%s" } `, app, name)

	fmt.Println(req.Body)

	resp, err := req.String()
	if err != nil {
		fmt.Println(err.Error())
	}

	resp = resp[1 : len(resp)-1]
	fmt.Println(resp)

	var a AuthResult
	json.Unmarshal([]byte(resp), &a)

	h.Username = a.Success.Username
}

func newPOSTRequest(targetURL string) httputil.HTTP {
	headers := httputil.Headers{}
	headers.JSON()

	return httputil.HTTP{
		TargetURL: targetURL,
		Method:    "POST",
		Headers:   headers,
	}
}

func newGETRequest(targetURL string) httputil.HTTP {
	return httputil.HTTP{
		TargetURL: targetURL,
		Method:    "GET",
	}
}
