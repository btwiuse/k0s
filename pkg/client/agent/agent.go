package agent

type Agent struct {
	ID        string `json:"id"`
	Connected int64  `json:"connected"`
	Hostname  string `json:"hostname"`
	Whoami    string `json:"whoami"`
	User      string `json:"user"`     // compat
	Username  string `json:"username"` // compat
	Pwd       string `json:"pwd"`
	OS        string `json:"os"`
	ARCH      string `json:"arch"`
	IP        string `json:"ip"`
}
