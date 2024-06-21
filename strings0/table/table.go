// three-legged-crow.zero
//
// @(#)table.go  Tuesday, May 09th, 2023
// Copyright(c) 2023, Leyton Goth. All rights reserved.

package table

import (
	"fmt"
	"strings"
)

// 表格分割符
var (
	HA = []string{"┌", "╔"} // 表头上左
	HB = []string{"┬", "╦"} // 表头上中
	HC = []string{"┐", "╗"} // 表头上右

	RowSeparator    = []string{"─", "═"} // 行分隔符
	ColumnSeparator = []string{"│", "║"} // 列分割符

	RA = []string{"├", "╠"} // 行左
	RB = []string{"┼", "╬"} // 行中
	RC = []string{"┤", "╣"} // 行右

	FA = []string{"└", "╚"} // 表底下左
	FB = []string{"┴", "╩"} // 表底下中
	FC = []string{"┘", "╝"} // 表底下右
)

type separatorIndex int // 分割符下标

const (
	singleLine separatorIndex = 0 // 单线
	doubleLine separatorIndex = 1 // 双线
)

// Table 输出文本版表格
//
//	┌───┬───┐
//	│   │   │
//	├───┼───┤
//	│   │   │
//	└───┴───┘
type Table struct {
	column         int
	header         []string
	rows           [][]string
	footer         []string
	columnsWidth   []int          // 标注每个列的宽度
	prettyLines    []string       // 打印输出使用
	separatorIndex separatorIndex // 当前使用分割符下标
}

// New 创建一个新表对象
func New(column int) *Table {
	if column <= 0 {
		column = 1
	}
	return &Table{
		column:         column,
		columnsWidth:   make([]int, column), // 每个列的宽度
		separatorIndex: singleLine,          // 默认单线
	}
}

// UseDoubleLine 使用双线画表格
//
//	╔═══╦═══╗
//	║   ║   ║
//	╠═══╬═══╣
//	║   ║   ║
//	╚═══╩═══╝
func (t *Table) UseDoubleLine() *Table {
	t.separatorIndex = doubleLine
	return t
}

// AddHeader 添加表头, 如果已经添加过, 再次添加会被新数据覆盖
func (t *Table) AddHeader(cells ...interface{}) *Table {
	if row := t.fillRow(cells...); row != nil {
		t.header = row
	}
	return t
}

// AppendRow 向表格尾部添加一行数据
func (t *Table) AppendRow(cells ...interface{}) *Table {
	row := t.fillRow(cells...)
	if row == nil {
		return t
	}
	t.rows = append(t.rows, row)
	return t
}

// AddFooter 添加表尾, 如果已经添加过, 再次添加会被新数据覆盖
func (t *Table) AddFooter(cells ...interface{}) *Table {
	if row := t.fillRow(cells...); row != nil {
		t.footer = row
	}
	return t
}

// AppendRows 添加多行
func (t *Table) AppendRows(rows [][]string) *Table {
	if len(rows) == 0 {
		return t
	}
	for i := 0; i < len(rows); i++ {
		min := len(rows[i])
		if t.column < min {
			min = t.column
		}
		row := make([]string, t.column)
		for j := 0; j < min; j++ {
			row[j] = rows[i][j]
			if len(row[j]) > t.columnsWidth[j] {
				t.columnsWidth[j] = len(row[j])
			}
		}
		t.rows = append(t.rows, row)
	}
	return t
}

// Print 打印到控制台
func (t *Table) Print() {
	t.prettyFill()
	for i := 0; i < len(t.prettyLines); i++ {
		fmt.Println(t.prettyLines[i])
	}
}

func (t *Table) prettyFill() {
	t.prettyLines = make([]string, 0, len(t.rows)+4) // 4: 表格可能表头, 表尾, 还有分割符, 一次性申请完成
	t.prettyLines = append(t.prettyLines, t.generateTopLine())
	if len(t.header) > 0 {
		t.prettyLines = append(t.prettyLines, t.generateRowLine(t.header))
		if len(t.rows) > 0 {
			t.appendRowSeparatorLine()
		}
	}
	for row := 0; row < len(t.rows); row++ {
		t.prettyLines = append(t.prettyLines, t.generateRowLine(t.rows[row]))
	}
	if len(t.footer) > 0 {
		if (len(t.header) > 0 && len(t.rows) == 0) || len(t.rows) > 0 {
			t.appendRowSeparatorLine()
		}
		t.prettyLines = append(t.prettyLines, t.generateRowLine(t.footer))
	}
	t.prettyLines = append(t.prettyLines, t.generateBottomLine())
}

func (t *Table) generateTopLine() string {
	line := make([]string, 2*t.column+1)
	line[0] = HA[t.separatorIndex]
	for i := 0; i < t.column; i++ {
		line[2*i+1] = strings.Repeat(RowSeparator[t.separatorIndex], t.columnsWidth[i]+2)
		line[2*i+2] = HB[t.separatorIndex]
	}
	line[2*t.column] = HC[t.separatorIndex]
	return strings.Join(line, "")
}

func (t *Table) generateBottomLine() string {
	line := make([]string, 2*t.column+1)
	line[0] = FA[t.separatorIndex]
	for i := 0; i < t.column; i++ {
		line[2*i+1] = strings.Repeat(RowSeparator[t.separatorIndex], t.columnsWidth[i]+2)
		line[2*i+2] = FB[t.separatorIndex]
	}
	line[2*t.column] = FC[t.separatorIndex]
	return strings.Join(line, "")
}

func (t *Table) generateRowLine(cells []string) string {
	line := make([]string, 2*t.column+1)
	line[0] = ColumnSeparator[t.separatorIndex]
	for i := 0; i < t.column; i++ {
		line[2*i+1] = padSpacesOnBothSides(cells[i], t.columnsWidth[i]+2)
		line[2*i+2] = ColumnSeparator[t.separatorIndex]
	}
	line[2*t.column] = ColumnSeparator[t.separatorIndex]
	return strings.Join(line, "")
}

func (t *Table) appendRowSeparatorLine() {
	line := make([]string, 2*t.column+1)
	line[0] = RA[t.separatorIndex]
	for i := 0; i < t.column; i++ {
		line[2*i+1] = strings.Repeat(RowSeparator[t.separatorIndex], t.columnsWidth[i]+2)
		line[2*i+2] = RB[t.separatorIndex]
	}
	line[2*t.column] = RC[t.separatorIndex]
	t.prettyLines = append(t.prettyLines, strings.Join(line, ""))
}

func (t *Table) fillRow(cells ...interface{}) []string {
	if len(cells) == 0 {
		return nil
	}
	min := len(cells)
	if t.column < min {
		min = t.column
	}
	row := make([]string, t.column)
	for i := 0; i < min; i++ {
		row[i] = fmt.Sprintf("%v", cells[i])
		if len(row[i]) > t.columnsWidth[i] {
			t.columnsWidth[i] = len(row[i])
		}
	}
	return row
}

// PrettyString 美化输出
func (t *Table) PrettyString() []string {
	t.prettyFill()
	return t.prettyLines
}

func padSpacesOnBothSides(text string, lineLen int) string {
	if len(text) >= lineLen {
		return text
	}
	if len(text) == 0 {
		return strings.Repeat(" ", lineLen)
	}
	var left = (lineLen - len([]rune(text))) / 2
	var right = lineLen - len([]rune(text)) - left
	return fmt.Sprintf("%s%s%s", strings.Repeat(" ", left), text, strings.Repeat(" ", right))
}
