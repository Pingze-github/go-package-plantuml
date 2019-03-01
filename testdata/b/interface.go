package b

import "sync"
import sub2 "github.com/Pingze-github/go-package-plantuml/testdata/b/sub"
import . "github.com/Pingze-github/go-package-plantuml/testdata/b/suba"

type B struct {}

type IA interface  {
	Add(a sub2.SubSA, locker sync.Locker, b B, subsa1 SubSa1)
}
