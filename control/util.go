package control

import (
    "io"
    "net"
    "os"
    "strings"
    "encoding/binary"
    "github.com/vuleetu/levelog"
)

type Callback func() error

var cmds = map[string]Callback{}

func Register(name string, callback Callback) {
    cmds[strings.ToLower(name)] = callback
}

func Start(addr string) {
    go start(addr)
}

var maximum_cmd_length uint32 = 1024

func SetMaximumCmdLenght(length uint32) {
    if length > 0 {
        maximum_cmd_length = length
    }
}

func start(addr string) {
    stat, err := os.Stat(addr)
    if err != nil {
        if !os.IsNotExist(err) {
            levelog.Error("Stat", addr, "failed")
            return
        }
    } else {
        if stat.IsDir() {
            levelog.Error("Address can not be a folder")
            return
        }

        err = os.Remove(addr)
        if err != nil {
            levelog.Error("Remove", addr, "failed")
            return
        }
    }

    l, err := net.Listen("unix", addr)
    if err != nil {
        levelog.Error("Listen at unix socket domain with address:", addr, "failed, error is:", err)
        return
    }

    for {
        conn, err := l.Accept()
        if err != nil {
            levelog.Error("Accept new connection from unix socket domain failed")
            continue
        }

        //will use sync block way to handle the conn instead of async
        processCmd(conn)
    }
}

func processCmd(conn net.Conn) {
    defer conn.Close()

    lenBuf := make([]byte, 4)
    _, err := io.ReadFull(conn, lenBuf)
    if err != nil {
        levelog.Error("Read length failed:", err)
        conn.Write([]byte("fail"))
        return
    }

    length := binary.BigEndian.Uint32(lenBuf)

    if length > maximum_cmd_length {
        levelog.Error("Length is greater than maximum cmd length")
        conn.Write([]byte("fail"))
        return
    }

    cmdBuf := make([]byte, length)
    _, err = io.ReadFull(conn, cmdBuf)
    if err != nil {
        levelog.Error("Read body failed:", err)
        conn.Write([]byte("fail"))
        return
    }

    if fun, ok := cmds[strings.ToLower(string(cmdBuf))]; ok {
        err = fun()
        if err != nil {
            levelog.Error("Command execute failed:", err)
            conn.Write([]byte("fail"))
            return
        }

        conn.Write([]byte("ok"))
        return
    }

    levelog.Warn("No callback for cmd", string(cmdBuf))
    conn.Write([]byte("fail"))
}
