package startermail

import (
	"embed"

	"github.com/bitwormhole/starter"
	"github.com/bitwormhole/starter-mail/gen/demo"
	"github.com/bitwormhole/starter-mail/gen/lib"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/collection"
)

const (
	myModuleName = "github.com/bitwormhole/starter-mail"
	myModuleVer  = "v0.1.0"
	myModuleRev  = 1
)

////////////////////////////////////////////////////////////////////////////////

//go:embed "src/main/resources"
var theModuleSrcMainRes embed.FS

// Module 导出模块【github.com/bitwormhole/starter-mail】
func Module() application.Module {

	mb := application.ModuleBuilder{}
	mb.Name(myModuleName)
	mb.Version(myModuleVer)
	mb.Revision(myModuleRev)

	mb.OnMount(lib.ExportConfigForStarterMailLib)
	mb.Resources(collection.LoadEmbedResources(&theModuleSrcMainRes, "src/main/resources"))
	mb.Dependency(starter.Module())

	return mb.Create()
}

////////////////////////////////////////////////////////////////////////////////

//go:embed "src/demo/resources"
var theModuleSrcDemoRes embed.FS

// ModuleForDemo 导出模块【github.com/bitwormhole/starter-mail#demo】
func ModuleForDemo() application.Module {

	parent := Module()

	mb := application.ModuleBuilder{}
	mb.Name(myModuleName + "#demo")
	mb.Version(myModuleVer)
	mb.Revision(myModuleRev)

	mb.OnMount(demo.ExportConfigForStarterMailDemo)
	mb.Resources(collection.LoadEmbedResources(&theModuleSrcDemoRes, "src/demo/resources"))
	mb.Dependency(parent)

	return mb.Create()
}

////////////////////////////////////////////////////////////////////////////////
