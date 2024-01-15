package application

type Application interface {
	// Run is a method that runs the application
	Run() (err error)
}
