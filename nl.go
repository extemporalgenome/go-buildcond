// +build !linux

package main

func init() {
	funcs = append(funcs, func() string { return "!linux" })
}
