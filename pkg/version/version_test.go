package version

import (
	"testing"

	"k0s.io/k0s/pkg"
)

func check(t *testing.T, name string, got string, expected string) bool {
	if got != expected {
		t.Errorf("check %s failed: got: %q, expected %q\n", name, got, expected)
	}
	return true
}

func TestDecode(t *testing.T) {
	inputs := []string{
		`{"GitCommit":"deadbeef","GitState":"dirty","GitBranch":"master","GitSummary":"summary","BuildDate":"now","Version":"0.0.1"}`,
	}
	expects := [][]func(v pkg.Version) bool{
		[]func(v pkg.Version) bool{
			func(v pkg.Version) bool { return check(t, "GitCommit", v.GetGitCommit(), "deadbeef") },
			func(v pkg.Version) bool { return check(t, "GitBranch", v.GetGitBranch(), "master") },
			func(v pkg.Version) bool { return check(t, "GitState", v.GetGitState(), "dirty") },
			func(v pkg.Version) bool { return check(t, "GitSummary", v.GetGitSummary(), "summary") },
			func(v pkg.Version) bool { return check(t, "BuildDate", v.GetBuildDate(), "now") },
			func(v pkg.Version) bool { return check(t, "Version", v.GetVersion(), "0.0.1") },
		},
	}
	for i, input := range inputs {
		var (
			v, err   = Decode([]byte(input))
			expected = expects[i]
		)
		if err != nil {
			t.Error(err)
		}
		for _, check := range expected {
			check(v)
		}
	}
}
