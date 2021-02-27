package main

import "github.com/jroimartin/gocui"

func openInputBox(g *gocui.Gui, v *gocui.View) error {
    maxX, maxY := g.Size()
    v, err := g.SetView("input", maxX/2-maxX/4, maxY/2, maxX/2+maxX/4, maxY/2+2)
    if err != nil {
        if err != gocui.ErrUnknownView {
            return err
        }
        vinput = v
        v.Editable = true
        v.Wrap = true
    }
    if _, err := g.SetCurrentView("input"); err != nil {
        return err
    }
    varray = append(varray, v)
    viewArr = append(viewArr, "input")
    curInputView = len(viewArr) - 1
    return nil
}

