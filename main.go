package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
)

// Response 图片上传成功后的响应体
type Response struct {
	Code     string `json:"code"`
	Filename string `json:"filename"`
	URL      string `json:"url"`
	Datetime string `json:"datetime"`
}

func main() {
	http.HandleFunc("/upload", uploadHandler)
	http.ListenAndServe(":8081", nil)
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	file, handler, err := r.FormFile("upload-file")
	response := Response{
		Code: "ERR",
		URL:  "",
	}

	if err != nil {
		json.NewEncoder(w).Encode(response)
		return
	}
	defer file.Close()

	f, err := os.OpenFile("./uploads/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		json.NewEncoder(w).Encode(response)
		return
	}
	defer f.Close()
	io.Copy(f, file)

	response.Code = "OK"
	response.URL = handler.Filename
	response.Filename = handler.Filename

	json.NewEncoder(w).Encode(response)
	return
}
