package main

import "github.com/demonlord1997/gocui"

func openInputBox(g *gocui.Gui, v *gocui.View) error {
    maxX, maxY := g.Size()
    v, err := g.SetView("input", maxX/2-maxX/4, maxY/2, maxX/2+maxX/4, maxY/2+2)
    if err != nil {
        if err != gocui.ErrUnknownView {
            return err
        }
        g_vinput = v
        v.Editable = true
        v.Wrap = true
    }
    if _, err := g.SetCurrentView("input"); err != nil {
        return err
    }
    g_varray = append(g_varray, v)
    g_viewArr = append(g_viewArr, "input")
    g_curInputView = len(g_viewArr) - 1
    return nil
}

