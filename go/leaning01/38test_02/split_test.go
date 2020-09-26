package split

import (
	"testing"
	"reflect"
)


func TestMoreSplit(t *testing.T) {
	got := S("abcd", "bc")
	want := []string{"a", "d"}

	if !reflect.DeepEqual(want, got) {
		t.Errorf("excepted:%v, got:%v", want, got)
	}
}
