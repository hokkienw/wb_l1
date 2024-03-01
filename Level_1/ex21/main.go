package main

import "fmt"

type LegacyPrinter struct {
}

func (l *LegacyPrinter) Print(msg string) {
	fmt.Printf("Legacy Printer: %s\n", msg)
}

type ModernPrinter struct {
	lp LegacyPrinter
}

func (m *ModernPrinter) Print(msg string) {
	m.lp.Print(msg)
	fmt.Println("Модернизация на подходе ...")
}

func main() {
	lp := LegacyPrinter{}
	mp := ModernPrinter{lp}
	mp.Print("Hello")
}
