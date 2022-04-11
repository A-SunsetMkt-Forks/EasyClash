package app

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"os"
	"time"

	"easyclash/config"
	"easyclash/utils"
)

type Msg struct {
	Code    int
	Message string
}

type App struct {
	ctx     context.Context
	ConfDir string
	store   *config.Store
	clash   *clash
	msg     []Msg
}

func NewApp() *App {
	path, _ := os.UserHomeDir()
	return &App{
		ConfDir: path + "/.easy_clash",
		store:   config.NewStore(),
		clash:   NewClash(),
	}
}

func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
	a.SaveRuleSet(a.GetRuleSet())
	a.StartClash()
}

func (a App) DomReady(ctx context.Context) {
}

func (a *App) Shutdown(ctx context.Context) {
	a.StopClash()
}

func (a *App) GetBaseConf() config.BaseConf {
	c := config.BaseConf{}
	_ = a.store.Load(a.ConfDir+"/base_conf.yml", &c)
	c.HomeDir = a.ConfDir
	return c
}

func (a *App) SaveBaseConf(c config.BaseConf) bool {
	c.Version = time.Now().UnixMicro()
	_ = a.store.Save(a.ConfDir+"/base_conf.yml", c)
	return true
}

func (a *App) GetProxy() *[]config.Proxy {
	c := new([]config.Proxy)
	_ = a.store.Load(a.ConfDir+"/proxy.yml", &c)
	return c
}

func (a *App) SaveProxy(c *[]config.Proxy) bool {
	defer a.updateVersion()
	_ = a.store.Save(a.ConfDir+"/proxy.yml", c)
	return true
}

func (a *App) GetProxyGroup() *[]config.ProxyGroup {
	c := new([]config.ProxyGroup)
	_ = a.store.Load(a.ConfDir+"/proxy_group.yml", &c)
	return c
}

func (a *App) SaveProxyGroup(c *[]config.ProxyGroup) bool {
	defer a.updateVersion()
	_ = a.store.Save(a.ConfDir+"/proxy_group.yml", c)
	return true
}

func (a *App) GetProxyProvider() *config.ProxyProviders {
	c := new(config.ProxyProviders)
	_ = a.store.Load(a.ConfDir+"/proxy_provider.yml", &c)
	return c
}

func (a *App) SaveProxyProvider(c *config.ProxyProviders) bool {
	defer a.updateVersion()
	_c := *c
	for k, v := range _c {
		v.Path = a.ConfDir + "/proxyproviders/" + k + ".yml"
		_c[k] = v
	}
	_ = a.store.Save(a.ConfDir+"/proxy_provider.yml", _c)
	return true
}

func (a *App) GetRule() *[]config.Rule {
	c := new([]config.Rule)
	_ = a.store.Load(a.ConfDir+"/rule.yml", &c)
	return c
}

func (a *App) SaveRule(c *[]config.Rule) bool {
	defer a.updateVersion()
	_ = a.store.Save(a.ConfDir+"/rule.yml", c)
	return true
}

func (a *App) GetRuleSet() *config.RuleProviders {
	c := new(config.RuleProviders)
	_ = a.store.Load(a.ConfDir+"/rule_set.yml", &c)
	return c
}

func (a *App) SaveRuleSet(c *config.RuleProviders) bool {
	defer a.updateVersion()
	_c := *c
	for k, v := range _c {
		v.Path = a.ConfDir + "/ruleset/" + k + ".yml"
		_c[k] = v
	}
	_ = a.store.Save(a.ConfDir+"/rule_set.yml", _c)
	return true
}

func (a *App) updateVersion() {
	b := a.GetBaseConf()
	b.Version = time.Now().UnixMicro()
	a.SaveBaseConf(b)
}

type Version struct {
	NewVersion  int64
	ProdVersion int64
}

func (a *App) HaveNewVersionConf() Version {
	b := a.GetBaseConf()
	c := new(config.ClashConfig)
	_ = a.store.Load(a.ConfDir+"/config.yml", &c)
	return Version{
		NewVersion:  b.Version,
		ProdVersion: c.Version,
	}
}

func (a *App) ClashIsRunning() bool {
	pid, _ := a.clash.isRunning()
	return pid > 0
}

func (a *App) StartClash() {
	a.StopClash()
	_, err := a.InitConf(a.ctx)
	if err != nil {
		return
	}

	err = a.clash.Start()
	if err != nil {
		a.addErr("Clash 核心程序启动失败", err)
		return
	}
	a.SetSystemProxy()
}

func (a *App) StopClash() {
	err := a.clash.Stop()
	if err != nil {
		a.addNotify("Clash 核心程序停止失败")
		return
	}
	err = utils.UnSetSystemProxy()
	if err != nil {
		a.addNotify("Clash 撤销系统代理失败")
		return
	}
}

func (a *App) SetSystemProxy() bool {
	c, _ := config.NewConf(a.ConfDir)
	err := utils.SetSystemProxy(c.Port, c.SocksPort)
	if err != nil {
		a.addNotify("Clash 设置系统代理失败")
		return false
	}
	return true
}

func (a *App) UnSetSystemProxy() bool {
	err := utils.UnSetSystemProxy()
	if err != nil {
		a.addNotify("Clash 撤销系统代理失败")
		return false
	}
	return true
}

func (a *App) InitConf(ctx context.Context) (*config.ClashConfig, error) {
	path, err := os.UserHomeDir()
	if err != nil {
		a.addErr("获取用户根目录失败", err)
		return nil, err
	}
	path = path + "/.easy_clash"
	if _, err := os.Stat(path); err != nil {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			a.addErr("初始化配置文件失败", err)
			return nil, err
		}
	}

	c, err := config.NewConf(a.ConfDir)
	if err != nil {
		a.addErr("搜集配置 失败", err)
		return nil, err
	}

	if err != nil {
		a.addErr("读取配置失败", err)
		return nil, err
	}
	if len(c.Proxies) == 0 && len(c.ProxyProviders) == 0 {
		a.addNotify("请先添加可用代理")
		return nil, err
	}
	flag := false
	for _, v := range c.ProxyGroups {
		if len(v.Use) > 0 || len(v.Proxies) > 0 {
			flag = true
		}
	}

	if !flag {
		a.addNotify("ProxyGroup 中未添加可用代理")
		return nil, err
	}

	err = config.CheckClashConf(a.ConfDir, c)
	if err != nil {
		a.addErr("配置有误", err)
		return nil, err
	}

	err = config.SaveClashConf(a.ConfDir, c)
	if err != nil {
		a.addErr("更新配置 失败", err)
		return nil, err
	}
	return c, nil
}

func (a *App) GetConfDir() string {
	return a.ConfDir
}

func (a *App) GetMsg() []Msg {
	defer func() {
		a.msg = []Msg{}
	}()
	return a.msg
}

func (a *App) addErr(msg string, err error) {
	a.msg = append(a.msg, Msg{500, fmt.Sprintf("%v", errors.Wrap(err, msg))})
}

func (a *App) addMsg(m Msg) {
	a.msg = append(a.msg, m)
}

func (a *App) addNotify(m string) {
	a.msg = append(a.msg, Msg{0, m})
}
