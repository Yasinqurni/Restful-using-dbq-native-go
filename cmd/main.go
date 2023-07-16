package main

import (
	"github-dbq/config"
	"github-dbq/db/seeder"
)

func main() {
	env, _ := config.LoadConfig()

	db := config.DbInit(env)

	seeder.StudentSeed(db)

}
