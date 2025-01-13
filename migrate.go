package drift

func MigrateUp() {
	db := ReadDatabaseConfiguration()
	db.Connect()
	db.RunMigrations("up", 0)
}

func MigrateDown(step int) {
	db := ReadDatabaseConfiguration()
	db.Connect()
	db.RunMigrations("down", step)
}

func Create(migration_name string) {
	createMigrationFile(migration_name)
}