package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/nowex35/event_management_app/datastore"
	// m "github.com/nowex35/event_management_app/datastore/model"
	"gorm.io/gen"
	// "gorm.io/gen/field"
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
	// g.ApplyBasic([]interface{}{
	// 	g.GenerateModel(m.TableNameTUser),
	// 	g.GenerateModel(m.TableNameTProject),
	// 	g.GenerateModel(
	// 		m.TableNameTUserAssign,
	// 		gen.FieldRelateModel(field.BelongsTo, "User", m.TUser{}, nil),
	// 		gen.FieldRelateModel(field.BelongsTo, "Project", m.TProject{}, nil),
	// 	),
	// 	g.GenerateModel(
	// 		m.TableNameTSession,
	// 		gen.FieldRelateModel(field.BelongsTo, "User", m.TUser{}, nil),
	// 	),
	// 	g.GenerateModel(
	// 		m.TableNameTUserEvent,
	// 		gen.FieldRelateModel(field.BelongsTo, "User", m.TUser{}, nil),
	// 		gen.FieldRelateModel(field.BelongsTo, "Event", m.TEvent{}, nil),
	// 	),
	// }...)
	g.Execute()
}
