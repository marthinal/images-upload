package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

const MaxUploadSize = 1024 * 1024

func main()  {
	http.HandleFunc("/images", uploadImage)
	err := http.ListenAndServe(":6090", nil)
	if err != nil {
		log.Fatal("Error starting the service")
	}
}

func uploadImage(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		w.Header().Set("Access-Control-Allow-Origin", "*")
		err := r.ParseMultipartForm(100000)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		for _, h := range r.MultipartForm.File["images"] {
			if h.Size > MaxUploadSize {
				http.Error(w, fmt.Sprintf("The uploaded image is too big: %s", h.Filename), http.StatusBadRequest)
				return
			}
			file, err := h.Open()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			defer file.Close()

			newFile, _ := os.Create("./images/" + h.Filename)
			_, err = io.Copy(newFile, file)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}