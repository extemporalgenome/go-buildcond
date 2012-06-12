// +build !linux !windows

package main

func init() {
	funcs = append(funcs, func() string { return "!linux !windows" })
}
