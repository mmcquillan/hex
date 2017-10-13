package models

// Rule Struct
type Rule struct {
	Name     string   `json:"rule"`
	Match    string   `json:"match"`
	Schedule string   `json:"schedule"`
	ACL      string   `json:"acl"`
	Format   bool     `json:"format"`
	Help     string   `json:"help"`
	Hide     bool     `json:"hide"`
	Active   bool     `json:"active"`
	Debug    bool     `json:"debug"`
	Actions  []Action `json:"actions"`
}
