package provision

import (
	"github.com/docker/machine/libmachine/drivers"
)

func init() {
	Register("Amazon Linux", &RegisteredProvisioner{
		New: NewAmznProvisioner,
	})
}

func NewAmznProvisioner(d drivers.Driver) Provisioner {
	return &AmznProvisioner{
		NewRedHatProvisioner("amzn", d),
	}
}

type AmznProvisioner struct {
	*RedHatProvisioner
}

func (provisioner *AmznProvisioner) String() string {
	return "amzn"
}
