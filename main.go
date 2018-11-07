package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/pkg/errors"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	if len(os.Args) == 1 {
		return errors.New("subcommand is missing")
	}

	switch os.Args[1] {
	case "getpairs":
		cli := New()
		res, err := cli.GetPairs(context.Background())
		if err != nil {
			return err
		}
		pairs := res.(*Pairs)
		for _, pair := range *pairs {
			fmt.Println(pair.CurrencyPair)
		}

	case "getprice":
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
	default:
		return errors.Errorf("invaild command: %q", os.Args[1])
	}
	return nil
}

//out outputs the json data encoded from go structure
func out(res Data) error {
	if err := json.NewEncoder(os.Stdout).Encode(res); err != nil {
		return err
	}
	return nil
}
