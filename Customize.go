package main

import (
    "fmt"
    "log"
    "github.com/demonlord1997/gocui"
    "gopkg.in/ini.v1"
)

func Customize(cfg  *ini.File) []*ini.Key {
    customQuickey := cfg.Section("customize").Keys()
    return  customQuickey
}

func SetCustomKey(cfg *ini.File, g *gocui.Gui)  {
    customQuickey := Customize(cfg)
    var convert_u_key interface{}
    for _, u_key := range customQuickey {
        if len(u_key.Value())==1 {
            convert_u_key = []rune(u_key.Value())[0]
        }else {
            switch u_key.Value() {
            case "<C-a>":
                convert_u_key = gocui.KeyCtrlA
            case "<C-b>":
                convert_u_key = gocui.KeyCtrlB
            case "<C-c>":
                convert_u_key = gocui.KeyCtrlC
            case "<C-d>":
                convert_u_key = gocui.KeyCtrlD
            case "<C-e>":
                convert_u_key = gocui.KeyCtrlE
            case "<C-f>":
                convert_u_key = gocui.KeyCtrlF
            case "<C-g>":
                convert_u_key = gocui.KeyCtrlG
            case "<C-h>":
                convert_u_key = gocui.KeyCtrlH
            case "<C-i>":
                convert_u_key = gocui.KeyCtrlI
            case "<C-j>":
                convert_u_key = gocui.KeyCtrlJ
            case "<C-k>":
                convert_u_key = gocui.KeyCtrlK
            case "<C-l>":
                convert_u_key = gocui.KeyCtrlL
            case "<C-m>":
                convert_u_key = gocui.KeyCtrlM
            case "<C-n>":
                convert_u_key = gocui.KeyCtrlN
            case "<C-o>":
                convert_u_key = gocui.KeyCtrlO
            case "<C-p>":
                convert_u_key = gocui.KeyCtrlP
            case "<C-q>":
                convert_u_key = gocui.KeyCtrlQ
            case "<C-r>":
                convert_u_key = gocui.KeyCtrlR
            case "<C-s>":
                convert_u_key = gocui.KeyCtrlS
            case "<C-t>":
                convert_u_key = gocui.KeyCtrlT
            case "<C-u>":
                convert_u_key = gocui.KeyCtrlU
            case "<C-v>":
                convert_u_key = gocui.KeyCtrlV
            case "<C-w>":
                convert_u_key = gocui.KeyCtrlW
            case "<C-x>":
                convert_u_key = gocui.KeyCtrlX
            case "<C-y>":
                convert_u_key = gocui.KeyCtrlY
            case "<C-z>":
                convert_u_key = gocui.KeyCtrlZ
            case "<f1>":
                convert_u_key = gocui.KeyF1
            case "<f2>":
                convert_u_key = gocui.KeyF2
            case "<f3>":
                convert_u_key = gocui.KeyF3
            case "<f4>":
                convert_u_key = gocui.KeyF4
            case "<f5>":
                convert_u_key = gocui.KeyF5
            case "<f6>":
                convert_u_key = gocui.KeyF6
            case "<f7>":
                convert_u_key = gocui.KeyF7
            case "<f8>":
                convert_u_key = gocui.KeyF8
            case "<f9>":
                convert_u_key = gocui.KeyF9
            case "<f10>":
                convert_u_key = gocui.KeyF10
            case "<f11>":
                convert_u_key = gocui.KeyF11
            case "<f12>":
                convert_u_key = gocui.KeyF12
            }
        }
        mapkeys(g, convert_u_key, u_key.Name())
    }
}

func mapkeys(g *gocui.Gui, convert_u_key interface{}, u_command string) {
    for _,k := range []string{"v1", "v2"} {
        if err := g.SetKeybinding(k, convert_u_key, gocui.ModNone,
        func (g *gocui.Gui, v *gocui.View) error {
            fmt.Println(u_command)
            err := quit(g, v)
            if err != nil {
                return err
            }
            return nil
        }); err != nil {
            log.Panicln(err)
        }
    }
}

