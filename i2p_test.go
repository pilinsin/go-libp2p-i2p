package i2p

import (
	"testing"
	"time"

	crdt "github.com/pilinsin/p2p-verse/crdt"
	ipfs "github.com/pilinsin/p2p-verse/ipfs"
	pubsub "github.com/pilinsin/p2p-verse/pubsub"
)

//go test -test.v=true .
//go test -test.v=true -timeout 1h .

func checkError(t *testing.T, err error, args ...interface{}) {
	if err != nil {
		args0 := make([]interface{}, len(args)+1)
		args0[0] = err
		copy(args0[1:], args)

		t.Fatal(args0...)
	}
}

func testPubSub(t *testing.T) {
	//closer, err := WaitForSamReady()
	//defer closer()
	//if err != nil{return}
	now := time.Now()
	pubsub.BaseTestPubSub(t, NewI2pHost)
	t.Log(time.Now().Sub(now).String())
}
func testIpfs(t *testing.T) {
	now := time.Now()
	ipfs.BaseTestIpfs(t, NewI2pHost)
	t.Log(time.Now().Sub(now).String())
}
func testAccess(t *testing.T) {
	now := time.Now()
	crdt.BaseTestAccessController(t, NewI2pHost)
	t.Log(time.Now().Sub(now).String())
}
func testLog(t *testing.T) {
	now := time.Now()
	crdt.BaseTestLogStore(t, NewI2pHost)
	t.Log(time.Now().Sub(now).String())
}
func testHash(t *testing.T) {
	now := time.Now()
	crdt.BaseTestHashStore(t, NewI2pHost)
	t.Log(time.Now().Sub(now).String())
}
func testSignature(t *testing.T) {
	now := time.Now()
	crdt.BaseTestSignatureStore(t, NewI2pHost)
	t.Log(time.Now().Sub(now).String())
}
func testTime(t *testing.T) {
	now := time.Now()
	crdt.BaseTestTimeLimit(t, NewI2pHost)
	t.Log(time.Now().Sub(now).String())
}
func testUpdatableSignature(t *testing.T) {
	now := time.Now()
	crdt.BaseTestUpdatableSignatureStore(t, NewI2pHost)
	t.Log(time.Now().Sub(now).String())
}

func TestI2p(t *testing.T) {
	rt := NewI2pRouter()
	checkError(t, rt.Start())

	t.Log("===== pubsub =====")
	testPubSub(t)
	t.Log("===== ipfs =====")
	testIpfs(t)
	t.Log("===== log =====")
	testLog(t)
	t.Log("===== hash =====")
	testHash(t)
	t.Log("===== signature =====")
	testSignature(t)
	t.Log("===== updatablesignature =====")
	testUpdatableSignature(t)
	t.Log("===== access =====")
	testAccess(t)
	t.Log("===== time =====")
	testTime(t)

	rt.Stop()
}
