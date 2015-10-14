// Espruingo
// A live code loader for Espruino - use your favorite Editor/IDE
// Copyright 2015 Ollie Phillips
// MIT license

package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"github.com/daviddengcn/go-colortext"
	"github.com/go-fsnotify/fsnotify"
	"github.com/jacobsa/go-serial/serial"
	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/js"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
)

// No attempt to load these modules will be made
var coreModules = []string{
	"CC3000",
	"http",
	"WIZnet",
}

func main() {
	var targetFile, device string
	// Handle arguments
	flag.Parse()
	if len(flag.Args()) < 1 {
		colorLog("! Usage: espruingo <target file> <port name>", "red")
		os.Exit(0)
	}
	if len(flag.Args()) < 2 {
		colorLog("! Please provide the <port name> argument", "red")
		os.Exit(0)
	} else {
		device = flag.Args()[1]
	}
	targetFile = flag.Args()[0]
	log.Println(device)
	// Set up connection
	options := serial.OpenOptions{
		PortName:        device,
		BaudRate:        115200,
		DataBits:        8,
		StopBits:        1,
		MinimumReadSize: 1,
	}

	s, err := serial.Open(options)
	if err != nil {
		colorLog("! No Espruino connected at "+device+"..", "red")
		os.Exit(0)
	}
	// Clean up
	defer s.Close()

	// Connected, so say so
	colorLog("! Espruino connected..", "green")
	ct.Foreground(ct.Blue, false)

	// Initialise watcher for targetFile
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan bool)

	// Clean up
	defer watcher.Close()

	// Monitor target file for changes
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				if event.Op&fsnotify.Write == fsnotify.Write {
					// File has changed, read it, send a reset() to Espurino and send the file line by line
					_, err = s.Write([]byte("\x03reset();\n"))
					colorLog("< Writing to board", "blue")

					// Open and scan file
					fContents := bytes.NewBuffer(nil)
					f, err := os.Open(targetFile)
					if err != nil {
						colorLog("! Target file could not be opened for reading..", "red")
						os.Exit(0)
						log.Fatal(err)
					}
					io.Copy(fContents, f)
					f.Close()

					// This is our script
					script := string(fContents.Bytes())

					// Load modules
					script = loadModules(script)

					// Minify script
					script = minifyScript(script)

					// Write to board
					script = "echo(0)\n" + script + "echo(1)\n"
					_, err = s.Write([]byte(script))
					if err != nil {
						colorLog("! Unable to write to Esrpuino..", "red")
						os.Exit(0)
						log.Fatal(err)
					}

				}
			case err = <-watcher.Errors:
				colorLog("! Unexpected error..", "red")
				os.Exit(0)
				log.Fatal(err)
			}
		}
	}()

	// Handle terminal output
	reader := bufio.NewReader(s)
	go func() {
		// Read buffer to terminal
		for {
			time.Sleep(time.Second / 100)
			reply, err := reader.ReadBytes('\n')

			n := len(reply)
			limit := n - 1
			if err != nil {
				colorLog("! Espruino disconnected..", "red")
				os.Exit(0)
			}

			output := cleanConsoleOutput(string(reply[:limit]))
			if !strings.Contains(output, "Console Moved") { // Adds nothing, remove it
				if output != "." && output != "" && output != "echo(0)" {
					colorLog("> "+output, "black")
				}
			}
		}
	}()

	// Add targetFile to watcher
	err = watcher.Add(targetFile)
	if err != nil {
		log.Fatal(err)
	}
	<-done
}

// Helper function to clean up the console output
func cleanConsoleOutput(buf string) string {
	output := strings.Replace(buf, "\b", "", -1)
	output = strings.Replace(output, ">", "", -1)
	output = strings.Replace(output, "=function () { ... }", "", -1)
	output = strings.Replace(output, "=undefined", "", -1)
	return output
}

// Helper function to colorize console output
func colorLog(msg string, color string) {
	ct.Background(ct.Black, true)
	switch color {
	case "red":
		ct.Foreground(ct.Red, false)
	case "blue":
		ct.Foreground(ct.Blue, false)
	case "green":
		ct.Foreground(ct.Green, false)
	case "magenta":
		ct.Foreground(ct.Magenta, false)
	case "black":
		ct.Foreground(ct.Black, false)
	}
	fmt.Println(msg)
	ct.ResetColor()
}

// Helper function to load modules and namespace them
func loadModules(script string) string {
	var moduleUri = "http://www.espruino.com/modules"
	var moduleName string
	var moduleJS string
	var moduleTargetUri string

	// Scan script for require statements, first set up regex
	r, _ := regexp.Compile("require\\(\"(.*)\"\\)")

	// With our matches
	for _, req := range r.FindAllString(script, -1) {
		// For each match
		moduleName = strings.Split(req, "\"")[1]
		moduleTargetUri = moduleUri + "/" + moduleName + ".min.js"
		// If a core module we don't need to load it
		if !contains(coreModules, moduleName) {
			// Load the module
			resp, err := http.Get(moduleTargetUri)
			if err != nil {
				colorLog("! Could not get a module:  "+moduleName, "red")
				os.Exit(0)
				log.Fatal(err)
			}
			defer resp.Body.Close()
			contents, err := ioutil.ReadAll(resp.Body)

			// Recursively load modules from the module
			contents = []byte(loadModules(string(contents)))

			// Process the module
			moduleJS = "var espruingo_" + moduleName + " = {};\n" + string(contents)
			moduleJS = strings.Replace(moduleJS, "exports.", "espruingo_"+moduleName+".", 1)
			script = strings.Replace(script, "require(\""+moduleName+"\")", "espruingo_"+moduleName, 1)

			// We need to add the module to the top of our script
			script = moduleJS + "\n\n" + script
		}

	}
	return script
}

// Helper function to minify code
func minifyScript(script string) string {
	m := minify.New()
	m.AddFunc("text/javascript", js.Minify)
	miny, err := minify.String(m, "text/javascript", script)
	if err != nil {
		log.Fatal("minify.String:", err)
	}
	return miny
}

// Helper contains function to see if in slice
func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
