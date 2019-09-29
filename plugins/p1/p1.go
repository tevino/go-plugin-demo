package main

import (
	"fmt"

	"github.com/tevino/go-plugin-demo/pinger"
)

type p1 struct{}

func (p1) Ping() error {
	fmt.Println("Calling p1.Ping()")
	return nil
}

// P exports Pinger implementation.
// NOTE: A public interface should be explicitly used here.
// To use a type defined within the plugin scope, the application of plugin must import the plugin at compile time.
var P pinger.Pinger = &p1{}

// Name exports a single variable
// NOTE: A const will not work, since a const is neither a variable nor a function thus can't be accessed by Plugin.Lookup().
var Name string = "P1"

func init() {
	fmt.Println("init function works when loading the plugin")
}

func main() {
	fmt.Println("main is not executed by default, it is not necessary to be defined")
}
