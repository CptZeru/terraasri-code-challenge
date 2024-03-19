package utils

import "testing"

func ThrowTestingErr(t *testing.T, expect interface{}, result interface{}) {
	t.Errorf("Expected: %v, Result: %v", expect, result)
}
