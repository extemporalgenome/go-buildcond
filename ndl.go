// +build !darwin !linux

package main

func init() {
	funcs = append(funcs, func() string { return "!darwin !linux" })
}
