// Code generated by "libovsdb.modelgen"
// DO NOT EDIT.

package sbdb

import "github.com/ovn-org/libovsdb/model"

const HAChassisTable = "HA_Chassis"

// HAChassis defines an object in HA_Chassis table
type HAChassis struct {
	UUID        string            `ovsdb:"_uuid"`
	Chassis     *string           `ovsdb:"chassis"`
	ExternalIDs map[string]string `ovsdb:"external_ids"`
	Priority    int               `ovsdb:"priority"`
}

func (a *HAChassis) GetUUID() string {
	return a.UUID
}

func (a *HAChassis) GetChassis() *string {
	return a.Chassis
}

func copyHAChassisChassis(a *string) *string {
	if a == nil {
		return nil
	}
	b := *a
	return &b
}

func equalHAChassisChassis(a, b *string) bool {
	if (a == nil) != (b == nil) {
		return false
	}
	if a == b {
		return true
	}
	return *a == *b
}

func (a *HAChassis) GetExternalIDs() map[string]string {
	return a.ExternalIDs
}

func copyHAChassisExternalIDs(a map[string]string) map[string]string {
	if a == nil {
		return nil
	}
	b := make(map[string]string, len(a))
	for k, v := range a {
		b[k] = v
	}
	return b
}

func equalHAChassisExternalIDs(a, b map[string]string) bool {
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

func (a *HAChassis) GetPriority() int {
	return a.Priority
}

func (a *HAChassis) DeepCopyInto(b *HAChassis) {
	*b = *a
	b.Chassis = copyHAChassisChassis(a.Chassis)
	b.ExternalIDs = copyHAChassisExternalIDs(a.ExternalIDs)
}

func (a *HAChassis) DeepCopy() *HAChassis {
	b := new(HAChassis)
	a.DeepCopyInto(b)
	return b
}

func (a *HAChassis) CloneModelInto(b model.Model) {
	c := b.(*HAChassis)
	a.DeepCopyInto(c)
}

func (a *HAChassis) CloneModel() model.Model {
	return a.DeepCopy()
}

func (a *HAChassis) Equals(b *HAChassis) bool {
	return a.UUID == b.UUID &&
		equalHAChassisChassis(a.Chassis, b.Chassis) &&
		equalHAChassisExternalIDs(a.ExternalIDs, b.ExternalIDs) &&
		a.Priority == b.Priority
}

func (a *HAChassis) EqualsModel(b model.Model) bool {
	c := b.(*HAChassis)
	return a.Equals(c)
}

var _ model.CloneableModel = &HAChassis{}
var _ model.ComparableModel = &HAChassis{}
