# ARSOWeatherImageAPI
API which converts ARSO weather image to data for given input.

# Example
```golang
area, location, err := arsoweatherimage.GetRainfallRateLevels("MB")
if err != nil {
  log.Fatal(err)
}

log.Println(area.Value, location.Value)
```

# Help
Only abbreviated names of places are supported. You can find them on the <a href="http://meteo.arso.gov.si/uploads/probase/www/observ/radar/si0-rm-anim.gif">GIF image</a>. The returned values are corresponding to the legend in the upper right of the image.

# How it works
It downloads the provided <a href="http://meteo.arso.gov.si/uploads/probase/www/observ/radar/si0-rm-anim.gif">GIF image</a> and looks for highest rainfall rate in given location and returns highest **in area** and **on location** rainfall rate level.
