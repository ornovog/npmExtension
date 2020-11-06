package jsonFlatFormat

type flatFormatData struct {
	PackageName          string
	DependencyToVersions stringToStringSlice
}

type stringToStringSlice map[string][]string

type stringToStringSet map[string]stringSet

type stringSet map[string]bool
