package main

import "github.com/jroimartin/gocui"

func cursorDown(g *gocui.Gui, v *gocui.View) error {
    if v != nil {
        cx, cy := v.Cursor()
        if err := v.SetCursor(cx, cy+1); err != nil {
            ox, oy := v.Origin()
            if err := v.SetOrigin(ox, oy+1); err != nil {
                return err
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
    if v == g_varray[0] || v == g_varray[1]{
        v.SetCursor(0, 9);
    } else if v == g_varray[2] {
        v.SetCursor(0, g_RFNumber-2);
    }
    return nil
}


func cursorTop(g *gocui.Gui, v *gocui.View) error {
    v.SetCursor(0, 0);
    return nil
}
