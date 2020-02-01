package balancediff

import "time"

type D struct {
	a              map[string]int64
	b              map[string]int64
	updated        time.Time
	lastNoDiffTime time.Time
	reported       bool
}

// SetBalances sets both maps of balances at when time
func (d *D) SetBalances(a, b map[string]int64, when time.Time) {
	if len(diff(d.a, a))+len(diff(d.b, b)) > 0 {
		d.reported = false
	}
	d.a = a
	d.b = b
	d.updated = when
	if len(diff(a, b)) == 0 && d.lastNoDiffTime.Before(when) {
		d.lastNoDiffTime = when
	}
}

// Diff returns the diff in balances at when time
func (d *D) Diff(when time.Time) map[string]int64 {
	if !d.reported && d.lastNoDiffTime.Before(when) {
		d.reported = true
		return diff(d.a, d.b)
	}
	return map[string]int64{}
}

func diff(a, b map[string]int64) map[string]int64 {
	diff := map[string]int64{}
	// Construct name keys as map
	keys := map[string]bool{}
	for name := range a {
		keys[name] = true
	}
	for name := range b {
		keys[name] = true
	}

	// Iterate over names
	for name := range keys {
		if b[name] != a[name] {
			diff[name] = a[name] - b[name]
		}
	}
	return diff
}
