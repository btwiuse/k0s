package header

type ClientHeader struct {
	Status string `json:"status"`
	Id     string `json:"id"`
	Err    string `json:"error,omitempty"`
}

type Header struct {
	Append bool   `json:"append"`
	Id     string `json"id"`

	BuildCode  string `json:"build_code"`
	DockerRepo string `json:"docker_repo"`
	GitBranch  string `json:"git_branch"`
	GitTag     string `json:"git_tag"`
	GitSha1    string `json:"git_sha1"`
	GitMsg     string `json:"git_msg"`
	IP         string `json:"ip"`
	Pwd        string `json:"pwd"`
	Whoami     string `json:"whoami"`
	Hostname   string `json:"hostname"`
	CreatedAt  string `json:"created_at"`

	GOARCH    string `json:"goarch"`
	GOOS      string `json:"goos"`
	GOROOT    string `json:"goroot"`
	GC        string `json:"gc"`
	GOVERSION string `json:"goversion"`
	NCPU      int    `json:"ncpu"`

	Commit string `json:"commit"`
	Built  string `json:"built"`
	Branch string `json:"branch"`
}
