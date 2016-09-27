/*
* uruk ocr server
* A Simple, small, powerful OCR web server
* Copyright (c) 2016 ali abdul ghani <alimiracle@riseup.net>
*    This Program is free software: you can redistribute it and/or modify
*    it under the terms of the GNU  General Public License as published by
*    the Free Software Foundation, either version 3 of the License, or
*    (at your option) any later version.
*    This Program is distributed in the hope that it will be useful,
*    but WITHOUT ANY WARRANTY; without even the implied warranty of
*    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
*    GNU General Public License for more details.
*    You should have received a copy of the GNU General Public License
*    along with this Program.  If not, see <http://www.gnu.org/licenses/>.
*/


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
type LANGCONFIG struct {
Lang string
}
type SERVER struct {
Port string
Url string
}

func upload(w http.ResponseWriter, r *http.Request) {
fileconf, err := os.Open("/etc/ocrconfig/lang.conf")
 if err != nil {
 log.Fatal(err)
 }

var all LANGCONFIG

 data, err := ioutil.ReadAll(fileconf)
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
            Languages: all.Lang,
    })

    fmt.Fprintf(w, out)
os.Remove(sc)
fileconf.Close()

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
