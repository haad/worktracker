package wtime

import (
	"testing"
	"time"
)

func TestGetDurantionString(t *testing.T) {
	testDurations := [...]int64{60, 3600, 86400, 604800}
	testDurationStrings := [...]string{"1m0s", "1h0m0s", "24h0m0s", "168h0m0s"}

	for i := 0; i < len(testDurations); i++ {
		t.Logf("Testing entry: %d", testDurations[i])
		if GetDurantionString(testDurations[i]) != testDurationStrings[i] {
			t.Errorf("GetDurantionString failed, expected %s got: %s", testDurationStrings[i], GetDurantionString(testDurations[i]))
		}
	}
}

func TestCompareStartDate(t *testing.T) {
	shortForm := "02/01/2006"
	dateForm := "01/2006"

	testEntryDates := [...]string{"02/11/2018", "05/10/2017", "24/09/2018", "12/01/2019", "29/12/2017", "29/12/2017", "28/02/2019"}
	testDateSpecs := [...]string{"<12/2018", ">10/2017", "=09/2018", "=02/2019", "<01/2018", "<01/2017", "=03/2019"}
	testResults := [...]bool{true, true, true, false, true, false, false}

	for i := 0; i < len(testEntryDates); i++ {
		t.Logf("Testing entry: %s", testEntryDates[i])
		sd, _ := time.Parse(shortForm, testEntryDates[i])

		if CompareStartDate(testDateSpecs[i], sd.Unix()) != testResults[i] {
			t.Errorf("GetStartEndMonth failed, expected %t got: %t", testResults[i], CompareStartDate(testDateSpecs[i], sd.Unix()))
		}
	}
	if CompareStartDate("@", time.Now().Unix()) != true {
		t.Errorf("GetStartEndMonth failed, expected %t got: %t", true,
			CompareStartDate(time.Now().Format(dateForm), time.Now().Unix()))
	}
	if CompareStartDate("@", time.Date(time.Now().Year()-2, 1, 0, 0, 0, 0, 0, time.UTC).Unix()) != false {
		t.Errorf("GetStartEndMonth failed, expected %t got: %t", false,
			CompareStartDate(time.Now().Format(dateForm), time.Now().Unix()))
	}

}
func TestGetStartEndMonth(t *testing.T) {
	const shortForm = "01/2006"
	var start int64
	var end int64

	testValues := [...]string{"11/2018", "10/2017", "09/2018", "01/2019", "12/2017"}
	testResults := [...]int64{2592000, 2678400, 2592000, 2678400, 2678400}

	for i := 0; i < len(testValues); i++ {
		t.Logf("Testing %s", testValues[i])
		sd, _ := time.Parse(shortForm, testValues[i])

		//t.Logf("Parsed time is: %s, Month: %d, Emonth: %d", sd.Format(time.ANSIC), sd.Year(), time.January)
		//t.Logf("Testing %s -- %s - %s", testValues[i], tm2.Format(shortForm), tm.Format(shortForm))
		if start, end = GetStartEndMonth(sd); end-start != testResults[i] {
			t.Errorf("GetStartEndMonth failed, expected %d got: %d", testResults[i], (end - start))
		}
	}
}
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
