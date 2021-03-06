// generated by github.com/elastic/go-elasticsearch/cmd/generator; DO NOT EDIT

package api

import (
	"fmt"
	"net/http"
	"net/url"
)

// FieldStats - this functionality is experimental and may be changed or removed completely in a future release. See https://www.elastic.co/guide/en/elasticsearch/reference/5.x/search-field-stats.html for more info.
//
// options: optional parameters. Supports the following functional options: WithIndex, WithAllowNoIndices, WithExpandWildcards, WithFields, WithIgnoreUnavailable, WithLevel, WithBody, WithErrorTrace, WithFilterPath, WithHuman, WithPretty, WithSourceParam, see the Option type in this package for more info.
func (a *API) FieldStats(options ...*Option) (*FieldStatsResponse, error) {
	req := &http.Request{
		URL: &url.URL{
			Scheme: a.transport.URL.Scheme,
			Host:   a.transport.URL.Host,
		},
		Method: "GET",
	}
	methodOptions := supportedOptions["FieldStats"]
	for _, option := range options {
		if _, ok := methodOptions[option.name]; !ok {
			return nil, fmt.Errorf("unsupported option: %s", option.name)
		}
		option.apply(req)
	}
	resp, err := a.transport.Do(req)
	return &FieldStatsResponse{resp}, err
}

// FieldStatsResponse is the response for FieldStats
type FieldStatsResponse struct {
	Response *http.Response
	// TODO: fill in structured response
}
