package i2p

import(
	"testing"

	"fmt"
	"time"
	pv "github.com/pilinsin/p2p-verse"
	pubsub "github.com/pilinsin/p2p-verse/pubsub"
)


//go test -test.v=true .
func TestPubSub(t *testing.T){
	//closer, err := WaitForSamReady()
	//defer closer()
	//if err != nil{return}

	N := 10

	b, err := NewI2pHost()
	checkError(t, err)
	bstrp, err := pv.NewBootstrap(b)
	defer bstrp.Close()
	bAddrInfo := bstrp.AddrInfo()
	t.Log("bootstrap AddrInfo: ", bAddrInfo)

	<-time.Tick(time.Second*5)

	ps0, err01 := pubsub.NewPubSub(NewI2pHost, bAddrInfo)
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

	<-time.Tick(time.Second*10)

	ps1, err11 := pubsub.NewPubSub(NewI2pHost, bAddrInfo)
	checkError(t, err11)

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

