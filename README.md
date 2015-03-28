# Ninja Sphere - LG TV Driver


[![MIT License](https://img.shields.io/badge/license-MIT-yellow.svg)](LICENSE)
[![Ninja Sphere](https://img.shields.io/badge/built%20by-ninja%20blocks-lightgrey.svg)](http://ninjablocks.com)
[![Ninja Sphere](https://img.shields.io/badge/works%20with-ninja%20sphere-8f72e3.svg)](http://ninjablocks.com)

---


### Introduction
This is a driver for LG Smart TVs, allowing them to be used as part of Ninja Sphere.



### Supported Sphere Protocols

| Name | URI | Supported Events | Supported Methods |
| ------ | ------------- | ---- | ----------- |
| volume | [http://schema.ninjablocks.com/protocol/volume](https://github.com/ninjasphere/schemas/blob/master/protocol/volume.json) | | set, volumeUp, volumeDown, mute, unmute, toggleMute |
| media-control | [http://schema.ninjablocks.com/protocol/media-control](https://github.com/ninjasphere/schemas/blob/master/protocol/media-control.json) | play, pause  | |
| on-off | [http://schema.ninjablocks.com/protocol/on-off](https://github.com/ninjasphere/schemas/blob/master/protocol/on-off.json) | state | turnOff |

#### Can't Do
* There is currently no way to get state back from the television.
* Turn on the TV. Turning off works, but the TV stops responding once it is in standby mode.
* Add a new TV via the config screen (dynamically added config screens are not yet supported by the sphere ui)
* TV must be turned on whilst the ninja is booted up for the first time (to allow the message to be sent for the tv to show the pin number on screen.) Once config is working, this will not be necessary

### Requirements

* Go 1.3

### Building

This project can be built with the debug.sh script, you will need to create the /data/sphere/user-autostart/drivers/driver-lg-tv/  folder, and make sure the compiled driver and the json file are put in it once you finish compiling.

### Installing

As config is not yet enabled in the sphere-ui the driver will need installing twice. The first time you install, it is necessary to have the LG TV turned on , and able to access the same network as the ninjasphere. Compile the driver and put it onto the ninjasphere, the driver will make the TV pop up a six digit Pin. Make a not of the pin, then edit line 16 of driver.go

var mydefaultpin =   <add your pin as a 6 digit integer here>

After that is done , recompile the driver, and re upload it to the spehere, it should now be possible to add the TV to the sphere app, and the access it from the spheramid.



### More Information

More information can be found on the [project site](http://github.com/mcmadhatter/driver-lg-tv) or by visiting the Ninja Blocks [forums](https://discuss.ninjablocks.com).

### Contributing Changes

To contribute code changes to the project, please clone the repository and submit a pull-request ([What does that mean?](https://help.github.com/articles/using-pull-requests/)). Changes, forks, pull reqs , suggestions etc are always welcome.

### License
This project is licensed under the MIT license, a copy of which can be found in the [LICENSE](LICENSE) file.

### Copyright
This work is Copyright (c) 2014-2015 - mcmadhatter.
