package main

import (
	"context"
	"flag"
	"io"
	"log"
	"os"

	"github.com/pkg/errors"
	"github.com/pongzu/zaif"
)

var out = flag.String("out", "", "output file. if user does not specify the output file, result will print out into stdout")

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	if len(os.Args) == 1 {
		return errors.New("specify the subcommand")
	}

	out, err := newOutputFile(*out)
	if err != nil {
		return errors.Wrap(err, "newOutputFile failed")
	}

	var (
		cli = zaif.New()
		ctx = context.Background()
	)

	switch os.Args[1] {
	case "getpairs":
		res, err := cli.GetPairs(ctx)
		if err != nil {
			return errors.Wrap(err, "GetPairs failed")
		}
		out.Write(res)
	case "getprice":
		res, err := cli.GetPrice(ctx, os.Args[2])
		if err != nil {
			return errors.Wrap(err, "GetPrice failed")
		}
		out.Write(res)
	case "getticker":
		res, err := cli.GetTicker(ctx, os.Args[2])
		if err != nil {
			return errors.Wrap(err, "GetTicker failed")
		}
		out.Write(res)
	case "getTrades":
		res, err := cli.GetTrades(ctx, os.Args[2])
		if err != nil {
			return errors.Wrap(err, "GetTrades failed")
		}
		out.Write(res)
	default:
		return errors.Errorf("invalid command: %q", os.Args[1])
	}
	return nil
}

func newOutputFile(out string) (io.Writer, error) {
	if out == "" {
		return os.Stdout, nil
	}

	f, err := os.Open(out)
	if err != nil {
		return nil, errors.Wrap(err, "os.Open failed")
	}
	return f, nil
}
