package myrds

import "testing"

func TestHasZSet(t *testing.T) {
	_ = InitRedis()
	// checkError("Rds init error:", err)

	HasZSet("name")
}
