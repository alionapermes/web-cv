package main

import (
	"bytes"
	"image/jpeg"
	"os"

	tcell "github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const (
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

var app *tview.Application

var navPos int
var navMap map[int]*tview.Box

func buildTextBlock(title, text string) (tview.Primitive, *tview.Box) {
  block := tview.NewTextView().
    SetDynamicColors(true).
    SetRegions(true).
    SetText(text)

  block.SetTitle(title).SetBorder(true)
  return block, block.Box
}

func buildExpBlock() (tview.Primitive, *tview.Box) {
  return buildTextBlock("Experience", expText)
}

func buildAboutBlock() (tview.Primitive, *tview.Box) {
  return buildTextBlock("About me", aboutText)
}

func buildEduBlock() (tview.Primitive, *tview.Box) {
  return buildTextBlock("Education", eduText)
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

  img := tview.NewImage().
    SetImage(photo).
    SetColors(tview.TrueColor)
  return img
}

func buildContactsBlock() (tview.Primitive, *tview.Box) {
  contactsBlock := tview.NewList().
    ShowSecondaryText(true).
    AddItem("Email:", email, 'e', func() {}).
    AddItem("Telegram:", linkTg, 't', func() {}).
    AddItem("GitHub:", linkGh, 'g', func() {})

  contactsBlock.SetTitle("Contacts").SetBorder(true)
  return contactsBlock, contactsBlock.Box
}

func navWrap(build func () (tview.Primitive, *tview.Box)) (tview.Primitive, *tview.Box) {
  primitive, box := build()

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

  return primitive, box
}

func main() {
  navPos = 0
  navMap = make(map[int]*tview.Box)

  app = tview.NewApplication()

  photoBlock := buildPhotoBlock()

  contactsBlock, box := navWrap(buildContactsBlock)
  navMap[len(navMap)] = box

  expBlock, box := navWrap(buildExpBlock)
  navMap[len(navMap)] = box

  aboutBlock, box := navWrap(buildAboutBlock)
  navMap[len(navMap)] = box

  eduBlock, box := navWrap(buildEduBlock)
  navMap[len(navMap)] = box


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

