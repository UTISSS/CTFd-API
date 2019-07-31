package ctfdapi


type ChallengeRequest struct {
    Name string `json:"name"`
    Category string `json:"category"`
    State string `json:"state"`
    Value string `json:"value"`
    Type string `json:"type"`
    Description string `json:"description"`
}

func (clt *Client) CreateChallenge(challenge *ChallengeRequest) (error) {
    req := Request{
            Method: "POST",
            Path: "api/v1/challenges",
            Data: challenge,}
    httpReq, err := clt.RawRequest(&req)
    if err != nil {
        return err
    }

    /*TODO: extract relevant info in resp */
    _, err = clt.GetResponse(httpReq, nil)
    return err
}
