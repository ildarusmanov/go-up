package commands

import (
	"log"
	"os"

	"github.com/ildarusmanov/go-up/internal/config"
	"github.com/ildarusmanov/go-up/internal/files"
	"github.com/ildarusmanov/go-up/internal/render"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var (
	initCmdArgPackage          *string
	initCmdArgAppTemplatesPath *string
	initCmdArgBaseAppTemplate  *string
)

// NewInitCommand initializes Init command
func NewInitCommand() (*cobra.Command, func()) {
	cmd := &cobra.Command{
		Use:   "init [name]",
		Short: "Create new application",
		Long:  ``,
		Args:  cobra.MinimumNArgs(1),
		Run:   runInitCmd,
	}

	return cmd, initInitCmd(cmd)
}

// initInitCmd initializes Init command
func initInitCmd(cmd *cobra.Command) func() {
	return func() {
		log.Printf("Load Init Command")
		initCmdArgPackage = cmd.PersistentFlags().StringP("pkg", "p", "", "application package")
		initCmdArgAppTemplatesPath = cmd.PersistentFlags().StringP("tpl-dir", "t", ApplicationTemplatesPath, "application templates source directory")
		initCmdArgBaseAppTemplate = cmd.PersistentFlags().StringP("tpl-name", "b", BaseApplicationTemplate, "application template name")
	}
}

// runInitCmd runs Init command
func runInitCmd(cmd *cobra.Command, args []string) {
	wd, err := os.Getwd()

	if len(args) < 1 {
		log.Fatalf("please, enter a new project name")
	}

	if err != nil {
		log.Fatalf("can not detect current directory: %s", err.Error())
	}

	td := *initCmdArgAppTemplatesPath + *initCmdArgBaseAppTemplate

	createNewProject(args[0], td, wd)
}

// createNewProject creates new project with template
func createNewProject(pname, tdir, wdir string) {
	if pname == "" {
		log.Fatal("Invalid project name")
	}

	pkg := *initCmdArgPackage

	if pkg == "" {
		log.Fatal("Invalid project name")
	}

	pdir := wdir + "/" + pname

	createProjectDirs(pdir, tdir)
	cfg := createGoupConfigFile(pkg, pdir)
	createBaseFiles(cfg, pdir, tdir)

	log.Println("Application initialized")
}

func createGoupConfigFile(pkg, pdir string) *config.GoupConfig {
	fpath := pdir + "/" + GoupConfigFile

	cfgFile, err := os.Create(fpath)

	if err != nil {
		log.Fatalf("Can not create %s: %s", fpath, err)
	}

	defer cfgFile.Close()

	log.Printf("New file %s successfully created", fpath)

	cfg := config.NewGoupConfig(pkg)
	enc := yaml.NewEncoder(cfgFile)

	defer enc.Close()

	if err := enc.Encode(cfg); err != nil {
		log.Fatalf("Can not write config file: %s", err)
	}

	return cfg
}

func createProjectDirs(pdir, tdir string) {
	dirs, err := files.GetFoldersPathsList(tdir)

	if err != nil {
		log.Fatal(err)
	}

	for _, dir := range dirs {
		if err := os.Mkdir(pdir+"/"+dir, os.FileMode(0777)); err != nil {
			log.Fatalf("Can not create %s: %s", dir, err)
		}

		log.Printf("New directory %s successfully created", dir)
	}

	log.Printf("Directories are created")
}

func createBaseFiles(cfg *config.GoupConfig, pdir, tdir string) {
	tpls, err := files.GetTemplatesPathsList(ApplicationTemplatesPath + BaseApplicationTemplate)

	if err != nil {
		log.Fatal(err)
	}

	if err := render.RenderTemplates(tdir, pdir, cfg, nil, tpls); err != nil {
		log.Fatal(err)
	}

	log.Printf("Templates are created")
}
