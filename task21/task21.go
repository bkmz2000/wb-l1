package main

import (
	"fmt"
	"strings"
)

type PrinterSettings struct {
	sep    string
	end    string
	hasSep bool
	hasEnd bool
}

func Sep(sep string) PrinterSettings {
	return PrinterSettings{sep, "", true, false}
}

func End(end string) PrinterSettings {
	return PrinterSettings{"", end, false, true}
}

func print(a ...interface{}) {
	sep := " "
	end := "\n"

	sepset := false
	endset := false

	for _, ob := range a {
		switch ob.(type) {
		case PrinterSettings:
			if ob.(PrinterSettings).hasSep {
				if !sepset {
					sep = ob.(PrinterSettings).sep
					sepset = true
				} else {
					panic("Multiple Seps")
				}
			}

			if ob.(PrinterSettings).hasEnd {
				if !endset {
					end = ob.(PrinterSettings).end
					endset = true
				} else {
					panic("Multiple Ends")
				}
			}
		}
	}

	ss := make([]string, 0)

	for i := range a {
		switch a[i].(type) {
		case PrinterSettings:
		default:
			ss = append(ss, fmt.Sprint(a[i]))
		}
	}

	fmt.Print(strings.Join(ss, sep), end)
}

func main() {
	print(1, 2, 3, Sep("#"), End("\n\n"))
	print(1, 2, 3, Sep("@"), End("\n-----\n"))
	print(1, 2, 3, Sep(""), End(""))
	print(1, 2, 3, Sep("\t"), End("\n"))
}
