package function

// AbsInt64 returns the absolute of the int64 value.
//
// DEPRECATED!!!
func AbsInt64(v int64) int64 {
	y := v >> 63       // y ← x >> 63
	return (v ^ y) - y // (x ⨁ y) - y
}
