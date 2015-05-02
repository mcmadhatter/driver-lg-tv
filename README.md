# Ninja Sphere - LG TV Driver


[![MIT License](https://img.shields.io/badge/license-MIT-yellow.svg)](LICENSE)
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

### Building < now optional >

This project can be built with the debug.sh script, you will need to create the /data/sphere/user-autostart/drivers/driver-lg-tv/  folder, and make sure the compiled driver and the json file are put in it once you finish compiling.

### Installing

* Create a new folder on your ninjasphere /data/sphere/user-autostart/drivers/driver-lg-tv   
* Add the driver-lg-tv and the package.json from github to this folder . 
* Make sure your TV is tuned on 
* Reboot the the spehere and got to http://ninjasphere.local  
* Click on the LG TVs button, then click New TV
* A pin number should appear on the TV screen
* Add the pin and a TV name to the config page and click save.

### More Information

More information can be found on the [project site](http://github.com/mcmadhatter/driver-lg-tv) or by visiting the Ninja Blocks [forums](https://discuss.ninjablocks.com).

### Contributing Changes

To contribute code changes to the project, please clone the repository and submit a pull-request ([What does that mean?](https://help.github.com/articles/using-pull-requests/)). Changes, forks, pull reqs , suggestions etc are always welcome.

### License
This project is licensed under the MIT license, a copy of which can be found in the [LICENSE](LICENSE) file.

### Copyright
This work is Copyright (c) 2014-2015 - mcmadhatter.
