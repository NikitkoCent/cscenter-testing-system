package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func rangesAreEqual(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func execParamsToString(params []string) string {
	result := fmt.Sprintf("%q", params)
	return result[1 : len(result) - 1]
}

func main() {
	configPath := flag.String("config", "", "Path to tests configuration JSON file")

	flag.Parse()

	if *configPath == "" {
		fmt.Fprintln(os.Stderr, "flag was not set: config")
		flag.Usage()
		os.Exit(1)
	}

	config, err := ReadTestsConfig(*configPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", *configPath, err)
		os.Exit(2)
	}

	log.Printf("Tested executable path: '%s'\n", *config.Executable)
	log.Println("Running the tests suite ...\n")

	var failedTests []string

	for _, test := range config.Tests {
		registerFailedTest := func(logParams ...interface{}) {
			failedTests = append(failedTests, test.Name)

			if len(logParams) == 0 {
				log.Println("\tTEST FAILED.\n")
			} else {
				log.Printf("\tTEST FAILED: %s.\n\n", fmt.Sprint(logParams...))
			}
		}

		log.Printf("\tTest name: '%s'\n", test.Name)
		log.Printf("\tReference results file path: '%s'", *test.ReferencePath)

		log.Println("\tObtaining reference results ...")

		var referenceBytes []byte

		if file, err := os.Open(*test.ReferencePath); err != nil {
			registerFailedTest(err)
			continue
		} else {
			referenceBytes, err = ioutil.ReadAll(file)

			file.Close()

			if err != nil {
				registerFailedTest(err)
				continue
			}
		}

		log.Printf("\tExecuting '%s %s' ...\n", *config.Executable, execParamsToString(test.ExecParams))

		out, err := exec.Command(*config.Executable, test.ExecParams...).Output()
		if err != nil {
			if exitError, ok := err.(*exec.ExitError); ok {
				exitCodeIsOk := false

				for _, permittedExitCode := range test.ExitCodes {
					if exitError.ExitCode() == permittedExitCode {
						exitCodeIsOk = true
						break
					}
				}

				if !exitCodeIsOk {
					registerFailedTest(exitError.ExitCode())
					continue
				}
			} else {
				registerFailedTest(err)
				continue
			}
		}

		log.Println("\tVerifying results ...")
		if !rangesAreEqual(referenceBytes, out) {
			registerFailedTest("incorrect results")
			continue
		}
		log.Println("\tPassed.\n")
	}

	if len(failedTests) > 0 {
		log.Printf("FAILED TESTS: %q\n", failedTests)
		os.Exit(3)
	} else {
		log.Println("All tests passed.")
	}
}
