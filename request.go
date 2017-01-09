package main

import (
	"strconv"
	"time"
)

type Request struct {
	Date    string
	Action  string
	Domain  string
	Type    string
	Value   string
	TTL     int64
	Profile string
}

func NewRequest(args []string) (*Request, error) {
	ttli64, err := strconv.ParseInt(args[4], 10, 64)
	if err != nil {
		return nil, err
	}

	req := &Request{
		Date:    time.Now().Format("2006/01/02 15:04:05 MST"),
		Action:  "UPSERT",
		Domain:  args[1],
		Type:    args[2],
		Value:   args[3],
		TTL:     ttli64,
		Profile: args[5],
	}

	return req, nil
}
