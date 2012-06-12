// +build !darwin

package main

func init() {
	funcs = append(funcs, func() string { return "!darwin" })
}
