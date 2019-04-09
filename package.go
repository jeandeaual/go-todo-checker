// +build go1.7

package main

import "go/build"

// IsBinaryPackage return whether or not the package is a binary-only package
func IsBinaryPackage(pkg *build.Package) bool {
	return pkg.BinaryOnly
}
