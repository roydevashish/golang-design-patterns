package main

import "github.com/roydevashish/golang-design-patterns/visitor/visitor"

func main() {
	clientList := []visitor.Client{
		&visitor.LawClient{Email: "law@gmail.com"},
		&visitor.RetailClient{Email: "retail@gmail.com"},
		&visitor.RestaurentClient{Email: "restaurent@gmail.com"},
	}
	for _, v := range clientList {
		v.Accept(&visitor.EmailVisitor{})
		v.Accept(&visitor.PDFExportVisitor{})
	}
}
