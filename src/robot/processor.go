package robot

import (
	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

/*
KEYUP KeyUp
*/
const KEYUP = "KeyUp"
/*
KEYDOWN KeyHold
*/
const KEYDOWN = "KeyHold"

type keydownMonitor struct {
  keys []uint16
  fn func(hook.Event)
}

var events = []keydownMonitor{}

func includes(keys []int, e hook.Event) bool {
  for _, key := range keys {
    if key == int(e.Rawcode) {
      return true
    }
  }
  return false
}

/*
Add Register keyboard event
*/
func Add(keys []uint16, fn func(hook.Event)) {
  events = append(events, keydownMonitor{keys, fn})
}
/*
Process start listen
*/
func Process() {
  robotgo.EventHook(hook.KeyDown, []string{"ctrl", "shift", "0"}, func(e hook.Event) {
    TapStr("你好")
  })

  defer robotgo.EventEnd()
  s := robotgo.EventStart()
  <-robotgo.EventProcess(s)
}

/*
TapStr tapString
*/
func TapStr(str string) {
  robotgo.TypeStr(str)
}
