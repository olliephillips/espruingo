# Espruingo
Espruingo is a live code loader for Espruino (http://www.espruino.com) - run once and it will monitor the target file, sending changes to Espruino on each file save. Use your favorite Editor/IDE.

Also provides console output from Espruino - plugin and connect to Espruino to monitor output.

Written in Go (1.5).

## Key Features
* Source file monitoring, files uploaded to Espurino board on save
* Basic module loading, fetches the minified versions of the modules from  http://www.espurino.com/modules. (Recursion not currently supported).
* Console output in terminal

## Install
Currently there are no binaries, so you need a Go environment and to download and install the source.
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

## License & Contributions
MIT licensed. Contributions welcome.

## Roadmap
* Maybe make cross platform binaries available.
* Minification
* Recursive module loading, handle require statements in modules