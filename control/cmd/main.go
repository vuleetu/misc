package main

import (
    "code.google.com/p/getopt"
    "encoding/binary"
    "io/ioutil"
    "net"
    "log"
    "fmt"
    "bytes"
)

var help = getopt.BoolLong("help", 'h', "Print help")
var sock = getopt.StringLong("sock", 's', "", "UNIX socket domain address")
var cmd = getopt.StringLong("cmd", 'c', "", "Command which will be sent")

func main() {
    getopt.Parse()
    if *help {
        getopt.Usage()
        return
    }

    if *sock == "" || *cmd == "" {
        getopt.Usage()
        return
    }

    conn, err := net.Dial("unix", *sock)
    if err != nil {
        log.Fatalln(err)
    }


    var buf bytes.Buffer
    err = binary.Write(&buf, binary.BigEndian, uint32(len(*cmd)))
    if err != nil {
        log.Fatalln(err)
    }
    buf.WriteString(*cmd)
    _, err = conn.Write(buf.Bytes())
    if err != nil {
        log.Fatalln(err)
    }

    data, err := ioutil.ReadAll(conn)
    if err != nil {
        log.Println("error is", err)
    }

    fmt.Println(string(data))
}

