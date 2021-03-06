package frizz

import (
	"path/filepath"
	"strings"

	"frizz.io/global"
	"frizz.io/pack"
	"frizz.io/utils"
	"frizz.io/validator"
	"github.com/dave/patsy"
	"github.com/dave/patsy/vos"
	"github.com/pkg/errors"
)

func New(local global.Package, imports ...global.Package) global.PackageContext {
	all := map[string]global.Package{}
	local.GetImportedPackages(all)
	for _, p := range imports {
		p.GetImportedPackages(all)
	}
	return pack.NewPackageContext(local.Path(), all)
}

func Package(p global.Package) (map[string]interface{}, error) {
	c := New(p)
	out := make(map[string]interface{})
	env := vos.Os()
	dir, err := patsy.Dir(env, p.Path())
	if err != nil {
		return nil, errors.Wrap(err, "getting dir from path")
	}
	files, err := filepath.Glob(filepath.Join(dir, "*.frizz.json"))
	if err != nil {
		return nil, errors.Wrap(err, "getting files")
	}
	for _, fpath := range files {
		_, fname := filepath.Split(fpath)
		name := strings.TrimSuffix(fname, ".frizz.json")

		v, err := utils.DecodeFile(fpath)
		if err != nil {
			return nil, errors.Wrapf(err, "decoding %s", fname)
		}

		r := pack.NewRootContext(c)
		if err := r.ParseImports(v); err != nil {
			return nil, errors.Wrapf(err, "parsing imports from %s", fname)
		}
		s := global.NewStack(name)
		d := pack.NewDataContext(c, r, s)
		i, null, err := pack.UnpackInterface(d, v)
		if err != nil {
			return nil, errors.Wrapf(err, "unpacking %s", fname)
		}
		if !null {
			out[name] = i
		}
	}
	return out, nil
}

func Unmarshal(context global.PackageContext, in []byte) (interface{}, error) {
	return pack.Unmarshal(context, in)
}

func Validate(context global.PackageContext, v interface{}) (valid bool, message string, err error) {
	return validator.Validate(context, v)
}
