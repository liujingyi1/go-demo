package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	//    "path/filepath"
	//    "strings"
	"bufio"

	"github.com/tealeg/xlsx"
)

//获取指定目录下的所有文件和目录
func ListDir(dirPth string) (files []string, files1 []string, err error) {
	//fmt.Println(dirPth)
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, nil, err
	}
	PthSep := string(os.PathSeparator)
	//    suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写

	for _, fi := range dir {

		if fi.IsDir() { // 忽略目录
			files1 = append(files1, dirPth+PthSep+fi.Name())
			ListDir(dirPth + PthSep + fi.Name())
		} else {
			//fmt.Println("s")
			files = append(files, dirPth+PthSep+fi.Name())
		}
	}
	return files, files1, nil
}

func getStdinInput(hint string) string {
	fmt.Print(hint)
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		return scanner.Text()
	}
	return ""
}

func findColByTitle(sheet *xlsx.Sheet, title string) int {
	titleRow := sheet.Rows[0]
	for titleIndex, col := range titleRow.Cells {
		if col.String() == title {
			return titleIndex
		}
	}

	return -1
}

func splitByDot(s rune) bool {
	if s == '.' {
		return true
	} else {
		return false
	}
}

func splitBy_(s rune) bool {
	if s == '_' {
		return true
	} else {
		return false
	}
}

type fileSort struct {
	filepath string
	num int
}

type fileSorts []fileSort

//Len()
func (s fileSorts) Len() int {
	return len(s)
}

//Less():从低到高排序
func (s fileSorts) Less(i, j int) bool {

	return s[i].num < s[j].num
}

//Swap()
func (s fileSorts) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func makeFileStruct(fileList []string) (fs fileSorts)  {
	for _, filename := range fileList {
		filesplits := strings.FieldsFunc(filename, splitByDot)
		if strings.Compare("csv", filesplits[len(filesplits)-1]) == 0 {
			var temp fileSort
			temp.filepath = filename
			_, filename := filepath.Split(filename)
			filenamesplits := strings.FieldsFunc(filename, splitBy_)
			atoinumber,_ := strconv.Atoi(filenamesplits[0])
			temp.num = atoinumber

			//fmt.Println(filenamesplits)
			fs = append(fs, temp)
		}
	}
	return fs
}

func PathExists(path string) (exist bool, isDir bool, err error) {
	fi, err := os.Stat(path)
	if err == nil {
		return true, fi.IsDir(), nil
	}
	if os.IsNotExist(err) {
		return false, false,nil
	}
	return false,false, err
}

func help() {
	fmt.Println("====================================================")
	fmt.Println("Usage for tool:")
	fmt.Println("         ./tool {directory path} {column} {.....})")
	fmt.Println("====================================================")

}

func main() {

	args := os.Args
	if len(args) < 3 || args == nil {
		help()
		return
	}

	absolutePath := args[1]
	columns := args[2: len(args)]
	fmt.Println(columns)

	if exist, isDir, _ := PathExists(absolutePath); !exist || !isDir{
		fmt.Printf("---- %s is not directory or not found ----\n", absolutePath)
		fmt.Println("---- check your directory at arg[1] ----")
		return
	}

	fmt.Println(absolutePath)

	start := time.Now()

	files, files1, _ := ListDir(absolutePath)

	for _, table := range files1 {
		temp, _, _ := ListDir(table)
		for _, temp1 := range temp {
			files = append(files, temp1)
		}
	}

	rawPath, _ := filepath.Split(absolutePath)
	fmt.Println(rawPath)
	//fmt.Println(rawFilename)

	fileSortsObject := makeFileStruct(files)

	sort.Sort(fileSortsObject)

	writeFile22(fileSortsObject, columns, absolutePath)

	cost := time.Since(start)
	fmt.Printf("\ncost=[%s]\n",cost)

	fmt.Printf("find your file at [%s]\n", rawPath+"/newFile.csv")
}

func findRow(source []string, dest []string) bool {
	for _, sou := range source  {
		for _,des := range dest {
			if strings.Compare(sou, des) == 0 {
				return true
			}
		}
	}
	return false
}

func writeFile22(fileList fileSorts, selectValue []string, absolutePath string) {

	newFile, err := os.OpenFile(absolutePath+"/../newFile.csv", os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Print("error....")
	}

	w := csv.NewWriter(newFile)

	tempSelect := []string{"",""}
	for _, aaa := range selectValue {
		tempSelect = append(tempSelect,aaa)
	}
	w.Write(tempSelect)
	for _, table1 := range fileList {
		fmt.Println(table1)

		cntb, err := ioutil.ReadFile(table1.filepath)
		if err != nil {
			fmt.Printf("error %s", err)
		}

		r2 := csv.NewReader(strings.NewReader(string(cntb)))
		ss, err := r2.ReadAll()
		if err != nil {
			fmt.Printf("err %s", err)
		}
		index := 0
		var newRow []string

		_, filenamewrite := filepath.Split(table1.filepath)
		newRow = append(newRow, filenamewrite, "")

		for _, row := range ss {

			foundRow := findRow(row,selectValue)
			if foundRow {
				index++
			}
			if index > len(selectValue) && foundRow {
				newRow = append(newRow, row[2]+" "+row[3])
			}
		}

		w.Write(newRow)
	}
	w.Flush()
	newFile.Close()
}
