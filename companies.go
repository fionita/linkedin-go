package linkedin

import (
	"errors"
	"fmt"
	"net/url"
)

// CompanyProfile gets company profile
func (cli *Client) CompanyProfile(id string, fields []string) (r map[string]interface{}, e error) {
	var opt map[string]interface{}

	if len(fields) > 0 {
		opt = map[string]interface{}{
			"fields": fields,
		}
	}

	r, e = cli.call("GET", "companies", id, "", opt)

	return r, e
}

// CompanyUpdates gets company updates
func (cli *Client) CompanyUpdates(id string, params map[string]string) (r map[string]interface{}, e error) {
	optionalParams := [3]string{"event-type", "count", "start"}
	v := url.Values{}
	for _, key := range optionalParams {
		val, ok := params[key]
		if ok == true {
			v.Add(key, val)
		}
	}

	path := fmt.Sprintf("/updates?%v", v.Encode())
	r, e = cli.call("GET", "companies", id, path, nil)

	return r, e
}

// CompanyUpdate get a specific company update
func (cli *Client) CompanyUpdate(id, key string) (r map[string]interface{}, e error) {
	if key == "" {
		e = errors.New("Update Key must be present")
		return
	}

	path := fmt.Sprintf("/updates/key=%v", key)

	r, e = cli.call("GET", "companies", id, path, nil)

	return r, e
}
