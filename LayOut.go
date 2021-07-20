package main

import (
    "fmt"
    "log"
    "bytes"
    "strings"
    "strconv"
    "os/exec"
    "gopkg.in/ini.v1"
    "github.com/demonlord1997/gocui"
)

/*
|---------------|-----------------------------|
|               |                             |
maxX = k+1.5k+2+2*x
x =maxX/2 -1 - 2.5k/2

*/

func layout(g *gocui.Gui) error {
    maxX, maxY := g.Size()
    var beginBoxPosX int = (maxX)/2 - 5*(g_folderLength)/4 + 1
    var beginBoxPosY int
    var iShowV3 = (maxY/2 > 8) && (maxX > 75)
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
            v.Title = "FOLDERS"
        }
        v.SelFgColor = gocui.ColorRed
        cfg, err := ini.Load(g_HOME+"/.config/dashboard-shell/config.ini")
        if err != nil {
            log.Fatal("Fail to open folder: ", err)
        }
        folderStrings := cfg.Section("folders").Key("name").String()
        g_folderlist = strings.Split(folderStrings, ",")
        g_folderMaxNumber = len(g_folderlist)
        for i := range g_fileNumber[0:g_folderMaxNumber] {
            showString := "["+ strconv.Itoa(i) +"]" + "\t "+ g_folderlist[i]
            //fmt.Println(showString)
            ssLength := len(showString)
            if ssLength >= g_folderLength{
                //folderName := strings.Split(g_folderlist[i], "/")
                //fnLength := len(folderName)
                showShortString := showString[0:(g_folderLength-13)] + ".." + showString[ssLength-9:ssLength]
                //showShortString := showString[0:5]+folderName[fnLength-1][0:]
                fmt.Fprintln(v, showShortString)
            }else{
                fmt.Fprintln(v, showString)
            }

            // ======================================
            //folderName := strings.Split(g_folderlist[i], "/")
            //fmt.Fprintln(v,"["+ strconv.Itoa(i) +"]" + "\t "+folderName[len(folderName)-1])
        }
    }
    if v, err := g.SetView("v2", beginBoxPosX+g_folderLength+1, beginBoxPosY, maxX-beginBoxPosX-1, beginBoxPosY+g_itermHeight); err != nil {
        if err != gocui.ErrUnknownView {
            return err
        }
        g_varray = append(g_varray, v)
        if iShowV3{
            v.Frame = false
        }else{
            v.Frame = true
            v.Title = "FILES"
        }

        //v.Highlight = false
        v.SelFgColor = gocui.ColorRed
        cfg, err := ini.Load(g_HOME+"/.config/dashboard-shell/config.ini")
        if err != nil {
            log.Fatal("Fail to read file: ", err)
        }
        fileStrings := cfg.Section("files").Key("name").String()
        g_filelist = strings.Split(fileStrings, ",")
        g_fileMaxNumber = len(g_filelist)

        for i := range g_fileNumber[0:g_fileMaxNumber] {
            //showString := "["+ strconv.itoa(i) +"]" + "\t "+ fileName
            showString := "["+ strconv.Itoa(i) +"]" + "\t "+ g_filelist[i]
            ssLength := len(showString)
            if ssLength >= 2*g_folderLength{
                fileName := strings.Split(g_filelist[i], "/")
                //fmt.Println(fileName)
                fnLength := len(fileName)
                showShortString := showString[0:(2*g_folderLength-len(fileName[fnLength-1])-5)] + ".../" + fileName[fnLength-1]
                fmt.Fprintln(v, showShortString)
            }else{
                fmt.Fprintln(v, showString)
            }
        }
    }

    if iShowV3{
        if v, err := g.SetView("v3", (maxX-72)/2+1, beginBoxPosY-2-7, (maxX+75)/2, beginBoxPosY); err != nil {
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

