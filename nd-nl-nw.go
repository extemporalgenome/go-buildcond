// +build !darwin,!linux,!windows

package main

func init() {
	funcs = append(funcs, func() string { return "!darwin,!linux,!windows" })
}
