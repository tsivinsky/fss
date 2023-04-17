package main

import (
	"flag"
	"log"
	"os"

	"github.com/tsivinsky/fss/cmd"
)

func main() {
	flag.Parse()

	arg := flag.Arg(0)
	if arg == "" {
		flag.Usage()
		os.Exit(1)
	}

	switch arg {
	case "new":
		err := cmd.CreateSession()
		if err != nil {
			log.Fatal(err)
		}
		break

	case "edit":
		err := cmd.EditSession()
		if err != nil {
			log.Fatal(err)
		}
		break

	case "rm":
		err := cmd.DeleteSession()
		if err != nil {
			log.Fatal(err)
		}

	default:
		err := cmd.StartSession()
		if err != nil {
			log.Fatal(err)
		}
		break
	}
}
