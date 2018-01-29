package models

// Rule Struct
type Rule struct {
	Id             string
	Name           string   `json:"rule" yaml:"rule"`
	Match          string   `json:"match" yaml:"match"`
	Schedule       string   `json:"schedule" yaml:"schedule"`
	URL            string   `json:"url" yaml:"url"`
	ACL            string   `json:"acl" yaml:"acl"`
	Channel        string   `json:"channel" yaml:"channel"`
	Format         bool     `json:"format" yaml:"format"`
	Threaded       bool     `json:"threaded" yaml:"threaded"`
	OutputFailOnly bool     `json:"output_fail_only" yaml:"output_fail_only"`
	OutputOnChange bool     `json:"output_on_change" yaml:"output_on_change"`
	GroupOutput    bool     `json:"group_output" yaml:"group_output"`
	Help           string   `json:"help" yaml:"help"`
	Hide           bool     `json:"hide" yaml:"hide"`
	Active         bool     `json:"active" yaml:"active"`
	Debug          bool     `json:"debug" yaml:"debug"`
	Actions        []Action `json:"actions" yaml:"actions"`
}
