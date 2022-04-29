package i2p

import(
	zero "github.com/eyedeekay/zerobundle"
)

func StartI2pRouter() error{
	return zero.ZeroAsFreestandingSAM()
}
// if i2p or i2pd are installed, StopI2pRouter does not stop the router.
func StopI2pRouter(){
	zero.StopZero()
}