package ctfdapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

/* TODO: should probably move everything except cookie to a config struct so that people can create multiple clients with slightly different configurations */
type Client struct {
	BaseURL    string
	UserAgent  string
	HttpClient *http.Client
	cookie     *http.Cookie
}

type Request struct {
	Method string
	Path   string
	Data   interface{}
}

func (clt *Client) JsonRawRequest(req *Request) (*http.Request, error) {
	if req == nil {
		return nil, fmt.Errorf("request is nil")
	}

	url := clt.BaseURL + req.Path
	data, err := Json.Marshal(req.Data)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequest(req.Method, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set("Content-Type", "application/json")
	if clt.UserAgent != "" {
		httpReq.Header.Set("User-Agent", clt.UserAgent)
	}
	if clt.cookie != nil {
		httpReq.AddCookie(clt.cookie)
	}

	return httpReq, nil
}

func (clt *Client) GetResponse(req *http.Request, respBody interface{}) (*http.Response, error) {
	resp, err := clt.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(respBody)
	return resp, err
}
