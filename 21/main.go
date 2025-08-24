package main

import "fmt"

// старая система
type OldSys struct{}

func (o *OldSys) OldMethod() string {
	return "старый метод"
}

// новый интерфейс
type NewInterface interface {
	NewMethod() string
}

// адаптер
type Adapter struct {
	old *OldSys
}

func (a *Adapter) NewMethod() string {
	return "адаптер работает: " + a.old.OldMethod()
}

// клиент
type Client struct{}

func (c *Client) UseInterface(n NewInterface) {
	fmt.Println(n.NewMethod())
}

func main() {
	old := &OldSys{}
	adapter := &Adapter{old: old}
	client := &Client{}

	client.UseInterface(adapter)
}
