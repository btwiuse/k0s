module k0s.io/pkg/fonts

go 1.19

replace k0s.io/pkg/console => ../console

require (
	github.com/lukesampson/figlet v0.0.0-20190211215653-8a3ef4a6ac42
	github.com/mattn/go-isatty v0.0.14
	k0s.io/pkg/console v0.0.0-00010101000000-000000000000
)

require (
	github.com/pkg/errors v0.9.1 // indirect
	golang.org/x/sys v0.0.0-20220708085239-5a0f0661e09d // indirect
)
