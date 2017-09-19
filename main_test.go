package main

import (
	"testing"
	"github.com/server-forecaster/task"
)

func TestAll(t *testing.T) {

	//tests.RunTests(t, api.GetTests())

	task.UpdateMatches()
}
