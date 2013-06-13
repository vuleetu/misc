package main

import (
    "flag"
    "encoding/binary"
    "io/ioutil"
    "net"
    "log"
    "fmt"
    "bytes"
)

var sock = flag.String("sock", "", "UNIX sock file path")
var cmd = flag.String("cmd", "", "Command will be sent")

func main() {
    flag.Parse()

    if *sock == "" || *cmd == "" {
        flag.PrintDefaults()
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

