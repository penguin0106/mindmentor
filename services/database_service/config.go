package main

type DatabaseServiceConfig struct {
	ConnectionString  string
	MigrationsEnabled bool
}

func ConfigureDatabaseService(connectionString string, migrationsEnabled bool) *DatabaseServiceConfig {
	return &DatabaseServiceConfig{
		ConnectionString:  connectionString,
		MigrationsEnabled: migrationsEnabled,
	}
}
