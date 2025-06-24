package repositories


type Database interface {
    Ping() error
    GetDB() interface{}  // Zwraca *sql.DB, *redis.Client, etc.
}

type PostgresDatabase interface {
    Database
    // Dodatkowe metody specyficzne dla PostgreSQL
}

type TimescaleDatabase interface {
    Database
    // Dodatkowe metody specyficzne dla Timescale
}

type RedisDatabase interface {
    Database
    // Dodatkowe metody specyficzne dla Redis
}