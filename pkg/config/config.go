package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
)

var (
	QSSHConfig     = &Config{}
	QSSHHomeDir    = os.Getenv("HOME") + "/.qssh/"
	QSSHConfigName = "config.json"
	QSSHConfigPath = path.Join(QSSHHomeDir, QSSHConfigName)
)

type Config struct {
	Machines []*Machine `json:"machines"`
}

func (c *Config) Get(host string) (*Machine, error) {
	for _, m := range QSSHConfig.Machines {
		if m.Host == host {
			return m, nil
		}
	}
	return nil, fmt.Errorf("not found machine")
}

func (c *Config) GetMachineByIndex(index int) (*Machine, error) {
	if index < 0 || index >= len(c.Machines) {
		return nil, fmt.Errorf("not found machine by index %v", index)
	}
	return c.Machines[index], nil
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
		QSSHConfig.Machines = append(QSSHConfig.Machines, machine)
	}
	return machine, c.save()
}

func (c *Config) Remove(host string) (*Machine, error) {
	var (
		machine *Machine
		index   = 0
	)
	for i, m := range QSSHConfig.Machines {
		if m.Host == host {
			index = i
			machine = m
			break
		}
	}
	if machine == nil {
		return nil, nil
	}
	QSSHConfig.Machines = append(
		QSSHConfig.Machines[:index],
		QSSHConfig.Machines[index+1:]...,
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
	if err := os.MkdirAll(QSSHHomeDir, 0644); err != nil {
		fmt.Printf("%v\n", err)
	}
	if _, err := os.Stat(QSSHConfigPath); os.IsNotExist(err) {
		QSSHConfig = &Config{}
		fmt.Printf("not found qssh config %s\n", QSSHConfigPath)
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
