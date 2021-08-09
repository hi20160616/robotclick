package main

import (
	"fmt"

	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

func main() {
	// robotgo.ScrollMouse(10, "up")
	// robotgo.MouseClick("left", true)
	// robotgo.MoveMouseSmooth(100, 200, 1.0, 100.0)
	// robotgo.TypeStr("Hello World. Winter is coming!")
	// robotgo.KeyTap("enter")

	// // screen
	// x, y := robotgo.GetMousePos()
	// fmt.Println("pos: ", x, y)
	//
	// color := robotgo.GetPixelColor(100, 200)
	// fmt.Println("color----", color)
	//
	// // bitmap
	// bitmap := robotgo.CaptureScreen(10, 20, 30, 40)
	// // use `defer robotgo.FreeBitmap(bit)` to free the bitmap
	// defer robotgo.FreeBitmap(bitmap)
	// fmt.Println("...", bitmap)
	//
	// fx, fy := robotgo.FindBitmap(bitmap)
	// fmt.Println("FindBitmap------", fx, fy)
	//
	// robotgo.SaveBitmap(bitmap, "test.png")

	// // handler
	// fpid, err := robotgo.FindIds("Google")
	// if err == nil {
	//         fmt.Println("pids...", fpid)
	//
	//         if len(fpid) > 0 {
	//                 robotgo.ActivePID(fpid[0])
	//
	//                 robotgo.Kill(fpid[0])
	//         }
	// }
	//
	// robotgo.ActiveName("chrome")
	//
	// isExist, err := robotgo.PidExists(100)
	// if err == nil && isExist {
	//         fmt.Println("pid exists is", isExist)
	//
	//         robotgo.Kill(100)
	// }
	//
	// abool := robotgo.ShowAlert("test", "robotgo")
	// if abool {
	//         fmt.Println("ok@@@ ", "ok")
	// }
	//
	// title := robotgo.GetTitle()
	// fmt.Println("title@@@ ", title)

	// event
	add()
	low()
	event()
}

func add() {
	fmt.Println("--- Please press ctrl + shift + q to stop hook ---")
	robotgo.EventHook(hook.KeyDown, []string{"q", "ctrl", "shift"}, func(e hook.Event) {
		fmt.Println("ctrl-shift-q")
		robotgo.EventEnd()
	})

	fmt.Println("--- Please press w---")
	robotgo.EventHook(hook.KeyDown, []string{"w"}, func(e hook.Event) {
		fmt.Println("w")
	})

	s := robotgo.EventStart()
	<-robotgo.EventProcess(s)
}

func low() {
	EvChan := hook.Start()
	defer hook.End()

	for ev := range EvChan {
		fmt.Println("hook: ", ev)
	}
}

func event() {
	ok := robotgo.AddEvents("q", "ctrl", "shift")
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
