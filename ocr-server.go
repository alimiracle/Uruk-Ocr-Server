package main

import (
    "html/template"
    "log"
    "net/http"
"fmt"
"os"
"time"
"io"
"crypto/md5"
"strconv"
    "github.com/otiai10/gosseract"
)

func upload(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
        crutime := time.Now().Unix()
        h := md5.New()
        io.WriteString(h, strconv.FormatInt(crutime, 10))
        token := fmt.Sprintf("%x", h.Sum(nil))

        t, _ := template.ParseFiles("upload.html")
        t.Execute(w, token)
    } else {
        r.ParseMultipartForm(32 << 20)
        file, handler, err := r.FormFile("uploadfile")
        if err != nil {
            fmt.Println(err)
            return
        }
        defer file.Close()
        f, err := os.OpenFile("test/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
        if err != nil {
            fmt.Println(err)
            return
        }
        defer f.Close()
        io.Copy(f, file)
sc := "test/"+handler.Filename
    out := gosseract.Must(gosseract.Params{
            Src:       sc,
            Languages: "eng",
    })

    fmt.Fprintf(w, out)
os.Remove(sc)
    }
}

func main() {
os.Mkdir("test", os.ModePerm)
http.HandleFunc("/", upload)
    err := http.ListenAndServe(":8080", nil) // setting listening port
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
