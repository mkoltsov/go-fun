 // Sample program to show how you can't always get the
 // address of a value.
 package main

 import "fmt"

 // duration is a type with a base type of int.
 type duration int

 // format pretty-prints the duration value.
 func (d duration) pretty() string {
     return fmt.Sprintf("Duration: %d", *d)
 }

 // main is the entry point for the application.
 func main() {
     duration(42).pretty()

     // ./listing.go:: cannot call pointer method on duration()
     // ./listing.go:: cannot take the address of duration()
 }