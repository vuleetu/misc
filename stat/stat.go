package gostat

import (
    "encoding/json"
    "github.com/vuleetu/web"
    "github.com/vuleetu/levelog"
)

var _stats = map[string]*Stat{}

func Run(addr string) {
    go run(addr)
}

func RegisterStat(stat *Stat) {
    _stats[stat.Name] = stat
}

func run(addr string) {
    s := web.NewServer()
    s.Handle(`/`, listroute)
    s.Handle(`/([^/]+)(/.*)?`, route)
    s.Run(addr)
}

func listroute(ctx *web.Context) {
    cmds := []*StatInfo{}
    for _, s := range _stats {
        si := StatInfo{s.Name, s.Description, s.Stream}
        cmds = append(cmds, &si)
    }

    bin, err := json.Marshal(cmds)
    if err != nil {
        ctx.Abort(400, err.Error())
        return
    }

    ctx.Write(bin)
}

func route(ctx *web.Context, cmd, other string) {
    levelog.Info(cmd, other)
    if stat, ok := _stats[cmd]; ok {
        stat.Fun(ctx)
        return
    }
    ctx.Abort(404, "Command not support")
}
