package main

import (
	"testing"
	"github.com/server-forecaster/util"
)

func TestAll(t *testing.T) {

	//tests.RunTests(t, api.GetTests())

	key := "patata99"
	secured, _ := util.Encrypt(key, "2:0")
	println("Secured ", secured)
	plain, _ := util.Decrypt(key, secured)
	println("Plain ", plain)

}
