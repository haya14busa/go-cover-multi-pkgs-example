package root

import "github.com/haya14busa/go-cover-multi-pkgs-example/sub"

func CoveredFromRoot() {
	_ = "ok"
}

func CoverSub() {
	sub.CoveredFromRoot()
	sub.CoveredFromSubAndRoot()
}
