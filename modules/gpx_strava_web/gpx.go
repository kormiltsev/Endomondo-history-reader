package gpx_strava_web

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/neighborhood999/gpx"
)

// sample using gpx lib
func FileTypeStravaGPX(fileadr string) {
	f, err := os.Open(fileadr)

	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	b, _ := ioutil.ReadAll(f)
	g, _ := gpx.ReadGPX(bytes.NewReader(b))

	fmt.Println("Creator: ", g.Creator)                                                        // type of file (app/web)
	fmt.Println("Name: ", g.Tracks[0].Name)                                                    // name of workout
	fmt.Println(fmt.Sprintf("Total distance: %.3f km", g.Distance()))                          // Get the total running distance
	fmt.Println(fmt.Sprintf("Total duration: %.1f sec", g.Duration()))                         // Get the total running duration
	fmt.Println(fmt.Sprintf("Pace per km: %d:%d", g.PaceInKM().Minutes, g.PaceInKM().Seconds)) // Get the running pace(km/min)
	fmt.Println(fmt.Sprintf("Start point is:\n  Latitude: %.6f\n  Longitude: %.6f", g.GetCoordinates()[0].Latitude, g.GetCoordinates()[0].Longitude))
}
