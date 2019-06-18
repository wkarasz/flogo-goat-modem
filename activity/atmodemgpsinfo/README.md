# 	AT Modem GPS Info - Activity

Communicate with AT Modem device using SIMCOM 7600X based devices.
Send raw/direct AT commands to SIMCOM chipset.

## Installation
Command for Flogo CLI:
```console
flogo install github.com/wkarasz/flogo-goat-modem/activity/atmodemgpsinfo
```

Link for Flogo Web UI:
```console
https://github.com/wkarasz/flogo-goat-modem/activity/atmodemgpsinfo
```

## Schema
Inputs and Outputs:
```json
{
  "inputs":[
    {
      "name": "devicePath",
      "type": "string",
      "required": true
    }
  ],
  "outputs": [
    {
      "name": "latitude",
      "type": "string"
    },
    {
      "name": "ns-indicator",
      "type": "string"
    },
    {
      "name": "longitude",
      "type": "string"
    },
    {
      "name": "ew-indicator",
      "type": "string"
    },
    {
      "name": "date",
      "type": "string"
    },
    {
      "name": "utctime",
      "type": "string"
    },
    {
      "name": "altitude",
      "type": "string"
    },
    {
      "name": "speed",
      "type": "string"
    },
    {
      "name": "course",
      "type": "string"
    }
  ]
}
```
## Inputs
| Input            | Description    |
|:-----------------|:---------------|
| devicePath       | The path to the AT modem device on the host system; e.g. /dev/ttyUSB0 |

# Outputs
| Output           | Description    |
|:-----------------|:---------------|
| latitude         | ddmm.mmmmmm    |
| ns-indicator     | north or south |
| longitude        | ddmm.mmmmmm    |
| ew-indicator     | east or west |
| date             | ddmmyy |
| utctime          | hhmmss.s |
| altitude         | altitude [meters] |
| speed            | speed [knots] |
| course           | course [degrees]|
