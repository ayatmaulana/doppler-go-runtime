# Doppler Go Runtime

Doppler is a secret management that I use as an alternative to Hashicorp Vautl, because it is free and can be used for small and medium projects.

But for now there is no golang package to use Doppler directly in our Golang Application Runtime, for that this package was created.

## 🍱&nbsp; Prerequesites

Here's what you have to prepare

- Golang > v1.18
- [Doppler Service Token](https://docs.doppler.com/docs/service-tokens)

## 🚀&nbsp; Installation

```go
$ go get github.com/ayatmaulana/doppler-go-runtime
```

## 🌴&nbsp; Usage

Before you use this package, make sure you have an `doppler.yaml` file in your root project folder with this content

```yaml
setup:
  project: your-project-name
  config: your-config-name
```

this package will auto discover your `doppler.yaml` file and use it to getting secret to Doppler Server

```Go
package main

import (
	"log"
	"os"

	"github.com/ayatmaulana/doppler-go-runtime"
)

func main(){
  doppler := dopplergoruntime.NewDopplerRuntime(dopplergoruntime.DopplerRuntimeConfig{
    Token: os.Getenv("DOPPLER_TOKEN"),
  })

  if err := doppler.Load(); err != nil {
    log.Println(err)
  }
}
```

if you want set the `project` and `config` programmatically you can also use the parameter

```Go
package main

import (
	"log"
	"os"

	"github.com/ayatmaulana/doppler-go-runtime"
)

func main(){
  doppler := dopplergoruntime.NewDopplerRuntime(dopplergoruntime.DopplerRuntimeConfig{
    Token: os.Getenv("DOPPLER_TOKEN"),
    Project: "your-project",
    Config: "your-config"
  })

  if err := doppler.Load(); err != nil {
    log.Println(err)
  }
}
```

and run with your Doppler Service Token like This

```bash
$ DOPPLER_TOKEN="dp.st.dev_xxxxxxxxxxxxxxxxx" go run main.go
```

## 💽&nbsp; Options

- `Token` - _string_
- `Project` - _string_
- `Config` - _string_
- `EnableDebug` - _boolean_

## 🤝&nbsp; Found a bug? Missing a specific feature?

Feel free to **file a new issue** with a respective title and description. If you already found a solution to your problem, **we would love to review your pull request**!.

## 📘&nbsp; License

This package is released under the under terms of the [ GNU GENERAL PUBLIC LICENSE ](LICENSE).
