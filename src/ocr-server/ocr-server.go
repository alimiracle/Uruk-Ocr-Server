package main

import (
"encoding/json"
    "html/template"
    "log"
    "net/http"
"fmt"
"os"
"time"
"io"
 "io/ioutil"
"crypto/md5"
"strconv"
    "github.com/otiai10/gosseract"
)
type CONFIG struct {
Tempfile string
Lang string
}
type SERVER struct {
Port string
Url string
}

func upload(w http.ResponseWriter, r *http.Request) {
file, err := os.Open("/etc/ocrconfig/server.conf")
 if err != nil {
 log.Fatal(err)
 }

var all CONFIG

 data, err := ioutil.ReadAll(file)
 if err != nil {
 log.Fatal(err)
 }
rede := json.Unmarshal(data, &all)
    if rede != nil {
 log.Fatal(rede)
}

    if r.Method == "GET" {
        crutime := time.Now().Unix()
        h := md5.New()
        io.WriteString(h, strconv.FormatInt(crutime, 10))
        token := fmt.Sprintf("%x", h.Sum(nil))

        t, _ := template.ParseFiles("/etc/ocrconfig/upload.htm")
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
file, err := os.Open("/etc/ocrconfig/server.conf")
 if err != nil {
 log.Fatal(err)
 }

var all SERVER

 data, err := ioutil.ReadAll(file)
 if err != nil {
 log.Fatal(err)
 }
rede := json.Unmarshal(data, &all)
    if rede != nil {
 log.Fatal(rede)
}
http.HandleFunc(all.Url, upload)
    er := http.ListenAndServe(all.Port, nil) // setting listening port
    if er != nil {
        log.Fatal("ListenAndServe: ", er)
    }
}
