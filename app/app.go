// Copyright (c) 2016 VMware
// Author: Tom Hite (thite@vmware.com)
//
// License: MIT (see https://github.com/tdhite/go-reminders/LICENSE).
//
package app

import (
	"flag"
	"log"
)

var Context = New()

// Initialize the flags processor with default values and help messages.
func initFlags() {
	const (
		listenPortDefault  = 8080
		listenPortUsage    = "port on which to listen for HTTP requests"
		contentRootDefault = "."
		contentRootUsage   = "path to (static content) templates, skeleton, etc."
		apiHostDefault     = "localhost"
		apiHostUsage       = "host to use for all API calls made internally."
	)

	flag.IntVar(&Context.ListenPort, "listenport", listenPortDefault, listenPortUsage)
	flag.IntVar(&Context.ListenPort, "l", listenPortDefault, listenPortUsage+" (shorthand)")
	flag.StringVar(&Context.APIHost, "apiHost", apiHostDefault, apiHostUsage)
	flag.StringVar(&Context.APIHost, "h", apiHostDefault, apiHostUsage+" (shorthand)")
	flag.StringVar(&Context.ContentRoot, "tplpath", contentRootDefault, contentRootUsage)
	flag.StringVar(&Context.ContentRoot, "t", contentRootDefault, contentRootUsage+" (shorthand)")
}

// Process application (command line) flags. Note this happens automatically.
// No need to explicitly call this function (in fact that is a bad idea).
func init() {
	initFlags()
	flag.Parse()
	log.Printf("Initialized app package.")
}