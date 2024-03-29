package config

import (
	"context"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/juiicesb/devops/build/cicd/internal/schema"
	"github.com/juiicesb/devops/pkg/devdeploy"
	"github.com/pkg/errors"
)

// RunSchemaMigrationsForTargetEnv executes schema migrations for the target environment.
func RunSchemaMigrationsForTargetEnv(log *log.Logger, awsCredentials devdeploy.AwsCredentials, targetEnv Env, isUnittest bool) error {

	cfg, err := NewConfig(log, targetEnv, awsCredentials)
	if err != nil {
		return err
	}

	infra, err := devdeploy.SetupInfrastructure(log, cfg)
	if err != nil {
		return err
	}

	connInfo, err := cfg.GetDBConnInfo(infra)
	if err != nil {
		return err
	}

	masterDb, err := sqlx.Open(connInfo.Driver, connInfo.URL())
	if err != nil {
		return errors.Wrap(err, "Failed to connect to db for schema migration.")
	}
	defer masterDb.Close()

	return schema.Migrate(context.Background(), masterDb, log, false)
}
