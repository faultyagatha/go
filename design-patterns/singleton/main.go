package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"sync"
)

type Database interface {
	GetPopulation(name string) int
}

// ---------------
// Private singleton
//  ---------------
type singletonDatabase struct {
	capitals map[string]int
}

// ---------------
// below is an example that
// breaks DIP
//  ---------------
func (db *singletonDatabase) GetPopulation(name string) int {
	return db.capitals[name]
}

// ---------------
// sync.Once ensures that something gets called only once
// Both init and sync.Once are thread-safe
// but only sync.Once is lazy
//  ---------------
var once sync.Once
var instance *singletonDatabase

func GetSingletonDatabase() *singletonDatabase {
	once.Do(func() {
		db := singletonDatabase{}
		caps, err := readData(".\\capitals.txt")
		if err == nil {
			db.capitals = caps
		}
		instance = &db
	})
	return instance
}

// init() — we could, but it's not lazy

func GetTotalPopulation(cities []string) int {
	result := 0
	for _, city := range cities {
		//singleton is hardcoded
		result += GetSingletonDatabase().GetPopulation(city)
	}
	return result
}

// ---------------
// below is an example that
// doesn't break DIP
//  ---------------
var instance2 Database

// ---------------
// here we do not hardcode the singleton
// but pass it to a func
//  ---------------
func GetTotalPopulationEx(db Database, cities []string) int {
	result := 0
	for _, city := range cities {
		result += db.GetPopulation(city)
	}
	return result
}

type DummyDatabase struct {
	dummyData map[string]int
}

func (d *DummyDatabase) GetPopulation(name string) int {
	//lazily fill in dummy data
	if len(d.dummyData) == 0 {
		d.dummyData = map[string]int{
			"alpha": 1,
			"beta":  2,
			"gamma": 3}
	}
	return d.dummyData[name]
}

// ---------------
// helper function that
// reads data from a file
//  ---------------
func readData(path string) (map[string]int, error) {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)

	file, err := os.Open(exPath + path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	result := map[string]int{}

	for scanner.Scan() {
		k := scanner.Text()
		scanner.Scan()
		v, _ := strconv.Atoi(scanner.Text())
		result[k] = v
	}

	return result, nil
}

func main() {
	db := GetSingletonDatabase()
	pop := db.GetPopulation("Seoul")
	fmt.Println("Pop of Seoul = ", pop)

	cities := []string{"Seoul", "Mexico City"}
	tp := GetTotalPopulation(cities)
	tp2 := GetTotalPopulationEx(GetSingletonDatabase(), cities)

	ok := tp == (17500000 + 17400000)   // testing on live data
	ok2 := tp2 == (17500000 + 17400000) // testing on live data
	fmt.Println(ok)
	fmt.Println(ok2)
}
