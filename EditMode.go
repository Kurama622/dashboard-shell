package main

import (
    "fmt"
    "github.com/demonlord1997/gocui"
)

func editMode(g *gocui.Gui, v *gocui.View) error {
    fmt.Println(g_editor)
    err := quit(g, v)
    if err != nil {
        return err
    }
    return nil
}
