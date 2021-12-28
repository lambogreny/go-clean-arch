package crm

import "github.com/augusto/imersao5-esquenta-go/entity/crm/cfr"

type CfrRepository interface {
	SelectCrm(owner string) ([]cfr.Account, error)
	CheckUpdateErp(id string) (bool, error)
	UpdateErp(account cfr.Account, owner string) error
	DeleteCrm(owner string, id string, tipo string) error
	InsertErp(account cfr.Account, owner string) error
}

type AccountRepository interface {
	SelectErp() ([]cfr.Cfr, error)
	CheckUpdateCrm(id string, owner string) (bool, error)
	UpdateCrm(cfr cfr.Cfr, owner string) error
	DeleteErp(id string, tipo string) error
	InsertCrm(cfr cfr.Cfr, owner string) error
}
