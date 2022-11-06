package main

import (
	"fmt"
	"os"
	"strings"
)

var Variables = make(map[string]string)

func SpitError(line, column int, err string) {

	Mlog_error(fmt.Sprintf(`Line: %d Column: %d | %s`, line+1, column+1, err))
	os.Exit(1)

}

func ProcessFunc(filename string, line, column int, funct string, params []string) {

	for i, v := range params {
		if strings.Contains(v, "-") && strings.Contains(v, "_") {
			SpitError(line, column, fmt.Sprintf(`Undefined Variable in param %d`, i+1))
		}
	}

	if funct == "Grab" {

		fmt.Println(funct, params)
		if params[0]+".mini" != filename {
			MinilangInstancev2(params[0] + ".mini")

		} else {

			SpitError(line, column, fmt.Sprintf(`Mentioned File "%s.mini" cannot be the same as %s`, params[0], filename))
		}
	} else if funct == "GrabOld" {

		fmt.Println(funct, params)
		MinilangInstance(params[0])
		fmt.Println(variables)

	} else if funct == "Var" {

		fmt.Println(funct, params)
		if len(params) == 2 {
			Variables[params[0]] = strings.Trim(params[1], " ")
			for k, v := range Variables {

				PrintLnBlue(fmt.Sprintf(`%s: %s`, k, v))

			}
		} else {
			SpitError(line, column, fmt.Sprintf(`Func %s Has too little or too many paramiters`, funct))
		}

	} else if funct == "log" {

		fmt.Println(funct, params)

	} else if funct == "log_error" {

		fmt.Println(funct, params)
		Mlog_error(params...)

	} else if funct == "" {

		// Spaces can pass

	} else {

		SpitError(line, column, fmt.Sprintf(`Func %s Does Not Exist`, funct))

	}

}
