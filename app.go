package resk

import (
	_ "github.com/Kinggo/account/apis/web"
	_ "github.com/Kinggo/account/core/accounts"
	"github.com/Kinggo/infra"
	"github.com/Kinggo/infra/base"
)

func init() {
	infra.Register(&base.PropsStarter{})
	infra.Register(&base.DbxDatabaseStarter{})
	infra.Register(&base.ValidatorStarter{})
	infra.Register(&base.IrisServerStarter{})
	infra.Register(&infra.WebApiStarter{})
	infra.Register(&base.EurekaStarter{})
	infra.Register(&base.HookStarter{})
}
