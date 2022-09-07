package toolkit

import (
	"fmt"
	"testing"
)

func TestTools_RandomString(t *testing.T) {
	fmt.Println("Starting Test")
	var testTools Tools
	s := testTools.RandomString(10)

	if len(s) != 10 {
		t.Error("Wrong length random string returned")
	}

}


