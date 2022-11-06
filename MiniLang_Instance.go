package main

import (
	"strings"
)

func RemoveFirstChar(input string) string {
	if len(input) <= 1 {
		return ""
	}
	return input[1:]
}

func MinilangInstancev2(filename string) {

	list := ReadByLine(filename)

	for line_num, line_val := range list {

		var Mode string
		var column_num2 int

		var Line_Func string
		var Line_Params []string
		var Line_Var string

		var Func_temp []string
		var Params_temp []string
		var Var_temp []string

		for column_num, column_val := range strings.Split(line_val, "") {

			if column_val == "(" && Mode != "param" {

				Mode = "param"

			}

			if column_val == "-" && Mode == "param" {

				Mode = "var"

			}

			if column_val == "_" && Mode == "var" {

				Mode = "param"

			}

			if Mode == "var" && Mode != "param" && column_val != "(" && column_val != ")" && column_val != `"` && column_val != "_" && column_val != "-" {

				Var_temp = append(Var_temp, column_val)

			}

			if Mode != "param" && Mode != "var" && column_val != "(" && column_val != ")" {

				Func_temp = append(Func_temp, column_val)

			}

			if Mode == "param" && column_val != "(" && column_val != ")" && column_val != `"` || Mode == "var" && column_val != "(" && column_val != ")" && column_val != `"` {

				Params_temp = append(Params_temp, column_val)

			}

			column_num2 = column_num

		}
		Line_Func = strings.Join(Func_temp, "")
		if len(Variables) != 0 {
			Line_Var = Variables[strings.Join(Var_temp, "")]
			Line_Params = strings.Split(strings.ReplaceAll(strings.Join(Params_temp, ""), "-"+strings.Join(Var_temp, "")+"_", Line_Var), ",")
		}
		if len(Variables) == 0 {
			Line_Params = strings.Split(strings.Join(Params_temp, ""), ",")
		}
		ProcessFunc(filename, line_num, column_num2, Line_Func, Line_Params)

	}

}

func main() {

	MinilangInstancev2("New.mini")

}
