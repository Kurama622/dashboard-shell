package main

import (
    "github.com/demonlord1997/gocui"
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
    g_varray[(g_active) % len(g_viewArr)].Highlight = false
    nextIndex := (g_active + 1) % len(g_viewArr)
    name := g_viewArr[nextIndex]
    if _, err := setCurrentViewOnTop(g, name); err != nil {
        return err
    }
    //g.Cursor = true
    //g.Highlight = true
    g_active = nextIndex
    g_varray[g_active].Highlight = true
    return nil
}

func preView(g *gocui.Gui, v *gocui.View) error {
    g_varray[(g_active) % len(g_viewArr)].Highlight = false
    curActive := g_active - 1
    if curActive < 0 {
        curActive = len(g_viewArr)
    }
    prevIndex := (curActive) % len(g_viewArr)
    name := g_viewArr[prevIndex]
    if _, err := setCurrentViewOnTop(g, name); err != nil {
        return err
    }
    //g.Cursor = true
    g_active = prevIndex
    g_varray[g_active].Highlight = true
    return nil
}


// dynamic nextView
func dynextView(g *gocui.Gui, disableCurrent bool) error {
    next := g_active  % len(g_viewArr)
    //if next > len(g_viewArr)-1 {
        //next = 0
    //}
    //if next < 0 {
        //next = len(g_viewArr)
    //}
    if _, err := g.SetCurrentView(g_viewArr[next]); err != nil {
        return err
    }
    g_curInputView = next
    return nil
}

func delView(g *gocui.Gui) error {
    if len(g_viewArr) <= 1 {
        return nil
    }
    g_varray =  g_varray[:len(g_varray)-1]
    if err := g.DeleteView(g_viewArr[g_curInputView]); err != nil {
        return err
    }
    g_viewArr = append(g_viewArr[:g_curInputView], g_viewArr[g_curInputView+1:]...)
    return dynextView(g, false)
}

func exitCurView(g *gocui.Gui, v *gocui.View) error{
    return delView(g)
}
