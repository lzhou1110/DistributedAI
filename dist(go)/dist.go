package main

import (
	"encoding/csv"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for {
		tmp := z - (z*z-x)/(2*z)
		if tmp == z || math.Abs(tmp-z) < 0.000000000001 {
			break
		}
		z = tmp
	}
	return z
}

func CalculateDist(fileName1 string, fileName2 string) {

	//read file 1
	readFile1, _ := os.OpenFile(fileName1, os.O_RDONLY, 0666)
	reader1 := csv.NewReader(readFile1)
	reader1.Read()
	//read file 2
	readFile2, _ := os.OpenFile(fileName2, os.O_RDONLY, 0666)
	reader2 := csv.NewReader(readFile2)
	reader2.Read()

	data1, err1 := reader1.Read()
	data2, err2 := reader2.Read()

	dateTemp1 := data1[2]
	dateTEMP1 := strings.Split(dateTemp1, " ")
	lastDate := dateTEMP1[0]
	todayValidNum := 0
	todaySum := 0.0

	siteName1 := data1[0]
	siteName2 := data2[0]
	DistfileName := siteName1 + siteName2 + "_distance.csv"
	write_file, _ := os.OpenFile(DistfileName, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	writer := csv.NewWriter(write_file)
	new_content := []string{"date", "value"}
	writer.Write(new_content)
	writer.Flush()

	for err1 != io.EOF && err2 != io.EOF {
		// day is the same for two dataset
		dateTemp1 := data1[2]
		dateTEMP1 := strings.Split(dateTemp1, " ")
		date := dateTEMP1[0]

		value1 := data1[3]
		value2 := data2[3]
		day := ""
		month := ""
		year := ""
		if value1 != "" && value2 != "" {
			value_f1, _ := strconv.ParseFloat(value1, 64)
			value_f2, _ := strconv.ParseFloat(value2, 64)
			todaySum = todaySum + (value_f1-value_f2)*(value_f1-value_f2)
			todayValidNum = todayValidNum + 1
		}
		if date != lastDate {
			if todayValidNum > 0 {

				todaySum = todaySum / float64(todayValidNum)
				todaySum = Sqrt(todaySum)

				times := strings.Split(lastDate, "/")
				day = times[0]
				month = times[1]
				year = times[2]

				new_content := []string{year + month + day, strconv.FormatFloat(todaySum, 'f', 5, 64)}
				writer.Write(new_content)
				writer.Flush()
			} else if todayValidNum == 0 {
				times := strings.Split(lastDate, "/")
				day = times[0]
				month = times[1]
				year = times[2]
				new_content := []string{year + month + day, ""}
				writer.Write(new_content)
				writer.Flush()
			}
			//fmt.Println(year+month+day)
			lastDate = date
			todaySum = 0
			todayValidNum = 0
		}
		data1, err1 = reader1.Read()
		data2, err2 = reader2.Read()
	}
	write_file.Close()

}

func main() {
	fileNames := []string{"PM10_KnC_CromwellRoad.csv", "PM10_KnC_EarlsCourtRd.csv", "PM10_KnC_NorthKen.csv", "PM10_KnC_NorthKenFDMS.csv"}

	fileNum := len(fileNames)
	for i := 0; i < fileNum; i = i + 1 {
		for j := i + 1; j < fileNum; j = j + 1 {
			CalculateDist(fileNames[i], fileNames[j])
		}
	}

}
