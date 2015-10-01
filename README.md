# Espruingo
Espruingo is a live code loader for Espruino - use your favorite Editor/IDE. Written in Go (1.5)

## Install
Currently, there are no binaries, so you need a Go environment and to download and install the source.
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
Uses the included esp.js file, assume the port name is mine on my MacBook (/dev/tty.usbmodemfa131)
``` espruingo esp.js /dev/tty.usbmodemfa131
```

## License & Contributions
MIT licensed. Welcome contributions

## Roadmap
Maybe make cross platform binaries available.