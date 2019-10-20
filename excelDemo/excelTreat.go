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

func split(s rune) bool {
	if s == '.' {
		return true
	} else {
		return false
	}
}

func split222(s rune) bool {
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

		filesplits := strings.FieldsFunc(filename, split)
		if strings.Compare("csv", filesplits[len(filesplits)-1]) == 0 {
			var temp fileSort
			temp.filepath = filename
			_, filename := filepath.Split(filename)
			filenamesplits := strings.FieldsFunc(filename, split222)
			atoinumber,_ := strconv.Atoi(filenamesplits[0])
			temp.num = atoinumber

			//fmt.Println(filenamesplits)
			fs = append(fs, temp)
		}
	}
	return fs
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func main() {
	//args := os.Args
	//if len(args) < 3 || args == nil {
	//	help()
	//	return
	//}
	//
	//absolutePath := args[1]
	//columns := args[2: len(args)]
	//fmt.Println(columns)
	//
	//if exist, _ := PathExists(absolutePath); !exist{
	//	fmt.Println("---- path not found ----")
	//	fmt.Println("---- check your path at arg[1] ----")
	//}

	start := time.Now()
	//some func or operation

	files, files1, _ := ListDir("F:\\yongyu\\jieguo")

	for _, table := range files1 {
		temp, _, _ := ListDir(table)
		for _, temp1 := range temp {
			files = append(files, temp1)
		}
	}

	fileSortsObject := makeFileStruct(files)
	fmt.Println(fileSortsObject)

	sort.Sort(fileSortsObject)
	//
	fmt.Println(fileSortsObject)

	writeFile22(fileSortsObject)

	cost := time.Since(start)
	fmt.Printf("cost=[%s]",cost)


	//writeFile11(files)
}

func help() {
	fmt.Println("====================================================")
	fmt.Println("Usage for tool:")
    fmt.Println("         ./tool {path} {column} {.....})")
	fmt.Println("====================================================")

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

func writeFile22(fileList fileSorts) {
	newFile, err := os.OpenFile("F:/yongyu/new99.csv", os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Print("error....")
	}
	w := csv.NewWriter(newFile)
	selectValue := []string{"Total VOI volume", "Object volume", "Percent object volume", "Object surface density","Structure model index", "Structure thickness", "Structure linear density"}
	selectUnit := []string{"TV", "Obj.V", "Obj.V/TV", "Obj.S/TV","SMI", "St.Th", "St.Li.Dn"}

	tempSelect := []string{"",""}
	tempSelectUtil := []string{"",""}
	for _, aaa := range selectValue {
		tempSelect = append(tempSelect,aaa)
	}
	for _, aaa := range selectUnit {
		tempSelectUtil = append(tempSelectUtil,aaa)
	}
	w.Write(tempSelect)
	w.Write(tempSelectUtil)
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

			//foundRow := false
			//for _, value := range row {
			//	found := false
			//	for _, sel := range selectValue {
			//		if strings.Compare(sel, value) == 0 {
			//			//fmt.Println("found")
			//			index++
			//			found = true
			//			foundRow = true
			//			break
			//		}
			//	}
			//	if found {
			//		break
			//	}
			//}
			if index > len(selectValue) && foundRow {
				newRow = append(newRow, row[2]+" "+row[3])
			}
		}

		w.Write(newRow)

	}
	w.Flush()
	newFile.Close()
}

//func writeFile11(fileList []string) {
//	newFile, err := os.OpenFile("F:/yongyu/new11.csv", os.O_WRONLY|os.O_CREATE, os.ModePerm)
//	if err != nil {
//		fmt.Print("error....")
//	}
//	w := csv.NewWriter(newFile)
//
//	selectValue := []string{"Total VOI volume", "Object volume", "Percent object volume", "Structure model index", "Structure thickness", "Structure linear density"}
//
//	var mut sync.Mutex
//	var wg sync.WaitGroup
//
//	w.Write(selectValue)
//
//	for _, table1 := range fileList {
//		wg.Add(1)
//		mut.Lock()
//
//		fs := strings.FieldsFunc(table1, split)
//		if strings.Compare("csv", fs[len(fs)-1]) == 0 {
//			fmt.Println(table1)
//			cntb, err := ioutil.ReadFile(table1)
//			if err != nil {
//				fmt.Printf("error %s", err)
//			}
//
//			r2 := csv.NewReader(strings.NewReader(string(cntb)))
//			ss, err := r2.ReadAll()
//			if err != nil {
//				fmt.Printf("err %s", err)
//			}
//
//			index := 0
//			w.Write([]string{"----", "------", "------", "------", "------", table1})
//			for _, row := range ss {
//				for _, value := range row {
//					found := false
//					for _, sel := range selectValue {
//						if strings.Compare(sel, value) == 0 {
//							//fmt.Println("found")
//							index++
//							if index > len(selectValue) {
//								fmt.Println(value)
//								w.Write(row)
//							}
//							found = true
//							break
//						}
//					}
//					if found {
//						break
//					}
//				}
//			}
//		}
//		mut.Unlock()
//		wg.Done()
//		//time.Sleep(time.Duration(2) * time.Second)
//	}
//	wg.Wait()
//	w.Flush()
//	newFile.Close()
//}
