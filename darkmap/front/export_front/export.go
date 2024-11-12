package export_front

import (
	"github.com/darklab8/fl-configs/configs/configs_mapped"
	"github.com/darklab8/fl-darkmap/darkmap/settings"
	"github.com/darklab8/fl-darkmap/darkmap/settings/logus"
	"github.com/darklab8/go-utils/utils/timeit"
	"github.com/darklab8/go-utils/utils/utils_logus"
)

type Export struct {
	Mapped *configs_mapped.MappedConfigs
}

func NewExport() *Export {
	e := &Export{}

	defer timeit.NewTimer("MappedConfigs creation").Close()

	freelancer_folder := settings.Env.FreelancerFolder
	if e.Mapped == nil {
		logus.Log.Debug("scanning freelancer folder", utils_logus.FilePath(freelancer_folder))
		e.Mapped = configs_mapped.NewMappedConfigs().Read(freelancer_folder)
	}

	e.export()

	return e
}

func (e *Export) export() {

}