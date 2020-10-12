package metcalfechain

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/shopspring/decimal"
)

const (

	testNodeAPI = "http://ip:port"
)

func Test_getBlockHeight(t *testing.T) {

	c := NewClient(testNodeAPI, true)

	r, err := c.getBlockHeight()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("height:", r)
	}

}

func Test_getBlockByHeight(t *testing.T) {

	c := NewClient(testNodeAPI, true)
	r, err := c.getBlockByHeight(95271)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r)
	}
}


func Test_getBlockHash(t *testing.T) {

	c := NewClient(testNodeAPI, true)

	height := uint64(95271)

	r, err := c.getBlockHash(height)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r)
	}

}

func Test_getBalance(t *testing.T) {

	c := NewClient(testNodeAPI, true)

	address := "MTCrKuJyN5YceCX9vHeRHZ5bifk9RtzspXGq"

	r, err := c.getBalance(address, true, 20000000)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r)
	}

	address = "rP9YxN6yjw5HJj5LeK55gtVr8RznEPLwRc"

	r, err = c.getBalance(address, true, 20000000)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r)
	}

}

func Test_getTransaction(t *testing.T) {

	c := NewClient(testNodeAPI, true)
	txid := "CFEF14A9A3A41D8E1DC46C32F3956554A7F8E5EDD259BC489EF2308ADFB50E68"
	r, err := c.getTransaction(txid, "MemoData")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r)
	}

}

func Test_convert(t *testing.T) {

	amount := uint64(5000000001)

	amountStr := fmt.Sprintf("%d", amount)

	fmt.Println(amountStr)

	d, _ := decimal.NewFromString(amountStr)

	w, _ := decimal.NewFromString("100000000")

	d = d.Div(w)

	fmt.Println(d.String())

	d = d.Mul(w)

	fmt.Println(d.String())

	r, _ := strconv.ParseInt(d.String(), 10, 64)

	fmt.Println(r)

	fmt.Println(time.Now().UnixNano())
}

func Test_getTransactionByAddresses(t *testing.T) {
	addrs := "ARAA8AnUYa4kWwWkiZTTyztG5C6S9MFTx11"

	c := NewClient(testNodeAPI, true)
	result, err := c.getMultiAddrTransactions("MemoData", 0, -1, addrs)

	if err != nil {
		t.Error("get transactions failed!")
	} else {
		for _, tx := range result {
			fmt.Println(tx.TxID)
		}
	}
}

func Test_tmp(t *testing.T) {

	c := NewClient(testNodeAPI, true)

	block, err := c.getBlockByHeight(48059631)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(block)
}
