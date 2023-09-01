package ui

import (
	"bytes"
	"image/jpeg"
	"os"

	tcell "github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"

  "github.com/alionapermes/web-cv/config"
  . "github.com/alionapermes/web-cv/ui/component"
)

type UI struct {
  app *tview.Application
  cfg *config.Config

  navPos int
  navMap map[int]Component

  contacts Component
  about Component
  education Component
  experience Component
}

func New(config *config.Config) *UI {
  ui := UI{
    app: tview.NewApplication(),
    cfg: config,
    navPos: 0,
    navMap: make(map[int]Component),
  }

  return ui.build()
}

func (ui *UI) Start() error {
  if err := ui.app.Run(); err != nil {
    ui.app.Stop()
    return err
  }

  return nil
}

func (ui *UI) navAdd(build func () Component) Component {
  component := build()
  ui.navMap[len(ui.navMap)] = component

  component.SetInputCapture(func(e *tcell.EventKey) *tcell.EventKey {
    mapLen := len(ui.navMap)

    switch key := e.Key(); key {
    case tcell.KeyTab: ui.navPos = (ui.navPos + 1) % mapLen
    case tcell.KeyBacktab: ui.navPos = (ui.navPos + mapLen - 1) % mapLen
    default: return e
    }

    ui.app.SetFocus(ui.navMap[ui.navPos])
    return nil
  })

  component.SetFocusFunc(func() {
    for _, b := range ui.navMap {
      b.SetBorderColor(tcell.ColorWhite)
    }
    component.SetBorderColor(tcell.ColorGreen)
  })

  return component
}

func (ui *UI) buildExpBlock() Component {
  return NewInfoBlock("Experience", ui.cfg.Texts.Experience)
}

func (ui *UI) buildAboutBlock() Component {
  return NewInfoBlock("About me", ui.cfg.Texts.About)
}

func (ui* UI) buildEduBlock() Component {
  return NewInfoBlock("Education", ui.cfg.Texts.Education)
}

func (ui *UI) buildPhotoBlock() Component {
  check := func(err error) {
    if err != nil {
      panic(err)
    }
  }

  data, err := os.ReadFile(ui.cfg.PhotoPath)
  check(err)

  photo, err := jpeg.Decode(bytes.NewReader(data))
  check(err)

  img := tview.NewImage().
    SetImage(photo).
    SetColors(tview.TrueColor)
  return img
}

func (ui *UI) buildContactsBlock() Component {
  contactsBlock := tview.NewList().
    ShowSecondaryText(true).
    AddItem("Email:", ui.cfg.Contacts.Email, 'e', func() {}).
    AddItem("Telegram:", ui.cfg.Contacts.Telegram, 't', func() {}).
    AddItem("GitHub:", ui.cfg.Contacts.GitHub, 'g', func() {})

  contactsBlock.SetTitle("Contacts").SetBorder(true)
  return contactsBlock
}

func (ui *UI) build() *UI {
  photoBlock := ui.buildPhotoBlock()

  contactsBlock := ui.navAdd(ui.buildContactsBlock)
  expBlock := ui.navAdd(ui.buildExpBlock)
  aboutBlock := ui.navAdd(ui.buildAboutBlock)
  eduBlock := ui.navAdd(ui.buildEduBlock)

  container := tview.NewFlex().SetDirection(tview.FlexColumn).
    AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
      AddItem(photoBlock, 0, 5, false).
      AddItem(contactsBlock, 0, 1, true), 0, 1, true).
    AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
      AddItem(expBlock, 0, 1, false).
      AddItem(aboutBlock, 0, 1, false).
      AddItem(eduBlock, 0, 1, false), 0, 1, false)
  
  ui.app.SetRoot(container, true)
  return ui
}

