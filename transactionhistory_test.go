package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"os"
	"io/ioutil"
	"encoding/json"
)

func TestWriteTransactionHistoryToFile(t *testing.T) {
	f, err := os.Create("/tmp/sample_box.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	tr := Transaction{
		"2018/10/24",
		"Moved in",
		Item{
			"HDMI cable",
			"For my STEAM Link",
		},
		1,
	}
	th := TransactionHistory{}
	th.Transactions = []Transaction{tr}

	th.Write(f)

	dat, err := ioutil.ReadFile("/tmp/sample_box.json")
	if err != nil {
		panic(err)
	}

	b, _ := json.Marshal(th)

	// Check what was written into the file
	assert.Equal(t, string(b), string(dat), "They should be equal")
}

func TestReadTransactionHistoryToFile(t *testing.T) {
	f, err := os.Create("/tmp/sample_box.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	tr := Transaction{
		"2018/10/24",
		"Moved in",
		Item{
			"HDMI cable",
			"For my STEAM Link",
		},
		1,
	}
	th := TransactionHistory{
		[]Transaction{tr},
	}
	th.Write(f)

	f2, err := os.Open("/tmp/sample_box.json")
	if err != nil {
		panic(err)
	}
	defer f2.Close()

	fileInfo, err := f2.Stat()
	if err != nil {
		panic(err)
	}

	newTh := TransactionHistory{[]Transaction{}}
	newTh.Read(f2, fileInfo.Size())



	assert.Equal(t, th, newTh)
}
