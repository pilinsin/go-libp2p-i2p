package i2p

import(
	"fmt"
	"log"
	"time"
	"runtime"
	"path/filepath"

	zero "github.com/eyedeekay/zerobundle"
	zeroim "github.com/eyedeekay/zerobundle/import"
)

func latestZero() string{
	switch runtime.GOOS {
	case "windows":
		return filepath.Join("router,i2p-zero.exe")
	case "darwin":
		return filepath.Join("router","bin", "launch.sh")
	default:
		return filepath.Join("router","bin", "i2p-zero")
	}
}

func StartI2pRouter() error{
	if err := zeroim.Unpack(""); err != nil{
		log.Println(err)
	}

	latest := latestZero()
	log.Println("latest zero version is:", latest)
	if !zero.CheckZeroIsRunning(){
		log.Println("zero doesn't appear to be running.", latest)
		if err := zero.StartZero(); err != nil{
			return err
		}
	}
	if ok, conn := zero.Available(); ok{
		log.Println("starting SAM")
		time.Sleep(3*time.Second)
		if err := zero.SAM(conn); err != nil{
			return err
		}
	}else{
		return fmt.Errorf("i2p router availability failure")
	}

	time.Sleep(time.Second)
	return nil

	//return zero.ZeroAsFreestandingSAM()
}
// if i2p or i2pd are installed, StopI2pRouter does not stop the router.
func StopI2pRouter(){
	zero.StopZero()
}
