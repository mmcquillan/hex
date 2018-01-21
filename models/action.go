package models

// Action Struct
type Action struct {
	Type           string            `json:"type"`
	Command        string            `json:"command"`
	HideOutput     bool              `json:"hide_output"`
	OutputToVar    bool              `json:"output_to_var"`
	OutputFailOnly bool              `json:"output_fail_only"`
	RunOnFail      bool              `json:"run_on_fail"`
	LastConfig     bool              `json:"last_config"`
	Config         map[string]string `json:"config"`
}
