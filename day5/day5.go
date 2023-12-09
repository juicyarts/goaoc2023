package day5

import (
	"strconv"
	"strings"
)

var stringToKeyMap = map[string]string{
	"seed-to-soil map:":            "Soil",
	"soil-to-fertilizer map:":      "Fertilizer",
	"fertilizer-to-water map:":     "Water",
	"water-to-light map:":          "Light",
	"light-to-temperature map:":    "Temperature",
	"temperature-to-humidity map:": "Humidity",
	"humidity-to-location map:":    "Location",
}

func XtoY(searchdata [][][]int, value int) int {
	sourceData := searchdata[1]
	destinationData := searchdata[0]

	for i, v := range sourceData {
		if (value >= v[0]) && (value <= v[1]) {
			var offset = destinationData[i][0] - v[0]
			return value + offset
		}
	}
	return value
}

type Almanac struct {
	seeds       []int
	seedRanges  [][]int
	soil        [][][]int
	fertilizer  [][][]int
	water       [][][]int
	light       [][][]int
	temperature [][][]int
	humidity    [][][]int
	location    [][][]int
}

func InputToAlmanac(input []string) Almanac {
	var currentGroup string
	var sds []int = []int{}
	var sdsRanges [][]int = [][]int{}
	var soil, fertilizer, water, light, temperature, humidity, location [][][]int = [][][]int{{}, {}}, [][][]int{{}, {}}, [][][]int{{}, {}}, [][][]int{{}, {}}, [][][]int{{}, {}}, [][][]int{{}, {}}, [][][]int{{}, {}}

OUTER:
	for _, line := range input {
		if strings.HasPrefix(line, "seeds: ") {
			seeds := strings.Split(strings.TrimPrefix(line, "seeds: "), " ")
			for seedIndex, seed := range seeds {
				seedAsInt, _ := strconv.Atoi(seed)

				if (seedIndex % 2) == 0 {
					nextSeedAsInt, _ := strconv.Atoi(seeds[seedIndex+1])
					sdsRanges = append(sdsRanges, []int{seedAsInt, seedAsInt + nextSeedAsInt - 1})

				}

				sds = append(sds, seedAsInt)
			}
			continue
		}

		for key, value := range stringToKeyMap {
			if strings.HasPrefix(line, key) {
				currentGroup = value
				continue OUTER
			}
		}

		if line != "" {
			row := strings.Split(line, " ")
			destinationRangeStart, _ := strconv.Atoi(row[0])
			sourceRangeStart, _ := strconv.Atoi(row[1])
			rangeLength, _ := strconv.Atoi(row[2])

			loDest, hiDest := destinationRangeStart, destinationRangeStart-1+rangeLength
			loSource, hiSource := sourceRangeStart, sourceRangeStart-1+rangeLength

			if currentGroup == "Soil" {
				soil = [][][]int{append(soil[0], []int{loDest, hiDest, rangeLength}), append(soil[1], []int{loSource, hiSource, rangeLength})}
			} else if currentGroup == "Fertilizer" {
				fertilizer = [][][]int{append(fertilizer[0], []int{loDest, hiDest, rangeLength}), append(fertilizer[1], []int{loSource, hiSource, rangeLength})}
			} else if currentGroup == "Water" {
				water = [][][]int{append(water[0], []int{loDest, hiDest, rangeLength}), append(water[1], []int{loSource, hiSource, rangeLength})}
			} else if currentGroup == "Light" {
				light = [][][]int{append(light[0], []int{loDest, hiDest, rangeLength}), append(light[1], []int{loSource, hiSource, rangeLength})}
			} else if currentGroup == "Temperature" {
				temperature = [][][]int{append(temperature[0], []int{loDest, hiDest, rangeLength}), append(temperature[1], []int{loSource, hiSource, rangeLength})}
			} else if currentGroup == "Humidity" {
				humidity = [][][]int{append(humidity[0], []int{loDest, hiDest, rangeLength}), append(humidity[1], []int{loSource, hiSource, rangeLength})}
			} else if currentGroup == "Location" {
				location = [][][]int{append(location[0], []int{loDest, hiDest, rangeLength}), append(location[1], []int{loSource, hiSource, rangeLength})}
			}
		}
	}

	return Almanac{
		seeds:       sds,
		seedRanges:  sdsRanges,
		soil:        soil,
		fertilizer:  fertilizer,
		water:       water,
		light:       light,
		temperature: temperature,
		humidity:    humidity,
		location:    location,
	}
}

func (almanac Almanac) GetLowestSeedLocation() int {
	var min int

	for _, v := range almanac.seeds {
		var seedLocation int

		for _, k := range [][][][]int{almanac.soil, almanac.fertilizer, almanac.water, almanac.light, almanac.temperature, almanac.humidity, almanac.location} {
			if seedLocation != 0 {
				seedLocation = XtoY(k, seedLocation)
			} else {
				seedLocation = XtoY(k, v)
			}
		}

		if min == 0 || seedLocation < min {
			min = seedLocation
		}
	}
	return min
}

func (almanac Almanac) GetLowestSeedLocationBySeedRange() int {
	var min int

	for _, seedRange := range almanac.seedRanges {

		loEnd, hiEnd := seedRange[0], seedRange[1]
		s := make([]int, hiEnd-loEnd)

		for seedIndex := range s {
			var seedLocation int

			for _, k := range [][][][]int{almanac.soil, almanac.fertilizer, almanac.water, almanac.light, almanac.temperature, almanac.humidity, almanac.location} {
				if seedLocation != 0 {
					seedLocation = XtoY(k, seedLocation)
				} else {
					seedLocation = XtoY(k, loEnd+seedIndex)
				}
			}

			if min == 0 || seedLocation < min {
				min = seedLocation
			}
		}

	}
	return min
}
