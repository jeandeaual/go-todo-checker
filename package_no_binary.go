// +build !go1.7

package main

import "go/build"

// IsBinaryPackage returns whether or not the package is a binary-only package
func IsBinaryPackage(pkg *build.Package) bool {
	// Binary-only packages are not supported by Go < 1.7
	return false
}
