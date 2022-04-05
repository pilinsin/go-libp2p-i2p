package i2p

import (
	"fmt"
	i2pd "github.com/eyedeekay/go-i2pd/goi2pd"
	sam3 "github.com/eyedeekay/sam3"
	"time"
)

func WaitForSamReady() (func(), error) {
	if _, err := sam3.NewSAM(SAMHost); err == nil {
		return func() {}, nil
	} else {
		fmt.Println(err)
	}

	closer := i2pd.InitI2PSAM()
	i2pd.StartI2P()

	time.Sleep(10 * time.Minute)
	_, err := sam3.NewSAM(SAMHost)
	if err == nil {
		closerFunc := func() {
			i2pd.StopI2P()
			closer()
		}
		return closerFunc, nil
	}

	i2pd.StopI2P()
	closer()
	return func() {}, err
}
