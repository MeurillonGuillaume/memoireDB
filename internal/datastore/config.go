package datastore

type Config struct {
	Type string `required:"true" help:"Which type of datastore should be used. Available options are 'memory' and 'persisted'"`
}
