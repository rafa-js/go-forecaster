package main

import (
	"testing"
	"github.com/server-forecaster/tests"
	"github.com/server-forecaster/tests/api"
)

func TestAll(t *testing.T) {

	tests.RunTests(t, api.GetTests())

}
