package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
)

var (
	QsshConfig     = &Config{}
	QsshHomeDir    = os.Getenv("HOME") + "/.qssh/"
	QsshConfigName = "config.json"
	QsshConfigPath = path.Join(QsshHomeDir, QsshConfigName)
)

type Config struct {
	Machines []*Machine `json:"machines"`
}

type Machine struct {
	User     string `json:"user,omitempty"`
	Password string `json:"password,omitempty"`
	Host     string `json:"host,omitempty"`
	Port     int    `json:"port,omitempty"`
	Key      string `json:"key,omitempty"`
}

func (c *Config) Get(host string) (*Machine, error) {
	for _, m := range QsshConfig.Machines {
		if m.Host == host {
			return m, nil
		}
	}
	return nil, fmt.Errorf("not found machine")
}

func (c *Config) List() ([]*Machine, error) {
	return c.Machines, nil
}

func (c *Config) Add(machine *Machine, force bool) (*Machine, error) {
	m, _ := c.Get(machine.Host)
	if m != nil {
		if force {
			m = machine
		} else {
			return nil, fmt.Errorf("machine already exists")
		}
	} else {
		QsshConfig.Machines = append(QsshConfig.Machines, machine)
	}
	return machine, c.save()
}

func (c *Config) Remove(host string) (*Machine, error) {
	var (
		machine *Machine
		index   = -1
	)
	for i, m := range QsshConfig.Machines {
		if m.Host == host {
			index = i
			machine = m
			break
		}
	}
	if machine == nil {
		return nil, nil
	}
	QsshConfig.Machines = append(
		QsshConfig.Machines[:index],
		QsshConfig.Machines[index+1:]...,
	)
	return machine, c.save()
}

func (c *Config) save() error {
	cfgData, err := json.MarshalIndent(QsshConfig, "", "	")
	if err != nil {
		return err
	}
	return os.WriteFile(QsshConfigPath, cfgData, 0644)
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

func init() {
	if err := os.MkdirAll(QsshHomeDir, 0644); err != nil {
		fmt.Printf("%v\n", err)
	}
	if _, err := os.Stat(QsshConfigPath); os.IsNotExist(err) {
		QsshConfig = &Config{}
		fmt.Printf("not found qssh config %s\n", QsshConfigPath)
		err = QsshConfig.save()
		if err != nil {
			panic(err)
		}
		return
	}
	cfg, err := load(QsshConfigPath)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	QsshConfig = cfg
}
