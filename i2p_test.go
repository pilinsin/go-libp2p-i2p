package i2p

import (
	"testing"

	pubsub "github.com/pilinsin/p2p-verse/pubsub"
	ipfs "github.com/pilinsin/p2p-verse/ipfs"
	crdt "github.com/pilinsin/p2p-verse/crdt"
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

	pubsub.BaseTestPubSub(t, NewI2pHost)
}
func testIpfs(t *testing.T){
	ipfs.BaseTestIpfs(t, NewI2pHost)
}
func testAccess(t *testing.T){
	crdt.BaseTestAccessController(t, NewI2pHost)
}
func testLog(t *testing.T){
	crdt.BaseTestLogStore(t, NewI2pHost)
}
func testHash(t *testing.T){
	crdt.BaseTestHashStore(t, NewI2pHost)
}
func testSignature(t *testing.T){
	crdt.BaseTestSignatureStore(t, NewI2pHost)
}
func testTime(t *testing.T){
	crdt.BaseTestTimeController(t, NewI2pHost)
}
func testUpdatableSignature(t *testing.T){
	crdt.BaseTestUpdatableSignatureStore(t, NewI2pHost)
}

func TestI2p(t *testing.T){
	rt := NewI2pRouter()
	checkError(t, rt.Start())

	t.Log("===== pubsub =====")
	testPubSub(t)
/*
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
*/

	rt.Stop()
}