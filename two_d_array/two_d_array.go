package two_d_array

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the hourglassSum function below.
func hourglassSum(arr [][]int32) int32 {
	maxX := 3
	maxY := 3
	var total int32 = -64

	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			// sum the top of hourglass
			sum := arr[y][x] + arr[y][x+1] + arr[y][x+2]

			// get the middle of hourglass
			sum += arr[y+1][x+1]

			// sum the bottom of hourglass
			sum += arr[y+2][x] + arr[y+2][x+1] + arr[y+2][x+2]

			// don't store result to keep space complexity down
			if total < sum {
				total = sum
			}
		}
	}

	return total
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	var arr [][]int32
	for i := 0; i < 6; i++ {
		arrRowTemp := strings.Split(readLine(reader), " ")

		var arrRow []int32
		for _, arrRowItem := range arrRowTemp {
			arrItemTemp, err := strconv.ParseInt(arrRowItem, 10, 64)
			checkError(err)
			arrItem := int32(arrItemTemp)
			arrRow = append(arrRow, arrItem)
		}

		if len(arrRow) != int(6) {
			panic("Bad input")
		}

		arr = append(arr, arrRow)
	}

	result := hourglassSum(arr)

	fmt.Fprintf(writer, "%d\n", result)

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
