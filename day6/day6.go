package day6

import (
	"strconv"
	"strings"
)

type Race struct {
	time     int
	distance int
}

func InputToRaces(input []string) []Race {
	var times []int = []int{}
	var distances []int = []int{}
	var races []Race = []Race{}

	for _, line := range input {
		if strings.HasPrefix(line, "Time: ") {
			timesFromString := strings.Split(strings.TrimPrefix(line, "Time: "), " ")
			for _, time := range timesFromString {
				if time == "" {
					continue
				}

				timeAsInt, _ := strconv.Atoi(time)
				times = append(times, timeAsInt)
			}
		}

		if strings.HasPrefix(line, "Distance: ") {
			distancesFromString := strings.Split(strings.TrimPrefix(line, "Distance: "), " ")
			for _, distance := range distancesFromString {
				if distance == "" {
					continue
				}

				distanceAsInt, _ := strconv.Atoi(distance)
				distances = append(distances, distanceAsInt)
			}
		}
	}

	for timeindex, time := range times {
		races = append(races, Race{time: time, distance: distances[timeindex]})
	}

	return races
}

func (race Race) NumberOfWaysToBeatRace() int {
	var timesBeaten int

	for speed := 0; speed < race.time; speed++ {
		reachableDistance := speed * (race.time - speed)
		if reachableDistance > race.distance {
			timesBeaten++
		}
	}

	return timesBeaten
}

func MultiplyNumberOfWaysToBeatRace(races []Race) int {
	var result = 1

	for i := 0; i < len(races); i++ {
		result *= races[i].NumberOfWaysToBeatRace()
	}

	return result
}

// Part 2
func InputToRace(input []string) []Race {
	var times []int = []int{}
	var distances []int = []int{}
	var races []Race = []Race{}

	for _, line := range input {
		if strings.HasPrefix(line, "Time: ") {
			timesFromString := strings.Replace(strings.TrimPrefix(line, "Time: "), " ", "", -1)
			timeAsInt, _ := strconv.Atoi(timesFromString)
			times = append(times, timeAsInt)
		}

		if strings.HasPrefix(line, "Distance: ") {
			distancesFromString := strings.Replace(strings.TrimPrefix(line, "Distance: "), " ", "", -1)
			distanceAsInt, _ := strconv.Atoi(distancesFromString)
			distances = append(distances, distanceAsInt)
		}
	}

	for timeindex, time := range times {
		races = append(races, Race{time: time, distance: distances[timeindex]})
	}

	return races
}
