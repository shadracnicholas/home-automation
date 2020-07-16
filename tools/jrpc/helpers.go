package main

import (
	"fmt"
	"net/http"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/shadracnicholas/home-automation/libraries/go/svcdef"
	"github.com/shadracnicholas/home-automation/tools/jrpc/imports"
)

var (
	reValidGoStruct           = regexp.MustCompile(`^[A-Z][a-zA-Z0-9]*`)
	reValidGoStructUnderscore = regexp.MustCompile("^[A-Z][a-zA-Z0-9_]*$")
)

// externalPackageName returns the name to use for the external package
func externalPackageName(opts *options) string {
	return strings.ReplaceAll(filepath.Base(opts.DefPath), ".", "")
}

func getMethod(r *svcdef.RPC) (string, error) {
	method, ok := r.Options["method"].(string)
	if !ok {
		return "", fmt.Errorf("method option not set on RPC %s", r.Name)
	}

	switch method {
	case http.MethodGet,
		http.MethodHead,
		http.MethodPost,
		http.MethodPut,
		http.MethodPatch,
		http.MethodDelete,
		http.MethodConnect,
		http.MethodOptions,
		http.MethodTrace:
		return method, nil
	}

	return "", fmt.Errorf("invalid method on RPC %s: %s", r.Name, method)
}

func getPath(r *svcdef.RPC) (string, error) {
	p, ok := r.Options["path"].(string)
	if !ok {
		return "", fmt.Errorf("path option not set on RPC %s", r.Name)
	}

	if p[0] != '/' {
		return "", fmt.Errorf("path on RPC %s must start with a /", r.Name)
	}

	return p, nil
}

type typeInfo struct {
	TypeName      string
	FullTypeName  string
	IsMessageType bool
	Repeated      bool
	Pointer       bool
}

// resolveTypeName will turn a fully-qualified type name into the go type and,
// if necessary, a path that needs to be imported.
func resolveTypeName(t *svcdef.Type, f *svcdef.File, im *imports.Manager) (*typeInfo, error) {
	var fullTypeName, importPath string
	var messageType bool
	var err error

	if t.Map { // map type
		key, err := resolveTypeName(t.MapKey, f, im)
		if err != nil {
			return nil, err
		}

		val, err := resolveTypeName(t.MapValue, f, im)
		if err != nil {
			return nil, err
		}

		// Nothing needs to be imported for a map type. If either
		// the key or value type needed an import, the recursive
		// call will have already dealt with that.
		// TODO: support nested maps properly
		fullTypeName = "map[" + key.FullTypeName + "]" + val.FullTypeName

	} else if strings.HasPrefix(t.Qualified, ".") { // local type (message is defined in the same def file)
		// Remove the first dot and replace any others with underscores
		fullTypeName = strings.ReplaceAll(t.Qualified[1:], ".", "_")
		messageType = true

		// By convention, the type will be defined in the external package
		importPath, err = resolver.Resolve(f.Path, packageDirExternal)
		if err != nil {
			return nil, err
		}

	} else if parts := strings.SplitN(t.Qualified, ".", 2); len(parts) == 2 { // imported type
		fullTypeName = strings.ReplaceAll(parts[1], ".", "_")
		messageType = true

		// Expect to find an import with an alias of parts[0], and again, by
		// convention, the type name will be defined in the external package.
		importPath, err = resolver.Resolve(f.Imports[parts[0]].Path, packageDirExternal)
		if err != nil {
			return nil, err
		}

	} else if data, ok := typeMap[t.Name]; ok { // "built-in" type
		fullTypeName, importPath = data.GoType, data.ImportPath

	} else {
		return nil, fmt.Errorf("invalid type %q", t.Name)
	}

	alias := im.Add(importPath)
	if alias != "" {
		fullTypeName = alias + "." + fullTypeName
	}

	// Type name can be used to instantiate the type as it
	// is not prepended with pointer or slice characters.
	typeName := fullTypeName

	if messageType {
		fullTypeName = "*" + fullTypeName
	}

	if t.Repeated {
		fullTypeName = "[]" + fullTypeName
	}

	if t.Optional && fullTypeName[0] != '*' { // Don't add a double * in the case of a non-repeated message type
		fullTypeName = "*" + fullTypeName
	}

	return &typeInfo{
		TypeName:      typeName,
		FullTypeName:  fullTypeName,
		IsMessageType: messageType,
		Repeated:      t.Repeated,
		Pointer:       fullTypeName[0] == '*',
	}, nil
}

func convertFieldName(name string) (goName string, jsonName string, err error) {
	// Make sure the field name is snake case
	re := regexp.MustCompile(`^[a-z][a-z0-9_]*$`)
	if !re.MatchString(name) {
		return "", "", fmt.Errorf("%s is an invalid jrpc field name", name)
	}

	return snakeToCamelCase(name), name, nil
}

type typeData struct {
	GoType     string
	ImportPath string
}

var typeMap = map[string]typeData{
	"any":     {"interface{}", ""},
	"bool":    {"bool", ""},
	"string":  {"string", ""},
	"int32":   {"int32", ""},
	"int64":   {"int64", ""},
	"uint8":   {"uint8", ""},
	"uint32":  {"uint32", ""},
	"uint64":  {"uint64", ""},
	"float32": {"float32", ""},
	"float64": {"float64", ""},
	"bytes":   {"[]byte", ""},
	"time":    {"Time", "time"},
}

func snakeToCamelCase(s string) string {
	var camel string
	var upper bool

	for i, c := range s {
		switch {
		case c == '_':
			upper = true
		case i == 0, upper:
			camel += strings.ToUpper(string(c))
			upper = false
		default:
			camel += string(c)
		}
	}

	return camel
}
