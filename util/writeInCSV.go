package util

import (
	"../class"
	"encoding/csv"
	"os"
	"strconv"
)

func WriteInCSV(data [][]int, time []float64, load []float64, sigma []float64, filename string) {
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM

	w := csv.NewWriter(f)
	w.Write([]string{"Tasks Number Limit", "Migration Number", "Upload Number", "Amount of Tasks",
		"Violation Tasks", "Execute Time", "Load", "Sigma"})
	for i := 0; i < len(data); i++ {
		w.Write([]string{strconv.Itoa(class.TASK_NUMBER_UPPER_LIMIT[i]), strconv.Itoa(data[i][0]), strconv.Itoa(data[i][1]),
			strconv.Itoa(data[i][2]), strconv.Itoa(data[i][3]), strconv.FormatFloat(time[i], 'f', 1, 64),
			strconv.FormatFloat(load[i], 'f', 3, 64), strconv.FormatFloat(sigma[i], 'f', 3, 64)})
	}
	w.Flush()
}
