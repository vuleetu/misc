package channel

import (
    "log"
    "reflect"
    "container/list"
)

func NewByteChan() (<-chan byte, chan<- byte) {
    r, w := NewChan(byte(0))
    return r.(chan byte), w.(chan byte)
}

func NewChan(typ interface{}) (interface{}, interface{}) {
    rc := reflect.MakeChan(reflect.ChanOf(reflect.BothDir, reflect.TypeOf(typ)), 0)
    wc := reflect.MakeChan(reflect.ChanOf(reflect.BothDir, reflect.TypeOf(typ)), 0)

    go loop(rc, wc)

    vprc := reflect.New(reflect.ChanOf(reflect.RecvDir, reflect.TypeOf(typ)))
    vprc.Elem().Set(rc)

    vpwc := reflect.New(reflect.ChanOf(reflect.SendDir, reflect.TypeOf(typ)))
    vpwc.Elem().Set(wc)

    return vprc.Elem().Interface(), vpwc.Elem().Interface()
}

func loop(rc reflect.Value, wc reflect.Value) {
    readCase := reflect.SelectCase{}
    readCase.Dir = reflect.SelectRecv
    readCase.Chan = wc

    writeCase := reflect.SelectCase{}
    writeCase.Dir = reflect.SelectSend
    writeCase.Chan = rc

    cases := []reflect.SelectCase{readCase, writeCase}

    l := list.New()

    for {
        tc := cases[:1]
        if l.Len() > 0 {
            writeCase.Send = reflect.ValueOf(l.Front().Value)
            cases[1] = writeCase
            tc = cases[:2]
        }


        idx, v, ok := reflect.Select(tc)
        if idx == 0 {
            if ok {
                l.PushBack(v.Interface())
            } else {
                //channel closed
                log.Println("Read channel closed")
                break
            }
        } else if idx == 1 {
            l.Remove(l.Front())
        }
    }
}
