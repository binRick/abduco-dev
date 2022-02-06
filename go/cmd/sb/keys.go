package main

import (
	"runtime"
	"time"

	"github.com/micmonay/keybd_event"
)

func Keys() {
	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		panic(err)
	}

	// For linux, it is very important to wait 2 seconds
	if runtime.GOOS == "linux" {
		//	time.Sleep(2 * time.Second)
	}

	kb.SetKeys(keybd_event.VK_BACKSLASH)
	kb.HasCTRL(true)

	// Or you can use Press and Release
	kb.Press()
	time.Sleep(10 * time.Millisecond)
	kb.Release()

}
