package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"os"
	"io/ioutil"
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
		"HDMI cable",
		"For my STEAM Link",
		1,
	}
	th := TransactionHistory{}
	th.Transactions = []Transaction{tr}

	th.Write(f)

	dat, err := ioutil.ReadFile("/tmp/sample_box.json")
	if err != nil {
		panic(err)
	}

	// Check what was written into the file
	assert.Equal(
		t,
		"{\""+
			"transactions\":["+
			"{"+
			"\"date\":\"2018/10/24\","+
			"\"reason\":\"Moved in\","+
			"\"itemName\":\"HDMI cable\","+
			"\"itemDescription\":\"For my STEAM Link\","+
			"\"amount\":1"+ "}"+
			"]"+
			"}", string(dat), "They should be equal")
}