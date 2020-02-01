package balancediff

import (
	"testing"
	"time"
)

func TestNoDiffBeforeTime(t *testing.T) {
	d := D{}
	when := time.Date(2020, 2, 1, 0, 0, 0, 0, time.UTC)
	d.SetBalances(map[string]int64{"1": 1}, map[string]int64{"1": 1}, when)

	when = when.Add(time.Minute) // One minute later
	d.SetBalances(map[string]int64{"1": 1}, map[string]int64{"1": 2}, when)
	testDiff(t, d.Diff(when.Add(-time.Minute)), map[string]int64{}) // No diff

	testDiff(t, d.Diff(when), map[string]int64{"1": -1}) // Give a diff
	testDiff(t, d.Diff(when), map[string]int64{})        // Second time, no diff

	when = when.Add(time.Minute) // One minute later
	d.SetBalances(map[string]int64{"1": 1}, map[string]int64{"1": 1, "2": -1}, when)
	// We stepped one minute back, but should get the latest diff
	testDiff(t, d.Diff(when.Add(-time.Minute)), map[string]int64{"2": 1})

	// Next time we check, there is no diff since we only want to notify once
	testDiff(t, d.Diff(when.Add(-time.Minute)), map[string]int64{})
}

func testDiff(t *testing.T, diff, desired map[string]int64) {
	t.Helper()
	if len(diff) != len(desired) {
		t.Fatalf("Diff should be %d long, but was  %d", len(desired), len(diff))
	}
	for k, v := range desired {
		if diff[k] != v {
			t.Fatalf("Diff of \"%s\" should be %d, but was %d", k, v, diff[k])
		}
	}
}
