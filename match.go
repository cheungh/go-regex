package main

import (
	"container/list"
	"fmt"
	"reflect"
	"strings"
	// "github.com/cheungh/stringutil"
)

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
	fmt.Println(reflect.TypeOf(listPattern))
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
			listPattern.PushBack(compileTypeStruct{key: "star", value: ""})
		}
	}
	return listPattern
}

func patternMatchAgainstStr(compiledPattern *list.List, matchedString string) {
	/*
						"ab+cd+"
	var matchedString = "habcdefg"
	*/
	// fmt.Println("string is ", matchedString)
	//var firstMatchIndex = -1
	var findStringAtIndex = 0
	var counter = 0
	var firstFoundIndex = -1
	var priorChar btye
	for e := compiledPattern.Front(); e != nil; e = e.Next() {
		//var s = "ab+cd+"
		//		  "fdhabcdefg"
		charPattern := e.Value.(compileTypeStruct).value

		if e.Value.(compileTypeStruct).key == "char" {
			for i := 0; i < len(charPattern) ; i++ {
				
				fmt.Printf("%c",charPattern[i])
				for j := findStringAtIndex; j < len(matchedString) ; j++ {
					if matchedString[j] == charPattern[i] {
						fmt.Printf("found %c at $d",charPattern[i], j)
						if firstFoundIndex < 0 {
							firstFoundIndex = j
						}
						findStringAtIndex = j
						fmt.Println()
						break
					}
				}
			}
			
			//fmt.Println()
		} else if e.Value.(compileTypeStruct).key == "plus" {
			if matchedString[findStringAtIndex+1] == 
			// +
		} else {
			// .*
		}
		//fmt.Print(e.Value.(compileTypeStruct).key, "::") // print out the elements
		//fmt.Println(e.Value.(compileTypeStruct).value)   // print out the elements
		counter++
	}
	if firstFoundIndex > -1 {
		// we found the match
		fmt.Println(matchedString[firstFoundIndex:findStringAtIndex+1])
	} else {
		fmt.Println("not found")
	}

}

func main() {

	//var s string = "ab.*b+c.*"
	var s = "ab+cd+"
	var matchedString = "fdhabbcdefg"
	var compiledPatternList = compileStringToPatternList(s)
	//for e := compiledPatternList.Front(); e != nil; e = e.Next() {
	//    fmt.Println(e.Value) // print out the elements
	//}
	patternMatchAgainstStr(compiledPatternList, matchedString)
}
