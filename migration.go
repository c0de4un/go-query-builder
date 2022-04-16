package crm

type IMigration interface {
	Run() error
}
