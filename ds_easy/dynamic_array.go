package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the dynamicArray function below.
 */
func dynamicArray(N int32, queries [][]int32) []int32 {

	var lastAnswer int32

	var sequenceOne []int32
	var sequenceTwo []int32

	var useSequence = [][]int32{sequenceOne, sequenceTwo}

	var returnSeq []int32

	for _, query := range queries {
		queryType := query[0]
		x := query[1]
		y := query[2]

		sequenceNum := (x ^ lastAnswer) % N

		if queryType == 1 {
			useSequence[sequenceNum] = append(useSequence[sequenceNum], y)
		}

		if queryType == 2 {
			elementIndex := y % int32(len(useSequence[sequenceNum]))
			lastAnswer = useSequence[sequenceNum][elementIndex]
			returnSeq = append(returnSeq, lastAnswer)
		}

	}

	return returnSeq
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nq := strings.Split(readLine(reader), " ")

	nTemp, err := strconv.ParseInt(nq[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	qTemp, err := strconv.ParseInt(nq[1], 10, 64)
	checkError(err)
	q := int32(qTemp)

	var queries [][]int32
	for queriesRowItr := 0; queriesRowItr < int(q); queriesRowItr++ {
		queriesRowTemp := strings.Split(readLine(reader), " ")

		var queriesRow []int32
		for _, queriesRowItem := range queriesRowTemp {
			queriesItemTemp, err := strconv.ParseInt(queriesRowItem, 10, 64)
			checkError(err)
			queriesItem := int32(queriesItemTemp)
			queriesRow = append(queriesRow, queriesItem)
		}

		if len(queriesRow) != int(3) {
			panic("Bad input")
		}

		queries = append(queries, queriesRow)
	}

	result := dynamicArray(n, queries)

	for resultItr, resultItem := range result {
		fmt.Fprintf(writer, "%d", resultItem)

		if resultItr != len(result)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
