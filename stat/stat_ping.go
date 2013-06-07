package stat

import (
    "github.com/vuleetu/web"
)

func init() {
    RegisterStat(&Stat{"ping", "Ping server", false, pingStat})
}

func pingStat(ctx *web.Context) {
    ctx.WriteString("pong")
}
