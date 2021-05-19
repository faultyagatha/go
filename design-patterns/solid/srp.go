package main

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"
)

var entryCount = 0

//Journal keeps track of journal entries
type Journal struct {
	entries []string
}

func (j *Journal) String() string {
	return strings.Join(j.entries, "\n")
}

func (j *Journal) AddEntry(text string) int {
	entryCount++
	entry := fmt.Sprintf("%d: %s", entryCount, text)
	j.entries = append(j.entries, entry)
	return entryCount
}

func (j *Journal) RemoveEntry(index int) {
	// TODO
}

//  ---------------
// breaks SRP and turns into a God Object
// in case other additional functionality is needed,
// it's better to implement as helper functions
// but not member functions (see below)
//  ---------------

func (j *Journal) Save(filename string) {
	_ = ioutil.WriteFile(filename,
		[]byte(j.String()), 0644)
}

func (j *Journal) Load(filename string) {

}

func (j *Journal) LoadFromWeb(url *url.URL) {

}

//  ---------------
// here starts a better solution that doesn't break SRP

var lineSeparator = "\n"

func SaveToFile(j *Journal, filename string) {
	_ = ioutil.WriteFile(filename,
		[]byte(strings.Join(j.entries, lineSeparator)), 0644)
}

// Persistence saves a journal to a file
type Persistence struct {
	lineSeparator string
}

func (p *Persistence) saveToFile(j *Journal, filename string) {
	_ = ioutil.WriteFile(filename,
		[]byte(strings.Join(j.entries, p.lineSeparator)), 0644)
}

func main_() {
	j := Journal{}
	j.AddEntry("I run six km today and then I ate a bug.")
	j.AddEntry("I ate a bug and run six km today.")
	fmt.Println(strings.Join(j.entries, "\n"))

	// separate function
	SaveToFile(&j, "journal.txt")

	// persistence
	p := Persistence{"\n"}
	p.saveToFile(&j, "journal.txt")
}
