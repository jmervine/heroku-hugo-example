package main

import (
	"net/http"
	"os"
	"path"
	"path/filepath"
)

func absPublic(root, public string) string {
	dir, err := filepath.Abs(path.Join(root, public))
	if err != nil {
		panic(err)
	}
	return dir
}

func main() {

	root := os.Getenv("ROOT")
	if root == "" {
		print("WARNING: 'ROOT' not set, using '" + root + "'\n")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
		print("WARNING: 'PORT' not set, using '" + port + "'\n")
	}

	public := absPublic(root, "public")

	print("Serving from '" + public + "'\n")
	print("Listneing on ':" + port + "'\n")
	panic(http.ListenAndServe(":"+port, http.FileServer(http.Dir(public))))
}
