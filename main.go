package main

import (
	"github.com/aleksmaus/gopkg/pkgutil"
)

const pkgFile = "/Users/amaus/work/gopkg/osquery-5.7.0.pkg"
const dstDir = "/Users/amaus/work/gopkg/out"

func main() {
	err := pkgutil.Expand(pkgFile, dstDir)
	if err != nil {
		panic(err)
	}
}
