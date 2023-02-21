package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func main() {
	// Loading local.env file
	err := godotenv.Load("local.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	//splitting the TELEPRESENCE_MOUNTS variable to get volume mount path for the index.html file
	mount_path := strings.Split(os.Getenv("TELEPRESENCE_MOUNTS"), ":")
	volume_path := os.Getenv("TELEPRESENCE_ROOT")
	// Concatinating the Volume path and mount path to get absolute path of index.html
	var index_file_path string = volume_path + mount_path[0] + "/index.html"
	fmt.Printf("%s", index_file_path)

	// sfile to the index.html and opening it
	sfile, err := os.Open(index_file_path)
	if err != nil {
		log.Fatal("open file error")
	}
	defer sfile.Close()
	// dfile will be the absoluet path where the index.html should be created
	dfile, err := os.Create("./static/index.html")
	if err != nil {
		log.Fatal("open file error")
	}
	defer dfile.Close()
	//Copying index.html from copied volume mount to static folder in root directory of code
	_, err = io.Copy(dfile, sfile)
	if err != nil {
		log.Fatal("copy file error")
	}
	defer dfile.Close()
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.ListenAndServe(":3000", nil)
}
