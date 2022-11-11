package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/maja42/goval"
)

var Variables = make(map[string]string)
var PreFunc = make(map[string]string)

func Mlog_error(params ...string) {

	PrintLnRed(strings.Join(params, ""))

}

func Mlog(params ...string) {

	fmt.Println(strings.Join(params, ""))

}

func SpitError(line, column int, err string) {

	Mlog_error(fmt.Sprintf(`Line: %d Column: %d | %s`, line+1, column+1, err))
	os.Exit(1)

}

func SpitConsoleError(err string) {

	Mlog_error(err)
	os.Exit(2)
}

func ProcessFunc(filename string, line, column int, funct string, params []string) {

	for i, v := range params {
		if strings.Contains(v, "-") && strings.Contains(v, "_") {
			SpitError(line, column, fmt.Sprintf(`Undefined Variable in param %d`, i+1))
		}
		params[i] = strings.Trim(v, " ")
	}

	if funct == "Grab" {

		if params[0]+".mini" != filename {
			MinilangInstancev2(params[0] + ".mini")

		} else {

			SpitError(line, column, fmt.Sprintf(`Mentioned File "%s.mini" cannot be the same as %s`, params[0], filename))
		}
	} else if funct == "Var" {

		if len(params) == 2 {
			Variables[params[0]] = strings.Trim(params[1], " ")
		} else {
			SpitError(line, column, fmt.Sprintf(`Func %s Has too little or too many paramiters`, funct))
		}
	} else if funct == "log" {

		PrintLnWhite(strings.Join(params, ""))

	} else if funct == "log_error" {

		Mlog_error(params...)

	} else if funct == "cmd" {

		p := "/C," + strings.Join(params, ",")
		Nf, _ := exec.Command("cmd", strings.Split(p, ",")...).Output()
		PrintLnGreen(fmt.Sprintf(`CMD: %s`, string(Nf)))

	} else if funct == "PreFunc" {

		if len(params) == 2 {
			PreFunc[params[0]] = strings.Trim(params[1], " ")
		} else {
			SpitError(line, column, fmt.Sprintf(`Func %s Has too little or too many paramiters`, funct))
		}

	} else if funct == "end" {

		os.Exit(0)

	} else if funct == "if" {

		os.Exit(0)

	} else if funct == "elif" {

		os.Exit(0)

	} else if funct == "EvalExpressions" {

		if len(params) == 1 {
			eval := goval.NewEvaluator()

			result, _ := eval.Evaluate(params[0], nil, nil)

			PrintLnYellow(fmt.Sprintf(`Eval: %v`, result))

		} else if len(params) == 2 {
			eval := goval.NewEvaluator()

			result, _ := eval.Evaluate(params[0], nil, nil)

			PrintLnYellow(fmt.Sprintf(`Eval: %v`, result))

			Variables[params[1]] = fmt.Sprintf(`%v`, result)

		} else {
			SpitError(line, column, fmt.Sprintf(`Func %s Has too little or too many paramiters`, funct))
		}

	} else if funct == "" {

		// Spaces can pass

	} else {
		if len(PreFunc) != 0 {
			l := 0
			for k, v := range PreFunc {
				if strings.Contains(k, funct) {
					MinilangInstancev2_Text(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(v, "{", "("), "}", ")"), "!", ","))
				} else if l == len(PreFunc) {
					SpitError(line, column, fmt.Sprintf(`Func %s Does Not Exist`, funct))
				}
				l += 1
			}
		} else {
			SpitError(line, column, fmt.Sprintf(`Func %s Does Not Exist`, funct))
		}

	}

}

func MinilangInstancev2_Text(Text string) {

	l := make([]string, 0)
	list := append(l, Text)

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
		ProcessFunc("MiniLang_func.go", line_num, column_num2, Line_Func, Line_Params)

	}

}
