package main

import (
    "fmt"
    "github.com/jroimartin/gocui"
)


func goToNumberPos0(g *gocui.Gui, v *gocui.View) error {
    if v == varray[0]{
        fmt.Println("cd " + folderlist[0])
    } else if v == varray[1] {
        fmt.Println(editor + " " + filelist[0])
    } else if v == varray[2] {
        fmt.Println("cd " + recentFolderList[0])
    }
    err := quit(g, v)
    if err != nil {
        return err
    }
    return nil
}

func goToNumberPos1(g *gocui.Gui, v *gocui.View) error {
    if v == varray[0]{
        fmt.Println("cd " + folderlist[1])
    }else if v == varray[1] {
        fmt.Println(editor + " " + filelist[1])
    } else if v == varray[2] {
        fmt.Println("cd " + recentFolderList[1])
    }
    err := quit(g, v)
    if err != nil {
        return err
    }
    return nil
}

func goToNumberPos2(g *gocui.Gui, v *gocui.View) error {
    if v == varray[0]{
        fmt.Println("cd " + folderlist[2])
    }else if v == varray[1] {
        fmt.Println(editor + " " + filelist[2])
    } else if v == varray[2] {
        fmt.Println("cd " + recentFolderList[2])
    }
    err := quit(g, v)
    if err != nil {
        return err
    }
    return nil
}

func goToNumberPos3(g *gocui.Gui, v *gocui.View) error {
    if v == varray[0]{
        fmt.Println("cd " + folderlist[3])
    }else if v == varray[1] {
        fmt.Println(editor + " " + filelist[3])
    } else if v == varray[2] {
        fmt.Println("cd " + recentFolderList[3])
    }
    err := quit(g, v)
    if err != nil {
        return err
    }
    return nil
}

func goToNumberPos4(g *gocui.Gui, v *gocui.View) error {
    if v == varray[0]{
        fmt.Println("cd " + folderlist[4])
    }else if v == varray[1] {
        fmt.Println(editor + " " + filelist[4])
    } else if v == varray[2] {
        fmt.Println("cd " + recentFolderList[4])
    }
    err := quit(g, v)
    if err != nil {
        return err
    }
    return nil
}

func goToNumberPos5(g *gocui.Gui, v *gocui.View) error {
    if v == varray[0]{
        fmt.Println("cd " + folderlist[5])
    }else if v == varray[1] {
        fmt.Println(editor + " " + filelist[5])
    } else if v == varray[2] {
        fmt.Println("cd " + recentFolderList[5])
    }
    err := quit(g, v)
    if err != nil {
        return err
    }
    return nil
}

func goToNumberPos6(g *gocui.Gui, v *gocui.View) error {
    if v == varray[0]{
        fmt.Println("cd " + folderlist[6])
    }else if v == varray[1] {
        fmt.Println(editor + " " + filelist[6])
    } else if v == varray[2] {
        fmt.Println("cd " + recentFolderList[6])
    }
    err := quit(g, v)
    if err != nil {
        return err
    }
    return nil
}

func goToNumberPos7(g *gocui.Gui, v *gocui.View) error {
    if v == varray[0]{
        fmt.Println("cd " + folderlist[7])
    }else if v == varray[1] {
        fmt.Println(editor + " " + filelist[7])
    } else if v == varray[2] {
        fmt.Println("cd " + recentFolderList[7])
    }
    err := quit(g, v)
    if err != nil {
        return err
    }
    return nil
}

func goToNumberPos8(g *gocui.Gui, v *gocui.View) error {
    if v == varray[0]{
        fmt.Println("cd " + folderlist[8])
    }else if v == varray[1] {
        fmt.Println(editor + " " + filelist[8])
    } else if v == varray[2] {
        fmt.Println("cd " + recentFolderList[8])
    }
    err := quit(g, v)
    if err != nil {
        return err
    }
    return nil
}

func goToNumberPos9(g *gocui.Gui, v *gocui.View) error {
    if v == varray[0]{
        fmt.Println("cd " + folderlist[9])
    }else if v == varray[1] {
        fmt.Println(editor + " " + filelist[9])
    } else if v == varray[2] {
        fmt.Println("cd " + recentFolderList[9])
    }
    err := quit(g, v)
    if err != nil {
        return err
    }
    return nil
}

func goToNumberPos10(g *gocui.Gui, v *gocui.View) error {
    fmt.Println("cd " + recentFolderList[10])
    err := quit(g, v)
    if err != nil {
        return err
    }
    return nil
}

func goToNumberPos11(g *gocui.Gui, v *gocui.View) error {
    fmt.Println("cd " + recentFolderList[11])
    err := quit(g, v)
    if err != nil {
        return err
    }
    return nil
}

func goToNumberPos12(g *gocui.Gui, v *gocui.View) error {
    fmt.Println("cd " + recentFolderList[12])
    err := quit(g, v)
    if err != nil {
        return err
    }
    return nil
}

func goToNumberPos13(g *gocui.Gui, v *gocui.View) error {
    fmt.Println("cd " + recentFolderList[13])
    err := quit(g, v)
    if err != nil {
        return err
    }
    return nil
}

func goToNumberPos14(g *gocui.Gui, v *gocui.View) error {
    fmt.Println("cd " + recentFolderList[14])
    err := quit(g, v)
    if err != nil {
        return err
    }
    return nil
}

func goToNumberPos15(g *gocui.Gui, v *gocui.View) error {
    fmt.Println("cd " + recentFolderList[15])
    err := quit(g, v)
    if err != nil {
        return err
    }
    return nil
}

func goToNumberPos16(g *gocui.Gui, v *gocui.View) error {
    fmt.Println("cd " + recentFolderList[16])
    err := quit(g, v)
    if err != nil {
        return err
    }
    return nil
}

func goToNumberPos17(g *gocui.Gui, v *gocui.View) error {
    fmt.Println("cd " + recentFolderList[17])
    err := quit(g, v)
    if err != nil {
        return err
    }
    return nil
}

func goToNumberPos18(g *gocui.Gui, v *gocui.View) error {
    fmt.Println("cd " + recentFolderList[18])
    err := quit(g, v)
    if err != nil {
        return err
    }
    return nil
}

func goToNumberPos19(g *gocui.Gui, v *gocui.View) error {
    fmt.Println("cd " + recentFolderList[19])
    err := quit(g, v)
    if err != nil {
        return err
    }
    return nil
}

