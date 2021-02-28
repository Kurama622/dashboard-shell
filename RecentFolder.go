package main

import (
    "fmt"
    "strconv"
    "strings"
    "io/ioutil"
    "github.com/jroimartin/gocui"
)

func openRecentFolder(g *gocui.Gui, v *gocui.View) error{
    maxX, maxY := g.Size()
    var RecentFolderNumber []int
    var showString string
    var ssLength int

    text, ioerr := ioutil.ReadFile(g_HOME+"/.config/dashboard-shell/RecFolder")
    if ioerr != nil{
        return ioerr
    }
    g_recentFolderList = strings.Split(string(text), "\n")
    g_RFNumber = len(g_recentFolderList)
    for n_rf:=0; n_rf<g_RFNumber-1; n_rf++ {
        RecentFolderNumber = append(RecentFolderNumber, n_rf)
    }

    v, err := g.SetView("RecentFolder", maxX/2-3*g_folderLength/2-2, maxY/2-g_RFNumber/2-1, maxX/2+3*g_folderLength/2+1, maxY/2+g_RFNumber/2+1)
    if err != nil {
        if err != gocui.ErrUnknownView {
            return err
        }
        v.Highlight = true
        v.SelFgColor = gocui.ColorRed
        //v.Frame = false
    }

    if _, err := g.SetCurrentView("RecentFolder"); err != nil {
        return err
    }

    for i := range RecentFolderNumber {
        if i<10 {
            showString = "["+ strconv.Itoa(i) +"] " + " \t"+ g_recentFolderList[i]
        }else{
            showString = "["+ strconv.Itoa(i) +"]" + " \t"+ g_recentFolderList[i]
        }
        ssLength = len(showString)
        if ssLength >= 3*g_folderLength{
            recentFolderName := strings.Split(g_recentFolderList[i], "/")
            rfnLength := len(recentFolderName)
            showShortString := showString[0:(3*g_folderLength-len(recentFolderName[rfnLength-1])-5)] + ".../" + recentFolderName[rfnLength-1]
            fmt.Fprintln(v, showShortString)
        }else{
            fmt.Fprintln(v, showString)
        }
    }

    g_varray = append(g_varray, v)
    g_viewArr = append(g_viewArr, "RecentFolder")
    g_curInputView = len(g_viewArr) - 1
    return nil
}

