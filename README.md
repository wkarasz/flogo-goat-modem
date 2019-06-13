# flogo-goat-modem
flogo-goat-modem is a library of Flogo activities designed for interfacing with an AT based modem, such as a Waveshare 7600X.  The underlying libraries is based on the work from github.com/warthog61/modem<br>
<br>
An AT-modem device is a requirement to use this library; e.g. Waveshare 7600X<br>
<br>
The library consists of pre-cast functions for SMS send commands and a direct device command function for everything else not yet implemented.
<br>
* Direct device command
  * [activity/atmodemdirect](activity/atmodemdirect/README.md)
* Pre-cast functions
  * [activity/atmodemsendsms](activity/atmodemsendsms/README.md)
