# Espruingo
Espruingo is a live code loader for Espruino - run once and it will monitor the target file, sending changes to Espruino on each file save. Use your favorite Editor/IDE. Written in Go (1.5)
Also provides console output from espruino. Plugin and connecto Espruino to monitor output.

## Install
Currently there are no binaries, so you need a Go environment and to download and install the source.
 - go get github.com/olliephillips/espruingo
 - go install github.com/olliephillips/espruingo

## Documentation
Intended to run as a binary, once installed, make sure in your PATH so you can run like this:
```
espruingo <file to watch> <port name>
```

If not in path, seems if you have a Go enviroment setup properly, you can run like this (my MacBook):
```
./espruingo <file to watch> <port name>

```

## Example
Using the included esp.js file, assuming on a Macbook or Linux you might start it like this:
```
espruingo esp.js /dev/tty.usbmodemfa131

```
Make changes and they are pushed to your board on file save

## License & Contributions
MIT licensed. Contributions welcome.

## Roadmap
Maybe make cross platform binaries available.