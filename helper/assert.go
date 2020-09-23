package helper

import (
	"reflect"
	"testing"
)

const logTrue bool = true

func AssertObject(t *testing.T, e, o interface{}) {
	if !reflect.DeepEqual(e, o) {
		t.Errorf("expect valuse is %v but %v", e, o)
	} else if logTrue {
		t.Log("Good~")
	}
}

func AssertString(t *testing.T, e, o string) {
	if e != o {
		t.Errorf("expect valuse is %d but %d", e, o)
	} else if logTrue {
		t.Log("Good~")
	}
}

func AssertInt(t *testing.T, e, o int) {
	if e != o {
		t.Errorf("expect valuse is %d but %d", e, o)
	} else if logTrue {
		t.Log("Good~")
	}
}

func AssertIntArr(t *testing.T, e, o []int) {
	for i := range e {
		if e[i] != o[i] {
			t.Errorf("expect valuse is %d but %d", e, o)
		}
	}
	if logTrue {
		t.Log("Good~")
	}
}

func AssertBool(t *testing.T, e, o bool) {
	var es, os string
	if e {
		es = "true"
	} else {
		es = "false"
	}
	if o {
		os = "true"
	} else {
		os = "false"
	}
	if e != o {
		t.Errorf("expect valuse is %s but %s", es, os)
	} else if logTrue {
		t.Log("Good~")
	}
}
