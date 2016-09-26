// Copyright (c) 2015 VMware
// Author: Tom Hite (thite@vmware.com)
//
// License: MIT (see https://github.com/tdhite/go-reminders/LICENSE).
//
package template

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/tdhite/q3-training-journal/journal"
	"github.com/tdhite/q3-training-journal/stats"
)

type Template struct {
	ContentRoot string
	APIHost     string
}

// Return a new Template object initialized -- convenience function.
func New(contentRoot string, apiHost string) Template {
	return Template{
		ContentRoot: contentRoot,
		APIHost:     apiHost,
	}
}

func init() {
	log.Println("Initialized Template.")
}

func (t *Template) generateAPIUrl(path string) string {
	return "http://" + t.APIHost + path
}

// Retrieve all Message entries from a topic via REST call.
func (t *Template) getTopic(topic string) []journal.Message {
	url := t.generateAPIUrl("/api/topic/" + topic + "?peekall=true")
	log.Println("url: " + url)

	res, err := http.Get(url)
	perror(err)
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	perror(err)

	var msgs []journal.Message
	if err := json.Unmarshal(body, &msgs); err != nil {
		log.Println(err)
	}

	return msgs
}

// Retrieve all Topic names via REST call.
func (t *Template) getAllTopics() *journal.Topics {
	url := t.generateAPIUrl("/api/topics")
	log.Println("url: " + url)

	res, err := http.Get(url)
	perror(err)
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	perror(err)

	topics := &journal.Topics{}
	topics.FromJson(body)

	return topics
}

// Add a topic message, to the journal via REST call.
func (t *Template) postTopic(topic string, msg journal.Message) {
	jsonData, err := json.Marshal(msg)
	perror(err)

	url := t.generateAPIUrl("/api/topic/" + topic)
	log.Println("url: " + url)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	perror(err)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	rsp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer rsp.Body.Close()

	_, err = ioutil.ReadAll(rsp.Body)
	perror(err)
}

// Retrieve stats info via REST call.
func (t *Template) getStatsHits() map[string]int {
	url := t.generateAPIUrl("/stats/hits")
	log.Println("url: " + url)

	res, err := http.Get(url)
	perror(err)
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	perror(err)

	data, err := stats.HitsFromJson(body)
	perror(err)

	return data
}
