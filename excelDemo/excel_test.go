package main

import (
	"encoding/csv"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestDog(t *testing.T) {
	newFile, err := os.OpenFile("F:/yongyu/jieguo/new.csv", os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		t.Log("error....")
	}
	w := csv.NewWriter(newFile)
	//w.Write([]string{"123", "234234", "345345", "234234"})

	cntb, err := ioutil.ReadFile("F:/yongyu/jieguo/03__rec_tra.ctan.csv")
	if err != nil {
		t.Logf("error %s", err)
	}
	r2 := csv.NewReader(strings.NewReader(string(cntb)))
	ss, err := r2.ReadAll()

	t.Log(err)

	for _, row := range ss {
		for _, value := range row {
			if strings.Compare("Total VOI volume", value) == 0 {
				t.Log(value)
				w.Write(row)
				break
			}
		}
	}

	w.Flush()
	newFile.Close()
}


func TestAAA(t *testing.T) {
	var a = 01
	var b = 022
	if a > b {
		t.Log(true)
	} else {
		t.Log(false)
	}
}