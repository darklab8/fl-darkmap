package settings

import (
	"fmt"

	_ "embed"

	"runtime/debug"

	"github.com/darklab8/fl-configs/configs/configs_settings"
	"github.com/darklab8/go-utils/utils/enverant"
	"github.com/darklab8/go-utils/utils/utils_settings"
)

type DarkstatEnvVars struct {
	utils_settings.UtilsEnvs
	configs_settings.ConfEnvVars
	SiteRoot   string
	AppVersion string
}

var Env DarkstatEnvVars

func init() {
	var GolangVersion string
	if info, ok := debug.ReadBuildInfo(); ok {
		fmt.Println(info.GoVersion)
		GolangVersion = info.GoVersion
	}
	env := enverant.NewEnverant()
	Env = DarkstatEnvVars{
		UtilsEnvs:   utils_settings.GetEnvs(env),
		ConfEnvVars: configs_settings.GetEnvs(env),
		SiteRoot:    env.GetStr("SITE_ROOT", enverant.OrStr("/")),
		AppVersion:  GolangVersion,
	}
	fmt.Sprintln("conf=", Env)
}
