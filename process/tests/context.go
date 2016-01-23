package tests

import (
	"reflect"
	"sync"

	"strings"

	"golang.org/x/net/context"
	"kego.io/context/cmdctx"
	"kego.io/context/envctx"
	"kego.io/context/jsonctx"
	"kego.io/context/sysctx"
	"kego.io/context/vosctx"
	"kego.io/context/wgctx"
	"kego.io/kerr"
)

type ContextBuilder struct {
	ctx context.Context
}

func NewContextBuilder() *ContextBuilder {
	return &ContextBuilder{context.Background()}
}

func Context(path string) *ContextBuilder {

	env := &envctx.Env{
		Path:    path,
		Aliases: map[string]string{},
	}

	ctx := context.Background()
	ctx = envctx.NewContext(ctx, env)

	return &ContextBuilder{ctx: ctx}
}

func ContextFrom(ctx context.Context) *ContextBuilder {
	return &ContextBuilder{ctx: ctx}
}

func (c *ContextBuilder) Ctx() context.Context {
	return c.ctx
}
func (c *ContextBuilder) Env() *envctx.Env {
	return envctx.FromContext(c.ctx)
}

func (c *ContextBuilder) initEnv() *envctx.Env {
	env := envctx.FromContextOrNil(c.ctx)
	if env == nil {
		env = &envctx.Env{Aliases: map[string]string{}}
		c.ctx = envctx.NewContext(c.ctx, env)
	}
	return env
}

func (c *ContextBuilder) initCmd() *cmdctx.Cmd {
	cmd := cmdctx.FromContextOrNil(c.ctx)
	if cmd == nil {
		cmd = &cmdctx.Cmd{}
		c.ctx = cmdctx.NewContext(c.ctx, cmd)
	}
	return cmd
}

func (c *ContextBuilder) initWg() *sync.WaitGroup {
	wg := wgctx.FromContextOrNil(c.ctx)
	if wg != nil {
		return wg
	}
	c.ctx = wgctx.NewContext(c.ctx)
	return wgctx.FromContext(c.ctx)
}

func (c *ContextBuilder) initJson() *jsonctx.JsonCache {
	jcache := jsonctx.FromContextOrNil(c.ctx)
	if jcache == nil {
		c.ctx = jsonctx.ManualContext(c.ctx)
		jcache = jsonctx.FromContext(c.ctx)
	}
	return jcache
}

func (c *ContextBuilder) initSys() *sysctx.SysCache {
	scache := sysctx.FromContextOrNil(c.ctx)
	if scache == nil {
		c.ctx = sysctx.NewContext(c.ctx)
		scache = sysctx.FromContext(c.ctx)
	}
	return scache
}
func (c *ContextBuilder) initSysPkg(path string) *sysctx.PackageInfo {
	scache := c.initSys()
	pi, ok := scache.Get(path)
	if !ok {
		pi = scache.Set(&envctx.Env{Path: path, Aliases: map[string]string{}})
	}
	return pi
}
func (c *ContextBuilder) initVos() vosctx.Vos {
	vos := vosctx.FromContextOrNil(c.ctx)
	if vos != nil {
		return vos
	}
	vos = &MockOs{
		EnvironmentVariables: map[string]string{},
	}
	c.ctx = vosctx.NewContext(c.ctx, vos)
	return vos
}

func (c *ContextBuilder) Cmd() *ContextBuilder {
	c.initCmd()
	return c
}

func (c *ContextBuilder) Wg() *ContextBuilder {
	c.initWg()
	return c
}
func (c *ContextBuilder) Log() *ContextBuilder {
	cmd := c.initCmd()
	cmd.Log = true
	return c
}

func (c *ContextBuilder) Path(path string) *ContextBuilder {
	env := c.initEnv()
	env.Path = path
	pi := c.initSysPkg(path)
	copyEnv(env, pi.Environment)
	return c
}

func (c *ContextBuilder) Alias(aliasPath string, aliasName string) *ContextBuilder {
	env := c.initEnv()
	env.Aliases[aliasPath] = aliasName
	if env.Path != "" {
		pi := c.initSysPkg(env.Path)
		copyEnv(env, pi.Environment)
	}
	return c
}

func copyEnv(from *envctx.Env, to *envctx.Env) {
	to.Path = from.Path
	to.Recursive = from.Recursive
	to.Hash = from.Hash
	to.Aliases = map[string]string{}
	for p, n := range from.Aliases {
		to.Aliases[p] = n
	}
}

func (c *ContextBuilder) Dir(dir string) *ContextBuilder {
	env := c.initEnv()
	env.Dir = dir
	return c
}

func (c *ContextBuilder) Recursive(state bool) *ContextBuilder {
	env := c.initEnv()
	env.Recursive = state
	return c
}

func (c *ContextBuilder) Jempty() *ContextBuilder {
	c.initJson()
	return c
}

func (c *ContextBuilder) Jauto() *ContextBuilder {
	jcache := c.initJson()
	jcache.Packages.InitAuto()
	jcache.Dummies.InitAuto()
	return c
}

func (c *ContextBuilder) Jsystem() *ContextBuilder {
	jcache := c.initJson()
	jcache.Packages.InitManual("kego.io/json", "kego.io/system")
	jcache.Dummies.InitAuto()
	return c
}

func (c *ContextBuilder) Jtype(name string, typ reflect.Type) *ContextBuilder {
	env := c.initEnv()
	return c.JtypePathRule(env.Path, name, typ, nil)
}
func (c *ContextBuilder) JtypeRule(name string, typ reflect.Type, rule reflect.Type) *ContextBuilder {
	env := c.initEnv()
	return c.JtypePathRule(env.Path, name, typ, rule)
}
func (c *ContextBuilder) JtypePath(path string, name string, typ reflect.Type) *ContextBuilder {
	return c.JtypePathRule(path, name, typ, nil)
}

func (c *ContextBuilder) JtypePathRule(path string, name string, typ reflect.Type, rule reflect.Type) *ContextBuilder {
	if path == "" {
		panic("must specify path")
	}

	jcache := c.initJson()

	p, ok := jcache.Packages.Get(path)
	if !ok {
		p = jcache.Packages.Set(path, 0)
	}
	isrule := false
	if strings.HasPrefix(name, jsonctx.RULE_PREFIX) {
		if rule != nil {
			panic(kerr.New("UBFYEAGXHJ", "rule specified!").Error())
		}
		isrule = true
		name = name[1:]
	}
	if isrule {
		p.Types.Set(name, &jsonctx.TypeInfo{
			Name: name,
			Rule: typ,
		})
	} else {
		p.Types.Set(name, &jsonctx.TypeInfo{
			Name: name,
			Type: typ,
			Rule: rule,
		})
	}

	return c

}

func (c *ContextBuilder) Sempty() *ContextBuilder {
	c.initSys()
	return c
}

func (c *ContextBuilder) Spkg(path string) *ContextBuilder {
	scache := c.initSys()
	if _, ok := scache.Get(path); !ok {
		scache.Set(&envctx.Env{Path: path})
	}
	return c
}

// Sauto runs parser.Parse, but to stop an import cycle we pass the parser.Parse function in as a
// parameter
func (c *ContextBuilder) Sauto(parseParse func(context.Context, string) (*sysctx.PackageInfo, error)) *ContextBuilder {
	env := c.initEnv()
	c.initSys()
	c.initCmd()
	pi, err := parseParse(c.ctx, env.Path)
	if err != nil {
		panic(err)
	}
	env.Hash = pi.Environment.Hash
	return c
}

// Ssystem runs parser.Parse on the system package, but to stop an import cycle we pass the
// parser.Parse function in as a parameter
func (c *ContextBuilder) Ssystem(parseParse func(context.Context, string) (*sysctx.PackageInfo, error)) *ContextBuilder {
	c.initEnv()
	c.initSys()
	c.initCmd()
	c.Jsystem() // need json types for system in order to parse
	_, err := parseParse(c.ctx, "kego.io/system")
	if err != nil {
		panic(err)
	}
	return c
}

func (c *ContextBuilder) Stype(name string, typ interface{}) *ContextBuilder {
	env := c.initEnv()
	return c.StypePath(env.Path, name, typ)
}

func (c *ContextBuilder) StypePath(path string, name string, typ interface{}) *ContextBuilder {

	if path == "" {
		panic("path not specified")
	}

	scache := c.initSys()

	p, ok := scache.Get(path)
	if !ok {
		p = scache.Set(&envctx.Env{Path: path})
	}

	p.Types.Set(name, typ)

	return c
}

func (c *ContextBuilder) AllPath(path string, name string, typ reflect.Type, rule reflect.Type, typTyp interface{}, ruleTyp interface{}) *ContextBuilder {
	c.JtypePathRule(path, name, typ, rule)
	c.StypePath(path, name, typTyp)
	c.StypePath(path, jsonctx.RULE_PREFIX+name, ruleTyp)
	return c
}

func (c *ContextBuilder) Dummy(iface reflect.Type, dummy reflect.Type) *ContextBuilder {
	jcache := c.initJson()
	jcache.Dummies.Set(iface, dummy)
	return c
}

func (c *ContextBuilder) OsVar(name string, value string) *ContextBuilder {
	vos := c.initVos()
	m, ok := vos.(*MockOs)
	if !ok {
		panic("must me *MockEnv")
	}
	m.EnvironmentVariables[name] = value
	return c
}

func (c *ContextBuilder) OsWd(dir string) *ContextBuilder {
	vos := c.initVos()
	m, ok := vos.(*MockOs)
	if !ok {
		panic("must me *MockEnv")
	}
	m.WorkingDirectory = dir
	return c
}