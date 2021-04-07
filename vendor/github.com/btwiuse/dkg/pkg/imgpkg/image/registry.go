package image

import (
	"fmt"
	"time"

	regauthn "github.com/google/go-containerregistry/pkg/authn"
	regname "github.com/google/go-containerregistry/pkg/name"
	regv1 "github.com/google/go-containerregistry/pkg/v1"
	regremote "github.com/google/go-containerregistry/pkg/v1/remote"
	regremtran "github.com/google/go-containerregistry/pkg/v1/remote/transport"
)

type RegistryOpts struct {
	CACertPaths []string
	VerifyCerts bool

	Username string
	Password string
	Token    string
	Anon     bool
}

type Registry struct {
	opts RegistryOpts
}

func NewRegistry(opts RegistryOpts) Registry {
	return Registry{opts}
}

func (i Registry) Generic(ref regname.Reference) (regv1.Descriptor, error) {
	opts, err := i.imageOpts()
	if err != nil {
		return regv1.Descriptor{}, err
	}

	desc, err := regremote.Get(ref, opts...)
	if err != nil {
		return regv1.Descriptor{}, err
	}

	return desc.Descriptor, nil
}

func (i Registry) Image(ref regname.Reference) (regv1.Image, error) {
	opts, err := i.imageOpts()
	if err != nil {
		return nil, err
	}

	return regremote.Image(ref, opts...)
}

func (i Registry) WriteImage(ref regname.Reference, img regv1.Image) error {
	opts, err := i.imageOpts()
	if err != nil {
		return err
	}

	err = i.retry(func() error { return regremote.Write(ref, img, opts...) })
	if err != nil {
		return fmt.Errorf("Writing image: %s", err)
	}

	return nil
}

func (i Registry) Index(ref regname.Reference) (regv1.ImageIndex, error) {
	opts, err := i.imageOpts()
	if err != nil {
		return nil, err
	}

	return regremote.Index(ref, opts...)
}

func (i Registry) WriteIndex(ref regname.Reference, idx regv1.ImageIndex) error {
	opts, err := i.imageOpts()
	if err != nil {
		return err
	}

	err = i.retry(func() error { return regremote.WriteIndex(ref, idx, opts...) })
	if err != nil {
		return fmt.Errorf("Writing image index: %s", err)
	}

	return nil
}

func (i Registry) ListTags(repo regname.Repository) ([]string, error) {
	opts, err := i.imageOpts()
	if err != nil {
		return nil, err
	}

	return regremote.List(repo, opts...)
}

func (i Registry) imageOpts() ([]regremote.Option, error) {
	var auth regremote.Option

        switch {
        case len(i.opts.Username) > 0:
                auth = regremote.WithAuth(&regauthn.Basic{Username: i.opts.Username, Password: i.opts.Password})
        case len(i.opts.Token) > 0:
                auth = regremote.WithAuth(&regauthn.Bearer{Token: i.opts.Token})
        case i.opts.Anon:
                auth = regremote.WithAuth(regauthn.Anonymous)
        default:
		auth = regremote.WithAuthFromKeychain(regauthn.DefaultKeychain)
        }

	return []regremote.Option{
		auth,
	}, nil
}

func (i Registry) retry(doFunc func() error) error {
	var lastErr error

	for i := 0; i < 5; i++ {
		lastErr = doFunc()
		if lastErr == nil {
			return nil
		}

		if tranErr, ok := lastErr.(*regremtran.Error); ok {
			if len(tranErr.Errors) > 0 {
				if tranErr.Errors[0].Code == regremtran.UnauthorizedErrorCode {
					return fmt.Errorf("Non-retryable error: %s", lastErr)
				}
			}
		}

		time.Sleep(1 * time.Second)
	}
	return fmt.Errorf("Retried 5 times: %s", lastErr)
}
