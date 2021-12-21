package crm

import "github.com/augusto/imersao5-esquenta-go/entity/crm/cfr"

type CfrRepository interface {
	Select() ([]cfr.Cfr, error)
}
