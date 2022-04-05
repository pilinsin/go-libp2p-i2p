package i2p

import (
	"testing"
)

//go test -test.v=true .
func TestI2pd(t *testing.T) {
	closer, err := WaitForSamReady()
	defer closer()
	checkError(t, err)
}
