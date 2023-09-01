package component

import (
	tcell "github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type DrawFunc = func(
  tcell.Screen, int, int, int, int) (int, int, int, int)

type InputCaptureHandler = func(*tcell.EventKey) *tcell.EventKey

type MouseCaptureHandler = func(
  tview.MouseAction,
  *tcell.EventMouse,
) (tview.MouseAction, *tcell.EventMouse)

type WrappedInputCaptureHandler = func(
  *tcell.EventKey,
  func(tview.Primitive))

type WrappedMouseCaptureHandler = func(
  tview.MouseAction,
  *tcell.EventMouse,
  func(tview.Primitive),
) (bool, tview.Primitive)

type Component interface {
  tview.Primitive

  DrawForSubclass(tcell.Screen, tview.Primitive)
  GetBackgroundColor() tcell.Color
  GetBorderAttributes() tcell.AttrMask
  GetBorderColor() tcell.Color
  GetDrawFunc() DrawFunc
  GetInnerRect() (int, int, int, int)
  GetInputCapture() InputCaptureHandler
  GetMouseCapture() MouseCaptureHandler
  GetTitle() string
  InRect(int, int) bool
  SetBackgroundColor(tcell.Color) *tview.Box
  SetBlurFunc(func()) *tview.Box
  SetBorder(bool) *tview.Box
  SetBorderAttributes(tcell.AttrMask) *tview.Box
  SetBorderColor(tcell.Color) *tview.Box
  SetBorderPadding(int, int, int, int) *tview.Box
  SetBorderStyle(tcell.Style) *tview.Box
  SetDrawFunc(DrawFunc) *tview.Box
  SetFocusFunc(func()) *tview.Box
  SetInputCapture(InputCaptureHandler) *tview.Box
  SetMouseCapture(MouseCaptureHandler) *tview.Box
  SetTitle(string) *tview.Box
  SetTitleAlign(int) *tview.Box
  SetTitleColor(tcell.Color) *tview.Box
  WrapInputHandler(WrappedInputCaptureHandler) WrappedInputCaptureHandler
  WrapMouseHandler(WrappedMouseCaptureHandler) WrappedMouseCaptureHandler
}

