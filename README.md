# ARSOWeatherImage
Converts ARSO weather image to string data for given location.

# Example
```golang
location, area, err := arsoweatherimage.RainfallRateLevels(arsoweatherimage.Maribor)
if err != nil {
  log.Fatal(err)
}

log.Println(location.Value, area.Value)
```

# How it works
<a href="http://meteo.arso.gov.si/uploads/probase/www/observ/radar/si0-rm-anim.gif">GIF image</a> is downloaded and scanned for the highest rainfall rate around the given location. Highest **on location** and **surrounding area** rainfall rates are returned.
