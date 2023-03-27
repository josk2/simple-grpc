package testutil

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestRandom(t *testing.T) {
	rand.NewSource(time.Now().Unix())
	for i := 0; i < 5; i++ {
		fmt.Print(fmt.Printf("Random #%d: %d", i, rand.Intn(6)+1))
	}
}
