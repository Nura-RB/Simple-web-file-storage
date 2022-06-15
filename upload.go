package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

func UploadFile(w http.ResponseWriter, r *http.Request) {
	file, handler, _ := r.FormFile("myFile")

	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	dst, err := os.Create(handler.Filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	parseFile(dst)
	dst.Close()
	os.Remove(dst.Name())
	fmt.Fprintf(w, "Successfully Uploaded File\n")
}

//--------------- Parsing and Saving the File ------------------------------------

func parseFile(dst *os.File) {
	file, err := os.Open(dst.Name())
	if err != nil {
		fmt.Println(err)
		return
	}
	err = os.Mkdir("storage/"+file.Name(), 0777)
	if err != nil {
		fmt.Println(err)
		return
	}
	data, _ := io.ReadAll(file)

	// filling by 0
	data = append(data, ' ')
	for len(data)%1048576 != 0 {
		data = append(data, '0')
	}

	// diving the file into parts by mb
	buf := 1048576 // 1mb
	name := 1
	for i := 1048576; i <= len(data); i += buf {
		filePart, err := os.Create("storage/" + file.Name() + "/" + strconv.Itoa(name))
		if err != nil {
			fmt.Println(err)
		}
		filePart.Write(data[i-buf : i])
		filePart.Close()

		name++
	}
	file.Close()
}
