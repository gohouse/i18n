package test

import (
	"fmt"
	"io/ioutil"
)

func GetAllFile(dirname string, s []string) ([]string, error) {
	rd, err := ioutil.ReadDir(dirname)
	if err != nil {
		fmt.Println("read dir fail:", err)
		return s, err
	}
	for _, fi := range rd {
		if fi.IsDir() {
			fullDir := dirname + "/" + fi.Name()
			s, err = GetAllFile(fullDir, s)
			if err != nil {
				fmt.Println("read dir fail:", err)
				return s, err
			}
		} else {
			fullName := dirname + "/" + fi.Name()
			s = append(s, fullName)
		}
	}
	return s, nil
}

func main() {
	//遍历打印所有的文件名
	var s []string
	s, _ = GetAllFile("/Users/fizz/go/src/github.com/gohouse/i18n/examples/language", s)

	fmt.Printf("slice: %v", s)
}