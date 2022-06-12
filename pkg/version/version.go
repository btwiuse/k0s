package version

import (
	"encoding/json"
	"runtime"

	"github.com/btwiuse/pretty"
	"k0s.io/pkg"
)

var (
	GitCommitString  string
	GitStateString   string
	GitBranchString  string
	GitSummaryString string
	BuildDateString  string
	VersionString    string
)

func GetVersion() pkg.Version {
	return &version{
		GitCommit:  GitCommitString,
		GitState:   GitStateString,
		GitBranch:  GitBranchString,
		GitSummary: GitSummaryString,
		BuildDate:  BuildDateString,
		Version:    VersionString,
		GoVersion:  runtime.Version(),
	}
}

type Version version

type version struct {
	GitCommit  string
	GitState   string
	GitBranch  string
	GitSummary string
	BuildDate  string
	Version    string
	GoVersion  string
}

func (v *version) GetGitCommit() string  { return v.GitCommit }
func (v *version) GetGitState() string   { return v.GitState }
func (v *version) GetGitBranch() string  { return v.GitBranch }
func (v *version) GetGitSummary() string { return v.GitSummary }
func (v *version) GetBuildDate() string  { return v.BuildDate }
func (v *version) GetVersion() string    { return v.Version }
func (v *version) GetGoVersion() string  { return v.GoVersion }
func (v *version) YAMLString() string    { return pretty.YAMLString(v) }
func (v *version) JsonString() string    { return pretty.JsonStringLine(v) }

func Decode(data []byte) (pkg.Version, error) {
	v := &version{}
	err := json.Unmarshal(data, v)
	if err != nil {
		return nil, err
	}
	return v, err
}
