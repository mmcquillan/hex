package models

// Rule Struct
type Rule struct {
	Id             string
	Name           string   `json:"rule"`
	Match          string   `json:"match"`
	Schedule       string   `json:"schedule"`
	URL            string   `json:"url"`
	ACL            string   `json:"acl"`
	Channel        string   `json:"channel"`
	Format         bool     `json:"format"`
	Threaded       bool     `json:"threaded"`
	OutputFailOnly bool     `json:"output_fail_only"`
	OutputOnChange bool     `json:"output_on_change"`
	GroupOutput    bool     `json:"group_output"`
	Help           string   `json:"help"`
	Hide           bool     `json:"hide"`
	Active         bool     `json:"active"`
	Debug          bool     `json:"debug"`
	Actions        []Action `json:"actions"`
}
