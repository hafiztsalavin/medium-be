package main

func main() {

	// Initialize news config
	// cfg := config.NewConfig()

	// Initialize DB repositories
	// db, err := postgres.NewPostgresRepo(&cfg.DatabaseConfig)
	// checkErr(err)

	// repository
	// tagRepo := postgres.NewTagRepository(db)
	// tagNews := postgres.NewNewsRepository(db)
}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}
