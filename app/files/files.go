package files

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func ReadFile(fileLocation string) string {
	fmt.Println("Reading file from", fileLocation)

	dirs, _ := os.ReadDir("/")
	for _, dir := range dirs {
		fmt.Println(dir.Name())
	}

	dirs, _ = os.ReadDir("/host_directories")
	for _, dir := range dirs {
		fmt.Println(dir.Name())
	}

	file, err := os.Open(fileLocation)
	if err != nil {
		log.Fatal("ERROR: Opening file --> " + err.Error())
	}

	defer func() {
		if err := file.Close(); err != nil {
			log.Fatal("ERROR: Closing file --> " + err.Error())
		}
	}()

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal("ERROR: Reading data from file -->" + err.Error())
	}

	return string(data)
}

func GetUserAndPasswordFromFileData(data string) (string, string) {
	fmt.Println("Getting user and password from file data")

	user := strings.Split(strings.Split(data, "\n")[0], "user=")[1]
	password := strings.Split(strings.Split(data, "\n")[1], "password=")[1]

	return strings.TrimSpace(user), strings.TrimSpace(password)
}
