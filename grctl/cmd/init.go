// Copyright (C) 2014-2018 Goodrain Co., Ltd.
// RAINBOND, Application Management Platform

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version. For any non-GPL usage of Rainbond,
// one or multiple Commercial Licenses authorized by Goodrain Co., Ltd.
// must be obtained first.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package cmd

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"github.com/goodrain/rainbond/node/nodem/client"

	"github.com/goodrain/rainbond/util"

	"github.com/Sirupsen/logrus"
	"github.com/goodrain/rainbond/builder/sources"
	"github.com/goodrain/rainbond/event"
	"github.com/goodrain/rainbond/grctl/clients"
	"github.com/urfave/cli" //"github.com/goodrain/rainbond/grctl/clients"
)

//NewCmdInit grctl init
func NewCmdInit() cli.Command {
	c := cli.Command{
		Name: "init",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "role",
				Usage: "Node identity property",
				Value: "manage,compute,gateway",
			},
			cli.StringFlag{
				Name:  "work_dir",
				Usage: "Installation configuration directory",
				Value: "/opt/rainbond/rainbond-ansible",
			},
			cli.StringFlag{
				Name:  "iip",
				Usage: "Internal IP",
				Value: "",
			},
			cli.StringFlag{
				Name:  "eip",
				Usage: "External IP",
				Value: "",
			},
			cli.StringFlag{
				Name:  "vip",
				Usage: "Virtual IP",
				Value: "",
			},
			cli.StringFlag{
				Name:  "rainbond-version",
				Usage: "Rainbond Install Version. default 5.1",
				Value: "5.1",
			},
			cli.StringFlag{
				Name:  "rainbond-repo",
				Usage: "Rainbond install repo",
				Value: "https://github.com/goodrain/rainbond-ansible.git",
			},
			cli.StringFlag{
				Name:  "install-type",
				Usage: "Install Type: online/offline",
				Value: "online",
			},
			cli.StringFlag{
				Name:  "deploy-type",
				Usage: "Deploy Type: onenode/multinode/thirdparty,默认onenode",
				Value: "onenode",
			},
			cli.StringFlag{
				Name:  "domain",
				Usage: "Application domain",
				Value: "",
			},
			cli.StringFlag{
				Name:  "pod-cidr",
				Usage: "Configuration pod-cidr",
				Value: "",
			},
			cli.StringFlag{
				Name:  "enable-feature",
				Usage: "New feature，disabled by default. default: windows",
				Value: "",
			},
			cli.StringFlag{
				Name:  "enable-online-images",
				Usage: "Get image online. default: offline",
				Value: "",
			},
			cli.StringFlag{
				Name:  "enable-ntp",
				Usage: "enable ntp config",
				Value: "",
			},
			cli.StringFlag{
				Name:  "storage",
				Usage: "Storage type, default:NFS",
				Value: "nfs",
			},
			cli.StringFlag{
				Name:  "network",
				Usage: "Network type, support calico/flannel/midonet,default: calico",
				Value: "calico",
			},
			cli.StringFlag{
				Name:  "enable-check",
				Usage: "enable check cpu/mem. default: enable/disable",
				Value: "enable",
			},
			cli.StringFlag{
				Name:  "storage-args",
				Usage: "Stores mount parameters",
				Value: "/grdata nfs rw 0 0",
			},
			cli.StringFlag{
				Name:  "config-file,f",
				Usage: "Global Config Path, default",
				Value: "/opt/rainbond/rainbond-ansible/scripts/installer/global.sh",
			},
			cli.StringFlag{
				Name:  "token",
				Usage: "Region Token",
				Value: "",
			},
			cli.StringFlag{
				Name:  "enable-exdb",
				Usage: "default disable external database",
				Value: "",
			},
			cli.StringFlag{
				Name:  "exdb-type",
				Usage: "external database type(mysql,postgresql)",
				Value: "",
			},
			cli.StringFlag{
				Name:  "exdb-host",
				Usage: "external database host",
				Value: "",
			},
			cli.StringFlag{
				Name:  "exdb-port",
				Usage: "external database port",
				Value: "3306",
			},
			cli.StringFlag{
				Name:  "exdb-user",
				Usage: "external database user",
				Value: "",
			},
			cli.StringFlag{
				Name:  "exdb-passwd",
				Usage: "external database password",
				Value: "",
			},
			cli.StringFlag{
				Name:  "excsdb-host",
				Usage: "external console database host",
				Value: "",
			},
			cli.StringFlag{
				Name:  "excsdb-port",
				Usage: "external console database port",
				Value: "3306",
			},
			cli.StringFlag{
				Name:  "excsdb-user",
				Usage: "external console database user",
				Value: "",
			},
			cli.StringFlag{
				Name:  "excsdb-passwd",
				Usage: "external console database password",
				Value: "",
			},
			cli.StringFlag{
				Name:  "enable-excsdb-only",
				Usage: "Additional support for the console to configure the database separately",
				Value: "",
			},
			cli.BoolFlag{
				Name:  "not-install-ui",
				Usage: "Whether or not to install the UI, normally no UI is installed when installing a second data center",
			},
			cli.BoolFlag{
				Name:   "test",
				Usage:  "use test shell",
				Hidden: true,
			},
			cli.StringFlag{
				Name:  "install_ssh_port",
				Usage: "new node ssh port",
				Value: "22",
			},
		},
		Usage: "grctl init cluster",
		Action: func(c *cli.Context) error {
			initCluster(c)
			return nil
		},
	}
	return c
}

//NewCmdInstallStatus install status
func NewCmdInstallStatus() cli.Command {
	c := cli.Command{
		Name: "install_status",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "taskID",
				Usage: "install_k8s,空则自动寻找",
			},
		},
		Usage: "获取task执行状态。grctl install_status",
		Action: func(c *cli.Context) error {
			taskID := c.String("taskID")
			if taskID == "" {
				tasks, err := clients.RegionClient.Tasks().List()
				if err != nil {
					logrus.Errorf("error get task list,details %s", err.Error())
					return nil
				}
				for _, v := range tasks {
					for _, vs := range v.Status {
						if vs.Status == "start" || vs.Status == "create" {
							//Status(v.ID)
							return nil
						}
					}
				}
			} else {
				//Status(taskID)
			}
			return nil
		},
	}
	return c
}
func updateConfigFile(path string, config map[string]interface{}) error {
	initConfig := make(map[string]interface{})
	var file *os.File
	var err error
	if ok, _ := util.FileExists(path); ok {
		file, err = os.OpenFile(path, os.O_RDWR, 0755)
		if err != nil {
			return err
		}
		defer file.Close()
		reader := bufio.NewReader(file)
		for {
			line, _, err := reader.ReadLine()
			if err != nil {
				break
			}
			if strings.Contains(string(line), "=") {
				keyvalue := strings.SplitN(string(line), "=", 1)
				if len(keyvalue) < 2 {
					break
				}
				initConfig[keyvalue[0]] = keyvalue[1]
			}
		}
	} else {
		file, err = util.OpenOrCreateFile(path)
		if err != nil {
			return err
		}
		defer file.Close()
	}
	for k, v := range config {
		initConfig[k] = v
	}
	for k, v := range initConfig {
		if k == "" {
			continue
		}
		switch v.(type) {
		case string:
			if v == "" {
				file.WriteString(fmt.Sprintf("%s=\"\"\n", k))
			} else {
				file.WriteString(fmt.Sprintf("%s=\"%s\"\n", k, v))
			}
		case bool:
			file.WriteString(fmt.Sprintf("%s=%t\n", k, v))
		default:
			continue
		}
	}
	return nil
}
func getConfig(c *cli.Context) map[string]interface{} {
	configs := make(map[string]interface{})
	configs["ROLE"] = c.String("role")
	//configs["work_dir"] = c.String("work_dir")
	configs["IIP"] = c.String("iip")
	configs["EIP"] = c.String("eip")
	configs["VIP"] = c.String("vip")
	configs["VERSION"] = c.String("rainbond-version")
	configs["INSTALL_TOKEN"] = c.String("token")
	configs["INSTALL_TYPE"] = c.String("install-type")
	configs["DEPLOY_TYPE"] = c.String("deploy-type")
	configs["DOMAIN"] = c.String("domain")
	configs["STORAGE"] = c.String("storage")
	configs["NETWORK_TYPE"] = c.String("network")
	configs["POD_NETWORK_CIDR"] = c.String("pod-cidr")
	configs["STORAGE_ARGS"] = c.String("storage-args")
	configs["ENABLE_CHECK"] = c.String("enable-check")
	configs["PULL_ONLINE_IMAGES"] = c.String("enable-online-images")
	configs["ENABLE_NTP"] = c.String("enable-ntp")
	configs["ENABLE_EXDB"] = c.String("enable-exdb")
	configs["EXDB_PASSWD"] = c.String("exdb-passwd")
	configs["EXDB_HOST"] = c.String("exdb-host")
	configs["EXDB_PORT"] = c.String("exdb-port")
	configs["EXDB_USER"] = c.String("exdb-user")
	configs["EXCSDB_ONLY_ENABLE"] = c.String("enable-excsdb-only")
	configs["EXCSDB_PASSWD"] = c.String("excsdb-passwd")
	configs["EXCSDB_HOST"] = c.String("excsdb-host")
	configs["EXCSDB_PORT"] = c.String("excsdb-port")
	configs["EXCSDB_USER"] = c.String("excsdb-user")
	configs["EXDB_TYPE"] = c.String("exdb-type")
	configs["INSTALL_SSH_PORT"] = c.String("install_ssh_port")
	configs["INSTALL_UI"] = !c.Bool("not-install-ui")
	return configs
}
func initCluster(c *cli.Context) {
	// check if the rainbond is already installed
	//fmt.Println("Checking install enviremant.")
	_, err := os.Stat("/opt/rainbond/.rainbond.success")
	if err == nil {
		println("Rainbond is already installed, if you want reinstall, then please delete the file: /opt/rainbond/.rainbond.success")
		return
	}
	role := c.String("role")
	roles := client.HostRule(strings.Split(role, ","))
	if err := roles.Validation(); err != nil {
		println(err.Error())
		return
	}
	if !roles.HasRule("manage") || !roles.HasRule("gateway") {
		println("first node must have manage and gateway role")
		return
	}
	// download source code from github if in online model
	if c.String("install-type") == "online" {
		fmt.Println("Download the installation configuration file remotely...")
		csi := sources.CodeSourceInfo{
			RepositoryURL: c.String("rainbond-repo"),
			Branch:        c.String("rainbond-version"),
		}
		os.RemoveAll(c.String("work_dir"))
		os.MkdirAll(c.String("work_dir"), 0755)
		_, err := sources.GitClone(csi, c.String("work_dir"), event.GetTestLogger(), 5)
		if err != nil {
			println(err.Error())
			return
		}
	}
	if err := updateConfigFile(c.String("config-file"), getConfig(c)); err != nil {
		showError("update config file failure " + err.Error())
	}
	// start setup script to install rainbond
	fmt.Println("Initializes the installation of the first node...")
	cmd := exec.Command("bash", "-c", fmt.Sprintf("cd %s ; ./setup.sh", c.String("work_dir")))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	err = cmd.Run()
	if err != nil {
		println(err.Error())
		return
	}
	ioutil.WriteFile("/opt/rainbond/.rainbond.success", []byte(c.String("rainbond-version")), 0644)
	return
}
