package main

import (
  "gorm.io/gen"
  "app/internal/models"
)

func main() {
  g := gen.NewGenerator(gen.Config{
     OutPath: "../dal", // output directory, default value is ./query
     Mode:    gen.WithDefaultQuery | gen.WithQueryInterface | gen.WithoutContext,
     FieldNullable: true,
     WithUnitTest: true,
  })
  g.ApplyBasic(models.Entities...)
  g.Execute()
}
