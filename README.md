### got

`go test -coverpkg=./... ./...` covers 100% lines, but we cannot know the fact from the output.

```sh
$ go test -coverpkg=./... ./...
ok      github.com/haya14busa/go-cover-multi-pkgs-example       0.001s  coverage: 83.3% of statements in ./...
ok      github.com/haya14busa/go-cover-multi-pkgs-example/sub   0.002s  coverage: 33.3% of statements in ./...
```

We also cannot get coverprofile.

```sh
# go test -coverprofile=coverage.out -coverpkg=./... ./...
cannot use test profile flag with multiple packages
```

### want

```sh
$ go test -coverpkg=./... ./...
ok      github.com/haya14busa/go-cover-multi-pkgs-example       0.001s  coverage: 83.3% of statements in ./...
ok      github.com/haya14busa/go-cover-multi-pkgs-example/sub   0.002s  coverage: 33.3% of statements in ./...
ok      github.com/haya14busa/go-cover-multi-pkgs-example/...   0.003s  coverage: 100.0% of statements in ./...
```

we can get coverprofile.

```sh
$ go test -coverprofile=coverage.out -coverpkg=./... ./...
```

of cource, we can produce html, and can confirm 100% lines are coverd.

```sh
$ go tool cover -html=coverage.out
```

### workaround using gocovmerge

- https://github.com/wadey/gocovmerge
- https://github.com/golang/go/issues/6909#issuecomment-233493644

```
make workaround
make testworkaround
```
