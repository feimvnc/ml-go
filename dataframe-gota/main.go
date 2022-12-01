package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-gota/gota/dataframe"
)

func main() {
	file, err := os.Open("./2008.csv")
	if err != nil {
		log.Fatal(err)
	}

	df := dataframe.ReadCSV(file)
	fmt.Println(df) // print df

	// print dataframe info
	fmt.Println(df.Dims())
	fmt.Println(df.Ncol())
	fmt.Println(df.Nrow())
	fmt.Println(df.Names())
	fmt.Println(df.Types())

	// describe dataframe
	fmt.Println(df.Describe())

	col1 := df.Select("Country")
	fmt.Println(col1)

	row1 := df.Subset(0)
	fmt.Println(row1)

	ds := df.Col("Country")
	fmt.Println(ds.HasNaN())
	fmt.Println(ds.IsNaN())

	df = df.Arrange(
		dataframe.Sort("Total"),
	)
	fmt.Println(df)

	df = df.Filter(
		dataframe.F{1, "Total", ">", 40000},
	)
	fmt.Println(df)

	// concat 2 df
	file2, err2 := os.Open("./2009.csv")
	if err2 != nil {
		log.Fatal(err)
	}

	df2 := dataframe.ReadCSV(file2)
	fmt.Println(df2) // print df

	df = df.Concat(df2)
	fmt.Println(df)

	df = df.Arrange(
		dataframe.RevSort("Total"),
	)
	fmt.Println(df)

	file3, err := os.Open("./sample.json")
	if err != nil {
		log.Fatal(err)
	}
	df3 := dataframe.ReadJSON(file3)
	// join 2 dataframe, value of year should match
	df = df.InnerJoin(df3, "year")
	fmt.Println(df)

	fmt.Println(df.Nrow())
	fmt.Println(df.Ncol())
	fmt.Println(df.Subset(0))

}
