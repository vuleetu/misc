package console

import (
    "fmt"
    "os"
    "strings"
)

var (
    Prompt = ">> "
    Descrition = "Console 0.1"
)

type Command struct {
    Name string
    Descrition string
    Callback func() error
}

var cmds = []*Command{}

func init() {
    Register(&Command{"help", "Print help information", PrintHelpInfo})
}

func Register(cmd *Command) {
    cmds = append(cmds, cmd)
}

func Start() {
    go loadConsole()
}

func Display(args ...interface{}) {
    fmt.Print(args...)
}

func loadConsole() {
    Display("\033[0;32;32m", Descrition, "\033[0m\n\n")
    var i = 1
    var cmd []byte
    buf := make([]byte, 1)
    for {
        cmd = []byte{}
        Display("\033[0;36m", i, Prompt, "\033[0m")

        for {
            _, err := os.Stdin.Read(buf)
            if err != nil {
                Display("Read command failed:", err, "\n")
                continue
            }

            if buf[0] == '\n' {
                break
            } else {
                cmd = append(cmd, buf...)
            }
        }

        if processCmd(strings.ToLower(string(cmd))) {
            i++
        }
    }
}

func processCmd(cmd string) bool {
    if cmd == "" {
        return false
    }

    for _, command := range cmds {
        if command.Name == cmd {
            if err := command.Callback(); err == nil {
                return true
            }

            return false
        }
    }

    Display("\033[0;32;31minvalid command:", cmd, "\033[0m\n")

    return false
}
func PrintHelpInfo() error {
    for _, command := range cmds {
        Display(" ", command.Name, "\t\t", command.Descrition, "\n")
    }

    return nil
}
