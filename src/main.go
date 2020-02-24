package main

import (
	"image/gif"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"rainfallrate/src/mappers/output"
	"rainfallrate/src/mappers/rainfall"
)

const dataURL = "http://www.arso.gov.si/vreme/napovedi%20in%20podatki/radar_anim.gif"

func main() {
	log.Println("Running rainfall rate app...")

	args := os.Args[1:]
	if len(args) < 3 {
		log.Println("Too few arguments. Set arguments and restart.")
		return
	}

	// Arguments.
	argsX, err := strconv.ParseInt(args[0], 10, 64)
	argsY, err := strconv.ParseInt(args[1], 10, 64)
	argsRadius, err := strconv.ParseInt(args[2], 10, 64)
	if err != nil {
		log.Println("Error reading arguments. Arguments must be integers.")
		return
	}

	// Calculate pixel window.
	xLocation := int(argsX)
	yLocation := int(argsY)
	x1 := int(argsX - argsRadius)
	y1 := int(argsY - argsRadius)
	x2 := int(argsX + argsRadius)
	y2 := int(argsY + argsRadius)

	// Initialize mappers.
	outputMapper := output.New()
	rainfallRateMapper := rainfall.New()

	for true {
		// Get gif data.
		resp, err := http.Get(dataURL)
		if err != nil {
			log.Println(err)
			return
		}

		// Decode gif data.
		dataGif, err := gif.DecodeAll(resp.Body)
		if err != nil {
			log.Println(err)
		}
		dataImages := dataGif.Image

		// Loop through each gif image.
		highestAreaRateLevel := 0.0
		highestLocationRateLevel := 0.0
		for _, item := range dataImages {
			if item != nil {
				for y := y1; y <= y2; y++ {
					for x := x1; x <= x2; x++ {
						r, g, b, _ := item.At(x, y).RGBA()
						level := rainfallRateMapper.GetLevelByRGBA(uint16(r), uint16(g), uint16(b))
						if x == xLocation && y == yLocation && highestLocationRateLevel < level {
							highestLocationRateLevel = level
						}
						if highestAreaRateLevel < level {
							highestAreaRateLevel = level
						}
					}
				}
			}
		}
		log.Println(outputMapper.ToOutputModel(highestAreaRateLevel, highestLocationRateLevel))

		// Sleep for the next five minutes
		time.Sleep(time.Duration(time.Minute * 5))
	}
}
