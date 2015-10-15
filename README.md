# Espruingo
Espruingo is a live code loader for Espruino (http://www.espruino.com) - run once and it will monitor the target file, sending changes to your Espruino board on each file save. Use your favorite Editor/IDE.

Also provides console output from Espruino - plugin and connect to Espruino to monitor output.

Written in Go (1.5).

## Key Features
* Source file monitoring, files uploaded to Espruino board on save
* Module loading, recursively fetches and loads minified versions of the modules from http://www.espurino.com/modules. 
* Minification before sending code to Espruino board
* Console output in terminal

## Install
You need a Go environment and to download and install the source. There are binaries, check the binaries folder and the README in there.
 - go get github.com/olliephillips/espruingo
 - go install github.com/olliephillips/espruingo

## Documentation
Intended to run as a binary, once installed, make sure in your PATH so you can run like this:
```
espruingo <file to watch> <port name>
```

If not in path, seems if you have a Go environment setup properly, you can run like this (my MacBook):
```
./espruingo <file to watch> <port name>

```

## Example
Using the included esp.js file, assuming on a Macbook or Linux you might start it like this:
```
espruingo esp.js /dev/tty.usbmodemfa131

```
Make changes to `esp.js` and they are pushed to your board on file save

## Changelog
v0.9.0 - Feature complete.

## License & Copyright
Copyright Ollie Phillips 2015. MIT licensed.

## Contributions
Contributions welcome. Please fork and create a new branch for your development.