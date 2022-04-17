package i2p

import (
	"testing"

	pubsub "github.com/pilinsin/p2p-verse/pubsub"
	ipfs "github.com/pilinsin/p2p-verse/ipfs"
	crdt "github.com/pilinsin/p2p-verse/crdt"
)

//go test -test.v=true .
//go test -test.v=true -timeout 1h .
/*
func checkError(t *testing.T, err error, args ...interface{}) {
	if err != nil {
		args0 := make([]interface{}, len(args)+1)
		args0[0] = err
		copy(args0[1:], args)

		t.Fatal(args0...)
	}
}
*/

func TestPubSub(t *testing.T) {
	//closer, err := WaitForSamReady()
	//defer closer()
	//if err != nil{return}

	pubsub.BaseTestPubSub(t, NewI2pHost)
}


func TestIpfs(t *testing.T){
	ipfs.BaseTestIpfs(t, NewI2pHost)
}
func TestAccess(t *testing.T){
	crdt.BaseTestAccessController(t, NewI2pHost)
}
func TestLog(t *testing.T){
	crdt.BaseTestLogStore(t, NewI2pHost)
}
func TestHash(t *testing.T){
	crdt.BaseTestHashStore(t, NewI2pHost)
}
func TestSignature(t *testing.T){
	crdt.BaseTestSignatureStore(t, NewI2pHost)
}
func TestTime(t *testing.T){
	crdt.BaseTestTimeController(t, NewI2pHost)
}
func TestUpdatableSignature(t *testing.T){
	crdt.BaseTestUpdatableSignatureStore(t, NewI2pHost)
}
