package main

import(
	"fmt"
	"time"
	"github.com/pilinsin/go-libp2p-i2p"
)

func main(){
	closer, err := i2p.WaitForSamReady()
	defer closer()
	if err != nil{return}

	h, err := i2p.NewI2pHost()
	fmt.Println(err)
	fmt.Println(h.ID())
	fmt.Println(h.Addrs())

	h2, err2 := i2p.NewI2pHost()
	fmt.Println(err2)
	fmt.Println(h2.ID())
	fmt.Println(h2.Addrs())
	
	h3, err3 := i2p.NewI2pHost()
	fmt.Println(err3)
	fmt.Println(h3.ID())
	fmt.Println(h3.Addrs())

	time.Sleep(10*time.Second)
	fmt.Println("finished")
}

