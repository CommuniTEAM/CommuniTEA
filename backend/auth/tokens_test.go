package auth_test

import (
	"strings"
	"testing"
)

// Placeholder b/c Codecov won't assess coverage for packages w/o a test
func TestTest(t *testing.T) {
	if strings.ToLower("TEST") != "test" {
		t.Fatalf("Test failed")
	}
}
