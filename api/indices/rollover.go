// generated by github.com/elastic/go-elasticsearch/cmd/generator; DO NOT EDIT

package indices

import (
	"fmt"
	"net/http"
	"net/url"
)

// Rollover - the rollover index API rolls an alias over to a new index when the existing index is considered to be too large or too old. See https://www.elastic.co/guide/en/elasticsearch/reference/5.x/indices-rollover-index.html for more info.
//
// alias: the name of the alias to rollover.
//
// options: optional parameters. Supports the following functional options: WithNewIndex, WithDryRun, WithMasterTimeout, WithTimeout, WithWaitForActiveShards, WithBody, WithErrorTrace, WithFilterPath, WithHuman, WithPretty, WithSourceParam, see the Option type in this package for more info.
func (i *Indices) Rollover(alias string, options ...*Option) (*RolloverResponse, error) {
	req := &http.Request{
		URL: &url.URL{
			Scheme: i.transport.URL.Scheme,
			Host:   i.transport.URL.Host,
		},
		Method: "POST",
	}
	methodOptions := supportedOptions["Rollover"]
	for _, option := range options {
		if _, ok := methodOptions[option.name]; !ok {
			return nil, fmt.Errorf("unsupported option: %s", option.name)
		}
		option.apply(req)
	}
	resp, err := i.transport.Do(req)
	return &RolloverResponse{resp}, err
}

// RolloverResponse is the response for Rollover
type RolloverResponse struct {
	Response *http.Response
	// TODO: fill in structured response
}
