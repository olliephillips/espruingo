// Espruingo
// A live code loader for Espruino - use your favorite Editor/IDE
// Copyright 2015 Ollie Phillips
// MIT license

package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/daviddengcn/go-colortext"
	"github.com/go-fsnotify/fsnotify"
	"github.com/jacobsa/go-serial/serial"
	"log"
	"os"
	"strings"
	"time"
)

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

	// Set up connection
	options := serial.OpenOptions{
		PortName:        device,
		BaudRate:        19200,
		DataBits:        8,
		StopBits:        1,
		MinimumReadSize: 4,
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
					_, err = s.Write([]byte("reset();\n"))

					colorLog("< Writing to board", "blue")
					// Open and scan file
					file, err := os.Open(targetFile)
					if err != nil {
						log.Fatal(err)
					}
					defer file.Close()

					scanner := bufio.NewScanner(file)

					for scanner.Scan() {
						//Send each line to Espruino
						_, err = s.Write([]byte(scanner.Text() + "\n"))
					}

					if err := scanner.Err(); err != nil {
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
