package pkg

import "github.com/evanw/esbuild/pkg/api"

func EsbuildJs(indexfile string, outfile string) (resulti api.BuildResult) {

	result := api.Build(api.BuildOptions{
		EntryPoints:       []string{indexfile},
		Outfile:           outfile,
		Platform:          api.PlatformBrowser,
		Target:            api.ES2017,
		EntryNames:        "[name]-[hash]",
		Sourcemap:         api.SourceMapLinked,
		Bundle:            true,
		AllowOverwrite:    true,
		MinifyWhitespace:  true,
		MinifyIdentifiers: true,
		Write:             false,
	})

	return result
}

func EsbuildCss(indexfile string, outfile string) (result api.BuildResult) {

	result = api.Build(api.BuildOptions{
		EntryPoints:       []string{indexfile},
		Outfile:           outfile,
		Loader:            map[string]api.Loader{".css": api.LoaderCSS},
		EntryNames:        "[name]-[hash]",
		Sourcemap:         api.SourceMapLinked,
		Bundle:            true,
		AllowOverwrite:    true,
		MinifyWhitespace:  true,
		MinifyIdentifiers: true,
		Write:             false,
	})

	return result
}
