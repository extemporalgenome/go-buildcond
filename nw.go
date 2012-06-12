// +build !windows

package main

func init() {
	funcs = append(funcs, func() string { return "!windows" })
}
