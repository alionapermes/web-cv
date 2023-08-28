package main

import (
  "os"
  "bytes"
  "image/jpeg"

  "github.com/rivo/tview"
)

const (
  email  = "alionapermes@gmail.com"
  linkTg = "t.me/alionapermes"
  linkGh = "github.com/alionapermes"

  cvText = `21 y.o. Golang developer with 2.5 years of backend experience`
)

func buildTextBlock(app *tview.Application) *tview.TextView {
  textView := tview.NewTextView().
    SetDynamicColors(true).
    SetRegions(true).
    SetChangedFunc(func() { app.Draw() }).
    SetText(cvText)

  textView.SetTitle("About me").SetBorder(true)
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

  return tview.NewImage().SetImage(photo) 
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

  photoBlock := buildPhotoBlock()
  contactsBlock := buildContactsBlock()
  textBlock := buildTextBlock(app)

  flex := tview.NewFlex().SetDirection(tview.FlexColumn).
    AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
      AddItem(photoBlock, 0, 4, false).
      AddItem(contactsBlock, 0, 1, true), 0, 1, true).
    AddItem(textBlock, 0, 2, false)

  if err := app.SetRoot(flex, true).Run(); err != nil {
    panic(err)
  }
}

