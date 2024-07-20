package cfg

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
)

var (
	QSSHConfig     = &Config{}
	QSSHHomeDir    = os.Getenv("HOME") + "/.qssh/"
	QSSHConfigName = "cfg.json"
	QSSHConfigPath = path.Join(QSSHHomeDir, QSSHConfigName)
)

type Config struct {
	Hosts []*Host `json:"hosts"`
}

func (c *Config) Get(hostName string) (*Host, error) {
	for _, m := range QSSHConfig.Hosts {
		if m.HostName == hostName {
			return m, nil
		}
	}
	return nil, fmt.Errorf("not found machine")
}

func (c *Config) List() ([]*Host, error) {
	return c.Hosts, nil
}

func (c *Config) Add(machine *Host, force bool) (*Host, error) {
	m, _ := c.Get(machine.HostName)
	if m != nil {
		if force {
			*m = *machine
		} else {
			return nil, fmt.Errorf("machine already exists")
		}
	} else {
		QSSHConfig.Hosts = append(QSSHConfig.Hosts, machine)
	}
	return machine, c.save()
}

func (c *Config) Remove(host string) (*Host, error) {
	var (
		machine *Host
		index   = -1
	)
	for i, m := range QSSHConfig.Hosts {
		if m.HostName == host {
			index = i
			machine = m
			break
		}
	}
	if machine == nil {
		return nil, nil
	}
	QSSHConfig.Hosts = append(
		QSSHConfig.Hosts[:index],
		QSSHConfig.Hosts[index+1:]...,
	)
	return machine, c.save()
}

func (c *Config) save() error {
	cfgData, err := json.MarshalIndent(QSSHConfig, "", "	")
	if err != nil {
		return err
	}
	return os.WriteFile(QSSHConfigPath, cfgData, 0644)
}

func init() {
	if err := os.MkdirAll(QSSHHomeDir, 0644); err != nil {
		fmt.Printf("%v\n", err)
	}
	if _, err := os.Stat(QSSHConfigPath); os.IsNotExist(err) {
		QSSHConfig = &Config{}
		fmt.Printf("not found qssh cfg %s\n", QSSHConfigPath)
		err = QSSHConfig.save()
		if err != nil {
			panic(err)
		}
		return
	}
	cfg, err := load(QSSHConfigPath)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	QSSHConfig = cfg
}

func load(p string) (*Config, error) {
	var cfg = &Config{}
	cfgData, err := os.ReadFile(p)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(cfgData, cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
