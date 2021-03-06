// generated by github.com/elastic/go-elasticsearch/cmd/generator; DO NOT EDIT

package tasks

import (
	"net/http"
)

// Option is a non-required API option that gets applied to an HTTP request.
type Option struct {
	name  string
	apply func(r *http.Request)
}

// WithActions - a comma-separated list of actions that should be cancelled. Leave empty to cancel all.
func WithActions(actions []string) *Option {
	return &Option{
		name: "WithActions",
		apply: func(r *http.Request) {
		},
	}
}

// WithErrorTrace - include the stack trace of returned errors.
func WithErrorTrace(errorTrace bool) *Option {
	return &Option{
		name: "WithErrorTrace",
		apply: func(r *http.Request) {
		},
	}
}

// WithFilterPath - a comma-separated list of filters used to reduce the respone.
func WithFilterPath(filterPath []string) *Option {
	return &Option{
		name: "WithFilterPath",
		apply: func(r *http.Request) {
		},
	}
}

// WithHuman - return human readable values for statistics.
func WithHuman(human bool) *Option {
	return &Option{
		name: "WithHuman",
		apply: func(r *http.Request) {
		},
	}
}

// WithNodeID - a comma-separated list of node IDs or names to limit the returned information; use "_local" to return information from the node you're connecting to, leave empty to get information from all nodes.
func WithNodeID(nodeID []string) *Option {
	return &Option{
		name: "WithNodeID",
		apply: func(r *http.Request) {
		},
	}
}

// WithParentNode - cancel tasks with specified parent node.
func WithParentNode(parentNode string) *Option {
	return &Option{
		name: "WithParentNode",
		apply: func(r *http.Request) {
		},
	}
}

// WithParentTask - cancel tasks with specified parent task id (node_id:task_number). Set to -1 to cancel all.
func WithParentTask(parentTask string) *Option {
	return &Option{
		name: "WithParentTask",
		apply: func(r *http.Request) {
		},
	}
}

// WithPretty - pretty format the returned JSON response.
func WithPretty(pretty bool) *Option {
	return &Option{
		name: "WithPretty",
		apply: func(r *http.Request) {
		},
	}
}

// WithSourceParam - the URL-encoded request definition. Useful for libraries that do not accept a request body for non-POST requests.
func WithSourceParam(sourceParam string) *Option {
	return &Option{
		name: "WithSourceParam",
		apply: func(r *http.Request) {
		},
	}
}

// WithTaskID - cancel the task with specified task id (node_id:task_number).
func WithTaskID(taskID string) *Option {
	return &Option{
		name: "WithTaskID",
		apply: func(r *http.Request) {
		},
	}
}

var (
	supportedOptions = map[string]map[string]struct{}{
		"Cancel": map[string]struct{}{
			"WithTaskID":      struct{}{},
			"WithActions":     struct{}{},
			"WithNodeID":      struct{}{},
			"WithParentNode":  struct{}{},
			"WithParentTask":  struct{}{},
			"WithErrorTrace":  struct{}{},
			"WithFilterPath":  struct{}{},
			"WithHuman":       struct{}{},
			"WithPretty":      struct{}{},
			"WithSourceParam": struct{}{},
		},
		"Get": map[string]struct{}{
			"WithTaskID":            struct{}{},
			"WithWaitForCompletion": struct{}{},
			"WithErrorTrace":        struct{}{},
			"WithFilterPath":        struct{}{},
			"WithHuman":             struct{}{},
			"WithPretty":            struct{}{},
			"WithSourceParam":       struct{}{},
		},
		"List": map[string]struct{}{
			"WithActions":           struct{}{},
			"WithDetailed":          struct{}{},
			"WithGroupBy":           struct{}{},
			"WithNodeID":            struct{}{},
			"WithParentNode":        struct{}{},
			"WithParentTask":        struct{}{},
			"WithWaitForCompletion": struct{}{},
			"WithErrorTrace":        struct{}{},
			"WithFilterPath":        struct{}{},
			"WithHuman":             struct{}{},
			"WithPretty":            struct{}{},
			"WithSourceParam":       struct{}{},
		},
	}
)
