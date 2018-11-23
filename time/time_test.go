package time

import (
	"testing"
	//"fmt"
	//"time"
)

func TestDehumanizeDuration(t *testing.T) {
	var v int64

	testValues := [...]string{"10h", "1w1d1h", "w1d1h", "1h30m", "1dh30m"}
	testResults := [...]int64{36000, 694800, 90000, 5400, 88200}

	for i := 0; i < len(testValues); i++ {
		t.Logf("Testing %s", testValues[i])
		if v, _ = DehumanizeDuration(testValues[i]); v != testResults[i] {
			t.Errorf("DehumanizeDuration failed, expected %d got: %d", testResults[i], v)
		}
	}

	t.Logf("Testing 10H")
	if _, err := DehumanizeDuration("10H"); err == nil {
		t.Errorf("DehumanizeDuration failed, expected error missing")
	}
}
