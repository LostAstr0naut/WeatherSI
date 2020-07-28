# RainfallRateSI
Library for checking rainfall rate image provided by ARSO.

# Help
The following inputs are supported:
- BO
- GO
- KP
- ID
- JE
- PO
- KR
- LJ
- KO
- TB
- SG
- NM
- CE
- KK
- MB
- PT
- MS

The returned values are corresponding to the legend in the upper right of the <a href="http://meteo.arso.gov.si/uploads/probase/www/observ/radar/si0-rm-anim.gif">GIF image</a>.

# How it works
It downloads the provided <a href="http://meteo.arso.gov.si/uploads/probase/www/observ/radar/si0-rm-anim.gif">GIF image</a> and looks for highest rainfall rate in given location and returns highest **area** and **location** rainfall rate level.
