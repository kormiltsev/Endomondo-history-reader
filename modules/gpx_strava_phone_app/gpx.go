//copied from working main (other way)

package gpx_strava_phone_app

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"strconv"
	"time"

	haversine "github.com/LucaTheHacker/go-haversine" //distance via gps
)

type trk struct {
	Name        string  `xml:"name"`
	Number      string  `xml:"namber"`
	TrackPoints []trkpt `xml:"trkpt"`
}

type trkpt struct {
	Lat  string  `xml:"lat,attr"`
	Lon  string  `xml:"lon,attr"`
	Ele  float64 `xml:"ele"`
	Time string  `xml:"time"`
}

type gpx struct {
	Vers    string `xml:"version,attr"`
	Creator string `xml:"creator,attr"`
	Donnow1 string `xml:"xmlns:xsi,attr"`
	Donnow2 string `xml:"xmlns,attr"`
	Donnow3 string `xml:"xmlns:topografix,attr"`
	Donnow4 string `xml:"xsi:schemaLocation,attr"`
	Track   trk    `xml:"trk"`
}

func StravaFromAndroid(catalogFilename string) {
	v := gpx{Vers: "none", Creator: "none"}

	//read from file
	data, err := ioutil.ReadFile(catalogFilename)
	if err != nil {
		fmt.Println("Cannot load catalog:", err)
	}

	//unmarshaling
	err = xml.Unmarshal([]byte(data), &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	//export data
	//fmt.Printf("version: %q\n", v.Vers)
	fmt.Printf("creator: %q\n", v.Creator)
	fmt.Println("Name: ", v.Track.Name)
	//fmt.Println("Date", v.Track.TrackPoints[0].Time)
	//fmt.Println("Track", v.Track)

	layout := "2006-01-02T15:04:05.000Z"
	tstart := v.Track.TrackPoints[0].Time
	tfin := v.Track.TrackPoints[len(v.Track.TrackPoints)-1].Time

	startdate, err := time.Parse(layout, tstart)
	if err != nil {
		fmt.Println(err)
	}

	finishdate, err := time.Parse(layout, tfin)
	if err != nil {
		fmt.Println(startdate)
	}

	duration := finishdate.Sub(startdate)

	fmt.Println("Workout starts: ", startdate)
	fmt.Println("Workout duration: ", duration)

	//distance
	fmt.Printf("Total distance: %.2fkm.", Workoutdistance(v.Track))
}

func Workoutdistance(workout trk) float64 {
	startlocation := haversine.Coordinates{
		Latitude:  0.0,
		Longitude: 0.0,
	}
	if s, err := strconv.ParseFloat(workout.TrackPoints[0].Lat, 64); err == nil {
		startlocation.Latitude = s
	}
	if s, err := strconv.ParseFloat(workout.TrackPoints[0].Lon, 64); err == nil {
		startlocation.Longitude = s
	}

	runlocation := startlocation
	nextlocation := startlocation
	distance := 0.0
	for _, dots := range workout.TrackPoints {
		if s, err := strconv.ParseFloat(dots.Lat, 64); err == nil {
			nextlocation.Latitude = s
		}
		if s, err := strconv.ParseFloat(dots.Lon, 64); err == nil {
			nextlocation.Longitude = s
		}

		distancer := haversine.Distance(
			runlocation,
			nextlocation,
		)
		runlocation = nextlocation
		distance += distancer.Kilometers()
	}
	return distance
}
