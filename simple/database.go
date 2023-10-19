package simple

type Database struct {
	Name string
}

type DatabasePostgreSQL Database
type DatabaseMongoDB Database

func NewDatabaseMongoDB() *DatabaseMongoDB {
	database := &Database{
		Name: "MongoDB",
	}
	return (*DatabaseMongoDB)(database)
}

func NewDatabasePostgreSQL() *DatabasePostgreSQL {
	database := &Database{
		Name: "PostgreSQL",
	}
	return (*DatabasePostgreSQL)(database)
}

type DatabaseRepository struct {
	DatabasePostgreSQL *DatabasePostgreSQL
	DatabaseMongoDB    *DatabaseMongoDB
}

func NewDatabaseRepository(databasePostgreSQL *DatabasePostgreSQL, databaseMongoDB *DatabaseMongoDB) *DatabaseRepository {
	return &DatabaseRepository{
		DatabasePostgreSQL: databasePostgreSQL,
		DatabaseMongoDB:    databaseMongoDB,
	}
}
