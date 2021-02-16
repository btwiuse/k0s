basicauth [![Build Status](https://drone.io/github.com/tango-contrib/basicauth/status.png)](https://drone.io/github.com/tango-contrib/basicauth/latest) [![](http://gocover.io/_badge/github.com/tango-contrib/basicauth)](http://gocover.io/github.com/tango-contrib/basicauth)
======

Middleware basicauth is a basic auth checker for [Tango](https://github.com/lunny/tango). 

## Installation

    go get github.com/tango-contrib/basicauth

## Simple Example

```Go
type AuthAction struct {}
func (a *AuthAction) Get() string {
    return "200"
}

func main() {
    tg := tango.Classic()
    tg.Use(basicauth.New(user, pass))
    tg.Get("/", new(AuthAction))
}
```

If you don't want some action to check auth, then
```Go
type NoAuthAction struct {
    basicauth.NoAuth
}
func (a *NoAuthAction) Get() string {
    return "200"
}

func main() {
    tg := tango.Classic()
    tg.Use(basicauth.New(user, pass))
    tg.Get("/", new(NoAuthAction))
}
```
will be ok.