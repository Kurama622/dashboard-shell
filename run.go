package main

import (
    "fmt"
    "github.com/demonlord1997/gocui"
)


func runFolderShell(g *gocui.Gui, v *gocui.View) error {
    var err error
    _, cy := v.Cursor()

    fmt.Println("cd " + g_folderlist[cy])
    err = quit(g, v)
    if err != nil {
        return err
    }
    return nil
}

func runRecentFolderShell(g *gocui.Gui, v *gocui.View) error {
    var err error
    _, cy := v.Cursor()

    fmt.Println("cd " + g_recentFolderList[cy])
    err = quit(g, v)
    if err != nil {
        return err
    }
    return nil
}
func runFileShell(g *gocui.Gui, v *gocui.View) error {
    var err error
    _, cy := v.Cursor()

    fmt.Println(g_editor+ " " + g_filelist[cy])
    err = quit(g, v)
    if err != nil {
        return err
    }
    return nil
}


func runShell(g *gocui.Gui, v *gocui.View) error {
    var l string
    var err error

    _, cy := v.Cursor()
    if l, err = v.Line(cy); err != nil {
        l = ""
    }

    fmt.Println(l)
    err = quit(g, v)
    if err != nil {
        return err
    }
    return nil
}

