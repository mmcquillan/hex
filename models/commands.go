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
	Schedule string
	Interval int
	Sampling int
	Remind   int
	Green    string
	Yellow   string
	Red      string
	Tags     string
}
