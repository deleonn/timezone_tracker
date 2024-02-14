package main

import (
	"fmt"
	"time"

	"github.com/getlantern/systray"
)

func main() {
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetTitle("Timezone Tracker")
	systray.SetTooltip("Timezone Tracker")

    pTime := systray.AddMenuItem("Loading...", "Current time")
    eTime := systray.AddMenuItem("Loading...", "Current time")

    pTime.Disable()
    eTime.Disable()

	// Update time every minute
	go func() {
		for {
            // Add your timezones here
            pst, _ := time.LoadLocation("America/Los_Angeles")
            est, _ := time.LoadLocation("America/New_York")
            cst, _ := time.LoadLocation("America/Chicago")

            now := time.Now().UTC()

            pstNow := now.In(pst)
            estNow := now.In(est)
            cstNow := now.In(cst)

            estDifference := estNow.Hour() - cstNow.Hour()
            pstDifference := pstNow.Hour() - cstNow.Hour()

            systray.SetTitle(fmt.Sprintf("%s", cstNow.Format("3:04 pm MST")))
            // You can add more menu items to display other timezones
            pTime.SetTitle(fmt.Sprintf("%s (%v)", pstNow.Format("3:04 pm MST"), pstDifference))
            eTime.SetTitle(fmt.Sprintf("%s (%v)", estNow.Format("3:04 pm MST"), estDifference))
            time.Sleep(1 * time.Second)
        }
	}()

    systray.AddSeparator()
	mQuitOrig := systray.AddMenuItem("Quit", "Quit the whole app")

	go func() {
		for {
			<-mQuitOrig.ClickedCh
			systray.Quit()
			return
		}
	}()
}

func onExit() {
	// Clean up here
}

