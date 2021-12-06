package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "8080"
)

func fileHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Println()
	file, header, err := r.FormFile("file") /*
			func (r *Request) FormFile(key string) (multipart.File, *multipart.FileHeader, error)
			type File interface {
			io.Reader
			io.ReaderAt
			io.Seeker
			io.Closer
		}

		type Closer interface {
			Close() error
		}

		type : *multipart.FileHeader
		type FileHeader struct {
			Filename string
			Header   textproto.MIMEHeader
			Size     int64
			// contains filtered or unexported fields
		}
	*/
	if err != nil {
		log.Fatal("Something went wrong while getting the form file, ", err)
	}
	defer file.Close()
	out, pathError := os.Create("tmp/uploadedFile") /*
		Create creates or truncates the named file. If the file already exists, it is truncated. If the file does not
		exist, it is created with mode 0666 (before umask). If successful, methods on the returned File can be used for
		I/O; the associated file descriptor has mode O_RDWR. If there is an error, it will be of type *PathError
	*/
	if pathError != nil {
		log.Fatal("Something went wrong while creating a file, ", pathError)
	}
	defer out.Close()
	_, copyFileError := io.Copy(out, file) /*
		function signature:-
		------------------
		func Copy(dst Writer, src Reader) (written int64, err error)

		Copy copies from src to dst until either EOF is reached on src or an error occurs. It returns the number
		of bytes copied and the first error encountered while copying, if any.

		A successful Copy returns err == nil, not err == EOF. Because Copy is defined to read from src until EOF,
		it does not treat an EOF from Read as an error to be reported.

		If src implements the WriterTo interface, the copy is implemented by calling src.WriteTo(dst).
		Otherwise, if dst implements the ReaderFrom interface, the copy is implemented by calling dst.ReadFrom(src).
	*/
	if copyFileError != nil {
		log.Fatal("Something went wrong while copying the form file to our local server file, ", copyFileError)
	}
	fmt.Fprint(rw, "File successfully loaded with the name "+header.Filename)
}

func main() {
	server_add := CONN_HOST + ":" + CONN_PORT
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		ParsedTemplate, _ := template.ParseFiles("templates/upload-file.html")
		ParsedTemplate.Execute(rw, nil)
	})
	http.HandleFunc("/upload", fileHandler)
	if err := http.ListenAndServe(server_add, nil); err != nil {
		log.Fatal("Something went wrong while starting the http server : ", err)
	}
}
