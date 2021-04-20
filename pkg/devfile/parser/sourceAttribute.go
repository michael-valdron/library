package parser

import (
	"fmt"
	v1 "github.com/devfile/api/v2/pkg/apis/workspaces/v1alpha2"
	"github.com/devfile/api/v2/pkg/attributes"
)

const ImportSourceAttribute = "library.devfile.io/imported-from"
const PluginOverrideAttribute = "library.devfile.io/plugin-override-from"
const ParentOverrideAttribute = "library.devfile.io/parent-override-from"

// addSourceAttributesForTemplateSpecContent adds an attribute 'library.devfile.io/imported-from=<source reference>'
//  to all elements of template spec content that support attributes.
func addSourceAttributesForTemplateSpecContent(sourceImportReference v1.ImportReference, template *v1.DevWorkspaceTemplateSpecContent) {
	for idx, component := range template.Components {
		if component.Attributes == nil {
			template.Components[idx].Attributes = attributes.Attributes{}
		}
		template.Components[idx].Attributes.PutString(ImportSourceAttribute, resolveImportReference(sourceImportReference))
	}
	for idx, command := range template.Commands {
		if command.Attributes == nil {
			template.Commands[idx].Attributes = attributes.Attributes{}
		}
		template.Commands[idx].Attributes.PutString(ImportSourceAttribute, resolveImportReference(sourceImportReference))
	}
	for idx, project := range template.Projects {
		if project.Attributes == nil {
			template.Projects[idx].Attributes = attributes.Attributes{}
		}
		template.Projects[idx].Attributes.PutString(ImportSourceAttribute, resolveImportReference(sourceImportReference))
	}
	for idx, project := range template.StarterProjects {
		if project.Attributes == nil {
			template.StarterProjects[idx].Attributes = attributes.Attributes{}
		}
		template.StarterProjects[idx].Attributes.PutString(ImportSourceAttribute, resolveImportReference(sourceImportReference))
	}
}

// addSourceAttributesForParentOverride adds an attribute 'library.devfile.io/imported-from=<source reference>'
//  to all elements of parent override that support attributes.
func addSourceAttributesForParentOverride(sourceImportReference v1.ImportReference, parentOverrides *v1.ParentOverrides) {
	for idx, component := range parentOverrides.Components {
		if component.Attributes == nil {
			parentOverrides.Components[idx].Attributes = attributes.Attributes{}
		}
		parentOverrides.Components[idx].Attributes.PutString(ParentOverrideAttribute, resolveImportReference(sourceImportReference))
	}
	for idx, command := range parentOverrides.Commands {
		if command.Attributes == nil {
			parentOverrides.Commands[idx].Attributes = attributes.Attributes{}
		}
		parentOverrides.Commands[idx].Attributes.PutString(ParentOverrideAttribute, resolveImportReference(sourceImportReference))
	}
	for idx, project := range parentOverrides.Projects {
		if project.Attributes == nil {
			parentOverrides.Projects[idx].Attributes = attributes.Attributes{}
		}
		parentOverrides.Projects[idx].Attributes.PutString(ParentOverrideAttribute, resolveImportReference(sourceImportReference))
	}
	for idx, project := range parentOverrides.StarterProjects {
		if project.Attributes == nil {
			parentOverrides.StarterProjects[idx].Attributes = attributes.Attributes{}
		}
		parentOverrides.StarterProjects[idx].Attributes.PutString(ParentOverrideAttribute, resolveImportReference(sourceImportReference))
	}

}

// addSourceAttributesForPluginOverride adds an attribute 'library.devfile.io/imported-from=<source reference>'
//  to all elements of plugin override that support attributes.
func addSourceAttributesForPluginOverride(sourceImportReference v1.ImportReference, pluginOverrides *v1.PluginOverrides) {
	for idx, component := range pluginOverrides.Components {
		if component.Attributes == nil {
			pluginOverrides.Components[idx].Attributes = attributes.Attributes{}
		}
		pluginOverrides.Components[idx].Attributes.PutString(PluginOverrideAttribute, resolveImportReference(sourceImportReference))
	}
	for idx, command := range pluginOverrides.Commands {
		if command.Attributes == nil {
			pluginOverrides.Commands[idx].Attributes = attributes.Attributes{}
		}
		pluginOverrides.Commands[idx].Attributes.PutString(PluginOverrideAttribute, resolveImportReference(sourceImportReference))
	}

}

// AddSourceAttributesForOverrideAndMerge adds an attribute record the import reference to all elements of template that support attributes.
func AddSourceAttributesForOverrideAndMerge(sourceImportReference v1.ImportReference, template interface{}) error {
	if template == nil {
		fmt.Errorf("cannot add source attributes to nil")
	}

	mainContent, isMainContent := template.(*v1.DevWorkspaceTemplateSpecContent)
	parentOverride, isParentOverride := template.(*v1.ParentOverrides)
	pluginOverride, isPluginOverride := template.(*v1.PluginOverrides)

	switch {
	case isMainContent:
		addSourceAttributesForTemplateSpecContent(sourceImportReference, mainContent)
	case isParentOverride:
		addSourceAttributesForParentOverride(sourceImportReference, parentOverride)
	case isPluginOverride:
		addSourceAttributesForPluginOverride(sourceImportReference, pluginOverride)
	default:
		return fmt.Errorf("unknown template type")
	}

	return nil
}
