// Package app is used to accept command line arguments when starting the container.
//
// Copyright (c) 2016 VMware
// Author: Luis M. Valerio (lvaleriocasti@vmware.com)
//
// License: MIT
//
package app

import (
	"github.com/vmtrain/data-manager/models"
	"github.com/vmtrain/data-manager/stats"
)

// Context is a struct to hold global application context variables.
type Context struct {
	ListenPort  int
	ContentRoot string
	APIHost     string
	Stats       stats.Stats
	backend     Backend
}

// New generates an AppContext struct
func New() *Context {
	// TODO Remove once we start integration with the blob service
	bcknd := NewBackend(
		NewMockDatastore(map[int]*models.Blob{}),
	)
	ctx := &Context{
		ListenPort:  80,
		ContentRoot: ".",
		APIHost:     "localhost",
		Stats:       stats.New(),
		backend:     bcknd,
	}
	return ctx
}
