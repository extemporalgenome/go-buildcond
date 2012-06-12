// +build !darwin !windows

package main

func init() {
	funcs = append(funcs, func() string { return "!darwin !windows" })
}
