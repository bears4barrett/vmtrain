// Package app is used to accept command line arguments when starting the container.
//
// Copyright (c) 2016 VMware
// Author: Luis M. Valerio (lvaleriocasti@vmware.com)
//
// License: MIT
//
package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func MakeRequest(url, httpMethod string, response interface{}, payload *bytes.Reader) error {
	client := &http.Client{}
	// Make call to server
	request, err := http.NewRequest(httpMethod, url, payload)
	if err != nil {
		return fmt.Errorf("Failed to generate request. Caused by: %+v", err)
	}
	resp, err := client.Do(request)
	if err != nil {
		return fmt.Errorf("Error generated when performing %s on %s. Caused by: %+v", httpMethod, url, err)
	}

	if resp == nil {
		return fmt.Errorf("nil response returned from endpoint %s", url)
	}

	if resp.StatusCode > 200 {
		return fmt.Errorf("Expected status '%d' but saw (%d)", 200, resp.StatusCode)
	}

	if response != nil {
		// Read and validate the request. The read on the request body is limited
		// to prevent malicious attacks on the server.
		body, err := ioutil.ReadAll(io.LimitReader(resp.Body, 1048576))
		if err != nil {
			return fmt.Errorf("Failed to read the response body")
		}
		defer resp.Body.Close()

		// Unmarshal JSON
		err = json.Unmarshal(body, response)
		if err != nil {
			return fmt.Errorf("Failed to unmarshal the response body")
		}
	}

	return nil
}
