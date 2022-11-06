package main

import (
	"fmt"
	"strconv"
	"strings"
)

var variables []string

func ProcessFunction(funct string, param string, line int, column int) {

	if funct == "namespace" {
		CreateDir(param)
		WriteFile(param+"/namespace.json", "{\n "+`"packageName":"`+param+`"`+"\n }")
	} else if funct == "Grab" {
		MinilangInstance(param)
	} else if funct == "Var" {
		variables = append(variables, param+" ,")
	} else if funct == "log" {
		Mlog(param)
	} else if funct == "log_error" {
		Mlog_error(param)
	} else {

		Mlog_error("Line: " + strconv.Itoa(line) + " Column: " + strconv.Itoa(column) + ` | Function "` + funct + `" Does Not Exist`)

	}

}

func Mlog_error(params ...string) {

	PrintLnRed(strings.Join(params, ""))

}

func Mlog(params ...string) {

	fmt.Println(strings.Join(params, ""))

}

func MinilangInstance(param string) {

	file := ReadByLine(param + ".mini")

	for line, v := range file {

		var function string
		var M string
		var Pstring []string
		var Fstring []string
		var param string

		for column, v2 := range strings.Split(v, "") {

			if v2 == "(" && M != "param" {

				M = "param"

			}
			if M == "param" && v2 != ")" && v2 != "(" {
				Pstring = append(Pstring, v2)
			}
			if M == "param" && v2 == ")" {

				param = strings.ReplaceAll(strings.Join(Pstring, ""), `"`, "")
				function = strings.Join(Fstring, "")
				ProcessFunction(function, param, line, column)

			}
			if M != "param" {
				Fstring = append(Fstring, v2)
			}

		}

	}

}
