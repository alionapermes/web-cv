package main

import (
	"bytes"
	"image/jpeg"
	"os"

	tcell "github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const (
  stateContacts = iota
  stateExp
  stateAbout
  stateEdu

  status = "21 y.o. Golang developer"
  email  = "alionapermes@gmail.com"
  linkTg = "t.me/alionapermes"
  linkGh = "github.com/alionapermes"

  aboutText = `  I'm interested in software architecture and studying best practices. Although I can't call my knowledge in this area sufficient for really big projects, I would like to develop in this direction`
  eduText   = `  2020-2025 Bachelor of Automation and Computer Engineering, Novosibirsk State Technical University, 3rd year student`
  expText   = `"Axiom", 2.5 years of backend development
  I'm responsible for the design and full technical support of one of the key projects (integration system with marketplaces (VK, Yandex, Ozon)) as well as several minor project including: internal applications for employees, integrations with mailing and chat services, automation of task distribution in CRM, etc
  Additionally, I sometimes develop new components and fix old ones on the frontend

Skills: PHP 7.4/8.2, MySQL, Git, JS (Vue3, JQuery)`
)

func buildExpBlock() *tview.TextView {
  textView := tview.NewTextView().
    SetDynamicColors(true).
    SetRegions(true).
    SetText(expText)

  textView.SetTitle("Experience").SetBorder(true)
  return textView
}

func buildAboutBlock() *tview.TextView {
  textView := tview.NewTextView().
    SetDynamicColors(true).
    SetRegions(true).
    SetText(aboutText)

  textView.SetTitle("About me").SetBorder(true)
  return textView
}

func buildEduBlock() *tview.TextView {
  textView := tview.NewTextView().
    SetDynamicColors(true).
    SetRegions(true).
    SetText(eduText)

  textView.SetTitle("Education").SetBorder(true)
  return textView
}

func buildPhotoBlock() *tview.Image {
  check := func(err error) {
    if err != nil {
      panic(err)
    }
  }

  data, err := os.ReadFile("assets/avatar.jpg")
  check(err)

  photo, err := jpeg.Decode(bytes.NewReader(data))
  check(err)

  return tview.NewImage().SetImage(photo).SetColors(tview.TrueColor) 
}

func buildContactsBlock() *tview.List {
  contactsBlock := tview.NewList().
    ShowSecondaryText(true).
    AddItem("Email:", email, 'e', func() {}).
    AddItem("Telegram:", linkTg, 't', func() {}).
    AddItem("GitHub:", linkGh, 'g', func() {})

  contactsBlock.SetTitle("Contacts").SetBorder(true)
  return contactsBlock
}

func main() {
  app := tview.NewApplication()

  contactsBlock := buildContactsBlock()
  photoBlock := buildPhotoBlock()
  aboutBlock := buildAboutBlock()
  eduBlock := buildEduBlock()
  expBlock := buildExpBlock()

  focusMap := []*tview.Box{contactsBlock.Box, expBlock.Box, aboutBlock.Box, eduBlock.Box}
  for i, box := range focusMap {
    (func(index int) {
      box.SetInputCapture(func(e *tcell.EventKey) *tcell.EventKey {
        mapLen := len(focusMap)

        var focusPos int
        switch key := e.Key(); key {
        case tcell.KeyTab: focusPos = (index + 1) % mapLen
        case tcell.KeyBacktab: focusPos = (index + mapLen - 1) % mapLen
        default: return e
        }
        app.SetFocus(focusMap[focusPos])
        return nil
      })
    })(i)

    _box := box
    _box.SetFocusFunc(func() {
      for _, b := range focusMap {
        b.SetBorderColor(tcell.ColorWhite)
      }
      _box.SetBorderColor(tcell.ColorGreen)
    })
  }

  flex := tview.NewFlex().SetDirection(tview.FlexColumn).
    AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
      AddItem(photoBlock, 0, 5, false).
      AddItem(contactsBlock, 0, 1, true), 0, 1, true).
    AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
      AddItem(expBlock, 0, 1, false).
      AddItem(aboutBlock, 0, 1, false).
      AddItem(eduBlock, 0, 1, false), 0, 1, false)

  if err := app.SetRoot(flex, true).Run(); err != nil {
    panic(err)
  }
}

