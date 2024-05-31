package utils

import (
	"flag"
	"fmt"
	"os"
)

func ParseFlagsAndCheckErrors(defineAndParseArgs func(), checkAndAppendErrs func(errs *[]error), exitCode int) {
	defineAndParseArgs()
	if !flag.Parsed() {
		flag.Parse()
	}

	errs := []error{}
	checkAndAppendErrs(&errs)
	if len(errs) > 0 {
		printErrorsAndUsageThenExit(errs, exitCode)
	}
}

func printErrorsAndUsageThenExit(errs []error, exitCode int) {
	if len(errs) > 0 {
		fmt.Fprintln(os.Stderr, "Error(s):")
		for _, err := range errs {
			fmt.Fprintln(os.Stderr, "*", err)
		}

		fmt.Fprintln(os.Stderr)
		flag.Usage()

		os.Exit(exitCode)
	}
}
