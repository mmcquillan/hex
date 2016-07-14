package models

type Command struct {
	Name     string
	Match    string
	Output   string
	Outputs  []string
	Cmd      string
	Args     string
	Help     string
	Process  bool
	HideHelp bool
	RunCheck bool
	Interval int
	Remind   int
	Green    string
	Yellow   string
	Red      string
}
