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

    text, ioerr := ioutil.ReadFile(HOME+"/.config/dashboard-shell/RecFolder")
    if ioerr != nil{
        return ioerr
    }
    recentFolderList = strings.Split(string(text), "\n")
    RFNumber := len(recentFolderList)
    for n_rf:=0; n_rf<RFNumber-1; n_rf++ {
        RecentFolderNumber = append(RecentFolderNumber, n_rf)
    }

    v, err := g.SetView("RecentFolder", maxX/2-3*folderLength/2-2, maxY/2-RFNumber/2-1, maxX/2+3*folderLength/2+1, maxY/2+RFNumber/2+1)
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
            showString = "["+ strconv.Itoa(i) +"] " + " \t"+ recentFolderList[i]
        }else{
            showString = "["+ strconv.Itoa(i) +"]" + " \t"+ recentFolderList[i]
        }
        ssLength = len(showString)
        if ssLength >= 3*folderLength{
            recentFolderName := strings.Split(recentFolderList[i], "/")
            rfnLength := len(recentFolderName)
            showShortString := showString[0:(3*folderLength-len(recentFolderName[rfnLength-1])-5)] + ".../" + recentFolderName[rfnLength-1]
            fmt.Fprintln(v, showShortString)
        }else{
            fmt.Fprintln(v, showString)
        }
    }

    varray = append(varray, v)
    viewArr = append(viewArr, "RecentFolder")
    curInputView = len(viewArr) - 1
    return nil
}

