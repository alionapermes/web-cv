package component

import "github.com/rivo/tview"

const (
  tPadding = 1
  bPadding = 1
  lPadding = 2
  rPadding = 2
)

type InfoBlock struct {
  *tview.TextView
}

func NewInfoBlock(title, text string) *InfoBlock {
  textView := tview.NewTextView().
    SetDynamicColors(true).
    SetRegions(true).
    SetText(text)

  textView.
    SetTitle(title).
    SetBorder(true).
    SetBorderPadding(tPadding, bPadding, lPadding, rPadding)
  return &InfoBlock{TextView: textView}
}

