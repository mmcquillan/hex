package models

// Action Struct
type Action struct {
	Type           string            `json:"type" yaml:"type"`
	Command        string            `json:"command" yaml:"command"`
	HideOutput     bool              `json:"hide_output" yaml:"hide_output"`
	OutputToVar    bool              `json:"output_to_var" yaml:"output_to_var"`
	OutputFailOnly bool              `json:"output_fail_only" yaml:"output_fail_only"`
	RunOnFail      bool              `json:"run_on_fail" yaml:"run_on_fail"`
	LastConfig     bool              `json:"last_config" yaml:"last_config"`
	Config         map[string]string `json:"config" yaml:"config"`
}
