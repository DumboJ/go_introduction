package main

import (
	"fmt"
	"time"
)

//error interface Error()string
type Logo_ocur struct {
	When time.Time
	what string
}

//Log_occur
func (log *Logo_ocur) Error() string {
	return fmt.Sprintf("at %v,%s", log.When, log.what)
}
func run() error {
	return &Logo_ocur{time.Now(), "jdbc connection error"}
}
func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
