package stat

import (
    "os"
    "bytes"
    "bufio"
    "strings"
    "os/exec"
    "strconv"
    "runtime"
    "encoding/json"
    "github.com/vuleetu/web"
    "github.com/vuleetu/levelog"
)

func init() {
    RegisterStat(&Stat{"sys", "Show basic server process info", false, processStat})
}

func processStat(ctx *web.Context) {
    PidStr := strconv.Itoa(os.Getpid())
    cmd := exec.Command("ps", "u", "-p", PidStr)
    var buf bytes.Buffer
    cmd.Stdout = &buf
    err := cmd.Run()
    if err != nil {
        levelog.Error(err)
        ctx.Abort(400, err.Error())
        return
    }

    r := bufio.NewReader(&buf)
    data, _, err := r.ReadLine()
    if err != nil {
        levelog.Error(err)
        ctx.Abort(400, err.Error())
        return
    }

    data, _, err = r.ReadLine()
    if err != nil {
        levelog.Error(err)
        ctx.Abort(400, err.Error())
        return
    }

    levelog.Info("Data is", string(data))
    ds := strings.Split(string(data), " ")
    nds := make([]string, 0, len(ds))
    for i := 0; i < len(ds); i++ {
        if ds[i] != "" {
            nds = append(nds, ds[i])
        }
    }

    bin, err := json.Marshal(processStatInfo{nds[2], nds[3], nds[4], nds[5], runtime.NumGoroutine()})
    if err != nil {
        levelog.Error(err)
        ctx.Abort(400, err.Error())
        return
    }

    ctx.Write(bin)
}
