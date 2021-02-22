package main

import (
    "fmt"
    "gopkg.in/ini.v1"
    "log"
    "strings"
    "bytes"
    "os/exec"
    "errors"
    "os"
    "runtime"
    "os/user"
)

func main() {
    HOME, _ := Home()
    cfg, err := ini.Load(HOME+"/Desktop/dashboard-shell/files.ini")
    if err != nil {
        log.Fatal("Fail to read file: ", err)
    }
    folderStrings := cfg.Section("files").Key("path").String()
    folderlist := strings.Split(folderStrings, ",")
    fmt.Println(folderlist)
    for _,c := range []string{"v1", "v2"} {
        fmt.Println(c)
    }
    //var itemNumber [10]int
}

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
