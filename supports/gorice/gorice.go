package gorice

import (
	"github.com/foolin/gin-template"
	"github.com/GeertJohan/go.rice"
)

/**
New gin template engine, default views root.
 */
func New(viewsRootBox *rice.Box) *gintemplate.TemplateEngine {
	return NewWithConfig(viewsRootBox, gintemplate.DefaultConfig)
}

/**
New gin template engine
Important!!! The viewsRootBox's name and config.Root must be consistent.
 */
func NewWithConfig(viewsRootBox *rice.Box, config gintemplate.TemplateConfig) *gintemplate.TemplateEngine {
	config.Root = viewsRootBox.Name()
	engine := gintemplate.New(config)
	engine.SetFileHandler(GoRiceFileHandler())
	return engine
}

func GoRiceFileHandler() gintemplate.FileHandler {
	return func(config gintemplate.TemplateConfig, tplFile string) (content string, err error) {
		// find a rice.Box
		templateBox, err := rice.FindBox(config.Root)
		if err != nil {
			return "", err
		}
		// get file contents as string
		return templateBox.String(tplFile + config.Extension)
	}
}
