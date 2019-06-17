# 	AT Modem Send SMS - Activity

Communicate with AT Modem device using SIMCOM 7600X based devices.
Send wrapped SMS AT commands to SIMCOM chipset.

## Installation
Command for Flogo CLI:
```console
flogo install github.com/wkarasz/flogo-goat-modem/activity/atmodemsmssend
```

Link for Flogo Web UI:
```console
https://github.com/wkarasz/flogo-goat-modem/activity/atmodemsmssend
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
    },
    { "name": "directCmd",
      "type": "string",
      "required": true
    },
    {
      "name": "recipientMobile",
      "type": "string",
      "required": true
    },
    {
      "name": "message",
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
| directCmd        | The raw command supported by the SIMCOM chipset; e.g. ATI<br>https://www.waveshare.com/w/upload/5/54/SIM7500_SIM7600_Series_AT_Command_Manual_V1.08.pdf|
| recipientMobile  | Cell phone number to send a text message |
| message          | Contents of text message }

# Outputs
| Output           | Description    |
|:-----------------|:---------------|
| result           | The result will contain a string response of the command or will contain an error message |
