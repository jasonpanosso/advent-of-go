package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string
var lines []string

type RangeMap struct {
	BaseA int
	BaseB int
	Range int
}

type FarmMaps struct {
	SeedSoil            *[]RangeMap
	SoilFertilizer      *[]RangeMap
	FertilizerWater     *[]RangeMap
	WaterLight          *[]RangeMap
	LightTemperature    *[]RangeMap
	TemperatureHumidity *[]RangeMap
	HumidityLocation    *[]RangeMap
}

var farmMaps = FarmMaps{
	SeedSoil:            &[]RangeMap{},
	SoilFertilizer:      &[]RangeMap{},
	FertilizerWater:     &[]RangeMap{},
	WaterLight:          &[]RangeMap{},
	LightTemperature:    &[]RangeMap{},
	TemperatureHumidity: &[]RangeMap{},
	HumidityLocation:    &[]RangeMap{},
}

var inputStrToFarmMap = map[string]*[]RangeMap{
	"seed-to-soil":            farmMaps.SeedSoil,
	"soil-to-fertilizer":      farmMaps.SoilFertilizer,
	"fertilizer-to-water":     farmMaps.FertilizerWater,
	"water-to-light":          farmMaps.WaterLight,
	"light-to-temperature":    farmMaps.LightTemperature,
	"temperature-to-humidity": farmMaps.TemperatureHumidity,
	"humidity-to-location":    farmMaps.HumidityLocation,
}

func init() {
	if len(input) == 0 {
		panic("empty input.txt file")
	}

	lines = strings.Split(input, "\n")
}

func main() {
	fmt.Println("== Day 05 ==")

	setupFarmMaps(lines)

	part1Output := part1(lines)
	fmt.Println("Part1 output:", part1Output)

	part2Output := part2(lines)
	fmt.Println("Part2 output:", part2Output)
}

func part1(lines []string) int {
	seeds := parseSeeds(lines[0])

	soils := getBaseBsForSlice(seeds, *farmMaps.SeedSoil)
	fertilizers := getBaseBsForSlice(soils, *farmMaps.SoilFertilizer)
	waters := getBaseBsForSlice(fertilizers, *farmMaps.FertilizerWater)
	lights := getBaseBsForSlice(waters, *farmMaps.WaterLight)
	temps := getBaseBsForSlice(lights, *farmMaps.LightTemperature)
	humidities := getBaseBsForSlice(temps, *farmMaps.TemperatureHumidity)
	locations := getBaseBsForSlice(humidities, *farmMaps.HumidityLocation)

	output := math.MaxInt

	for _, location := range locations {
		if location < output {
			output = location
		}
	}

	return output
}

func part2(lines []string) int {
	seeds := parseSeeds(lines[0])

	// have to do soils differently, getBaseBsForSlice second impl that takes num
	// start range and end range : )
	soils := getBaseBsForSlice(seeds, *farmMaps.SeedSoil)
	fertilizers := getBaseBsForSlice(soils, *farmMaps.SoilFertilizer)
	waters := getBaseBsForSlice(fertilizers, *farmMaps.FertilizerWater)
	lights := getBaseBsForSlice(waters, *farmMaps.WaterLight)
	temps := getBaseBsForSlice(lights, *farmMaps.LightTemperature)
	humidities := getBaseBsForSlice(temps, *farmMaps.TemperatureHumidity)
	locations := getBaseBsForSlice(humidities, *farmMaps.HumidityLocation)

	output := math.MaxInt

	for _, location := range locations {
		if location < output {
			output = location
		}
	}

	return output
}

func stringToIntArray(str string) []int {
	words := strings.Fields(str)
	var numbers []int

	for _, word := range words {
		number, err := strconv.Atoi(word)
		if err != nil {
			continue
		}

		numbers = append(numbers, number)
	}

	return numbers
}

func setupFarmMaps(strs []string) {
	for i := 0; i < len(strs); i++ {
		if !strings.HasSuffix(strs[i], "map:") {
			continue
		}

		words := strings.Fields(strs[i])

		_, exists := inputStrToFarmMap[words[0]]
		if !exists {
			fmt.Printf("Unexpected input.txt format, %v not in listed map types", words[0])
			panic("")
		}

		var rangeMaps []RangeMap

		i++
		for {
			if strs[i] == "" {
				break
			}

			nums := stringToIntArray(strs[i])
			if len(nums) != 3 {
				fmt.Printf("Unexpected input.txt format, map definition did not contain three numbers: %v", strs[i])
				panic("")
			}

			rangeMaps = append(rangeMaps, RangeMap{BaseA: nums[1], BaseB: nums[0], Range: nums[2]})

			i++
		}

		*inputStrToFarmMap[words[0]] = append(*inputStrToFarmMap[words[0]], rangeMaps...)
	}
}

func parseSeeds(str string) []int {
	if !strings.HasPrefix(str, "seeds") {
		panic("Unexpected input.txt format")
	}

	seeds := stringToIntArray(str)
	return seeds
}

func getBaseBsForSlice(nums []int, rangeMaps []RangeMap) []int {
	var validBaseBs []int

	for _, num := range nums {
		var baseBs []int
		for _, rangeMap := range rangeMaps {
			if num >= rangeMap.BaseA && num <= rangeMap.BaseA+rangeMap.Range {
				diff := num - rangeMap.BaseA
				baseB := rangeMap.BaseB + diff
				baseBs = append(baseBs, baseB)
			}
		}

		if len(baseBs) == 0 {
			validBaseBs = append(validBaseBs, num)
		} else {
			validBaseBs = append(validBaseBs, baseBs...)
		}
	}

	return validBaseBs
}
