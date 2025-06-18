package task2

import (
	"bufio"
	"errors"
	"io"
	"log"
	"os"
)

func ReadFile(fileName string) (result string, err error) {
	file, err := os.Open(fileName)
	
	if err != nil {
		log.Printf("error with opening file: %s", os.ErrNotExist.Error())

		if errors.Is(err, os.ErrNotExist) {
			return "", os.ErrNotExist
		} else if errors.Is(err, os.ErrPermission){
			return "", os.ErrPermission
		}

		return "", err
	}

	defer func() {
		if err := file.Close(); err != nil {
			log.Printf("error with closing file: %s", err.Error())
		}
	}()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err.Error() != io.EOF.Error() {
				log.Printf("error with reading file: %s", err.Error())
			}
			break
		}

		result += line
	}

	return result, nil
}