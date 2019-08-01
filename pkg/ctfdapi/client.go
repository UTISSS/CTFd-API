package ctfdapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Client struct {
	Config     ClientConfig
	HttpClient *http.Client
	cookie     *http.Cookie
}

type ClientConfig struct {
	BaseURL   string
	UserAgent string
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

	url := clt.Config.BaseURL + req.Path
	data, err := json.Marshal(req.Data)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequest(req.Method, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set("Content-Type", "application/json")
	if clt.Config.UserAgent != "" {
		httpReq.Header.Set("User-Agent", clt.Config.UserAgent)
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
