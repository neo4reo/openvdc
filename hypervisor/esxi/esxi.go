// +build linux

package esxi

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"os/exec"
	"strconv"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/axsh/openvdc/hypervisor"
	"github.com/axsh/openvdc/hypervisor/util"
	"github.com/axsh/openvdc/model"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	cli "github.com/vmware/govmomi/govc/cli"
	_ "github.com/vmware/govmomi/govc/datastore"
	_ "github.com/vmware/govmomi/govc/device"
	_ "github.com/vmware/govmomi/govc/device/floppy"
	_ "github.com/vmware/govmomi/govc/device/serial"
	_ "github.com/vmware/govmomi/govc/vm"
	_ "github.com/vmware/govmomi/govc/vm/guest"
	"golang.org/x/crypto/ssh"
)

type BridgeType int

const (
	None BridgeType = iota // 0
	Linux
	OVS
)

var settings struct {
	ScriptPath      string
	EsxiUser        string
	EsxiPass        string
	EsxiIp          string
	EsxiDatacenter  string
	EsxiInsecure    bool
	EsxiHostSshkey  string
	EsxiVmUser      string
	EsxiVmPass      string
	EsxiVmDatastore string
	EsxiUrl         string
	BridgeName      string
	BridgeType      BridgeType
}

func (t BridgeType) String() string {
	switch t {
	case Linux:
		return "linux"
	case OVS:
		return "ovs"
	default:
		return "none"
	}
}

type EsxiMachine struct {
	SerialConsolePort int
	Nics              []Nic
}

type Nic struct {
	IfName       string
	Index        string
	Ipv4Addr     string
	MacAddr      string
	Bridge       string
	BridgeHelper string
	Type         string
}

type EsxiHypervisorProvider struct {
}

type EsxiHypervisorDriver struct {
	hypervisor.Base
	template  *model.EsxiTemplate
	machine   *EsxiMachine
	imageName string
	hostName  string
	vmName    string
}

func (p *EsxiHypervisorProvider) Name() string {
	return "esxi"
}

func init() {
	hypervisor.RegisterProvider("esxi", &EsxiHypervisorProvider{})
	viper.SetDefault("hypervisor.esxi-insecure", true)
}

func (p *EsxiHypervisorProvider) LoadConfig(sub *viper.Viper) error {
	if sub.IsSet("bridges.name") {
		settings.BridgeName = sub.GetString("bridges.name")
		if sub.IsSet("bridges.type") {
			switch sub.GetString("bridges.type") {
			case "linux":
				settings.BridgeType = Linux
			case "ovs":
				settings.BridgeType = OVS
			default:
				return errors.Errorf("Unknown bridges.type value: %s", sub.GetString("bridges.type"))
			}
		}
	} else if sub.IsSet("bridges.linux.name") {
		log.Warn("bridges.linux.name is obsolete option")
		settings.BridgeName = sub.GetString("bridges.linux.name")
		settings.BridgeType = Linux
	} else if sub.IsSet("bridges.ovs.name") {
		log.Warn("bridges.ovs.name is obsolete option")
		settings.BridgeName = sub.GetString("bridges.ovs.name")
		settings.BridgeType = OVS
	}

	if sub.GetString("hypervisor.esxi-user") == "" {
		return errors.Errorf("Missing configuration hypervisor.exsi-user")
	}
	settings.EsxiUser = sub.GetString("hypervisor.esxi-user")

	if sub.GetString("hypervisor.esxi-pass") == "" {
		return errors.Errorf("Missing configuration hypervisor.exsi-pass")
	}
	settings.EsxiPass = sub.GetString("hypervisor.esxi-pass")

	if sub.GetString("hypervisor.esxi-ip") == "" {
		return errors.Errorf("Missing configuration hypervisor.exsi-ip")
	}
	settings.EsxiIp = sub.GetString("hypervisor.esxi-ip")

	if sub.GetString("hypervisor.esxi-datacenter") == "" {
		return errors.Errorf("Missing configuration hypervisor.exsi-datacenter")
	}
	settings.EsxiDatacenter = sub.GetString("hypervisor.esxi-datacenter")

	if sub.GetString("hypervisor.esxi-host-sshkey") == "" {
		return errors.Errorf("Missing configuration hypervisor.esxi-host-sshkey")
	}
	settings.EsxiHostSshkey = sub.GetString("hypervisor.esxi-host-sshkey")

	if sub.GetString("hypervisor.esxi-vm-datastore") == "" {
		return errors.Errorf("Missing configuration hypervisor.exsi-vm-datastore")
	}
	settings.EsxiVmDatastore = sub.GetString("hypervisor.esxi-vm-datastore")

	settings.ScriptPath = sub.GetString("hypervisor.script-path")
	settings.EsxiInsecure = sub.GetBool("hypervisor.esxi-insecure")
	settings.EsxiVmUser = sub.GetString("hypervisor.esxi-vm-user")
	settings.EsxiVmPass = sub.GetString("hypervisor.esxi-vm-pass")

	esxiInfo := fmt.Sprintf("%s:%s@%s", settings.EsxiUser, settings.EsxiPass, settings.EsxiIp)

	u, err := url.Parse("https://" + esxiInfo + "/sdk")
	if err != nil {
		return errors.Wrap(err, "Failed to parse url for ESXi server")
	}
	settings.EsxiUrl = u.String()

	return nil
}

func (p *EsxiHypervisorProvider) CreateDriver(instance *model.Instance, template model.ResourceTemplate) (hypervisor.HypervisorDriver, error) {
	EsxiTmpl, ok := template.(*model.EsxiTemplate)
	if !ok {
		return nil, errors.Errorf("template type is not *model.WmwareTemplate: %T, template")
	}
	instanceIdx, _ := strconv.Atoi(strings.TrimPrefix(instance.GetId(), "i-"))
	driver := &EsxiHypervisorDriver{
		Base: hypervisor.Base{
			Log:      log.WithFields(log.Fields{"Hypervisor": "esxi", "instance_id": instance.GetId()}),
			Instance: instance,
		},
		template: EsxiTmpl,
		machine:  &EsxiMachine{SerialConsolePort: (15000 + instanceIdx)},
		vmName:   instance.GetId(),
	}

	return driver, nil
}

func (d *EsxiHypervisorDriver) log() *log.Entry {
	return d.Base.Log
}

func join(separator byte, args ...string) string {
	argLength := len(args)
	currentArg := 0
	var buf bytes.Buffer
	for _, arg := range args {
		currentArg = currentArg + 1
		buf.WriteString(arg)
		if currentArg == argLength {
			separator = 0
		}
		if separator > 0 {
			buf.WriteByte(separator)
		}
	}
	return buf.String()
}

func runCmd(cmd string, args []string) error {
	c := exec.Command(cmd, args...)
	if err := c.Run(); err != nil {
		return errors.Errorf("failed to execute command :%s %s", cmd, args)
	}
	return nil
}

func esxiRunCmd(cmdList ...[]string) error {
	for _, args := range cmdList {
		a := []string{
			args[0],
			join('=', "-dc", settings.EsxiDatacenter),
			join('=', "-k", strconv.FormatBool(settings.EsxiInsecure)),
			join('=', "-u", settings.EsxiUrl),
		}
		for _, arg := range args[1:] {
			a = append(a, arg)
		}
		if rc := cli.Run(a); rc != 0 {
			return errors.Errorf("Failed api request: %s", args[0])
		}
	}
	return nil
}

// wrappers for esxi api syntax
func storageImg(imgName string) string {
	path := join('/', imgName, imgName)
	return join('.', path, "vmx")
}

func vmUserDetails() string {
	userDetails := join(':', settings.EsxiVmUser, settings.EsxiVmPass)
	return join('=', "-l", userDetails)
}

func (d *EsxiHypervisorDriver) vmPath() string {
	return join(0, "-vm.path=[", settings.EsxiVmDatastore, "]", storageImg(d.vmName))
}

type Image struct {
	Path      string
	Format    string
	Size      int
	baseImage string
}

func (d *EsxiHypervisorDriver) MetadataDrivePath() string {
	return "/tmp/metadrive.img"
}

func (d *EsxiHypervisorDriver) MetadataDriveDatamap() map[string]interface{} {
	metadataMap := make(map[string]interface{})
	metadataMap["hostname"] = d.vmName
	for idx, nic := range d.machine.Nics {
		if nic.Type == "veth" {
			iface := make(map[string]interface{})
			iface["ifname"] = nic.IfName
			iface["ipv4"] = nic.Ipv4Addr
			iface["mac"] = nic.MacAddr
			metadataMap[fmt.Sprintf("nic-%02d", idx)] = iface
		}
	}
	return metadataMap
}

func (m *EsxiMachine) AddNICs(nics []Nic) {
	for _, nic := range nics {
		m.Nics = append(m.Nics, nic)
	}
}

func (d *EsxiHypervisorDriver) CreateInstance() error {

	var nics []Nic
	for idx, iface := range d.template.GetInterfaces() {
		nics = append(nics, Nic{
			IfName:   fmt.Sprintf("%s_%02d", d.vmName, idx),
			Type:     iface.Type,
			Ipv4Addr: iface.Ipv4Addr,
			MacAddr:  iface.Macaddr,
			Bridge:   settings.BridgeName,
		})
	}

	d.machine.AddNICs(nics)

	if err := util.CreateMetadataDisk(d); err != nil {
		return err
	}
	if err := util.MountMetadataDisk(d); err != nil {
		return err
	}
	if err := util.WriteMetadata(d); err != nil {
		return err
	}
	if err := util.UmountMetadataDisk(d); err != nil {
		return err
	}

	// Create new folder
	err := esxiRunCmd(
		[]string{"datastore.mkdir", join('=', "-ds", settings.EsxiVmDatastore), d.vmName},
	)
	if err != nil {
		return err
	}

	err = esxiRunCmd(
		[]string{"datastore.upload", join('=', "-ds", settings.EsxiVmDatastore), d.MetadataDrivePath(), fmt.Sprintf("%s/metadrive.img", d.vmName)},
	)
	if err != nil {
		return err
	}

	if err := os.Remove(d.MetadataDrivePath()); err != nil {
		return errors.Errorf("Unable to remove metadrive: %s", d.MetadataDrivePath())
	}

	key, err := ioutil.ReadFile(settings.EsxiHostSshkey)
	if err != nil {
		return errors.Errorf("Unable to read the specified ssh private key: %s", settings.EsxiHostSshkey)
	}
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		return errors.Errorf("Unable to parse the specified ssh private key: %s", settings.EsxiHostSshkey)
	}

	conn, err := ssh.Dial("tcp", strings.Join([]string{settings.EsxiIp, "22"}, ":"), &ssh.ClientConfig{
		User: settings.EsxiUser,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
	})

	if err != nil {
		return errors.Errorf("Unable establish ssh connection to %s", settings.EsxiIp)
	}
	defer conn.Close()

	session, err := conn.NewSession()
	if err != nil {
		return errors.Errorf("Unable to open a session on connection %d", conn)
	}
	defer session.Close()

	var out bytes.Buffer
	var stderr bytes.Buffer
	session.Stdout = &out
	session.Stderr = &stderr

	// Ssh into esxiHost and use "vmkfstools" to clone vmdk"
	basePath := join('/', "/vmfs", "volumes", settings.EsxiVmDatastore, "CentOS7", "CentOS7.vmdk")
	newPath := join('/', "/vmfs", "volumes", settings.EsxiVmDatastore, d.vmName, "CentOS7.vmdk")

	if err := session.Run(join(' ', "vmkfstools -i", basePath, newPath, "-d thin")); err != nil {
		return errors.Errorf(stderr.String(), "Error cloning vmdk")
	}

	// TODO: don't hardcode the base image
	// NOTE: serial port devices starts from 9000 on the current driver.

	datastore := join('=', "-ds", settings.EsxiVmDatastore)
	err = esxiRunCmd(
		[]string{"datastore.cp", datastore, storageImg("CentOS7"), storageImg(d.vmName)},
		[]string{"vm.register", datastore, storageImg(d.vmName)},
		[]string{"vm.change", join('=', "-name", d.vmName), d.vmPath()},
		[]string{"device.floppy.insert", fmt.Sprintf("-vm=%s", d.vmName), fmt.Sprintf("%s/metadrive.img", d.vmName)},
		[]string{"device.serial.add", d.vmPath()},
		[]string{"device.serial.connect", d.vmPath(), "-device=serialport-9000", join(':', "telnet://", strconv.Itoa(d.machine.SerialConsolePort))},
	)
	if err != nil {
		return err
	}

	return nil
}

func (d *EsxiHypervisorDriver) DestroyInstance() error {
	return esxiRunCmd(
		[]string{"datastore.rm", fmt.Sprintf("-ds=%s", settings.EsxiVmDatastore), fmt.Sprintf("%s/metadrive.img", d.vmName)},
		[]string{"vm.destroy", d.vmPath()},
	)
}

func (d *EsxiHypervisorDriver) StartInstance() error {
	port := strconv.Itoa(d.machine.SerialConsolePort)

	return esxiRunCmd(
		[]string{"device.serial.connect", d.vmPath(), "-device=serialport-9000", join(':', "telnet://", port)},
		[]string{"vm.power", "-on=true", "-suspend=false", d.vmPath()},
	)
}

func (d *EsxiHypervisorDriver) StopInstance() error {
	return esxiRunCmd(
		[]string{"vm.power", "-suspend=true", d.vmPath()},
		[]string{"device.serial.disconnect", d.vmPath(), fmt.Sprintf("-device=serialport-9000")},
	)
}

func (d EsxiHypervisorDriver) RebootInstance() error {
	// Linux, this should be doable through api call.
	return esxiRunCmd(
		[]string{"guest.start", vmUserDetails(), d.vmPath(), "/sbin/reboot"},
	)
}

func (d *EsxiHypervisorDriver) Recover(instanceState model.InstanceState) error {
	//Todo: handle recovery
	return nil
}
