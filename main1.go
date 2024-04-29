package main

//
//import (
//	"fmt"
//	"sort"
//)
//
//func main() {
//	populationForCity := map[string]int{"Zistanbul": 12610000, "Karachi": 10620000, "Mumbai": 12690000, "Shanghai": 13680000}
//	for city, population := range populationForCity {
//		fmt.Printf("%-10s %8d\n", city, population)
//	}
//	fmt.Println()
//	cities := make([]string, 0, len(populationForCity))
//	for city := range populationForCity {
//		cities = append(cities, city)
//	}
//	sort.Strings(cities)
//	for _, city := range cities {
//		fmt.Printf("%-10s %8d\n", city, populationForCity[city])
//	}
//
//	cityForPopulation := make(map[int]string, len(populationForCity))
//	for city, population := range populationForCity {
//		cityForPopulation[population] = city
//	}
//	fmt.Println(cityForPopulation)
//
//}
