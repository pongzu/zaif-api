package main

import "encoding/json"

// Data is an interfdce for data from api
type Data interface {
	unmshl(json.RawMessage) error
}
