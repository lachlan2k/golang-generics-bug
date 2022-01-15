package foo

import (
	"brokenthing/internal/bar"
	"fmt"
)

// I'm pretty sure that the pointer layout might get confused here?
func doFunkyStuff(test bar.BarInterface) {
	// UB doesn't trigger if this print isn't here:
	fmt.Printf("Test\n")

	/*
	 * Toggle the following to see broken vs working behaviour:
	 */

	test.UnmarshalBroken(`{"json_field":"test"}`)
	// test.UnmarshalWorking(`{"json_field":"test"}`)

	return
}

func TriggerUB() {
	var test bar.ConcreteStruct
	doFunkyStuff(&test)

	return
}
