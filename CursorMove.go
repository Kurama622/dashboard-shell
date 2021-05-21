package main

import (
    //"fmt"
    "strconv"
    "github.com/demonlord1997/gocui"
)

func cursorDown(g *gocui.Gui, v *gocui.View) error {
    var l string
    var err error
    var prel string
    var i int
    if v != nil {
        cx, cy := v.Cursor()
        prel, _ = v.Line(cy)
        for i=1; ;i++{
            l, err = v.Line(cy+i)
            if err != nil {
                l = ""
            }
            if prel != l{
                break
            }
        }
        if l != ""{
            if err = v.SetCursor(cx, cy+i); err != nil {
                ox, oy := v.Origin()
                if err = v.SetOrigin(ox, oy+i); err != nil {
                    return err
                }
            }
        }
    }
    return nil
}

func cursorUp(g *gocui.Gui, v *gocui.View) error {
    if v != nil {
        ox, oy := v.Origin()
        cx, cy := v.Cursor()
        if err := v.SetCursor(cx, cy-1); err != nil && oy > 0 {
            if err := v.SetOrigin(ox, oy-1); err != nil {
                return err
            }
        }
    }
    return nil
}

func cursorBottom(g *gocui.Gui, v *gocui.View) error {
    var l string
    var err error
    if v == g_varray[0]{
        if l, err = v.Line(g_folderMaxNumber-1); err != nil {
            l = ""
        }
        curNumber, _ := strconv.Atoi(l[1:2])
        diffNumber := g_folderMaxNumber-1 - curNumber
        if diffNumber != 0 {
            v.SetCursor(0, g_folderMaxNumber-1);
            v.MoveCursor(0, diffNumber, false)
        }else{
            v.SetCursor(0, g_folderMaxNumber-1);
        }
    }else if v == g_varray[1]{
        if l, err = v.Line(g_fileMaxNumber-1); err != nil {
            l = ""
        }
        curNumber, _ := strconv.Atoi(l[1:2])
        diffNumber := g_fileMaxNumber-1 - curNumber
        if diffNumber != 0 {
            v.SetCursor(0, g_fileMaxNumber-1);
            v.MoveCursor(0, diffNumber, false)
        }else{
            v.SetCursor(0, g_fileMaxNumber-1);
        }
    }else if v == g_varray[2] {
        v.SetCursor(0, g_RFNumber-2);
    }
    return nil
}


func cursorTop(g *gocui.Gui, v *gocui.View) error {
    var l string
    var err error

    if v == g_varray[0] || v == g_varray[1]{
        if l, err = v.Line(0); err != nil {
            l = ""
        }
        curNumber, _ := strconv.Atoi(l[1:2])
        if curNumber != 0 {
            v.SetCursor(0, 0);
            v.MoveCursor(0, -curNumber, false)
        }else{
            v.SetCursor(0, 0);
        }
    } else if v == g_varray[2] {
        v.SetCursor(0, 0);
    }
    return nil
}
