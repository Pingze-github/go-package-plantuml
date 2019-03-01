package main

import (
	"github.com/Pingze-github/go-package-plantuml/testdata/a"
)

func m1(a1 a.IA) {
	a1.Add()
}

func main() {
	m1(&a.SA{})
}

