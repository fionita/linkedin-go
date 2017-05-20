package linkedin

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const apiURL = "https://api.linkedin.com/v1"

// Config struct
type Config struct {
	AccessToken string
}

// Client struct
type Client struct {
	conf *Config
}

// Init - initialize the client
func Init(conf *Config) (*Client, error) {
	if conf.AccessToken == "" {
		return nil, fmt.Errorf("%v", "Access token is required")
	}

	return &Client{conf}, nil
}

func (li *Client) call(verb string, endpoint string, id string, path string, options map[string]interface{}) (r map[string]interface{}, e error) {
	if endpoint == "people" && id == "" {
		id = "~"
	} else if id == "" {
		e = errors.New("Id must be present")
	}
	var fields string
	if v, ok := options["fields"]; ok {
		fields = ":(" + strings.Join(v.([]string), ",") + ")"
		delete(options, "fields")
	}

	url := fmt.Sprintf(apiURL + "/" + endpoint + "/" + id + fields + path)

	body, err := json.Marshal(options)
	if err != nil {
		return nil, err
	}
	bodyStr := []byte(body)

	req, err := http.NewRequest(verb, url, bytes.NewBuffer(bodyStr))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+li.conf.AccessToken)
	req.Header.Add("x-li-format", "json")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	data, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(data, &r)

	if err != nil {
		return nil, err
	}

	if _, error := r["errorCode"]; error {
		err = errors.New(string(data))
		return nil, err
	}

	return r, e
}
