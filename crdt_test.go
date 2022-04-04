package i2p

import(
	"testing"
	"time"

	p2pcrypto "github.com/libp2p/go-libp2p-core/crypto"
	pv "github.com/pilinsin/p2p-verse"
	crdt "github.com/pilinsin/p2p-verse/crdt"
)

func TestTimeController(t *testing.T){
	b, err := NewI2pHost()
	checkError(t, err)
	bstrp, err := pv.NewBootstrap(b)
	defer bstrp.Close()
	bAddrInfo := bstrp.AddrInfo()
	t.Log("bootstrap AddrInfo: ", bAddrInfo)


	priv, pub, _ := p2pcrypto.GenerateEd25519Key(nil)
	pid := crdt.PubKeyToStr(pub)
	accesses := make(chan string)
	go func(){
		defer close(accesses)
		accesses <- pid
	}()
	va := crdt.NewVerse(NewI2pHost, "tc/c", false, false, bAddrInfo)
	ac, err := va.NewAccessController("ac", accesses)
	checkError(t, err)

	begin := time.Now()
	end := begin.Add(time.Hour)
	eps := time.Minute*2
	cool := time.Second*10
	n := 1
	vt := crdt.NewVerse(NewI2pHost, "tc/t", false, false, bAddrInfo)
	tc, err := vt.NewTimeController("tc", begin, end, eps, cool, n)
	checkError(t, err)

	opts0 := &crdt.StoreOpts{Priv: priv, Pub: pub, Ac: ac, Tc: tc}
	v0 := crdt.NewVerse(NewI2pHost, "tc/ta", false, false, bAddrInfo)
	db0, err := v0.NewStore("us", "updatableSignature", opts0)
	checkError(t, err)
	defer db0.Close()
	t.Log("db0 generated")


	v1 := crdt.NewVerse(NewI2pHost, "tc/tb", false, false, bAddrInfo)
	var db1 crdt.IStore
	for{
		db1, err = v1.LoadStore(db0.Address(), "updatableSignature")
		if err == nil{break}
		if err.Error() == "load error: sync timeout"{
			t.Log(err, ", now reloading...")
			time.Sleep(time.Second*10)
			continue
		}
		checkError(t, err)
	}
	defer db1.Close()
	t.Log("db1 generated")


	checkError(t, db0.Put("aaa", []byte("meow meow ^.^")))
	t.Log("put done")
	//wait for db1.tc.AutoGrant()
	time.Sleep(time.Minute*2)


	checkError(t, db1.Sync())
	v10, err := db1.Get(crdt.PubKeyToStr(opts0.Pub)+"/aaa")
	checkError(t, err)
	t.Log("db1.Get:", string(v10))


	t.Log("finished")
}