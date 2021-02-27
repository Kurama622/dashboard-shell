package main

import (
    "fmt"
    "log"
    "bytes"
    "strings"
    "strconv"
    "os/exec"
    "gopkg.in/ini.v1"
    "github.com/jroimartin/gocui"
)

func layout(g *gocui.Gui) error {
    maxX, maxY := g.Size()
    var beginBoxPosX int
    var beginBoxPosY int
    var iShowV3 = (maxY/2 > 8) && (maxX > 75)
    beginBoxPosX = maxX/2 - 3*folderLength/2 - 2
    if iShowV3 {
        beginBoxPosY = maxY/2 - 1
    }else{
        beginBoxPosY = (maxY-itermHeight)/2
    }

    if v, err := g.SetView("v1", beginBoxPosX, beginBoxPosY, beginBoxPosX+folderLength, beginBoxPosY+itermHeight); err != nil {
        if err != gocui.ErrUnknownView {
            return err
        }
        varray = append(varray, v)
        //fmt.Println(active)
        v.Highlight = true
        if iShowV3{
            v.Frame = false
        }else{
            v.Frame = true
        }
        v.SelFgColor = gocui.ColorRed
        cfg, err := ini.Load(HOME+"/.config/dashboard-shell/config.ini")
        if err != nil {
            log.Fatal("Fail to open folder: ", err)
        }
        folderStrings := cfg.Section("folders").Key("name").String()
        folderlist = strings.Split(folderStrings, ",")

        for i := range fileNumber {
            folderName := strings.Split(folderlist[i], "/")
            fmt.Fprintln(v,"["+ strconv.Itoa(i) +"]" + "\t "+folderName[len(folderName)-1])
        }
        //v.Title = "FOLDERS"
    }
    if v, err := g.SetView("v2", beginBoxPosX+folderLength+2, beginBoxPosY, maxX-beginBoxPosX-1, beginBoxPosY+itermHeight); err != nil {
        if err != gocui.ErrUnknownView {
            return err
        }
        varray = append(varray, v)
        if iShowV3{
            v.Frame = false
        }else{
            v.Frame = true
        }

        //v.Highlight = false
        v.SelFgColor = gocui.ColorRed
        cfg, err := ini.Load(HOME+"/.config/dashboard-shell/config.ini")
        if err != nil {
            log.Fatal("Fail to read file: ", err)
        }
        fileStrings := cfg.Section("files").Key("name").String()
        filelist = strings.Split(fileStrings, ",")

        for i := range fileNumber {
            //showString := "["+ strconv.itoa(i) +"]" + "\t "+ fileName
            showString := "["+ strconv.Itoa(i) +"]" + "\t "+ filelist[i]
            ssLength := len(showString)
            if ssLength >= 2*folderLength{
                fileName := strings.Split(filelist[i], "/")
                fnLength := len(fileName)
                showShortString := showString[0:(2*folderLength-len(fileName[fnLength-1])-5)] + ".../" + fileName[fnLength-1]
                fmt.Fprintln(v, showShortString)
            }else{
                fmt.Fprintln(v, showString)
            }
        }
        //v.Title = "FILES"
    }

    if iShowV3{
        if v, err := g.SetView("v3", (maxX-72)/2, beginBoxPosY-2-7, (maxX+75)/2, beginBoxPosY); err != nil {
            if err != gocui.ErrUnknownView{
                return err
            }
            v.Frame = false
            cmd := exec.Command("/bin/bash", "-c", "echo $(date '+%Y.%m.%d') | figlet -f banner")
            var out bytes.Buffer
            cmd.Stdout = &out
            err := cmd.Run()
            if err != nil {
                log.Fatal(err)
            }
            fmt.Fprintln(v,out.String())
        }
    }

    // 使item可选
    if selectView == 0{
        if _, err := setCurrentViewOnTop(g, "v1"); err != nil {
            log.Panicln(err)
        }
        selectView = 1
    }
    return nil
}

