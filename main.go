package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
)

func main() {
	if err := run(); err != nil {
		log.Println(err)
	}

}

func run() error {
	switch os.Args[1] {
	case "getrate":
		cli := New()
		res, err := cli.GetPrice(context.Background(), os.Args[2])
		if err != nil {
			return err
		}
		if err := out(res); err != nil {
			return err
		}

	case "getticker":
		cli := New()
		res, err := cli.GetTicker(context.Background(), os.Args[2])
		if err != nil {
			return err
		}
		if err := out(res); err != nil {
			return err
		}
	}
	return nil
}

func out(res Data) error {
	if err := json.NewEncoder(os.Stdout).Encode(res); err != nil {
		return err
	}
	return nil
}
