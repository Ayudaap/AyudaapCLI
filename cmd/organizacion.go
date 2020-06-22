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

	"syreclabs.com/go/faker"
	"syreclabs.com/go/faker/locales"

	"Ayudaap.org/models"
	"Ayudaap.org/repository"
	"github.com/spf13/cobra"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// organizacionCmd represents the organizacion command
var organizacionCmd = &cobra.Command{
	Use:   "organizacion",
	Short: "Inicializa los datos de prueba para una organizacion",
	Long:  `Inicializa datos de prueba`,
	Run: func(cmd *cobra.Command, args []string) {
		borrar, _ := cmd.Flags().GetBool("borrar")
		if borrar {
			borrarOrganizaciones()
		} else {
			InicializarOrganizaciones()
		}
	},
}

func init() {
	inicializarCmd.AddCommand(organizacionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// organizacionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// organizacionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	organizacionCmd.Flags().BoolP("borrar", "b", false, "--borrar o -b para borrar los datos actuales de la BD")
}

var organizaciones []models.Organizacion

// Inicializa la lista de organizaciones
func InicializarOrganizaciones() {
	faker.Locale = locales.En
	rand.Seed(50)
	total := rand.Intn(42)

	for i := 0; i <= total; i++ {

		organizaciones = append(organizaciones, models.Organizacion{
			ID:   primitive.NewObjectID(),
			Tipo: models.OrganizacionNoGubernamental,
			Domicilio: models.Direccion{
				ID:             primitive.NewObjectID(),
				Calle:          faker.Address().StreetName(),
				NumeroExterior: faker.Address().BuildingNumber(),
				CodigoPostal:   faker.Address().Postcode(),
				Colonia:        faker.Address().City(),
				Estado:         faker.Address().State(),
			},
			Auditoria: models.Auditoria{
				CreatedAt: primitive.Timestamp{T: uint32(time.Now().Unix())},
				UpdatedAt: primitive.Timestamp{T: uint32(time.Now().Unix())},
			},
			Nombre: faker.Company().Name(),
		})

		rd := faker.RandomInt(1, 7)
		for j := 0; j <= rd; j++ {
			rndVal := faker.RandomInt(0, 1)
			var principal bool = false
			if rndVal == 1 {
				principal = true
			}

			organizaciones[i].Domicilio.Directorio = append(organizaciones[i].Domicilio.Directorio, models.Directorio{
				Alias:             fmt.Sprintf("%s %s", faker.Name().Prefix(), faker.Name().LastName()),
				CorreoElectronico: faker.Internet().Email(),
				Nombre:            faker.Name().Name(),
				Telefono:          faker.PhoneNumber().PhoneNumber(),
				EsPrincipal:       principal,
				ID:                primitive.NewObjectID(),
			})
		}
	}

	guardarOrganizacionesInicializer()

	fmt.Printf("\n%d creados\n", len(organizaciones))
}

// Inicializa la base de datos
func guardarOrganizacionesInicializer() {
	orgRepo := new(repository.OrganizacionesRepository)

	for _, org := range organizaciones {
		orgRepo.InsertOrganizacion(org)
	}
}

func borrarOrganizaciones() {
	orgRepo := new(repository.OrganizacionesRepository)

	err := orgRepo.PurgarOrganizaciones()
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Collecion purgada")
}
