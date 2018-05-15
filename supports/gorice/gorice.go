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
	engine.SetFileHandler(FileHandler(viewsRootBox))
	return engine
}

/**
 Support go.rice file handler
 */
func FileHandler(viewsRootBox *rice.Box) gintemplate.FileHandler {
	return func(config gintemplate.TemplateConfig, tplFile string) (content string, err error) {
		// get file contents as string
		return viewsRootBox.String(tplFile + config.Extension)
	}
}
