package linkedin

import (
	"fmt"
	"net/url"
)

func (cli *Client) CompanyProfile(id string, fields []string) (r map[string]interface{}, e error) {
	var opt map[string]interface{}

	if len(fields) > 0 {
		opt = map[string]interface{}{
			"fields": fields,
		}
	}

	r, e = cli.Call("GET", "companies", id, "", opt)

	return r, e
}

func (cli *Client) CompanyUpdates(id string, params map[string]string) (r map[string]interface{}, e error) {
	v := url.Values{}
	for key, val := range params {
		v.Add(key, val)
	}

	path := fmt.Sprintf("/updates?%v", v.Encode())
	r, e = cli.Call("GET", "companies", id, path, nil)

	return r, e
}
