package main

// import block
import (
	"fmt"
	"os"

	fileutil "frontendmasters.com/go/files/fileUtils"
)

var (
	a int
	n string
)

func main() {

	var price = 34.04

	stringPrice := fmt.Sprintf("%.2f",price)
	fmt.Println(stringPrice)
	rootPath, _ := os.Getwd();
	filePath := rootPath + "/data/text.txt"

	content, err := fileutil.ReadTextFiles(filePath)

	if err != nil {
		fmt.Printf("Error Panic!!! %v", err)

	} else {
		fmt.Println(content)
		newContent := fmt.Sprintf("Original: %v\n Double Original: %v%v", content,content,content)
		fileutil.WriteToFile(filePath + ".output.txt", newContent)
	}
}