// three-legged-crow.zero
//
// @(#)table_test.go  Tuesday, May 09th, 2023
// Copyright(c) 2023, Leyton Goth. All rights reserved.

package table

import (
	"testing"
)

// TestPrint1 ...
func TestPrint1(t *testing.T) {
	tab := New(0)
	tab.Print()
}

// TestPrint2 ...
func TestPrint2(t *testing.T) {
	tab := New(1)
	tab.AddHeader("only", "header", "only!")
	tab.Print()
}

// TestPrint3 ...
func TestPrint3(t *testing.T) {
	tab := New(2)
	tab.AppendRow("only", "one row", "only!")
	tab.Print()
}

// TestPrint4 ...
func TestPrint4(t *testing.T) {
	tab := New(3)
	tab.AddFooter("only", "footer", "only!")
	tab.Print()
}

// TestPrint5 ...
func TestPrint5(t *testing.T) {
	tab := New(9)
	tab.AddHeader("only", "footer", "only!")
	tab.AppendRow(1, 2, 3, 4, 5, "", "", 12345)
	tab.Print()
}

// TestPrint6 ...
func TestPrint6(t *testing.T) {
	tab := New(3)
	tab.AddFooter("only", "footer", "only!")
	tab.AppendRow(1, 2, 3, 4, 5)
	tab.Print()
}

// TestPrint7 ...
func TestPrint7(t *testing.T) {
	tab := New(4)
	tab.AddHeader("this", "is", "table's", "header")
	tab.AppendRow("hey", "row")
	tab.AddFooter("yes", "footer")
	tab.Print()
}

// // TestPrint8 ...
func TestPrint8(t *testing.T) {
	tab := New(5)
	tab.AddHeader("Hello", "World", "This", "Is", "My", "First", "Table").AddFooter("No", "No~", "Y")
	tab.AppendRow("A", 1, 3, "xxx", "b", "")
	tab.AppendRow("B", 4, 8, "yy", "zz")
	tab.AppendRows([][]string{
		{"a", "b", "", "b", "", "b", ""},
		{"a3 this is a long long text row", "b3", "b", "", "b", "", "b", "", "b", ""},
		{"a4", "b4", "this ia an another long long row", "ba4", "aaaa4", "b4fff"},
	})
	tab.AppendRows([][]string{
		{},
		{},
		{},
	})
	tab.AddFooter("x")
	tab.Print()
}

// TestPrint9 ...
func TestPrint9(t *testing.T) {
	tab := New(8)
	tab.AddHeader("th☂is", "is", "tab☓☓☓☓le's", "header")
	tab.AppendRow("hey", "ro☓w")
	tab.AppendRows([][]string{
		{"☓", "", "", "", "◎◎", "", "", "", "", ""},
		{"111111333333"},
		{},
	})
	tab.AddFooter("yes", "footer", "☛", "☛☛", "", "☛")
	tab.Print()
}

// TestPrint10 ...
func TestPrint10(t *testing.T) {
	tab := New(2).UseDoubleLine()
	tab.AddHeader(" ", " ")
	tab.AddFooter(" ")
	tab.Print()
}

// TestPrint11 ...
func TestPrint11(t *testing.T) {
	tab := New(8).UseDoubleLine()
	tab.AddHeader("th☂is", "is", "tab☓☓☓☓le's", "header")
	tab.AppendRow("hey", "ro☓w")
	tab.AppendRows([][]string{
		{"☓", "", "", "", "◎◎", "", "", "", "", ""},
		{"111111333333"},
		{},
	})
	tab.AddFooter("yes", "footer", "☛", "☛☛", "", "☛")
	tab.Print()
}

// TestPrint12 ...
func TestPrint12(t *testing.T) {
	tab := New(1)
	tab.AppendRows([][]string{
		{"SUCCESS"},
		{},
		{"Merge Request was created or updated:"},
		{"Writing objects: 100% (7/7), 759 bytes | 189.00 KiB/s, done."},
	})
	tab.Print()
}
