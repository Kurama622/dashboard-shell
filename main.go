package main

import (
    "log"
    "errors"
    "runtime"
    "os/user"
    "os"
    "strings"
    "os/exec"
    "bytes"
    "gopkg.in/ini.v1"
    "github.com/jroimartin/gocui"
)

var (
    curInputView = 2
    viewArr = []string{"v1", "v2"}
    active  = 0
    selectView = 0
    varray = []*gocui.View{}
    vinput *gocui.View
    fileNumber = [10]int{0,1,2,3,4,5,6,7,8,9}
    folderlist []string
    recentFolderList []string
    filelist []string
    HOME string
    editor string
    folderLength = 22
    itermHeight = 11
    //selectNum = 0
)


func main() {
    HOME, _ = Home()
    cfg, err := ini.Load(HOME+"/.config/dashboard-shell/config.ini")
    if err != nil {
        log.Fatal("Fail to read file: ", err)
    }
    editor = cfg.Section("editor").Key("name").String()
    g, err := gocui.NewGui(gocui.OutputNormal)
    g.InputEsc = true
    if err != nil{
        log.Panicln(err)
    }

    defer g.Close()
    g.SetManagerFunc(layout)
    quicKeys(g)
}


func quit(g *gocui.Gui, v *gocui.View) error {
    return gocui.ErrQuit
}


// ==============GET $HOME ==================
func Home() (string, error) {
    user, err := user.Current()
    if nil == err {
        return user.HomeDir, nil
    }
    // cross compile support
    if "windows" == runtime.GOOS {
        return homeWindows()
    }
    // Unix-like system, so just assume Unix
    return homeUnix()
}

func homeUnix() (string, error) {
    // First prefer the HOME environmental variable
    if home := os.Getenv("HOME"); home != "" {
        return home, nil
    }
    // If that fails, try the shell
    var stdout bytes.Buffer
    cmd := exec.Command("sh", "-c", "eval echo ~$USER")
    cmd.Stdout = &stdout
    if err := cmd.Run(); err != nil {
        return "", err
    }

    result := strings.TrimSpace(stdout.String())
    if result == "" {
        return "", errors.New("blank output when reading home directory")
    }
    return result, nil
}

func homeWindows() (string, error) {
    drive := os.Getenv("HOMEDRIVE")
    path := os.Getenv("HOMEPATH")
    home := drive + path
    if drive == "" || path == "" {
        home = os.Getenv("USERPROFILE")
    }
    if home == "" {
        return "", errors.New("HOMEDRIVE, HOMEPATH, and USERPROFILE are blank")
    }

    return home, nil
}
