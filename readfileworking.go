package main

import (
	"bufio"
	"container/list"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

type resultStruct struct {
	isFound     bool
	substrIndex int
	wordFound   string
}

type compileTypeStruct struct {
	key   string
	value string
}

func compileStringToPatternList(strPattern string) *list.List {
	var i, j int
	// let listPattern be the regex compile pattern to be return
	listPattern := list.New()
	// 1. string contain '.*'?
	patternArray := strings.Split(strPattern, ".*")
	//fmt.Println(reflect.TypeOf(listPattern))
	for i = 0; i < len(patternArray); i++ {

		// 2. string pattern contains '+'
		if strings.Contains(patternArray[i], "+") {
			plusArray := strings.Split(patternArray[i], "+")
			for j = 0; j < len(plusArray); j++ {
				listPattern.PushBack(compileTypeStruct{key: "char", value: plusArray[j]})

				if j != len(plusArray)-1 {
					// add start operator
					listPattern.PushBack(compileTypeStruct{key: "plus", value: ""})
				}
			}
		} else {
			listPattern.PushBack(compileTypeStruct{key: "char", value: patternArray[i]})
		}
		if i != len(patternArray)-1 {
			// add start operator
			listPattern.PushBack(compileTypeStruct{key: "plus", value: ""})
		}
	}
	return listPattern
}

func findRegexInFile(path string, listResult *list.List, compiledPattern *list.List) {
	file := strings.NewReader(path)

	var pattern = "aa"
	
	// make buffered channels here
	jobs := make(chan string)
	results := make(chan resultStruct)

	// we need a wait group, not sure.
	wg := new(sync.WaitGroup)

	// start up some workers that will block and wait?
	for w := 1; w <= 20; w++ {
		wg.Add(1)
		go matchPattern(jobs, results, wg, pattern)
	}

	// Go over a file line by line and queue up a ton of work
	go func() {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			jobs <- scanner.Text()
		}
		close(jobs)
	}()

	// Now collect all the results...
	// But first, make sure we close the result channel when everything was processed
	go func() {
		wg.Wait()
		close(results)
	}()

	// Add up the results from the results channel.
	// counts := 0

	for element := range results {
		if element.isFound == true {
			// fmt.Println("found", element)
			listResult.PushBack(element)
		}
	}
}

func matchPattern(jobs <-chan string, results chan<- resultStruct, wg *sync.WaitGroup, pattern string) {
	// Decreasing internal counter for wait-group as soon as goroutine finishes
	defer wg.Done()

	// eventually I want to have a []string channel to work on a chunk of lines not just one line of text
	for j := range jobs {
		indexFound := strings.Index(j, pattern)
		if indexFound > 0 {
			results <- resultStruct{true, indexFound, ""}
		} else {
			results <- resultStruct{false, indexFound, ""}
		}
	}
}

func main() {
	// An artificial input source.  Normally this is a file passed on the command line.
	file, err := os.Open("../large.log")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var numOfLineToProcessOnce = 100
	var counter = 0
	var counterIdx int
	var multiLineString = ""
	listResult := list.New()

	listCompiledPattern := compileStringToPatternList("abc+def")
	/*for e := listPa.Front(); e != nil; e = e.Next() {
		//key := v.Value(compileTypeStruct{key string}).key
		//getKey := (e.Value)
		fmt.Print(e.Value.(compileTypeStruct).key, "::") // print out the elements
		fmt.Println(e.Value.(compileTypeStruct).value)   // print out the elements

	}*/
	for scanner.Scan() {
		counterIdx = counter % numOfLineToProcessOnce
		multiLineString += scanner.Text() + "\n"
		if counterIdx == 0 && counter != 0 {
			findRegexInFile(multiLineString, listResult, listCompiledPattern)

			multiLineString = ""
		}
		counter++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for e := listResult.Front(); e != nil; e = e.Next() {
		//key := v.Value(compileTypeStruct{key string}).key
		//getKey := (e.Value)
		fmt.Println(e.Value) // print out the elements
	}

}
