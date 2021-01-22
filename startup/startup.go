package startup

import "github.com/skeyic/neuron/app/model"

func StartUp() error {
	// Load
	for _, e := range []error{
		model.TheUsersMaster.Init(),
	} {
		if e != nil {
			return e
		}
	}

	return nil
}
