package component

import "github.com/rivo/tview"

type InfoBlock struct {
  *tview.TextView
}

func NewInfoBlock(title, text string) *InfoBlock {
  textView := tview.NewTextView().
    SetDynamicColors(true).
    SetRegions(true).
    SetText(text)

  textView.SetTitle(title).SetBorder(true)
  return &InfoBlock{TextView: textView}
}

