package pattern

import (
	"testing"
)

func TestSet(t *testing.T) {
	Singleton().Set("item1", "value1")
}

func TestGet(t *testing.T) {
	keyOne := "key1"
	valueOne := "value1"

	valueTwo := "value2"
	keyTwo := "key2"

	Singleton().Set(keyOne, valueOne)
	Singleton().Set(keyTwo, valueTwo)

	if v, _ := Singleton().Get(keyOne); v != valueOne {
		t.Errorf("%s is different from %s", v, valueOne)
	}

	if v, _ := Singleton().Get(keyTwo); v != valueTwo {
		t.Errorf("%s is different from %s", v, valueTwo)
	}
}
