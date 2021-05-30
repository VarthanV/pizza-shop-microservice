package migrations

import "context"

type MigrationService interface {
	RunMigrations(ctx context.Context) 
}
