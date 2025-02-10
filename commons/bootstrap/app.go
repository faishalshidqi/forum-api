package bootstrap

type Application struct {
	Env *Env
	DB  Database
}

func App() Application {
	app := Application{}
	app.Env = NewEnv()
	app.DB = *NewPSQLDatabase(app.Env)
	return app
}
