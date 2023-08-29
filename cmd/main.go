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

type uiComponent interface {
  tview.TextView | tview.Image | tview.List
}

var app *tview.Application

var navPos int
var navMap map[int]*tview.Box

func buildExpBlock() (*tview.TextView, *tview.Box) {
  textView := tview.NewTextView().
    SetDynamicColors(true).
    SetRegions(true).
    SetText(expText)

  textView.SetTitle("Experience").SetBorder(true)
  return textView, textView.Box
}

func buildAboutBlock() (*tview.TextView, *tview.Box) {
  textView := tview.NewTextView().
    SetDynamicColors(true).
    SetRegions(true).
    SetText(aboutText)

  textView.SetTitle("About me").SetBorder(true)
  return textView, textView.Box
}

func buildEduBlock() (*tview.TextView, *tview.Box) {
  textView := tview.NewTextView().
    SetDynamicColors(true).
    SetRegions(true).
    SetText(eduText)

  textView.SetTitle("Education").SetBorder(true)
  return textView, textView.Box
}

func buildPhotoBlock() (*tview.Image, *tview.Box) {
  check := func(err error) {
    if err != nil {
      panic(err)
    }
  }

  data, err := os.ReadFile("assets/avatar.jpg")
  check(err)

  photo, err := jpeg.Decode(bytes.NewReader(data))
  check(err)

  img := tview.NewImage().
    SetImage(photo).
    SetColors(tview.TrueColor)
  return img, img.Box
}

func buildContactsBlock() (*tview.List, *tview.Box) {
  contactsBlock := tview.NewList().
    ShowSecondaryText(true).
    AddItem("Email:", email, 'e', func() {}).
    AddItem("Telegram:", linkTg, 't', func() {}).
    AddItem("GitHub:", linkGh, 'g', func() {})

  contactsBlock.SetTitle("Contacts").SetBorder(true)
  return contactsBlock, contactsBlock.Box
}

func navWrap[C uiComponent](build func () (*C, *tview.Box)) (*C) {
  component, box := build()

  box.SetInputCapture(func(e *tcell.EventKey) *tcell.EventKey {
    mapLen := len(navMap)

    switch key := e.Key(); key {
    case tcell.KeyTab: navPos = (navPos + 1) % mapLen
    case tcell.KeyBacktab: navPos = (navPos + mapLen - 1) % mapLen
    default: return e
    }

    app.SetFocus(navMap[navPos])
    return nil
  })

  box.SetFocusFunc(func() {
    for _, b := range navMap {
      b.SetBorderColor(tcell.ColorWhite)
    }
    box.SetBorderColor(tcell.ColorGreen)
  })

  return component
}

func main() {
  navPos = 0
  navMap = make(map[int]*tview.Box)

  app = tview.NewApplication()

  contactsBlock := navWrap(buildContactsBlock)
  photoBlock := navWrap(buildPhotoBlock)
  aboutBlock := navWrap(buildAboutBlock)
  eduBlock := navWrap(buildEduBlock)
  expBlock := navWrap(buildExpBlock)

  navMap[len(navMap)] = contactsBlock.Box
  navMap[len(navMap)] = expBlock.Box
  navMap[len(navMap)] = aboutBlock.Box
  navMap[len(navMap)] = eduBlock.Box

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

