package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/nowex35/event_management_app/datastore"
	"gorm.io/gen"
)

func main() {

	var err error

	err = godotenv.Load(".env")
	if err != nil {
		fmt.Println(err)
	}

	err = datastore.Init()
	if err != nil {
		fmt.Println(err)
	}

	g := gen.NewGenerator(gen.Config{
		OutPath:           "./datastore/gen/",
		Mode:              gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
		FieldWithIndexTag: true,
		FieldWithTypeTag:  true,
		FieldNullable:     true,
	})
	g.UseDB(datastore.Client)
	g.ApplyBasic(g.GenerateAllTable()...)
	g.Execute()
}
