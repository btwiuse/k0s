kubectl version info

```
~/k0s/pkg/version$ k0s hub &
~/k0s/pkg/version$ k0s kubectl -s http://127.0.0.1:8000 version -o yaml
clientVersion:
  buildDate: "1970-01-01T00:00:00Z"
  compiler: gc
  gitCommit: $Format:%H$
  gitTreeState: ""
  gitVersion: v0.0.0-master+$Format:%H$
  goVersion: go1.19
  major: ""
  minor: ""
  platform: linux/amd64
kustomizeVersion: v4.5.7
serverVersion:
  buildDate: "2021-09-23T22:46:18Z"
  compiler: gc
  gitCommit: 10bca343e85180f668522fe252552da20220cb7a
  gitTreeState: clean
  gitVersion: v1.23.0
  goVersion: go1.16.8
  major: "1"
  minor: "23"
  platform: linux/amd64
```

version info offered by this package:
```
~/k0s/pkg/version$ ./demo 
{
  "GitCommit": "184c6d3584743d2b08b872f9628cc1048699ee26",
  "GitState": "dirty",
  "GitBranch": "master",
  "GitSummary": "release-413-20230111-8-g184c6d358-dirty",
  "BuildDate": "2023-02-08T20:13:14Z",
  "Version": "v0.1.6",
  "GoVersion": "go1.20",
  "Platform": "linux/amd64",
  "Major": "0",
  "Minor": "1",
  "Compiler": "gc"
}
```
