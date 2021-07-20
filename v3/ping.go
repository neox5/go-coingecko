package v3

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Ping() (*PingResponse, error) {
	c := http.Client{}

	resp, err := c.Get(baseUrl + "/ping")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("error reading response body")
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bad request (status=%d, body=%v)", resp.StatusCode, resp)
	}

	pR := &PingResponse{}
	err = json.Unmarshal(bodyBytes, pR)
	if err != nil {
		return nil, errors.New("error unmarshaling reponse to json")
	}

	return pR, nil
}
