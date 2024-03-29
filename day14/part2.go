package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Reindeer struct {
	speed      int
	flyTime    int
	restTime   int
	distance   int
	points     int
	flying     bool
	timeInMode int
}

func main() {
	reindeers, err := readReindeerDetails("day14/input.txt")
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	simulateRaceWithPoints(&reindeers, 2503)
	maxPoints := findMaxPoints(reindeers)
	fmt.Println(maxPoints)
}

func readReindeerDetails(filename string) ([]Reindeer, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var reindeers []Reindeer
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Fields(scanner.Text())
		speed, _ := strconv.Atoi(parts[3])
		flyTime, _ := strconv.Atoi(parts[6])
		restTime, _ := strconv.Atoi(parts[13])

		reindeers = append(reindeers, Reindeer{speed: speed, flyTime: flyTime, restTime: restTime, flying: true})
	}

	return reindeers, scanner.Err()
}

func simulateRaceWithPoints(reindeers *[]Reindeer, totalSeconds int) {
	for i := 0; i < totalSeconds; i++ {
		maxDistance := 0
		for j := range *reindeers {
			reindeer := &(*reindeers)[j]
			if reindeer.flying {
				reindeer.distance += reindeer.speed
			}
			reindeer.timeInMode++
			if reindeer.flying && reindeer.timeInMode == reindeer.flyTime || !reindeer.flying && reindeer.timeInMode == reindeer.restTime {
				reindeer.flying = !reindeer.flying
				reindeer.timeInMode = 0
			}
			if reindeer.distance > maxDistance {
				maxDistance = reindeer.distance
			}
		}
		for j := range *reindeers {
			reindeer := &(*reindeers)[j]
			if reindeer.distance == maxDistance {
				reindeer.points++
			}
		}
	}
}

func findMaxPoints(reindeers []Reindeer) int {
	maxPoints := 0
	for _, reindeer := range reindeers {
		if reindeer.points > maxPoints {
			maxPoints = reindeer.points
		}
	}
	return maxPoints
}
