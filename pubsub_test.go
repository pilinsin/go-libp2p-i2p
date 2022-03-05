package i2p

import(
	"testing"

	"fmt"
	"time"
	"context"
	pv "github.com/pilinsin/p2p-verse"
	pubsub "github.com/pilinsin/p2p-verse/pubsub"
)
func checkError(t *testing.T, err error, args ...interface{}){
	if err != nil{
		args0 := make([]interface{}, len(args)+1)
		args0[0] = err
		copy(args0[1:], args)

		t.Fatal(args0...)
	}
}

//go test -test.v=true .
func TestPubSub(t *testing.T){
	//closer, err := i2p.WaitForSamReady()
	//defer closer()
	//if err != nil{return}

	N := 10

	b, err := NewI2pHost()
	checkError(t, err)
	bstrp, err := pv.NewBootstrap(b, "pubsub:ejvoaenvaeo;vn;aeo")
	defer bstrp.Close()
	bAddrInfo := bstrp.AddrInfo()
	t.Log("bootstrap AddrInfo: ", bAddrInfo)

	h0, err0 := NewI2pHost()
	checkError(t, err0)

	ps0, err01 := pubsub.NewPubSub(context.Background(), h0, nil, bAddrInfo)
	checkError(t, err01)

	tpc0, err02 := ps0.JoinTopic("test topic")
	checkError(t, err02)

	go func(){
		defer tpc0.Close()
		itr := 0
		for{
			mess, err := tpc0.GetAll()
			t.Log(itr, err)
			if err == nil && len(mess)>0{
				itr += len(mess)
				for _, mes := range mess{
					t.Log(string(mes.Data))
				}
			}
	
			if itr >= N{return}
		}
	}()

	<-time.Tick(time.Second*5)

	h1, err1 := NewI2pHost()
	checkError(t, err1)

	ps1, err11 := pubsub.NewPubSub(context.Background(), h1, nil, bAddrInfo)
	checkError(t, err11)


	t.Log("h0 id             : ", h0.ID())
	t.Log("h0 address        : ", h0.Addrs())
	t.Log("h0 connected peers:", h0.Network().Peers())
	t.Log("h1 id             : ", h1.ID())
	t.Log("h0 address        :", h1.Addrs())
	t.Log("h1 connected peers:", h1.Network().Peers())

	tpc1, err12 := ps1.JoinTopic("test topic")
	checkError(t, err12)
	defer tpc1.Close()
	t.Log("topic peers list  :", tpc1.ListPeers())
	for i := 0; i<N; i++{
		tpc1.Publish([]byte(fmt.Sprintln("message ", i)))
	}


	<-time.Tick(10*time.Second)
	t.Log("finished")
}

