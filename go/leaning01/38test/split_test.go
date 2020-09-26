package split

import (
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	got := S("a:b:c", ":")
	want := []string{"a", "b", "c"}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("execpted:%v, got:%v", want, got)
	}
}
