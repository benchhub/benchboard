// Package host collects hardware info and metrics of the host machine
package host

type Procfile interface {
	Path() string
}
