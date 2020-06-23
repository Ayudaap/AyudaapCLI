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
	"log"
	"math/rand"
	"time"

	"Ayudaap.org/models"
	"Ayudaap.org/repository"
	"github.com/spf13/cobra"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"syreclabs.com/go/faker"
	"syreclabs.com/go/faker/locales"
)

// proyectoCmd represents the proyecto command
var proyectoCmd = &cobra.Command{
	Use:   "proyecto",
	Short: "Inicializa un nuevo proyecto",
	Long:  `Inicializa un nuevo objeto de tipo proyecto`,
	Run: func(cmd *cobra.Command, args []string) {
		borrar, _ := cmd.Parent().PersistentFlags().GetBool("borrar")
		if borrar {
			borrarProyecto()
		} else {
			inicializarProyecto()
		}
	},
}

func init() {
	inicializarCmd.AddCommand(proyectoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// proyectoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// proyectoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

var proyectos []models.Proyecto

// Inicializa la lista de organizaciones
func inicializarProyecto() {
	faker.Locale = locales.En
	rand.Seed(50)
	total := rand.Intn(42)

	for i := 0; i <= total; i++ {

		auditoria := models.Auditoria{
			CreatedAt: primitive.Timestamp{T: uint32(time.Now().Unix())},
			UpdatedAt: primitive.Timestamp{T: uint32(time.Now().Unix())},
		}

		var activo bool = false
		if j := faker.RandomInt(0, 1); j == 0 {
			activo = true
		}

		proyectos = append(proyectos, models.Proyecto{
			ID:        primitive.NewObjectID(),
			Actividad: faker.Company().CatchPhrase(),
			Activo:    activo,
			Area: models.Area{
				ID:          primitive.NewObjectID(),
				Nombre:      faker.Commerce().Department(),
				Descripcion: faker.RandomString(1),
				Auditoria:   auditoria,
			},
			CapacidadesRequeridas: faker.Commerce().Department(),
			Auditoria:             auditoria,
			Costo:                 faker.Commerce().Price(),
			Nombre:                faker.Commerce().ProductName(),
			Objetivo:              faker.Hacker().SaySomethingSmart(),
			VoluntariosRequeridos: faker.RandomInt(3, 9),
		})
	}

	guardarProyectoInicializer()
	fmt.Printf("\n%d creados\n", len(proyectos))
}

// Inicializa la base de datos
func guardarProyectoInicializer() {
	proyRepo := new(repository.ProyectosRepository)
	for _, proy := range proyectos {
		proyRepo.InsertProyecto(proy)
	}
}

func borrarProyecto() {
	proyRepo := new(repository.ProyectosRepository)

	err := proyRepo.PurgarProyectos()
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Collecion purgada")
}
