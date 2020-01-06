package freebsdpkg

// Shell commands
type Scripts struct {
	PostInstall   string `json:"post-install"`   // Shell command(s)
	PostDeinstall string `json:"post-deinstall"` // Shell command(s)
}

type Dependencies map[string]struct { // for example libgdiplus
	Origin  string // category/package for example x11-toolkits/libgdiplus
	Version string // "5.6_2"
}

type Options struct {
	TEST string `json:"TEST"` // "off"
}

type Files map[string]string // [filename]checksum

// format for +MANIFEST file
type Manifest struct {
	Name         string       `json:"name"`   // Package name
	Origin       string       `json:"origin"` // category/package
	Version      string       `json:"version"`
	Comment      string       `json:"comment"` // Short description
	Maintainer   string       `json:"maintainer"`
	Url          string       `json:"www"`    // Project WWW address
	Abi          string       `json:"abi"`    // "FreeBSD:13:*"
	Architecture string       `json:"arch"`   // "freebsd:13:*"
	Prefix       string       `json:"prefix"` // "/usr/local"
	Flatsize     int          `json:"flatsize"`
	LicenseLogic string       `json:"licenselogic"` // "single"
	Licenses     []string     `json:"licenses"`     // List of licenses
	Description  string       `json:"desc"`         // Long description
	Dependencies Dependencies `json:"deps"`
	Categories   []string     `json:"categories"`
	Options      Options      `json:"options"`
	Files        Files        `json:"files"` // [filename]checksum
	Scripts      Scripts      `json:"scripts"`
}

// format for +MANIFEST_COMPACT file
type ManifestCompact struct {
	Name         string       `json:"name"`   // Package name
	Origin       string       `json:"origin"` // category/package
	Version      string       `json:"version"`
	Comment      string       `json:"comment"` // Short description
	Maintainer   string       `json:"maintainer"`
	Url          string       `json:"www"`    // Project WWW address
	Abi          string       `json:"abi"`    // "FreeBSD:13:*"
	Architecture string       `json:"arch"`   // "freebsd:13:*"
	Prefix       string       `json:"prefix"` // "/usr/local"
	Flatsize     int          `json:"flatsize"`
	LicenseLogic string       `json:"licenselogic"` // "single"
	Licenses     []string     `json:"licenses"`     // List of licenses
	Description  string       `json:"desc"`         // Long description
	Dependencies Dependencies `json:"deps"`
	Categories   []string     `json:"categories"`
	Options      Options      `json:"options"`
}

func (m *Manifest) ToCompact() ManifestCompact {
	return ManifestCompact{
		Name:         m.Name,
		Origin:       m.Origin,
		Version:      m.Version,
		Comment:      m.Comment,
		Maintainer:   m.Maintainer,
		Url:          m.Url,
		Abi:          m.Abi,
		Architecture: m.Architecture,
		Prefix:       m.Prefix,
		Flatsize:     m.Flatsize,
		LicenseLogic: m.LicenseLogic,
		Licenses:     m.Licenses,
		Description:  m.Description,
		Dependencies: m.Dependencies,
		Categories:   m.Categories,
		Options:      m.Options,
	}
}
