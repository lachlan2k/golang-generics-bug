package main

import (
	"brokenthing/internal/foo"
)

// For some reason, I can only consistently replicate this cross-package
func main() {
	foo.TriggerUB()
}
