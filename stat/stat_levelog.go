package stat

import (
    "github.com/vuleetu/web"
    "github.com/vuleetu/levelog"
)

func LogStat(log *levelog.LevelLogger) func(*web.Context) {
    return func(ctx *web.Context) {
        levelogStat(ctx, log)
    }
}

type StreamLog struct {
    ctx *web.Context
    ch chan byte
}

func (s *StreamLog) Write(data []byte) (n int, err error) {
    n, err = s.ctx.Write(data)
    s.ctx.Flusher().Flush()
    if err != nil {
        s.ch <- 1
    }
    return
}

func levelogStat(ctx *web.Context, log *levelog.LevelLogger) {
    ctx.WriteString("<pre>Now loading logs\n")
    ch := make(chan byte, 1)
    log.AddTraceLog(&StreamLog{ctx, ch}, "all")
    select {
    case  <-ch:
        levelog.Info("Error happened")
    }
}
