package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type bot struct {
	number      int
	givesLowTo  *bot
	givesHighTo *bot
	holdsLow    int
	holdsHigh   int
	isOutput    bool
}

type botHelpers interface {
	receive() bool
}

func (b bot) receive(value int) bot {
	if b.holdsLow == 0 {
		b.holdsLow = value
	} else if value > b.holdsLow {
		b.holdsHigh = value
	} else if value < b.holdsLow {
		b.holdsHigh = b.holdsLow
		b.holdsLow = value
	}
	return b
}

func main() {
	data, _ := ioutil.ReadFile("input.txt")

	// First build all the outputs.
	outputSet := make(map[int]bool)
	for _, line := range strings.Split(string(data), "\n") {
		if strings.Contains(line, "low to output") && strings.Contains(line, "high to output") {
			lowOutput, _ := strconv.Atoi(strings.Fields(line)[6])
			highOutput, _ := strconv.Atoi(strings.Fields(line)[11])
			outputSet[lowOutput] = true
			outputSet[highOutput] = true
		} else if strings.Contains(line, "low to output") {
			lowOutput, _ := strconv.Atoi(strings.Fields(line)[6])
			outputSet[lowOutput] = true
		}
	}
	// Outputs are built as bots just for simplicity.
	var outputList = make([]bot, len(outputSet))
	for outputNumber, _ := range outputSet {
		var newOutput bot
		newOutput.number = outputNumber
		newOutput.isOutput = true
		outputList[outputNumber] = newOutput
	}

	// Populate the botSet to determine which bot numbers to make.
	// Use a map in lieu of a set.
	botSet := make(map[int]bool)
	for _, line := range strings.Split(string(data), "\n") {
		if line == "" {
			continue
		}
		if line[0] == 'b' {
			original, _ := strconv.Atoi(strings.Fields(line)[1])
			lowTo, _ := strconv.Atoi(strings.Fields(line)[6])
			highTo, _ := strconv.Atoi(strings.Fields(line)[11])
			botSet[original] = true
			botSet[lowTo] = true
			botSet[highTo] = true

		} else if line[0] == 'v' {
			original, _ := strconv.Atoi(strings.Fields(line)[5])
			botSet[original] = true
		}
	}
	// Actually build the list of bots.
	var botList = make([]bot, len(botSet))
	for botNumber, _ := range botSet {
		var newBot bot
		newBot.number = botNumber
		newBot.holdsLow = 0
		newBot.holdsHigh = 0
		botList[botNumber] = newBot
	}
	// Setup the connections between bots and record the values they hold at the start.
	for _, line := range strings.Split(string(data), "\n") {
		if line == "" {
			continue
		}
		if line[0] == 'b' && strings.Contains(line, "low to bot") {
			botNumber, _ := strconv.Atoi(strings.Fields(line)[1])
			givesLowTo, _ := strconv.Atoi(strings.Fields(line)[6])
			botList[botNumber].givesLowTo = &botList[givesLowTo]
		} else if line[0] == 'b' && strings.Contains(line, "low to output") {
			botNumber, _ := strconv.Atoi(strings.Fields(line)[1])
			givesLowTo, _ := strconv.Atoi(strings.Fields(line)[6])
			botList[botNumber].givesLowTo = &outputList[givesLowTo]
		}
		if line[0] == 'b' && strings.Contains(line, "high to bot") {
			botNumber, _ := strconv.Atoi(strings.Fields(line)[1])
			givesHighTo, _ := strconv.Atoi(strings.Fields(line)[11])
			botList[botNumber].givesHighTo = &botList[givesHighTo]
		} else if line[0] == 'b' && strings.Contains(line, "high to output") {
			botNumber, _ := strconv.Atoi(strings.Fields(line)[1])
			givesHighTo, _ := strconv.Atoi(strings.Fields(line)[11])
			botList[botNumber].givesHighTo = &outputList[givesHighTo]
		}
		if line[0] == 'v' {
			value, _ := strconv.Atoi(strings.Fields(line)[1])
			botNumber, _ := strconv.Atoi(strings.Fields(line)[5])
			botList[botNumber] = botList[botNumber].receive(value)
		}
	}
	// Now walk through the bots and redistribute values until we compare the targets.
Outer:
	for {
	Inner:
		for _, botInstance := range botList {
			if botInstance.holdsLow > 0 && botInstance.holdsHigh > 0 {
				lowBot := *botInstance.givesLowTo
				lowBot = lowBot.receive(botInstance.holdsLow)
				if lowBot.isOutput {
					outputList[lowBot.number] = lowBot
				} else {
					botList[lowBot.number] = lowBot
				}
				highBot := *botInstance.givesHighTo
				highBot = highBot.receive(botInstance.holdsHigh)
				if highBot.isOutput {
					outputList[highBot.number] = highBot
				} else {
					botList[highBot.number] = highBot
				}
				botInstance.holdsLow = 0
				botInstance.holdsHigh = 0
				botList[botInstance.number] = botInstance
				break Inner
			}
		}
		if outputList[0].holdsLow*outputList[1].holdsLow*outputList[2].holdsLow != 0 {
			fmt.Println("Part 2 Solution:", outputList[0].holdsLow*outputList[1].holdsLow*outputList[2].holdsLow)
			break Outer
		}
	}
}
