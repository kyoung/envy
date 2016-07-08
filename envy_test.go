package envy

import (
	"fmt"
	"os"
	"testing"
)

func TestParseRow(t *testing.T) {
	var expectations = []struct {
		i string // input
		k string // output key
		v string // output value
	}{
		{"this=that", "this", "that"},
		{"oneNine123=3412", "oneNine123", "3412"},
		{"FOO=BAR", "FOO", "BAR"},
	}
	for _, exp := range expectations {
		k, v, err := getKeyValue([]byte(exp.i))
		if err != nil {
			fmt.Println(err)
			t.FailNow()
		}
		if exp.k != k || exp.v != v {
			t.Errorf("Input %s expected %s, %s. Got %s, %s.", exp.i, exp.k, exp.v, k, v)
		}
	}
}

func TestLoadEnvVars(t *testing.T) {
	loadEnvVars("./test-fixtures/.env")
	grappa := os.Getenv("GRAPPA")
	if grappa != "leftovers" {
		t.Fail()
	}
}
