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
      "name": "result",
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
| result           | The result will contain a string response of the command or will contain an error message |
