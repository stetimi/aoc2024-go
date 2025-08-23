package days

import (
	u "aoc2024-go/utils"
)

const space = -1

type Disk struct {
	data         []int
	spaceIndices []int
	spaceLengths []int
}

func Day9(contents []byte) u.Answers {
	diskMap := string(contents)
	part1 := day9Part1(diskMap)
	return u.Part1OnlyIntAnswers(part1)
}

func day9Part1(diskMap string) int {
	disk := toDisk(diskMap, func(index, length int) {})
	return disk.checksum(disk.defragmentByByte())
}

func toDisk(diskMap string, onSpace func(index, length int)) *Disk {
	var data []int
	var spaceIndices []int
	var spaceLengths []int
	arrayIndex := 0

	for i, ch := range diskMap {
		length := u.MustParseInt(string(ch))
		if i%2 == 0 {
			id := i / 2
			for range length {
				data = append(data, id)
			}
		} else {
			onSpace(arrayIndex, length)
			spaceIndices = append(spaceIndices, arrayIndex)
			spaceLengths = append(spaceLengths, length)
			for j := 0; j < length; j++ {
				data = append(data, space)
			}
			arrayIndex += length
		}
	}
	return &Disk{
		data:         data,
		spaceIndices: spaceIndices,
		spaceLengths: spaceLengths,
	}
}

func (d *Disk) defragmentByByte() int {
	readPtr := len(d.data) - 1
	writePtr := 0
	for writePtr < readPtr {
		for writePtr < readPtr && d.data[writePtr] != space {
			writePtr++
		}
		for readPtr > writePtr && d.data[readPtr] == space {
			readPtr--
		}
		if writePtr < readPtr {
			d.data[writePtr] = d.data[readPtr]
			d.data[readPtr] = space
			writePtr++
			readPtr--
		}
	}
	return writePtr
}

func (d *Disk) checksum(limit int) int {
	sum := 0
	for i := range min(limit, len(d.data)) + 1 {
		if d.data[i] != space {
			sum += i * d.data[i]
		}
	}
	return sum
}
