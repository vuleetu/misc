package daemon

import (
    "os"
)

//#include <unistd.h>
//#include <string.h>
//#include <sys/types.h>
//#include <sys/stat.h>
import "C"

func Start() {
    daemon1()
}

func daemon1() {
    if C.daemon(1, 0) != 0 {
        os.Exit(-1)
    }
}

func daemon2() {
    p := C.fork()
    if p > 0 {
        //main process exit
        os.Exit(0)
    } else if p < 0 {
        os.Exit(1)
    }

    C.setsid()

    p = C.fork()
    if p > 0 {
        //main process exit
        os.Exit(0)
    } else if p < 0 {
        os.Exit(1)
    }

    ////change working directory
    //C.chdir(C.CString("/tmp"))
    //C.umask(0)
}
