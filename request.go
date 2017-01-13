package main

import (
	"errors"
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

var records = []string{
	"A",
	"AAAA",
	"CNAME",
	"MX",
	"PTR",
	"TXT",
}

func contains(keyword string, list []string) bool {
	for _, v := range list {
		if keyword == v {
			return true
		}
	}
	return false
}

func NewRequest(args []string) (*Request, error) {
	ttli64, err := strconv.ParseInt(args[3], 10, 64)
	if err != nil {
		return nil, err
	}

	if !contains(args[1], records) {
		return nil, errors.New("Invalid record type")
	}

	if ttli64 > 86400 {
		return nil, errors.New("TTL can't be more than 86400 seconds")
	}

	req := &Request{
		Date:    time.Now().Format("2006/01/02 15:04:05 MST"),
		Action:  "UPSERT",
		Domain:  args[0],
		Type:    args[1],
		Value:   args[2],
		TTL:     ttli64,
		Profile: args[4],
	}

	if req.Type == "TXT" {
		req.Value = "\"" + req.Value + "\""
	}

	return req, nil
}
