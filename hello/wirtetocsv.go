package main

import (
	"encoding/csv"
	"os"
	"strconv"
)

func main() {
    f, err := os.Create("clusterings.csv")
    if err != nil {
        panic(err)
    }
    defer f.Close()

    f.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM

    w := csv.NewWriter(f)
    w.Write([]string{"X","Y","Clusterings"})
    c := GetClusterings()
    for i:=0; i<len(c); i++{
    	for j:=0; j<len(c[i]); j++{
    		w.Write([]string{strconv.Itoa(c[i][j].Lx), strconv.Itoa(c[i][j].Ly), strconv.Itoa(i)})
		}
	}
    w.Flush()
}