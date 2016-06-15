package models

type Command struct {
	Name     string
	Match    string
	Output   string
	Outputs  []string
	Cmd      string
	Args     string
	Help     string
	HideHelp bool
	RunCheck bool
	Interval int
	Green    string
	Yellow   string
	Red      string
}
