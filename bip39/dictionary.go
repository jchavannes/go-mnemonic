package bip39

import (
	"fmt"
	"log"
	"math"
	"strings"
)

//we keep all the dics in the memory for fast access
//lazy inits, and global vars :(
//TODO keep 1 copy of each word instead of 2 in memory
var dict map[string][]string
var reverseDict map[string]map[string]int

func dictionaryWordToIndex(w string) (int, error) {
	//make sure we have the dic
	_, err := dictionary()
	if err != nil {
		return 0, err
	}
	lang := "english"

	rev, ok := reverseDict[lang]
	if ok == false {
		return 0, fmt.Errorf("cannot find %v reversed", lang)
	}

	index, ok := rev[w]
	if ok == false {
		return 0, fmt.Errorf("word %v don't exists in the %v dictionary",
			w, lang)
	}
	return index, nil
}

func dictionaryIndexToWord(i int) (string, error) {
	size := int(math.Pow(2, wordBits))

	if i < 0 || i > size-1 {
		return "", fmt.Errorf("invalid index %v, must be 0-%v",
			i, size-1)
	}

	dict, err := dictionary()
	if err != nil {
		return "", err
	}

	return dict[i], nil
}

func dictionary() ([]string, error) {
	if dict == nil {
		dict = make(map[string][]string, 1)
		reverseDict = make(map[string]map[string]int, 1)
	}
	lang := "english"
	res, ok := dict[lang]
	if ok {
		return res, nil
	}

	size := int(math.Pow(2, wordBits))

	dict[lang] = make([]string, size)
	reverseDict[lang] = make(map[string]int, size)

	words := strings.Split(strings.TrimSpace(EnglishWords), "\n")
	i := 0
	for _, word := range words {
		dict[lang][i] = word
		reverseDict[lang][word] = i
		i++
	}

	if i != size {
		log.Fatalf("incomplete dictionary %v, exp lines %v, got %v",
			lang, i, size)
	}

	return dict[lang], nil
}
