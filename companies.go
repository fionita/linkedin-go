package linkedin

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
