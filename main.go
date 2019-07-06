package main

import (
	"encoding/json"
	"io"
	"math/rand"
	"net/http"
	"os"
	"time"
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

	file, _, err := r.FormFile("upload-file")
	response := Response{
		Code: "ERR",
		URL:  "",
	}

	if err != nil {
		json.NewEncoder(w).Encode(response)
		return
	}
	defer file.Close()

	filename := randStr(10)
	f, err := os.OpenFile("./uploads/"+filename+".png", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		json.NewEncoder(w).Encode(response)
		return
	}
	defer f.Close()
	io.Copy(f, file)

	response.Code = "OK"
	response.URL = filename + ".png"
	response.Filename = filename + ".png"

	json.NewEncoder(w).Encode(response)
	return
}

// 生成随机字符串
func randStr(num int) string {

	letters := []rune("abcdefghijklmnpqrstuvwxyzABCDEFGHIJKLMNPQRSTUVWXYZ123456789")
	str := make([]rune, num)

	rand.Seed(time.Now().UnixNano())
	for i := range str {
		str[i] = letters[rand.Intn(len(letters))]
	}

	return string(str)
}
