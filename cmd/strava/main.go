package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	androidstrava "github.com/kormiltsev/strava/modules/gpx_strava_phone_app"
	endo "github.com/kormiltsev/strava/modules/tcx_endomondo"
	"github.com/neighborhood999/gpx"
)

const dir = "./etc/*.gpx"     // search for any Garmin files
const endodir = "./etc/*.tcx" // -/- Endomondo (or other tcx) files

func main() {
	//Endomondo
	en, err := filepath.Glob(endodir)
	if err != nil {
		fmt.Println("Error parsing Endo (TCX) file: ", err)
	}
	// manage every file
	for _, endofilename := range en {
		tcx, err := endo.ParseFile(endofilename)
		if err != nil {
			fmt.Println("Error parsing TCX file: ", err)
		}
		fmt.Println("Endomondo file found: ", endofilename)
		fmt.Println("Total Duration: ", tcx.Activities[0].TotalDuration())
		fmt.Println("Av heartbeat: ", tcx.Activities[0].AverageHeartbeat())
		fmt.Println("Av Pace: ", tcx.Activities[0].AveragePace())
	}

	// looking for files
	m, err := filepath.Glob(dir)
	if err != nil {
		fmt.Println("Error parsing GPX file: ", err)
	}
	// manage files
	for _, filename := range m {
		f, err := os.Open(filename)
		if err != nil {
			fmt.Println("Error open GPX file", err)
		}

		defer f.Close()

		b, _ := ioutil.ReadAll(f)
		g, _ := gpx.ReadGPX(bytes.NewReader(b))
		switch g.Creator {
		case "Strava Android Application":
			androidstrava.StravaFromAndroid(filename)
		case "StravaGPX":
			fmt.Println("Creator: ", g.Creator)                                                        // type of file (app/web)
			fmt.Println("Name: ", g.Tracks[0].Name)                                                    // name of workout
			fmt.Println(fmt.Sprintf("Total distance: %.3f km", g.Distance()))                          // Get the total running distance
			fmt.Println(fmt.Sprintf("Total duration: %.1f sec", g.Duration()))                         // Get the total running duration
			fmt.Println(fmt.Sprintf("Pace per km: %d:%d", g.PaceInKM().Minutes, g.PaceInKM().Seconds)) // Get the running pace(km/min)
			fmt.Println(fmt.Sprintf("Start point is:\n  Latitude: %.6f\n  Longitude: %.6f", g.GetCoordinates()[0].Latitude, g.GetCoordinates()[0].Longitude))
		}
	}
}
