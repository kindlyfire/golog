package theme

import (
	"fmt"
	"os"
	"path"

	"github.com/kindlyfire/golog/models"
	"github.com/kindlyfire/golog/modules/config"
	"github.com/kindlyfire/golog/modules/db"
	col "github.com/logrusorgru/aurora"
	ini "gopkg.in/ini.v1"
)

var (
	// ThemeName is the slug of the current theme
	ThemeName = "golog"

	// ThemeDir is the directory containing the theme's files
	ThemeDir = "themes/golog/"

	// ThemeInfo is the info.ini file of the theme, parsed
	ThemeInfo *ini.File
)

// Load loads information about the theme
func Load() {
	// Get the active theme from the database
	optActiveTheme := models.Option{}
	defaultOptActiveTheme := models.Option{Name: "active theme", Value: "golog"}
	db.Db.Where("name = ?", "active theme").Attrs(defaultOptActiveTheme).FirstOrCreate(&optActiveTheme)

	err := ActivateTheme(optActiveTheme.Value)

	if err != nil {
		err2 := ActivateTheme("golog")

		if err2 != nil {
			fmt.Println(col.Sprintf(col.Red("Could not load default theme, could not load backup theme.")))
			fmt.Println(col.Sprintf(col.Red("    %s"), err))
			fmt.Println(col.Sprintf(col.Red("    %s"), err2))
			os.Exit(1)
		} else {
			fmt.Println("Loaded backup theme `golog`.")
		}
	}
}

// ActivateTheme activates a theme using it's slug
func ActivateTheme(slug string) error {
	cfg, err := ini.Load(path.Join(config.WorkingDirectory, "themes", slug, "info.ini"))

	if err != nil {
		return err
	}

	ThemeName = slug
	ThemeDir = path.Join("themes", slug)
	ThemeInfo = cfg

	return nil
}
