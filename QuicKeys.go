package main

import (
    "log"
    "github.com/demonlord1997/gocui"
)

func quicKeys(g *gocui.Gui){
    // exit inputbox
    for _,c := range []string{"input", "RecentFolder"} {
        if err := g.SetKeybinding(c, gocui.KeyEsc, gocui.ModNone, exitCurView); err != nil {
            log.Panicln(err)
        }
    }

    for _,c := range []string{"v1", "v2"} {
        if err := g.SetKeybinding(c, gocui.KeyEsc, gocui.ModNone, quit); err != nil {
            log.Panicln(err)
        }
        if err := g.SetKeybinding(c, 'l', gocui.ModNone, nextView); err != nil {
            log.Panicln(err)
        }
        if err := g.SetKeybinding(c, gocui.KeyArrowRight, gocui.ModNone, nextView); err != nil {
            log.Panicln(err)
        }
        if err := g.SetKeybinding(c, 'h', gocui.ModNone, preView); err != nil {
            log.Panicln(err)
        }
        if err := g.SetKeybinding(c, gocui.KeyArrowLeft, gocui.ModNone, preView); err != nil {
            log.Panicln(err)
        }
        if err := g.SetKeybinding(c, 'o', gocui.ModNone, openInputBox); err != nil {
            log.Panicln(err)
        }
        if err := g.SetKeybinding(c, 'r', gocui.ModNone, openRecentFolder); err != nil {
            log.Panicln(err)
        }
        if err := g.SetKeybinding(c, 'q', gocui.ModNone, quit); err != nil {
            log.Panicln(err)
        }
        if err := g.SetKeybinding(c, gocui.KeyTab, gocui.ModNone, nextView); err != nil {
            log.Panicln(err)
        }
    }

    for _,c := range []string{"v1", "v2", "RecentFolder"} {
        if err := g.SetKeybinding(c, 'j', gocui.ModNone, cursorDown); err != nil {
            log.Panicln(err)
        }
        if err := g.SetKeybinding(c, 'k', gocui.ModNone, cursorUp); err != nil {
            log.Panicln(err)
        }
        if err := g.SetKeybinding(c, 'Q', gocui.ModNone, quit); err != nil {
            log.Panicln(err)
        }
        if err := g.SetKeybinding(c, gocui.KeyArrowDown, gocui.ModNone, cursorDown); err != nil {
            log.Panicln(err)
        }
        if err := g.SetKeybinding(c, gocui.KeyArrowUp, gocui.ModNone, cursorUp); err != nil {
            log.Panicln(err)
        }
        if err := g.SetKeybinding(c, 'g', gocui.ModNone, cursorTop); err != nil {
            log.Panicln(err)
        }
        if err := g.SetKeybinding(c, 'G', gocui.ModNone, cursorBottom); err != nil {
            log.Panicln(err)
        }
        if err := g.SetKeybinding(c, '0', gocui.ModNone, goToNumberPos0); err != nil {
            log.Panicln(err)
        }
        if err := g.SetKeybinding(c, '1', gocui.ModNone, goToNumberPos1); err != nil {
            log.Panicln(err)
        }
        if err := g.SetKeybinding(c, '2', gocui.ModNone, goToNumberPos2); err != nil {
            log.Panicln(err)
        }
        if err := g.SetKeybinding(c, '3', gocui.ModNone, goToNumberPos3); err != nil {
            log.Panicln(err)
        }
        if err := g.SetKeybinding(c, '4', gocui.ModNone, goToNumberPos4); err != nil {
            log.Panicln(err)
        }
        if err := g.SetKeybinding(c, '5', gocui.ModNone, goToNumberPos5); err != nil {
            log.Panicln(err)
        }
        if err := g.SetKeybinding(c, '6', gocui.ModNone, goToNumberPos6); err != nil {
            log.Panicln(err)
        }
        if err := g.SetKeybinding(c, '7', gocui.ModNone, goToNumberPos7); err != nil {
            log.Panicln(err)
        }
        if err := g.SetKeybinding(c, '8', gocui.ModNone, goToNumberPos8); err != nil {
            log.Panicln(err)
        }
        if err := g.SetKeybinding(c, '9', gocui.ModNone, goToNumberPos9); err != nil {
            log.Panicln(err)
        }
    }
    if err := g.SetKeybinding("RecentFolder", 'q', gocui.ModNone, goToNumberPos10); err != nil {
        log.Panicln(err)
    }
    if err := g.SetKeybinding("RecentFolder", 'w', gocui.ModNone, goToNumberPos11); err != nil {
        log.Panicln(err)
    }
    if err := g.SetKeybinding("RecentFolder", 'e', gocui.ModNone, goToNumberPos12); err != nil {
        log.Panicln(err)
    }
    if err := g.SetKeybinding("RecentFolder", 'r', gocui.ModNone, goToNumberPos13); err != nil {
        log.Panicln(err)
    }
    if err := g.SetKeybinding("RecentFolder", 't', gocui.ModNone, goToNumberPos14); err != nil {
        log.Panicln(err)
    }
    if err := g.SetKeybinding("RecentFolder", 'y', gocui.ModNone, goToNumberPos15); err != nil {
        log.Panicln(err)
    }
    if err := g.SetKeybinding("RecentFolder", 'u', gocui.ModNone, goToNumberPos16); err != nil {
        log.Panicln(err)
    }
    if err := g.SetKeybinding("RecentFolder", 'i', gocui.ModNone, goToNumberPos17); err != nil {
        log.Panicln(err)
    }
    if err := g.SetKeybinding("RecentFolder", 'o', gocui.ModNone, goToNumberPos18); err != nil {
        log.Panicln(err)
    }
    if err := g.SetKeybinding("RecentFolder", 'p', gocui.ModNone, goToNumberPos19); err != nil {
        log.Panicln(err)
    }

    // run shell command
    if err := g.SetKeybinding("v1", gocui.KeyEnter, gocui.ModNone, runFolderShell); err != nil {
        log.Panicln(err)
    }
    if err := g.SetKeybinding("v2", gocui.KeyEnter, gocui.ModNone, runFileShell); err != nil {
        log.Panicln(err)
    }
    if err := g.SetKeybinding("input", gocui.KeyEnter, gocui.ModNone, runShell); err != nil {
        log.Panicln(err)
    }
    if err := g.SetKeybinding("RecentFolder", gocui.KeyEnter, gocui.ModNone, runRecentFolderShell); err != nil {
        log.Panicln(err)
    }


    if err := g.SetKeybinding("", gocui.KeyCtrlQ, gocui.ModNone, chooseV1); err != nil {
        log.Panicln(err)
    }
    if err := g.SetKeybinding("", gocui.KeyCtrlW, gocui.ModNone, chooseV2); err != nil {
        log.Panicln(err)
    }

    if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
        log.Panicln(err)
    }
}
