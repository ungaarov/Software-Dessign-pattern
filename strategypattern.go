package main

import "fmt"

type SendData interface {
	Send()
}

type Macbook struct{}

func (spm Macbook) Send() {
	fmt.Println("The data from the macbook was successfully saved")
}

type Applewatch struct{}

func (sum Applewatch) Send() {
	fmt.Println("Data in apple watch is synchronized")
}

type Iphone struct{}

func (sum Iphone) Send() {
	fmt.Println("The data from the iPhone was saved in the cloud")
}

type icloud struct {
	SendData SendData
}

func NewIcloud(dr SendData) *icloud {
	return &icloud{
		SendData: dr,
	}
}

func (t *icloud) DataSend() {
	t.SendData.Send()
}

func (t *icloud) SetData(dr SendData) {
	t.SendData = dr
}

func main() {
	newIcloud := NewIcloud(Macbook{})
	newIcloud.DataSend()
	newIcloud.SetData(Applewatch{})
	newIcloud.DataSend()
	newIcloud.SetData(Iphone{})
	newIcloud.DataSend()
}
