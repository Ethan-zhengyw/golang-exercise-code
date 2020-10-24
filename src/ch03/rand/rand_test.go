package rand_test

import "testing"
import . "ch03/rand"

func TestRandom(t *testing.T) {
	Random()
}

func TestSeed(t *testing.T) {
	Seed(1)
}
