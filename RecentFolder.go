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
    var n string

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
        v.Title = "RECENT FOLDER"
        //v.Frame = false
    }

    if _, err := g.SetCurrentView("RecentFolder"); err != nil {
        return err
    }

    for i := range RecentFolderNumber {
        if i<10 {
            showString = "["+ strconv.Itoa(i) +"] " + " \t"+ g_recentFolderList[i]
        }else{
            switch i{
                case 10:
                    n = "q"
                case 11:
                    n = "w"
                case 12:
                    n = "e"
                case 13:
                    n = "r"
                case 14:
                    n = "t"
                case 15:
                    n = "y"
                case 16:
                    n = "u"
                case 17:
                    n = "i"
                case 18:
                    n = "o"
                case 19:
                    n = "p"
            }
            showString = "["+ n +"] " + " \t"+ g_recentFolderList[i]
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

