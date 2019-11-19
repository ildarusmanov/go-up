package commands

import (
	"errors"
	"log"
	"os"
	"strings"

	"github.com/ildarusmanov/go-up/internal/config"
	"github.com/ildarusmanov/go-up/internal/files"
	"github.com/ildarusmanov/go-up/internal/render"
	"github.com/spf13/cobra"
)

var (
	generateCmdArgCfgDir *string
	ErrTemplateNotFound  = errors.New("template not found")
)

func NewGenerateCommand() (*cobra.Command, func()) {
	cmd := &cobra.Command{
		Use:   "generate [template name] [destination name]",
		Short: "Generate new item by template with name",
		Long:  ``,
		Args:  cobra.MinimumNArgs(1),
		Run:   runGenerateCmd,
	}

	return cmd, initGenerateCmd(cmd)
}

func runGenerateCmd(cmd *cobra.Command, args []string) {
	if len(args) < 2 {
		log.Fatal("Please, enter a template name and destination name")
	}

	wd, err := os.Getwd()

	if err != nil {
		log.Fatalf("can not detect current directory: %s", err.Error())
	}

	tname := args[0]
	dname := args[1]

	generateNewItem(tname, dname, wd)
}

func initGenerateCmd(cmd *cobra.Command) func() {
	return func() {
		log.Printf("Load Generate Command")
		generateCmdArgCfgDir = cmd.PersistentFlags().StringP("cfg", "c", "", "application config directory")
	}
}

func generateNewItem(tname, dname, wdir string) {
	goupCfg, err := config.LoadGoupConfig(*generateCmdArgCfgDir)

	if err != nil {
		log.Fatal(err)
	}

	templatesConfig, err := config.LoadTemplatesConfig()

	if err != nil {
		log.Fatal(err)
	}

	tcfg, err := getTemplateConfig(tname, wdir, goupCfg, templatesConfig)

	if err != nil {
		log.Fatal(err)
	}

	tdir, err := tcfg.GetTemplatesDirectory(wdir)

	if err != nil {
		log.Fatal(err)
	}

	createDirs(wdir, tdir, dname)
	createTplFiles(tcfg, wdir, tdir, dname)
}

func getTemplateConfig(
	tname, wdir string,
	goupCfg *config.GoupConfig,
	tplsCfg *config.TemplatesConfig,
) (*config.TemplatesConfigItem, error) {
	for _, item := range tplsCfg.Templates {
		if item.Name == tname {
			err := item.LoadConfig(wdir, goupCfg)

			if err != nil {
				return nil, err
			}

			return item, nil
		}
	}

	return nil, ErrTemplateNotFound
}

func createDirs(pdir, tdir, dname string) {
	dirs, err := files.GetFoldersPathsList(tdir)

	if err != nil {
		log.Fatal(err)
	}

	for _, dir := range dirs {
		if err := os.Mkdir(pdir+"/"+dir, os.FileMode(0777)); err != nil {

			log.Fatalf("Can not create %s: %s", strings.ReplaceAll(dir, "{name}", dname), err)
		}

		log.Printf("New directory %s successfully created", dir)
	}

	log.Printf("Directories are created")
}

func createTplFiles(payload interface{}, pdir, tdir, dname string) {
	tpls, err := files.GetTemplatesPathsList(tdir)

	for i, tname := range tpls {
		tpls[i] = strings.ReplaceAll(tname, "{name}", dname)
	}

	if err != nil {
		log.Fatal(err)
	}

	if err := render.RenderTemplates(tdir, pdir, payload, nil, tpls); err != nil {
		log.Fatal(err)
	}

	log.Printf("Templates are created")
}
