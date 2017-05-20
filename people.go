package linkedin

// PeopleProfile gets a people profile
func (cli *Client) PeopleProfile(id string, fields []string) (r map[string]interface{}, e error) {
	var opt map[string]interface{}

	if len(fields) > 0 {
		opt = map[string]interface{}{
			"fields": fields,
		}
	}

	r, e = cli.call("GET", "people", id, "", opt)

	return r, e
}

// PeopleShare shares to profile
func (cli *Client) PeopleShare(data map[string]interface{}) (r map[string]interface{}, e error) {
	r, e = cli.call("POST", "people", "", "/shares", data)
	return r, e
}
