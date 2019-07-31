package ctfdapi

import (
    "fmt"
    "net/http"
    "io"
    "encoding/json"
    "bytes"
)

/* TODO: should probably move everything except cookie to a config struct so that people can create multiple clients with slightly different configurations 
  Also, I should probably actually use the URL struct */
type Client struct {
    BaseURL string
    UserAgent string
    HttpClient *http.Client
    cookie *http.Cookie
}


type Request struct {
    Method string
    Path string
    Data interface{}
}

/* TODO: Rename this to JsonRawRequest, since we will have to create form encoded requests as well */
func (clt *Client) RawRequest(req *Request) (*http.Request, error) {
    if req == nil {
        return nil, fmt.Errorf("request is nil")
    }
    url := clt.BaseURL + req.Path

    /*TODO: Once this becomes just JsonRawRequest, we can just encode the Data */
    var buf io.ReadWriter
    if req.Data != nil {
        buf = new(bytes.Buffer)
        if err := json.NewEncoder(buf).Encode(req.Data); err != nil {
            return nil, err
        }
    }
    httpReq, err := http.NewRequest(req.Method, url, buf)
    if err != nil {
        return nil, err
    }

    if req.Data != nil {
        httpReq.Header.Set("Content-Type", "application/json")
    }
    httpReq.Header.Set("User-Agent", clt.UserAgent)
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
