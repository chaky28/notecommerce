package helpers

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

func GetHostIp() string {
	bash := exec.Command("hostname", "-I")
	output, err := bash.Output()
	if err != nil {
		log.Fatal("ERROR: Running command to get ip address --> ", err.Error())
	}

	return strings.Split(string(output), " ")[0]
}

func IsRunningInDockerContainer() bool {
	///Checks if application is running within a docker container

	if _, err := os.Stat("/.dockerenv"); err == nil {
		return true
	}

	return false
}
