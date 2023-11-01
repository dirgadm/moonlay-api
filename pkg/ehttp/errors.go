package ehttp

import (
	"encoding/json"
	"path/filepath"
	"regexp"
	"strings"
)

func ErrorOutput(field string, message string) error {
	o := new(Errors)
	o.Failure(field, message)
	return o.error()
}

type Errors struct {
	Valid           bool              // state of validation
	messages        map[string]string // compiled error messages
	failureMessages map[string]string // failing error messages
	customMessages  map[string]string // custom messages
	failureKeys     []string
}

func (o *Errors) Error() string {
	res, err := json.Marshal(o.Messages())
	if err != nil {
		res = []byte("")
	}
	return string(res)
}

func (o *Errors) error() *Errors {
	for i := range o.failureMessages {
		if c := o.customMessages[i]; c != "" {
			o.Failure(i, c)
			continue
		}
		re := regexp.MustCompile("[^a-z.]")
		ix := re.ReplaceAllString(i, "*")
		if c := o.customMessages[ix]; c != "" {
			o.Failure(i, c)
		}
	}

	res := make(map[string]string)
	for _, i := range o.failureKeys {
		k := strings.TrimSuffix(i, filepath.Ext(i))
		if _, ok := res[k]; !ok {
			res[k] = o.failureMessages[i]
		}
	}

	o.messages = res
	return o
}

func (o *Errors) Messages() map[string]string {
	return o.messages
}

func (o *Errors) Failure(k string, e string) {
	if o.failureMessages == nil {
		o.failureMessages = make(map[string]string)
	}

	o.Valid = false
	o.failureKeys = append(o.failureKeys, k)
	o.failureMessages[k] = e
}
