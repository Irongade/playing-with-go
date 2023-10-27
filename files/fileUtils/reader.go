package fileutils

import "os"

func ReadTextFiles(filename string) (string, error) {
	content, err := os.ReadFile(filename)

	if err != nil {
		// error occured

		// 3 ways to handle the error

		// return ""

		// or
		// panic("AHAHAHAHAHAHAHAHHAH")

		// or change the structure of the function

		return "", err

	} else {
		// no error occured

		return string(content), nil
	}
}