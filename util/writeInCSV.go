package util

import (
	"../class"
	"encoding/csv"
	"os"
	"strconv"
)

func WriteInCSV(data [][]int, time []float64, filename string) {
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM

	w := csv.NewWriter(f)
	w.Write([]string{"Tasks Number Limit", "Migration Number", "Upload Number", "Amount of Tasks", "Violation Tasks", "Execute Time"})
	for i := 0; i < len(data); i++ {
		w.Write([]string{strconv.Itoa(class.TASK_NUMBER_UPPER_LIMIT[i]), strconv.Itoa(data[i][0]), strconv.Itoa(data[i][1]),
			strconv.Itoa(data[i][2]), strconv.Itoa(data[i][3]), strconv.FormatFloat(time[i], 'E', -1, 64)})
	}
	w.Flush()
}
