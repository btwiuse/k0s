package header

type SlaveHeader struct {
	Status string `json:"status"`
	Id     string `json:"id"`
	Err    string `json:"error,omitempty"`
}

type AgentHeader struct {
	Append bool   `json:"append"`
	Id     string `json"id"`
	Nonce  string `json"nonce"`

	Pwd      string `json:"pwd"`
	Whoami   string `json:"whoami"`
	Hostname string `json:"hostname"`
}
