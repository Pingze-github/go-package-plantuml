package b

import (
	sub2 "github.com/Pingze-github/go-package-plantuml/testdata/b/sub"
	a "sync"
	"github.com/Pingze-github/go-package-plantuml/testdata/b/suba"
)

type SB struct {
}

func (this  SB) Add(a sub2.SubSA, locker a.Locker, b B, subsa1 suba.SubSa1){}



