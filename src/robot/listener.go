package robot

import (
  "fmt"

  "github.com/go-vgo/robotgo"
  hook "github.com/robotn/gohook"
)

/*
Add Event
*/
func Add() {
  fmt.Println("--- Please press ctrl + shift + q to stop hook ---")
  robotgo.EventHook(hook.KeyDown, []string{"0", "ctrl", "alt"}, func(e hook.Event) {
    fmt.Println("ctrl + alt + 0")
    robotgo.EventEnd()
  })

  s := robotgo.EventStart()
  <- robotgo.EventProcess(s)
}

/*
Low xx
*/
func Low() {
  evChan := hook.Start()
  defer hook.End()
  for ev := range evChan {
    fmt.Println("hook: ", ev)
  }
}

/*
Event listener
*/
func Event() {
  ok := robotgo.AddEvents("0", "ctrl", "alt")
  if ok {
    fmt.Println("add events...")
  }

  keve := robotgo.AddEvent("k")
  if keve {
    fmt.Println("you press... ", "k")
  }

  mleft := robotgo.AddEvent("mleft")
  if mleft {
    fmt.Println("you press... ", "mouse left button")
  }
}
