# RainfallRateSI
Client for checking rainfall rate image provided by ARSO (Slovenian government).

# How it works
Client downloads <a href="http://www.arso.gov.si/vreme/napovedi%20in%20podatki/radar_anim.gif">GIF image</a> every five minutes and processes it, as that's the rate that ARSO updates it at. It looks for highest rainfall rate in given area (`x`, `y` and `radious`) and prints out seperately **area** and **location** rainfall rate level in `float` (same as the numbers on the upper right legend on the image).

# Usage
**Build** binary from `src` folder and **run** it with **x**, **y** and **radious** arguments. Arguments are **not** optional and need to be provided in **pixels** and also **must be relative to the entire image**. 

**Example:** ```/bin/rainfallrate {x} {y} {radious}```
