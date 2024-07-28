package linker

import (
	"time"

	"github.com/darklab8/fl-configs/configs/configs_export"
	"github.com/darklab8/fl-configs/configs/configs_mapped"
	"github.com/darklab8/fl-darkcore/darkcore/builder"
	"github.com/darklab8/fl-darkcore/darkcore/core_static"

	"github.com/darklab8/fl-darkmap/darkmap/front"
	"github.com/darklab8/fl-darkmap/darkmap/front/static_front"
	"github.com/darklab8/fl-darkmap/darkmap/front/urls"
	"github.com/darklab8/fl-darkmap/darkmap/settings"
	"github.com/darklab8/fl-darkmap/darkmap/settings/logus"
	"github.com/darklab8/fl-darkmap/darkmap/types"
	"github.com/darklab8/go-utils/utils/timeit"
	"github.com/darklab8/go-utils/utils/utils_logus"
)

type Linker struct {
	mapped  *configs_mapped.MappedConfigs
	configs *configs_export.Exporter
}

type LinkOption func(l *Linker)

func NewLinker(opts ...LinkOption) *Linker {
	l := &Linker{}
	for _, opt := range opts {
		opt(l)
	}

	timeit.NewTimerF(func(m *timeit.Timer) {
		freelancer_folder := settings.Env.FreelancerFolder
		if l.configs == nil {
			l.mapped = configs_mapped.NewMappedConfigs()
			logus.Log.Debug("scanning freelancer folder", utils_logus.FilePath(freelancer_folder))
			l.mapped.Read(freelancer_folder)
			l.configs = configs_export.NewExporter(l.mapped)
		}
	}, timeit.WithMsg("MappedConfigs creation"))
	return l
}

func (l *Linker) Link() *builder.Builder {
	var build *builder.Builder
	timeit.NewTimerF(func(m *timeit.Timer) {
		timeit.NewTimerF(func(m *timeit.Timer) {
			staticPrefix := "static/"
			siteRoot := settings.Env.SiteRoot
			params := types.GlobalParams{
				Buildpath:         "",
				SiteRoot:          siteRoot,
				StaticRoot:        siteRoot + staticPrefix,
				OppositeThemeRoot: siteRoot + "dark.html",
				Timestamp:         time.Now().UTC(),
			}

			static_files := []builder.StaticFile{
				builder.NewStaticFileFromCore(core_static.HtmxJS),
				builder.NewStaticFileFromCore(core_static.HtmxPreloadJS),
				builder.NewStaticFileFromCore(core_static.SortableJS),
				builder.NewStaticFileFromCore(core_static.ResetCSS),
				builder.NewStaticFileFromCore(core_static.FaviconIco),

				builder.NewStaticFileFromCore(static_front.CommonCSS),
				builder.NewStaticFileFromCore(static_front.CustomCSS),
				builder.NewStaticFileFromCore(static_front.CustomJS),
			}

			build = builder.NewBuilder(params, static_files)
		}, timeit.WithMsg("building creation"))

		// var data *configs_export.Exporter
		// timeit.NewTimerF(func(m *timeit.Timer) {
		// 	data = l.configs.Export()
		// }, timeit.WithMsg("exporting data"))

		build.RegComps(
			builder.NewComponent(
				urls.Index,
				front.Index(),
			),
		)
	})
	return build
}
