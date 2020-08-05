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

	"github.com/spf13/cobra"
)

//nombreOpcion Nombre del elemento
var nombreOpcion string

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new TipoArchivo (repositorio / ruta / modelo ) ",
	Short: "Crea un nuevo archivo elemento del proyecto",
	Long:  `Agrega un nuevo componente de programación puede ser Modelo, Ruta o Repositorio`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "repositorio":
			newRepositorio()
			break
		case "ruta":
			newRuta()
			break
		case "modelo":
			newModelo()
			break
		}
	},
}

func init() {
	rootCmd.AddCommand(newCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	newCmd.Flags().StringVarP(&nombreOpcion, "nombre", "n", "r", "Nombre del componente")
}

func newRepositorio() {
	fmt.Printf("nuevo Repo %s\n", nombreOpcion)
}

func newRuta() {
	fmt.Printf("nueva ruta %s\n", nombreOpcion)
}

func newModelo() {
	fmt.Printf("nuevo modelo %s\n", nombreOpcion)
}