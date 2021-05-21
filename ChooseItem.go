package main

import (
    "fmt"
    "github.com/demonlord1997/gocui"
)


func goToNumberPos0(g *gocui.Gui, v *gocui.View) error {
    if v == g_varray[0]{
        fmt.Println("cd " + g_folderlist[0])
    } else if v == g_varray[1] {
        fmt.Println(g_editor + " " + g_filelist[0])
    } else if v == g_varray[2] {
        fmt.Println("cd " + g_recentFolderList[0])
    }
    err := quit(g, v)
    if err != nil {
        return err
    }
    return nil
}

func goToNumberPos1(g *gocui.Gui, v *gocui.View) error {
    if v == g_varray[0]{
        fmt.Println("cd " + g_folderlist[1])
    }else if v == g_varray[1] {
        fmt.Println(g_editor + " " + g_filelist[1])
    } else if v == g_varray[2] {
        fmt.Println("cd " + g_recentFolderList[1])
    }
    err := quit(g, v)
    if err != nil {
        return err
    }
    return nil
}

func goToNumberPos2(g *gocui.Gui, v *gocui.View) error {
    if v == g_varray[0]{
        fmt.Println("cd " + g_folderlist[2])
    }else if v == g_varray[1] {
        fmt.Println(g_editor + " " + g_filelist[2])
    } else if v == g_varray[2] {
        fmt.Println("cd " + g_recentFolderList[2])
    }
    err := quit(g, v)
    if err != nil {
        return err
    }
    return nil
}

func goToNumberPos3(g *gocui.Gui, v *gocui.View) error {
    if v == g_varray[0]{
        fmt.Println("cd " + g_folderlist[3])
    }else if v == g_varray[1] {
        fmt.Println(g_editor + " " + g_filelist[3])
    } else if v == g_varray[2] {
        fmt.Println("cd " + g_recentFolderList[3])
    }
    err := quit(g, v)
    if err != nil {
        return err
    }
    return nil
}

func goToNumberPos4(g *gocui.Gui, v *gocui.View) error {
    if v == g_varray[0]{
        fmt.Println("cd " + g_folderlist[4])
    }else if v == g_varray[1] {
        fmt.Println(g_editor + " " + g_filelist[4])
    } else if v == g_varray[2] {
        fmt.Println("cd " + g_recentFolderList[4])
    }
    err := quit(g, v)
    if err != nil {
        return err
    }
    return nil
}

func goToNumberPos5(g *gocui.Gui, v *gocui.View) error {
    if v == g_varray[0]{
        fmt.Println("cd " + g_folderlist[5])
    }else if v == g_varray[1] {
        fmt.Println(g_editor + " " + g_filelist[5])
    } else if v == g_varray[2] {
        fmt.Println("cd " + g_recentFolderList[5])
    }
    err := quit(g, v)
    if err != nil {
        return err
    }
    return nil
}

func goToNumberPos6(g *gocui.Gui, v *gocui.View) error {
    if v == g_varray[0]{
        fmt.Println("cd " + g_folderlist[6])
    }else if v == g_varray[1] {
        fmt.Println(g_editor + " " + g_filelist[6])
    } else if v == g_varray[2] {
        fmt.Println("cd " + g_recentFolderList[6])
    }
    err := quit(g, v)
    if err != nil {
        return err
    }
    return nil
}

func goToNumberPos7(g *gocui.Gui, v *gocui.View) error {
    if v == g_varray[0]{
        fmt.Println("cd " + g_folderlist[7])
    }else if v == g_varray[1] {
        fmt.Println(g_editor + " " + g_filelist[7])
    } else if v == g_varray[2] {
        fmt.Println("cd " + g_recentFolderList[7])
    }
    err := quit(g, v)
    if err != nil {
        return err
    }
    return nil
}

func goToNumberPos8(g *gocui.Gui, v *gocui.View) error {
    if v == g_varray[0]{
        fmt.Println("cd " + g_folderlist[8])
    }else if v == g_varray[1] {
        fmt.Println(g_editor + " " + g_filelist[8])
    } else if v == g_varray[2] {
        fmt.Println("cd " + g_recentFolderList[8])
    }
    err := quit(g, v)
    if err != nil {
        return err
    }
    return nil
}

func goToNumberPos9(g *gocui.Gui, v *gocui.View) error {
    if v == g_varray[0]{
        fmt.Println("cd " + g_folderlist[9])
    }else if v == g_varray[1] {
        fmt.Println(g_editor + " " + g_filelist[9])
    } else if v == g_varray[2] {
        fmt.Println("cd " + g_recentFolderList[9])
    }
    err := quit(g, v)
    if err != nil {
        return err
    }
    return nil
}

func goToNumberPos10(g *gocui.Gui, v *gocui.View) error {
    fmt.Println("cd " + g_recentFolderList[10])
    err := quit(g, v)
    if err != nil {
        return err
    }
    return nil
}

func goToNumberPos11(g *gocui.Gui, v *gocui.View) error {
    fmt.Println("cd " + g_recentFolderList[11])
    err := quit(g, v)
    if err != nil {
        return err
    }
    return nil
}

func goToNumberPos12(g *gocui.Gui, v *gocui.View) error {
    fmt.Println("cd " + g_recentFolderList[12])
    err := quit(g, v)
    if err != nil {
        return err
    }
    return nil
}

func goToNumberPos13(g *gocui.Gui, v *gocui.View) error {
    fmt.Println("cd " + g_recentFolderList[13])
    err := quit(g, v)
    if err != nil {
        return err
    }
    return nil
}

func goToNumberPos14(g *gocui.Gui, v *gocui.View) error {
    fmt.Println("cd " + g_recentFolderList[14])
    err := quit(g, v)
    if err != nil {
        return err
    }
    return nil
}

func goToNumberPos15(g *gocui.Gui, v *gocui.View) error {
    fmt.Println("cd " + g_recentFolderList[15])
    err := quit(g, v)
    if err != nil {
        return err
    }
    return nil
}

func goToNumberPos16(g *gocui.Gui, v *gocui.View) error {
    fmt.Println("cd " + g_recentFolderList[16])
    err := quit(g, v)
    if err != nil {
        return err
    }
    return nil
}

func goToNumberPos17(g *gocui.Gui, v *gocui.View) error {
    fmt.Println("cd " + g_recentFolderList[17])
    err := quit(g, v)
    if err != nil {
        return err
    }
    return nil
}

func goToNumberPos18(g *gocui.Gui, v *gocui.View) error {
    fmt.Println("cd " + g_recentFolderList[18])
    err := quit(g, v)
    if err != nil {
        return err
    }
    return nil
}

func goToNumberPos19(g *gocui.Gui, v *gocui.View) error {
    fmt.Println("cd " + g_recentFolderList[19])
    err := quit(g, v)
    if err != nil {
        return err
    }
    return nil
}

