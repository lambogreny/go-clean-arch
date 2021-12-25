package crm

import "github.com/augusto/imersao5-esquenta-go/entity/crm/cfr"

type CfrRepository interface {
	SelectCrm(owner string) ([]cfr.Account, error)
	CheckUpdateErp(id string) (bool, error)
	UpdateErp(account cfr.Account, owner string) error
	DeleteCrm(owner string, id string) error //#TODO
	InsertErp(account cfr.Account, owner string) error
}
