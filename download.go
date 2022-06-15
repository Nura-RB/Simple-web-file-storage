package utils

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

func Download(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
		return
	}
	fileName := r.FormValue("getFile")

	var data []byte
	name := 1
	for {
		file, err := os.Open("storage/" + fileName + "/" + strconv.Itoa(name))
		if err != nil {
			break
		}
		buf, _ := io.ReadAll(file)
		data = append(data, buf...)
		file.Close()
		name++
	}

	for i := len(data) - 1; data[i] == 0; i-- {
		data = data[:len(data)-1]
	}
	data = data[:len(data)-1]
	downlFile, _ := os.Create(fileName)
	downlFile.Write(data)
	file, _ := ioutil.ReadFile(fileName)
	w.Write(file)
	downlFile.Close()
	err = os.Remove(fileName)
	if err != nil {
		fmt.Println(err)
	}
}
