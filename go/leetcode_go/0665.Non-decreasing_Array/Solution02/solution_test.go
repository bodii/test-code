package main

import (
	"os"
	"testing"
)

func testFunc(nums []int, wantResult bool, t *testing.T) {
	t.Helper()
	if result := checkPossibility(nums); result != wantResult {
		t.Errorf(
			"checkPossibility func params: %v, \nresult: %v.\n!! but want result: %v\n",
			nums, result, wantResult)
	}
}

func Test1(t *testing.T) {
	nums := []int{3, 4, 2, 3}
	succResult := false
	testFunc(nums, succResult, t)
}

func Test2(t *testing.T) {
	nums := []int{5, 7, 1, 8}
	succResult := true
	testFunc(nums, succResult, t)
}

func Test3(t *testing.T) {
	nums := []int{4, 2, 3}
	succResult := true
	testFunc(nums, succResult, t)
}

func Test4(t *testing.T) {
	nums := []int{0, 4, 2, 3, 4}
	succResult := true
	testFunc(nums, succResult, t)
}

func Test5(t *testing.T) {
	nums := []int{4, 2, 1}
	succResult := false
	testFunc(nums, succResult, t)
}

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}
