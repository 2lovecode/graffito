package top100

import "testing"

func Test_Lc32(t *testing.T) {
	data := map[string]int{
		"(()":    2,
		")()())": 4,
		"()(()":  2,
		"":       0,
	}
	for k, v := range data {
		vv := longestValidParentheses(k)
		if vv != v {
			t.Errorf("expect: %d, got: %d", v, vv)
		} else {
			t.Log("ok")
		}
	}

	for k, v := range data {
		vv := longestValidParentheses2(k)
		if vv != v {
			t.Errorf("expect: %d, got: %d", v, vv)
		} else {
			t.Log("ok")
		}
	}
	for k, v := range data {
		vv := longestValidParentheses3(k)
		if vv != v {
			t.Errorf("expect: %d, got: %d", v, vv)
		} else {
			t.Log("ok")
		}
	}
}
