package main

// Import any additional packages
import (
	"fmt"
	"os"
	"github.com/anishathalye/porcupine"
)

// Create event/operation type.
// ref: https://github.com/anishathalye/porcupine?tab=readme-ov-file#usage
type registerInput struct {

}

// Create a model : sequential specification of event/operation
// ref: https://github.com/anishathalye/porcupine?tab=readme-ov-file#usage
var registerModel = porcupine.Model{
	// Define the initlization state for the event/operation
	Init: func() interface{} {

	},

	// Define the state transition step for the event/operation
	Step: func(state, input, output interface{}) (bool, interface{}) {

	},

	// Define the event/operation for logging
	DescribeOperation: func(input, output interface{}) string {

	},
}

// Create histories and check for its linerarization against the sequential specification
// of the events/operations part of the histories.
func main() {
	// 1. Create the history. List of events/operations.

	// 2. Check the events against the model. 
	// Hints: ref porcupine.CheckEvents and porcupine.CheckEventsVerbose

	// 3. Visualize the history and linerization points.
	// Hints: porcupine.Visualize
}
