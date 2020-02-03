
[![GoDoc](https://godoc.org/github.com/philpearl/scratchbuild?status.svg)](https://godoc.org/github.com/philpearl/scratchbuild)

A library and tiny app for directly building very simple docker images & pushing them to a repository without involving the docker daemon, and with no complex dependencies.

```go
	o := scratchbuild.Options{
		Dir:      "./testdata",
		Name:     "philpearl/test",
		BaseURL:  "https://index.docker.io",
		Tag:      "latest",
		User:     "philpearl",
		Password: "sekret",
	}

	b := &bytes.Buffer{}
	if err := scratchbuild.TarDirectory("./testdata", b); err != nil {
		log.Fatalf("failed to tar layer. %s", err)
	}

	c := scratchbuild.New(&o)

	token, err := c.Auth()
	if err != nil {
		log.Fatalf("failed to authorize. %s", err)
	}
	c.Token = token

	if err := c.BuildImage(&scratchbuild.ImageConfig{
		Entrypoint: []string{"/app"},
	}, b.Bytes()); err != nil {
		log.Fatalf("failed to build and send image. %s", err)
	}
```