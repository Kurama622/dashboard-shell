package main

import (
    "github.com/jroimartin/gocui"
)


func setCurrentViewOnTop(g *gocui.Gui, name string) (*gocui.View, error) {
    if _, err := g.SetCurrentView(name); err != nil {
        return nil, err
    }
    return g.SetViewOnTop(name)
}

func chooseV1(g *gocui.Gui, v *gocui.View) error {
    if _, err := g.SetCurrentView("v1"); err != nil {
        return err
    }
    g.Cursor = true
    v.Highlight = true
    return nil
}
func chooseV2(g *gocui.Gui, v *gocui.View) error {
    if _, err := g.SetCurrentView("v2"); err != nil {
        return err
    }
    v.Highlight = true
    g.Cursor = true
    return nil
}

func nextView(g *gocui.Gui, v *gocui.View) error {
    varray[(active) % len(viewArr)].Highlight = false
    nextIndex := (active + 1) % len(viewArr)
    name := viewArr[nextIndex]
    if _, err := setCurrentViewOnTop(g, name); err != nil {
        return err
    }
    g.Cursor = true
    //g.Highlight = true
    active = nextIndex
    varray[active].Highlight = true
    return nil
}

func preView(g *gocui.Gui, v *gocui.View) error {
    varray[(active) % len(viewArr)].Highlight = false
    curActive := active - 1
    if curActive < 0 {
        curActive = len(viewArr)
    }
    prevIndex := (curActive) % len(viewArr)
    name := viewArr[prevIndex]
    if _, err := setCurrentViewOnTop(g, name); err != nil {
        return err
    }
    g.Cursor = true
    active = prevIndex
    varray[active].Highlight = true
    return nil
}


// dynamic nextView
func dynextView(g *gocui.Gui, disableCurrent bool) error {
    next := curInputView + 1
    if next > len(viewArr)-1 {
        next = 0
    }
    if _, err := g.SetCurrentView(viewArr[next]); err != nil {
        return err
    }
    curInputView = next
    return nil
}

func delView(g *gocui.Gui) error {
    if len(viewArr) <= 1 {
        return nil
    }
    varray =  varray[:len(varray)-1]
    if err := g.DeleteView(viewArr[curInputView]); err != nil {
        return err
    }
    viewArr = append(viewArr[:curInputView], viewArr[curInputView+1:]...)
    return dynextView(g, false)
}

func exitCurView(g *gocui.Gui, v *gocui.View) error{
    return delView(g)
}
