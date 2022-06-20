package i2p

import (
	"errors"
	"os"
	"os/exec"
	"strings"
	"time"

	sam3 "github.com/eyedeekay/sam3"
)

func HasI2pRouter() bool {
	cmd := exec.Command("i2prouter", "status")
	stat, err := cmd.Output()
	return err != nil || !strings.Contains(string(stat), "Command 'i2prouter' not found")
}
func IsI2pRunning() bool {
	cmd := exec.Command("i2prouter", "status")
	stat, err := cmd.Output()
	return err == nil && strings.Contains(string(stat), "I2P Service is running")
}
func IsSamRunning() bool {
	_, err := sam3.NewSAM(SAMHost)
	//sam.Close -> invalid mamory error
	return err == nil
}

type I2pRouter struct {
	isOwner bool
}

func NewI2pRouter() *I2pRouter {
	return &I2pRouter{}
}
func (rt *I2pRouter) Start() error {
	if has := HasI2pRouter(); !has {
		return errors.New("i2p is not installed")
	}
	if ok := IsI2pRunning(); !ok {
		cmd := exec.Command("i2prouter", "start")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Start(); err != nil {
			return err
		}
		rt.isOwner = true
	}
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for {
		<-ticker.C
		if IsSamRunning() {
			return nil
		}
	}
}
func (rt *I2pRouter) Stop() {
	if rt.isOwner {
		cmd := exec.Command("i2prouter", "stop")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		rt.isOwner = false
		cmd.Start()
	}
}
