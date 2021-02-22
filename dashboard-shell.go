package main

import (
    "fmt"
    "log"
    "strconv"
    "errors"
    "runtime"
    "os/user"
    "os"
    "strings"
    "os/exec"
    "bytes"
    "gopkg.in/ini.v1"
    "github.com/jroimartin/gocui"
)

var (
    curInputView = 2
    viewArr = []string{"v1", "v2"}
    active  = 0
    selectView = 0
    varray = []*gocui.View{}
    vinput *gocui.View
    itemNumber = [10]int{0,1,2,3,4,5,6,7,8,9}
    folderlist []string
    filelist []string
    HOME string
    editor string
    folderLength = 22
    itermHeight = 11
    //selectNum = 0
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

func exitInputBox(g *gocui.Gui, v *gocui.View) error{
    return delView(g)
}


func goToNumberPos0(g *gocui.Gui, v *gocui.View) error {
    if v == varray[0]{
        fmt.Println("cd " + folderlist[0])
    } else if v == varray[1] {
        fmt.Println(editor + " " + filelist[0])
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
    }
    err := quit(g, v)
    if err != nil {
        return err
    }
    return nil
}


func runFolderShell(g *gocui.Gui, v *gocui.View) error {
    var err error
    _, cy := v.Cursor()

    fmt.Println("cd " + folderlist[cy])
    err = quit(g, v)
    if err != nil {
        return err
    }
    return nil
}

func runFileShell(g *gocui.Gui, v *gocui.View) error {
    var err error
    _, cy := v.Cursor()

    fmt.Println(editor+ " " + filelist[cy])
    err = quit(g, v)
    if err != nil {
        return err
    }
    return nil
}


func runShell(g *gocui.Gui, v *gocui.View) error {
    var l string
    var err error

    _, cy := v.Cursor()
    if l, err = v.Line(cy); err != nil {
        l = ""
    }

    fmt.Println(l)
    err = quit(g, v)
    if err != nil {
        return err
    }
    return nil
}


func main() {
    HOME, _ = Home()
    cfg, err := ini.Load(HOME+"/.config/dashboard-shell/config.ini")
    if err != nil {
        log.Fatal("Fail to read file: ", err)
    }
    editor = cfg.Section("settings").Key("editor").String()
    g, err := gocui.NewGui(gocui.OutputNormal)
    g.InputEsc = true
    if err != nil{
        log.Panicln(err)
    }

    defer g.Close()
    g.SetManagerFunc(layout)

    // exit TUI
    if err := g.SetKeybinding("", 'q', gocui.ModNone, quit); err != nil {
        log.Panicln(err)
    }
    if err := g.SetKeybinding("", 'Q', gocui.ModNone, quit); err != nil {
        log.Panicln(err)
    }
    // exit inputbox
    if err := g.SetKeybinding("input", gocui.KeyEsc, gocui.ModNone, exitInputBox); err != nil {
        log.Panicln(err)
    }

    for _,c := range []string{"v1", "v2"} {
        if err := g.SetKeybinding(c, gocui.KeyEsc, gocui.ModNone, quit); err != nil {
            log.Panicln(err)
        }
        if err := g.SetKeybinding(c, 'j', gocui.ModNone, cursorDown); err != nil {
            log.Panicln(err)
        }
        if err := g.SetKeybinding(c, 'k', gocui.ModNone, cursorUp); err != nil {
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
        if err := g.SetKeybinding(c, 'l', gocui.ModNone, nextView); err != nil {
            log.Panicln(err)
        }
        if err := g.SetKeybinding(c, 'h', gocui.ModNone, preView); err != nil {
            log.Panicln(err)
        }
        if err := g.SetKeybinding(c, 'o', gocui.ModNone, openInputBox); err != nil {
            log.Panicln(err)
        }
        if err := g.SetKeybinding(c, gocui.KeyTab, gocui.ModNone, nextView); err != nil {
            log.Panicln(err)
        }
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


func quit(g *gocui.Gui, v *gocui.View) error {
    return gocui.ErrQuit
}


func layout(g *gocui.Gui) error {
    maxX, maxY := g.Size()
    var beginBoxPosX int
    var beginBoxPosY int
    var iShowV3 = (maxY/2 > 8) && (maxX > 73)
    beginBoxPosX = maxX/2 - 3*folderLength/2 - 1
    if iShowV3 {
        beginBoxPosY = maxY/2
    }else{
        beginBoxPosY = (maxY-itermHeight)/2
    }

    if v, err := g.SetView("v1", beginBoxPosX, beginBoxPosY, beginBoxPosX+folderLength, beginBoxPosY+itermHeight); err != nil {
        if err != gocui.ErrUnknownView {
            return err
        }
        varray = append(varray, v)
        //fmt.Println(active)
        v.Highlight = true
        if iShowV3{
            v.Frame = false
        }else{
            v.Frame = true
        }
        v.SelFgColor = gocui.ColorRed
        cfg, err := ini.Load(HOME+"/.config/dashboard-shell/folders.ini")
        if err != nil {
            log.Fatal("Fail to open folder: ", err)
        }
        folderStrings := cfg.Section("folders").Key("path").String()
        folderlist = strings.Split(folderStrings, ",")

        for i := range itemNumber {
            folderName := strings.Split(folderlist[i], "/")
            fmt.Fprintln(v,"["+ strconv.Itoa(i) +"]" + "\t "+folderName[len(folderName)-1])
        }
        //v.Title = "FOLDERS"
    }
    if v, err := g.SetView("v2", beginBoxPosX+folderLength+2, beginBoxPosY, maxX-beginBoxPosX, beginBoxPosY+itermHeight); err != nil {
        if err != gocui.ErrUnknownView {
            return err
        }
        varray = append(varray, v)
        if iShowV3{
            v.Frame = false
        }else{
            v.Frame = true
        }

        //v.Highlight = false
        v.SelFgColor = gocui.ColorRed
        cfg, err := ini.Load(HOME+"/.config/dashboard-shell/files.ini")
        if err != nil {
            log.Fatal("Fail to read file: ", err)
        }
        fileStrings := cfg.Section("files").Key("path").String()
        filelist = strings.Split(fileStrings, ",")

        for i := range itemNumber {
            //showString := "["+ strconv.itoa(i) +"]" + "\t "+ fileName
            showString := "["+ strconv.Itoa(i) +"]" + "\t "+ filelist[i]
            ssLength := len(showString)
            if ssLength >= 2*folderLength{
                fileName := strings.Split(filelist[i], "/")
                fnLength := len(fileName)
                showShortString := showString[0:(2*folderLength-len(fileName[fnLength-1])-5)] + ".../" + fileName[fnLength-1]
                fmt.Fprintln(v, showShortString)
            }else{
                fmt.Fprintln(v, showString)
            }
        }
        //v.Title = "FILES"
    }

    if iShowV3{
        if v, err := g.SetView("v3", (maxX-72)/2, maxY/2-2-8, (maxX+73)/2, maxY/2); err != nil {
            if err != gocui.ErrUnknownView{
                return err
            }
            v.Frame = false
            cmd := exec.Command("/bin/bash", "-c", "echo $(date '+%Y.%m.%d') | figlet -f banner")
            var out bytes.Buffer
            cmd.Stdout = &out
            err := cmd.Run()
            if err != nil {
                log.Fatal(err)
            }
            fmt.Fprintln(v,out.String())
        }
    }

    // 使item可选
    if selectView == 0{
        if _, err := setCurrentViewOnTop(g, "v1"); err != nil {
            log.Panicln(err)
        }
        selectView = 1
    }
    return nil
}

// ==============GET $HOME ==================
func Home() (string, error) {
    user, err := user.Current()
    if nil == err {
        return user.HomeDir, nil
    }

    // cross compile support
    if "windows" == runtime.GOOS {
        return homeWindows()
    }

    // Unix-like system, so just assume Unix
    return homeUnix()
}

func homeUnix() (string, error) {
    // First prefer the HOME environmental variable
    if home := os.Getenv("HOME"); home != "" {
        return home, nil
    }

    // If that fails, try the shell
    var stdout bytes.Buffer
    cmd := exec.Command("sh", "-c", "eval echo ~$USER")
    cmd.Stdout = &stdout
    if err := cmd.Run(); err != nil {
        return "", err
    }

    result := strings.TrimSpace(stdout.String())
    if result == "" {
        return "", errors.New("blank output when reading home directory")
    }

    return result, nil
}

func homeWindows() (string, error) {
    drive := os.Getenv("HOMEDRIVE")
    path := os.Getenv("HOMEPATH")
    home := drive + path
    if drive == "" || path == "" {
        home = os.Getenv("USERPROFILE")
    }
    if home == "" {
        return "", errors.New("HOMEDRIVE, HOMEPATH, and USERPROFILE are blank")
    }

    return home, nil
}


// -----------print image shape -----------------------------------
            //hTopVal := "─"
            //htop := strings.Repeat(hTopVal, maxX/2-beginBoxPosX-2)
            //fillVal := " "
            //fmt.Fprintln(v, "┌"+htop+"┐")
            //for i := 1; i <= (maxY/2-2-1)/2; i++  {
                //fillSpace1_1 := strings.Repeat(fillVal, i+1)
                //fillSpace1_2 := strings.Repeat(fillVal, maxX/2-beginBoxPosX-2-i-1-1)
                //fmt.Fprintln(v, "│"+fillSpace1_1+"-"+fillSpace1_2+"│")
            //}
            //fillSpace1_1 := 

            //img := `
   //$i{)|\]l"      ',>[1{}!.
 //$;?(t-+~>I^       ,--_/Qm].
//.!>r_    :1u0J|xJJr;   ']Jn,
//$l)v"  IubokMbOYQpWd\, ./L/"
//.<)Jn1zY{ll}< $i]<iI?uuCZn~.
 //$i1uXj   I"+v+. -zI :x{~I.
   //$,!.  $! $I$.. <?"',!'
    //$l::l,'.  ,->I"::>":'
    //;l>^I:  $",$!:'"<;
      //^' II. 'I,. ,}[.
      //'I  II.  ;_}+~"
       //':^ ";!' >{!
         //$. ,I~I]_I.
            //' :>-(-'
        //`

            //fmt.Fprintln(v, img)
