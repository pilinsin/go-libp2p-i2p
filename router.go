package i2p

import(
	"fmt"
	"log"
	"time"
	"runtime"
	"path/filepath"
	"os"
	"os/exec"

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

func latestZeroDir() string{
	var platform string
	switch runtime.GOOS{
	case "windows":
		platform = "win-gui"
	case "darwin":
		platform = "mac-gui"
	default:
		platform = "linux"
	}

	return filepath.Join(".", "i2p-zero-"+platform+"."+zeroim.ZERO_VERSION)
}
func baseArgs() string{
	args := "--i2p.dir.base=base"
	os.MkdirAll("base", 0755)
	return args
}
func configArgs() string{
	args := "--i2p.dir.config=config"
	os.MkdirAll("config", 0755)
	return args
}
var cmd *exec.Cmd
func commandZero() (*exec.Cmd, error){
	if err := zeroim.Unpack(""); err != nil{
		log.Println(err)
	}

	latest := latestZero()
	latestAbs, err := filepath.Abs(filepath.Join(".", latestZeroDir(), latest))
	if err != nil{return nil, err}
	cmd = exec.Command(latestAbs, baseArgs(), configArgs())
	curAbsDir, _ := filepath.Abs(".")
	cmd.Dir = curAbsDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd, nil
}
func startZero() error{
	if zeroCmd, err := commandZero(); err != nil{
		return err
	}else{
		return zeroCmd.Start()
	}
}
func startI2pZeroRouter() error{
	if err := zeroim.Unpack(""); err != nil{
		log.Println(err)
	}

	latest := latestZero()
	log.Println("latest zero version is:", latest)
	if !zero.CheckZeroIsRunning(){
		log.Println("zero doesn't appear to be running.", latest)
		if err := startZero(); err != nil{
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
}
func isI2pRouterAvailable() error{
	if ok, conn := zero.Available(); !ok{
		return fmt.Errorf("i2p router in not available now")
	}else{
		log.Println("starting SAM")
		time.Sleep(3*time.Second)
		if err := zero.SAM(conn); err != nil{
			return err
		}
	}
	return nil
}

func StartI2pRouter() error{
	switch runtime.GOOS {
	case "android", "ios":
		return isI2pRouterAvailable()
	default:
		return startI2pZeroRouter()
	}
}
// if i2p or i2pd are installed, StopI2pRouter does not stop the router.
func StopI2pRouter(){
	if cmd == nil{return}
	switch runtime.GOOS {
	case "windows":
		cmd.Process.Signal(os.Kill)
	default:
		cmd.Process.Signal(os.Interrupt)
	}
}
