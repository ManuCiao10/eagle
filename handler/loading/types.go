package loading

import (
	"github.com/eagle/handler/profile"
	"github.com/eagle/handler/settings"
)

type Settings struct {
	Settings settings.Settings `json:"settings"`
}

type Env struct {
	Env settings.Env `json:"env"`
}

type Profiles struct {
	Profiles []profile.Profile `json:"profiles"`
}

type Proxies struct {
	ID        string   `json:"id"`
	ProxyList []string `json:"proxy_list"`
}

type Config struct {
	// Accounts        Accounts
	// Tasks           Tasks
	// QuicktaskGroups map[int][]QuicktaskGroup
	Proxies  Proxies  `json:"proxies"`
	Env      Env      `json:"env"`
	Settings Settings `json:"settings"`
	Profiles Profiles `json:"profiles"`
}

var (
	ProxiesIdx = make(map[int]*Proxies)
)
