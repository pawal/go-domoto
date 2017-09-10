package domoto

import (
	"fmt"

	resty "gopkg.in/resty.v0"
)

// Call Domoticz API
// input is a map of query parameters for the API
func (c *Config) Call(qp *map[string]string) (resp *resty.Response, err error) {
	r := resty.R().
		SetHeader("Accept", mediaType).
		SetHeader("Content-Type", mediaType).
		SetQueryParams(*qp)
	if c.secret != "" {
		r.SetHeader("Authorization", "Basic "+c.secret)
	}
	resp, err = r.Get(fmt.Sprintf("%s/%s", c.BaseURL, "json.htm"))
	return
}
