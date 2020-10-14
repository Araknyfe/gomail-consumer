package main

import "log"

// FailOnError Fatal error when fail
func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
