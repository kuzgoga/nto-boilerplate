package main

import (
  "gorm.io/gen"
  "app/internal/models"
)

func main() {
  // Initialize the generator with configuration
  g := gen.NewGenerator(gen.Config{
     OutPath: "../dal", // output directory, default value is ./query
     Mode:    gen.WithDefaultQuery | gen.WithQueryInterface | gen.WithoutContext,
     FieldNullable: true,
  })

  // Generate default DAO interface for those specified structs
  g.ApplyBasic(models.Entities...)

  // Execute the generator
  g.Execute()
}
