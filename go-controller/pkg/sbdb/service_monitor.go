// Code generated by "libovsdb.modelgen"
// DO NOT EDIT.

package sbdb

import "github.com/ovn-org/libovsdb/model"

const ServiceMonitorTable = "Service_Monitor"

type (
	ServiceMonitorProtocol = string
	ServiceMonitorStatus   = string
)

var (
	ServiceMonitorProtocolTCP   ServiceMonitorProtocol = "tcp"
	ServiceMonitorProtocolUDP   ServiceMonitorProtocol = "udp"
	ServiceMonitorStatusOnline  ServiceMonitorStatus   = "online"
	ServiceMonitorStatusOffline ServiceMonitorStatus   = "offline"
	ServiceMonitorStatusError   ServiceMonitorStatus   = "error"
)

// ServiceMonitor defines an object in Service_Monitor table
type ServiceMonitor struct {
	UUID        string                  `ovsdb:"_uuid"`
	ExternalIDs map[string]string       `ovsdb:"external_ids"`
	IP          string                  `ovsdb:"ip"`
	LogicalPort string                  `ovsdb:"logical_port"`
	Options     map[string]string       `ovsdb:"options"`
	Port        int                     `ovsdb:"port"`
	Protocol    *ServiceMonitorProtocol `ovsdb:"protocol"`
	SrcIP       string                  `ovsdb:"src_ip"`
	SrcMAC      string                  `ovsdb:"src_mac"`
	Status      *ServiceMonitorStatus   `ovsdb:"status"`
}

func (a *ServiceMonitor) GetUUID() string {
	return a.UUID
}

func (a *ServiceMonitor) GetExternalIDs() map[string]string {
	return a.ExternalIDs
}

func copyServiceMonitorExternalIDs(a map[string]string) map[string]string {
	if a == nil {
		return nil
	}
	b := make(map[string]string, len(a))
	for k, v := range a {
		b[k] = v
	}
	return b
}

func equalServiceMonitorExternalIDs(a, b map[string]string) bool {
	if (a == nil) != (b == nil) {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if w, ok := b[k]; !ok || v != w {
			return false
		}
	}
	return true
}

func (a *ServiceMonitor) GetIP() string {
	return a.IP
}

func (a *ServiceMonitor) GetLogicalPort() string {
	return a.LogicalPort
}

func (a *ServiceMonitor) GetOptions() map[string]string {
	return a.Options
}

func copyServiceMonitorOptions(a map[string]string) map[string]string {
	if a == nil {
		return nil
	}
	b := make(map[string]string, len(a))
	for k, v := range a {
		b[k] = v
	}
	return b
}

func equalServiceMonitorOptions(a, b map[string]string) bool {
	if (a == nil) != (b == nil) {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if w, ok := b[k]; !ok || v != w {
			return false
		}
	}
	return true
}

func (a *ServiceMonitor) GetPort() int {
	return a.Port
}

func (a *ServiceMonitor) GetProtocol() *ServiceMonitorProtocol {
	return a.Protocol
}

func copyServiceMonitorProtocol(a *ServiceMonitorProtocol) *ServiceMonitorProtocol {
	if a == nil {
		return nil
	}
	b := *a
	return &b
}

func equalServiceMonitorProtocol(a, b *ServiceMonitorProtocol) bool {
	if (a == nil) != (b == nil) {
		return false
	}
	if a == b {
		return true
	}
	return *a == *b
}

func (a *ServiceMonitor) GetSrcIP() string {
	return a.SrcIP
}

func (a *ServiceMonitor) GetSrcMAC() string {
	return a.SrcMAC
}

func (a *ServiceMonitor) GetStatus() *ServiceMonitorStatus {
	return a.Status
}

func copyServiceMonitorStatus(a *ServiceMonitorStatus) *ServiceMonitorStatus {
	if a == nil {
		return nil
	}
	b := *a
	return &b
}

func equalServiceMonitorStatus(a, b *ServiceMonitorStatus) bool {
	if (a == nil) != (b == nil) {
		return false
	}
	if a == b {
		return true
	}
	return *a == *b
}

func (a *ServiceMonitor) DeepCopyInto(b *ServiceMonitor) {
	*b = *a
	b.ExternalIDs = copyServiceMonitorExternalIDs(a.ExternalIDs)
	b.Options = copyServiceMonitorOptions(a.Options)
	b.Protocol = copyServiceMonitorProtocol(a.Protocol)
	b.Status = copyServiceMonitorStatus(a.Status)
}

func (a *ServiceMonitor) DeepCopy() *ServiceMonitor {
	b := new(ServiceMonitor)
	a.DeepCopyInto(b)
	return b
}

func (a *ServiceMonitor) CloneModelInto(b model.Model) {
	c := b.(*ServiceMonitor)
	a.DeepCopyInto(c)
}

func (a *ServiceMonitor) CloneModel() model.Model {
	return a.DeepCopy()
}

func (a *ServiceMonitor) Equals(b *ServiceMonitor) bool {
	return a.UUID == b.UUID &&
		equalServiceMonitorExternalIDs(a.ExternalIDs, b.ExternalIDs) &&
		a.IP == b.IP &&
		a.LogicalPort == b.LogicalPort &&
		equalServiceMonitorOptions(a.Options, b.Options) &&
		a.Port == b.Port &&
		equalServiceMonitorProtocol(a.Protocol, b.Protocol) &&
		a.SrcIP == b.SrcIP &&
		a.SrcMAC == b.SrcMAC &&
		equalServiceMonitorStatus(a.Status, b.Status)
}

func (a *ServiceMonitor) EqualsModel(b model.Model) bool {
	c := b.(*ServiceMonitor)
	return a.Equals(c)
}

var _ model.CloneableModel = &ServiceMonitor{}
var _ model.ComparableModel = &ServiceMonitor{}
