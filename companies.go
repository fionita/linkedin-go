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
func (cli *Client) CompanyUpdate(id, key, filter string) (r map[string]interface{}, e error) {
	if key == "" {
		e = errors.New("Update Key must be present")
		return
	}

	path := fmt.Sprintf("/updates/key=%v", key)
	filters := [2]string{"update-comments", "likes"}

	for _, f := range filters {
		if f == filter {
			path = fmt.Sprintf("%v/%v", path, f)
		}
	}

	r, e = cli.call("GET", "companies", id, path, nil)

	return r, e
}

// CompanyShare creates a company share
func (cli *Client) CompanyShare(id string, data map[string]interface{}) (r map[string]interface{}, e error) {
	r, e = cli.call("POST", "companies", id, "/shares", data)
	return r, e
}

// CompanyAddComment adds a comment on behalf of a company
func (cli *Client) CompanyAddComment(id, key, comment string) (r map[string]interface{}, e error) {
	if key == "" {
		e = errors.New("Update Key must be present")
		return
	}

	path := fmt.Sprintf("/updates/key=%v/update-comments-as-company", key)
	data := map[string]interface{}{
		"comment": comment,
	}

	r, e = cli.call("POST", "companies", id, path, data)
	return r, e
}
