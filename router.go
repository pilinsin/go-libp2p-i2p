package i2p

import(
	"time"
	"os"
	"os/exec"
	"strings"

	sam3 "github.com/eyedeekay/sam3"
)

func IsI2pRunning() bool{
	cmd := exec.Command("i2prouter", "status")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	stat, err := cmd.Output()
	return err == nil && strings.HasPrefix(string(stat), "I2P Service is running")
}
func IsSamRunning() bool{
	_, err := sam3.NewSAM(SAMHost)
	//sam.Close -> invalid mamory error
	return err == nil
}


var isI2pOwner bool

func StartI2pRouter() error{
	if ok := IsI2pRunning(); !ok{
		cmd := exec.Command("i2prouter", "start")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Start(); err != nil{return err}
		isI2pOwner = true
	}

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for{
		select{
		case <-ticker.C:
			if IsSamRunning(){
				return nil
			}
		default:
		}
	}
}

func StopI2pRouter(){
	if isI2pOwner{
		cmd := exec.Command("i2prouter", "stop")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		isI2pOwner = false
		cmd.Start()
	}
}

