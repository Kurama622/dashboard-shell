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
    beginBoxPosX = maxX/2 - 3*g_folderLength/2 - 2
    if iShowV3 {
        beginBoxPosY = maxY/2 - 1
    }else{
        beginBoxPosY = (maxY-g_itermHeight)/2
    }

    if v, err := g.SetView("v1", beginBoxPosX, beginBoxPosY, beginBoxPosX+g_folderLength, beginBoxPosY+g_itermHeight); err != nil {
        if err != gocui.ErrUnknownView {
            return err
        }
        g_varray = append(g_varray, v)
        //fmt.Println(active)
        v.Highlight = true
        v.Wrap=true
        if iShowV3{
            v.Frame = false
        }else{
            v.Frame = true
        }
        v.SelFgColor = gocui.ColorRed
        cfg, err := ini.Load(g_HOME+"/.config/dashboard-shell/config.ini")
        if err != nil {
            log.Fatal("Fail to open folder: ", err)
        }
        folderStrings := cfg.Section("folders").Key("name").String()
        g_folderlist = strings.Split(folderStrings, ",")

        for i := range g_fileNumber {
            folderName := strings.Split(g_folderlist[i], "/")
            fmt.Fprintln(v,"["+ strconv.Itoa(i) +"]" + "\t "+folderName[len(folderName)-1])
        }
        //v.Title = "FOLDERS"
    }
    if v, err := g.SetView("v2", beginBoxPosX+g_folderLength+2, beginBoxPosY, maxX-beginBoxPosX-1, beginBoxPosY+g_itermHeight); err != nil {
        if err != gocui.ErrUnknownView {
            return err
        }
        g_varray = append(g_varray, v)
        if iShowV3{
            v.Frame = false
        }else{
            v.Frame = true
        }

        //v.Highlight = false
        v.SelFgColor = gocui.ColorRed
        cfg, err := ini.Load(g_HOME+"/.config/dashboard-shell/config.ini")
        if err != nil {
            log.Fatal("Fail to read file: ", err)
        }
        fileStrings := cfg.Section("files").Key("name").String()
        g_filelist = strings.Split(fileStrings, ",")

        for i := range g_fileNumber {
            //showString := "["+ strconv.itoa(i) +"]" + "\t "+ fileName
            showString := "["+ strconv.Itoa(i) +"]" + "\t "+ g_filelist[i]
            ssLength := len(showString)
            if ssLength >= 2*g_folderLength{
                fileName := strings.Split(g_filelist[i], "/")
                fnLength := len(fileName)
                showShortString := showString[0:(2*g_folderLength-len(fileName[fnLength-1])-5)] + ".../" + fileName[fnLength-1]
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
    if g_selectView == 0{
        if _, err := setCurrentViewOnTop(g, "v1"); err != nil {
            log.Panicln(err)
        }
        g_selectView = 1
    }
    return nil
}

