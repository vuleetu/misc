package daemon

import (
    "os"
    "log"
)

//#include <unistd.h>
//#include <string.h>
//#include <sys/types.h>
//#include <sys/stat.h>
import "C"

func Start() {
    p := C.fork()
    log.Println(p)
    if p > 0 {
        //main process exit
        os.Exit(0)
    } else if p < 0 {
        os.Exit(1)
    }

    C.setsid()

    p = C.fork()
    log.Println(p)
    if p > 0 {
        //main process exit
        os.Exit(0)
    } else if p < 0 {
        os.Exit(1)
    }

    //change working directory
    C.chdir(C.CString("/tmp"))
    C.umask(0)
}
