/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// inicializarCmd represents the inicializar command
var inicializarCmd = &cobra.Command{
	Use:   "inicializar",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fstatus, _ := cmd.Flags().GetBool("float")
		if fstatus {
			addFloat(args)
		} else {
			addInt(args)
		}
	},
}

func init() {
	rootCmd.AddCommand(inicializarCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// inicializarCmd.PersistentFlags().String("foo", "", "A help for foo")
	inicializarCmd.PersistentFlags().BoolP("borrar", "b", false, "--borrar o -b para borrar los datos actuales de la BD")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// inicializarCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	inicializarCmd.Flags().BoolP("float", "f", false, "Add floating numbers")
}

func addInt(args []string) {
	if len(args) <= 1 {
		fmt.Println("Los argumentos deben ser mayores a 1")
	}
	var sum int

	for _, val := range args {
		item, err := strconv.Atoi(val)
		if err != nil {
			fmt.Println(err.Error())
		}
		sum = sum + item
	}
	fmt.Printf("Total %d\n", sum)
}

func addFloat(args []string) {
	if len(args) <= 1 {
		fmt.Println("Los argumentos deben ser mayores a 1")
	}
	var sum float64
	for _, fval := range args {
		// convert string to float64
		ftemp, err := strconv.ParseFloat(fval, 64)
		if err != nil {
			fmt.Println(err)
		}
		sum = sum + ftemp
	}
	fmt.Printf("Total %f\n", sum)
}
