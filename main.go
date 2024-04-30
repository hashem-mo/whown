package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"sync"
)

type StoreArgs struct {
	Concurrency int
	Help        bool
	Input       string
	Function	string
}



func intakeFunction(chanInput chan string, input string) ([]string) {
    readFile, err := os.Open(input)
    if err != nil {
        log.Fatalln(err) // Return the error to the caller
    }
    defer readFile.Close() // Ensure the file is closed when the function returns

    var lines []string // Create a slice to hold the lines
    fileScanner := bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)

    for fileScanner.Scan() {
        chanInput <-  fileScanner.Text() // Append each line to the slice
    }

    if err := fileScanner.Err(); err != nil {
        log.Fatalln(err)
    }

    return lines // Return the slice containing all the lines
}









func main() {

	args := StoreArgs{}
	storeCommand := flag.NewFlagSet("main", flag.ContinueOnError)
	storeCommand.StringVar(&args.Function, "f", "whois", "What function to run ssl/whois")
	storeCommand.IntVar(&args.Concurrency, "c", 100, "How many goroutines running concurrently")
	storeCommand.BoolVar(&args.Help, "h", false, "print usage!")
	storeCommand.StringVar(&args.Input, "i", "NONE", "A file with IPs/CIDRs or domains on each line")

	storeCommand.Parse(os.Args[1:])



	inputChannel := make(chan string)


	var inputwg sync.WaitGroup

	if args.Function == "whois"{
	for i := 0; i < args.Concurrency; i++ {
		inputwg.Add(1)
		go func() {
			defer inputwg.Done()
			for line := range inputChannel {
				getOrg(line)

			}
		}()
	}
	} else if args.Function == "ssl"{
		for i := 0; i < args.Concurrency; i++ {
			inputwg.Add(1)
			go func() {
				defer inputwg.Done()
				for line := range inputChannel {
					getOrganizationFromSSL(line)
	
				}
			}()
		}
	} else {
		log.Fatalln("You provided invalid function")
	}



	intakeFunction(inputChannel, args.Input)
	close(inputChannel)
	inputwg.Wait()
	

}