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
	"io/ioutil"
	"log"
	"strings"

	"github.com/spf13/cobra"
)

//nombreOpcion
var nombreOpcion, salida string

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
		default:
			log.Printf("Opcion invalida")
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
	newCmd.Flags().StringVarP(&salida, "salida", "s", "r", "ruta de salida")
}

func newRepositorio() {
	fmt.Printf("nuevo Repo %s\n", nombreOpcion)

	archivo, err := ioutil.ReadFile("/Users/erikvillegas/go/src/AyudaapCLI/plantillas/repositorio.tmp")
	if err != nil {
		panic(err)
	}

	nombreMinusculas := strings.ToLower(nombreOpcion)
	nombreInicial := nombreOpcion[0:1]
	nombreInicial = strings.ToLower(nombreInicial)

	newContents := strings.Replace(string(archivo), "<%nombre%>", nombreOpcion, -1)
	newContents = strings.Replace(newContents, "<%nombreMinusculas%>", nombreMinusculas, -1)
	newContents = strings.Replace(newContents, "<%nombreInicial%>", nombreInicial, -1)
	err = ioutil.WriteFile(salida, []byte(newContents), 0764)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Nueva modelo generado: %s \nruta: %s \n", nombreOpcion, salida)

}

func newRuta() {
	archivo, err := ioutil.ReadFile("/Users/erikvillegas/go/src/AyudaapCLI/plantillas/ruta.tmp")

	if err != nil {
		panic(err)
	}

	nombreMinuscula := strings.ToLower(nombreOpcion)

	newContents := strings.Replace(string(archivo), "<%nombreRepo%>", nombreOpcion, -1)
	newContents = strings.Replace(newContents, "<%nombreMinuscula%>", nombreMinuscula, -1)

	err = ioutil.WriteFile(salida, []byte(newContents), 0764)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Nueva ruta generada: %s \nruta: %s\n", nombreOpcion, salida)
}

func newModelo() {

	archivo, err := ioutil.ReadFile("/Users/erikvillegas/go/src/AyudaapCLI/plantillas/modelo.tmp")
	if err != nil {
		panic(err)
	}

	newContents := strings.Replace(string(archivo), "<%nombre%>", nombreOpcion, -1)
	err = ioutil.WriteFile(salida, []byte(newContents), 0764)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Nueva modelo generado: %s \nruta: %s \n", nombreOpcion, salida)
}
