package helpers

import (
	"log"
	"os/exec"
	"strings"
)

func GetHostIp() string {
	bash := exec.Command("hostname", "-I")
	if err := bash.Run(); err != nil {
		log.Fatal("ERROR: Running commando to get ip address --> ", err.Error())
	}

	output, err := bash.Output()
	if err != nil {
		log.Fatal("ERROR: Getting terminal output --> ", err.Error())
	}

	return strings.Split(string(output), " ")[0]
}
