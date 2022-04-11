package config

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"

	"easyclash/utils"
)

type Proxy struct {
	// public
	Enable     bool   `yaml:"enable"      json:"enable"`
	CanDel     bool   `yaml:"cal-del"     json:"canDel"`
	CanEdit    bool   `yaml:"can-edit"    json:"canEdit"`
	CanDisable bool   `yaml:"can-disable" json:"canDisable"`
	Name       string `yaml:"name"        json:"name"`
	Type       string `yaml:"type"        json:"type"`
	Server     string `yaml:"server"      json:"server"`
	Port       int    `yaml:"port"        json:"port"`

	// for type eq vmess
	UUID    string `yaml:"uuid,omitempty"    json:"uuid,omitempty"`
	AlterId int    `yaml:"alterId,omitempty" json:"alterId,omitempty"`
	Cipher  string `yaml:"cipher,omitempty"  json:"cipher,omitempty"`

	// for type eq trojan
	Password string `json:"password,omitempty" json:"password,omitempty"`
	Sni      string `json:"sni,omitempty"      json:"sni,omitempty"`
	Udp      bool   `json:"udp,omitempty"      json:"udp,omitempty"`
}

type Rule struct {
	Type       string `json:"type"`
	Token      string `json:"token"`
	ProxyGroup string `json:"proxyGroup"`
	Enable     bool   `json:"enable"`
	CanDel     bool   `yaml:"cal-del"     json:"canDel"`
	CanEdit    bool   `yaml:"can-edit"    json:"canEdit"`
	CanDisable bool   `yaml:"can-disable" json:"canDisable"`
}

type RuleSet struct {
	Type       string `json:"type"               yaml:"type"`
	Behavior   string `json:"behavior"           yaml:"behavior"`
	ProxyGroup string `json:"proxyGroup"         yaml:"proxy-group"`
	Url        string `json:"url,omitempty"      yaml:"url,omitempty"`
	Path       string `json:"path,omitempty"     yaml:"path,omitempty"`
	Interval   int    `json:"interval,omitempty" yaml:"interval,omitempty"`
	Enable     bool   `json:"enable"             yaml:"enable"`
	CanDel     bool   `yaml:"cal-del"            json:"canDel"`
	CanEdit    bool   `yaml:"can-edit"           json:"canEdit"`
	CanDisable bool   `yaml:"can-disable"        json:"canDisable"`
}

type RuleProviders map[string]RuleSet

type ProxyGroup struct {
	Name       string   `yaml:"name"               json:"name"`
	Type       string   `yaml:"type"               json:"type"`
	Url        string   `yaml:"url,omitempty"      json:"url,omitempty"`
	Interval   int      `yaml:"interval,omitempty" json:"interval,omitempty"`
	Proxies    []string `yaml:"proxies"            json:"proxies"`
	Use        []string `yaml:"use"                json:"use"`
	Enable     bool     `yaml:"enable"             json:"enable"`
	CanDel     bool     `yaml:"cal-del"            json:"canDel"`
	CanEdit    bool     `yaml:"can-edit"           json:"canEdit"`
	CanDisable bool     `yaml:"can-disable"        json:"canDisable"`
}

type BaseConf struct {
	Version            int64  `yaml:"version"             json:"version"`
	Port               int    `yaml:"port"                json:"port"`
	SocksPort          int    `yaml:"socks-port"          json:"socksPort"`
	RedirPort          int    `yaml:"redir-port"          json:"redirPort"`
	AllowLan           bool   `yaml:"allow-lan"           json:"allowLan"`
	ExternalController string `yaml:"external-controller" json:"externalController"`
	HomeDir            string `json:"homeDir"             json:"homeDir"`
}

type ProxyProvider struct {
	Type        string `json:"type"               yaml:"type"`
	Url         string `json:"url,omitempty"      yaml:"url,omitempty"`
	Path        string `json:"path,omitempty"     yaml:"path,omitempty"`
	Interval    int    `json:"interval,omitempty" yaml:"interval,omitempty"`
	Enable      bool   `json:"enable"             yaml:"enable"`
	CanDel      bool   `yaml:"cal-del"            json:"canDel"`
	CanEdit     bool   `yaml:"can-edit"           json:"canEdit"`
	CanDisable  bool   `yaml:"can-disable"        json:"canDisable"`
	HealthCheck struct {
		Enable   bool   `yaml:"enable"             json:"enable"`
		Lazy     bool   `yaml:"lazy"               json:"lazy"`
		Interval int    `json:"interval,omitempty" yaml:"interval,omitempty"`
		Url      string `json:"url,omitempty"      yaml:"url,omitempty"`
	} `yaml:"health-check" json:"healthCheck"`
}

type ProxyProviders map[string]ProxyProvider

type Config struct {
	Version            int64          `yaml:"version"             json:"version"`
	Port               int            `yaml:"port"                json:"port"`
	SocksPort          int            `yaml:"socks-port"          json:"socksPort"`
	RedirPort          int            `yaml:"redir-port"          json:"redirPort"`
	AllowLan           bool           `yaml:"allow-lan"           json:"allowLan"`
	ExternalController string         `yaml:"external-controller" json:"externalController"`
	Proxies            []Proxy        `yaml:"proxies"`
	ProxyGroups        []ProxyGroup   `yaml:"proxy-groups"`
	ProxyProviders     ProxyProviders `yaml:"proxy-providers"`
	Rules              []Rule         `yaml:"rules"`
	RuleProviders      RuleProviders  `yaml:"rule-providers"`
}

type ClashConfig struct {
	Version            int64          `yaml:"version"             json:"version"`
	Port               int            `yaml:"port"                json:"port"`
	SocksPort          int            `yaml:"socks-port"          json:"socksPort"`
	RedirPort          int            `yaml:"redir-port"          json:"redirPort"`
	AllowLan           bool           `yaml:"allow-lan"           json:"allowLan"`
	ExternalController string         `yaml:"external-controller" json:"externalController"`
	Proxies            []Proxy        `yaml:"proxies"`
	ProxyGroups        []ProxyGroup   `yaml:"proxy-groups"`
	ProxyProviders     ProxyProviders `yaml:"proxy-providers"`
	Rules              []string       `yaml:"rules"`
	RuleProviders      RuleProviders  `yaml:"rule-providers"`
}

func UserConfig(home string) (*Config, error) {
	s := NewStore()
	c := DefaultBase
	_ = s.Load(home+"/base_conf.yml", &c)

	c1 := new([]Proxy)
	_ = s.Load(home+"/proxy.yml", &c1)

	c2 := new([]ProxyGroup)
	_ = s.Load(home+"/proxy_group.yml", &c2)

	c3 := new([]Rule)
	_ = s.Load(home+"/rule.yml", &c3)

	c4 := new(RuleProviders)
	_ = s.Load(home+"/rule_set.yml", &c4)

	c5 := new(ProxyProviders)
	_ = s.Load(home+"/proxy_provider.yml", &c5)

	_c := &Config{
		Version:            c.Version,
		Port:               c.Port,
		SocksPort:          c.SocksPort,
		RedirPort:          c.RedirPort,
		AllowLan:           c.AllowLan,
		ExternalController: c.ExternalController,
		Proxies:            *c1,
		ProxyGroups:        *c2,
		Rules:              *c3,
		RuleProviders:      *c4,
		ProxyProviders:     *c5,
	}

	return _c, nil
}

func TransConf(c *Config) *ClashConfig {

	filterRP := func(rp RuleProviders) RuleProviders {
		_rp := RuleProviders{}
		for k, v := range rp {
			if v.Enable {
				_rp[k] = v
			}
		}
		return _rp
	}

	filterRule := func(arr []string) string {
		var tmp []string
		for _, v := range arr {
			if v != "" {
				tmp = append(tmp, v)
			}
		}
		return strings.Join(tmp, ",")
	}

	filterP := func(rp []Proxy) []Proxy {
		var _rp []Proxy
		for _, v := range rp {
			if v.Enable {
				_rp = append(_rp, v)
			}
		}
		return _rp
	}

	filterPG := func(rp []ProxyGroup) []ProxyGroup {
		var _rp []ProxyGroup
		for _, v := range rp {
			if v.Enable {
				_rp = append(_rp, v)
			}
		}
		return _rp
	}

	filterPP := func(rp ProxyProviders) ProxyProviders {
		_rp := ProxyProviders{}
		for k, v := range rp {
			if v.Enable {
				_rp[k] = v
			}
		}
		return _rp
	}

	cc := new(ClashConfig)
	cc.Version = c.Version
	cc.Port = c.Port
	cc.SocksPort = c.SocksPort
	cc.RedirPort = c.RedirPort
	cc.ExternalController = c.ExternalController
	cc.Proxies = filterP(c.Proxies)
	cc.ProxyGroups = filterPG(c.ProxyGroups)
	cc.RuleProviders = filterRP(c.RuleProviders)
	cc.ProxyProviders = filterPP(c.ProxyProviders)

	var ruleSet []Rule
	for k, v := range cc.RuleProviders {
		pg := v.ProxyGroup
		if pg == "" {
			pg = "DIRECT"
		}
		ruleSet = append(ruleSet, Rule{
			Type:       "RULE-SET",
			Token:      k,
			ProxyGroup: v.ProxyGroup,
		})
	}

	rule := c.Rules

	defaultRule := rule[len(rule)-2:]
	customRule := rule[:len(rule)-2]

	c.Rules = []Rule{}

	if customRule != nil {
		c.Rules = append(c.Rules, customRule...)
	}

	if ruleSet != nil {
		c.Rules = append(c.Rules, ruleSet...)
	}

	c.Rules = append(c.Rules, defaultRule...)

	var rs []string
	for _, v := range c.Rules {
		tokens := []string{v.Type, v.Token, v.ProxyGroup}
		rs = append(rs, filterRule(tokens))
	}

	cc.Rules = rs
	return cc
}

func NewConf(home string) (*ClashConfig, error) {
	c, err := UserConfig(home)
	if err != nil {
		return nil, err
	}
	_c := TransConf(c)
	return _c, nil
}

func SaveClashConf(home string, c *ClashConfig) error {
	return saveClashConf(filepath.Join(home, "config.yml"), c)
}

func saveClashConf(path string, c *ClashConfig) error {
	d, err := yaml.Marshal(c)
	if err != nil {
		return err
	}

	f, _ := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	defer f.Close()
	_, err = f.Write(d)
	if err != nil {
		return err
	}
	return nil
}

func CheckClashConf(home string, c *ClashConfig) error {
	err := saveClashConf(filepath.Join(home, "config_tmp.yml"), c)
	if err != nil {
		return err
	}
	out, err := utils.Exec(home+"/easy_clash_core", "-t", "-f", home+"/config_tmp.yml")
	if err != nil {
		return errors.New(out)
	}

	return nil

}
