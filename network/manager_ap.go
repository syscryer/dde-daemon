package main

import (
	nm "dbus/org/freedesktop/networkmanager"
	"dlib/dbus"
)

type apSecType uint32

const (
	apSecNone apSecType = iota
	apSecWep
	apSecPsk
	apSecEap
)

type accessPoint struct {
	Ssid         string
	Secured      bool
	SecuredInEap bool
	Strength     uint8
	Path         dbus.ObjectPath
}

func NewAccessPoint(apPath dbus.ObjectPath) (ap accessPoint, err error) {
	calcStrength := func(s uint8) uint8 {
		switch {
		case s <= 10:
			return 0
		case s <= 25:
			return 25
		case s <= 50:
			return 50
		case s <= 75:
			return 75
		case s <= 100:
			return 100
		}
		return 0
	}

	nmAp, err := nmNewAccessPoint(apPath)
	if err != nil {
		return
	}

	ap = accessPoint{
		Ssid:         string(nmAp.Ssid.Get()),
		Secured:      getApSecType(nmAp) != apSecNone,
		SecuredInEap: getApSecType(nmAp) == apSecEap,
		Strength:     calcStrength(nmAp.Strength.Get()),
		Path:         nmAp.Path,
	}
	return
}

func getApSecType(ap *nm.AccessPoint) apSecType {
	return doParseApSecType(ap.Flags.Get(), ap.WpaFlags.Get(), ap.RsnFlags.Get())
}

func doParseApSecType(flags, wpaFlags, rsnFlags uint32) apSecType {
	r := apSecNone

	if (flags&NM_802_11_AP_FLAGS_PRIVACY != 0) && (wpaFlags == NM_802_11_AP_SEC_NONE) && (rsnFlags == NM_802_11_AP_SEC_NONE) {
		r = apSecWep
	}
	if wpaFlags != NM_802_11_AP_SEC_NONE {
		r = apSecPsk
	}
	if rsnFlags != NM_802_11_AP_SEC_NONE {
		r = apSecPsk
	}
	if (wpaFlags&NM_802_11_AP_SEC_KEY_MGMT_802_1X != 0) || (rsnFlags&NM_802_11_AP_SEC_KEY_MGMT_802_1X != 0) {
		r = apSecEap
	}
	return r
}

// GetAccessPoints return all access point's dbus path of target device.
func (m *Manager) GetAccessPoints(path dbus.ObjectPath) (aps []dbus.ObjectPath, err error) {
	dev, err := nmNewDeviceWireless(path)
	if err != nil {
		return
	}
	aps, err = dev.GetAccessPoints()
	return
}

// GetAccessPointProperty return access point's detail information.
// TODO
// func (m *Manager) GetAccessPointProperty(apPath dbus.ObjectPath) (apJSON string, err error) {
func (m *Manager) GetAccessPointProperty(apPath dbus.ObjectPath) (ap accessPoint, err error) {
	ap, err = NewAccessPoint(apPath)
	// if err != nil {
	// return
	// }
	// apJSON, err = marshalJSON(ap)
	return
}

func (m *Manager) ActivateConnectionForAccessPoint(apPath, devPath dbus.ObjectPath) (uuid string, err error) {
	logger.Debugf("ActivateConnectionForAccessPoint: apPath=%s, devPath=%s", apPath, devPath)
	// if there is no connection for current access point, create one
	ap, err := nmNewAccessPoint(apPath)
	if err != nil {
		return
	}
	cpath, ok := nmGetWirelessConnectionBySsid(ap.Ssid.Get())
	if ok {
		logger.Debug("activate connection") // TODO test
		uuid = nmGetConnectionUuid(cpath)
		_, err = nmActivateConnection(cpath, devPath)
	} else {
		logger.Debug("add and activate connection") // TODO test
		uuid = newUUID()
		data := newWirelessConnectionData(string(ap.Ssid.Get()), uuid, []byte(ap.Ssid.Get()), getApSecType(ap))
		_, _, err = nmAddAndActivateConnection(data, devPath)
	}
	return
}

// CreateConnectionByAccessPoint create connection for access point and return the uuid.
func (m *Manager) CreateConnectionForAccessPoint(apPath dbus.ObjectPath) (uuid string, err error) {
	logger.Debug("CreateConnectionForAccessPoint: apPath", apPath)
	uuid, err = m.GetConnectionUuidByAccessPoint(apPath)
	if len(uuid) != 0 {
		// connection already exists
		return
	}

	// create connection
	ap, err := nmNewAccessPoint(apPath)
	if err != nil {
		return
	}
	// TODO FIXME
	secType := getApSecType(ap)
	if secType == apSecEap {
		logger.Debug("ignore wireless connection:", string(ap.Ssid.Get()))
		return "", dbus.NewNoObjectError(apPath)
	}

	uuid = newWirelessConnection(string(ap.Ssid.Get()), []byte(ap.Ssid.Get()), getApSecType(ap))
	return
}

// TODO
func (m *Manager) EditConnectionForAccessPoint(apPath dbus.ObjectPath, devPath dbus.ObjectPath) (session *ConnectionSession, err error) {
	// // if is read only connection(default system connection created by
	// // network manager), create a new connection
	// // TODO
	// cpath, err := nmGetConnectionByUuid(uuid)
	// if err != nil {
	// 	return
	// }
	// connData, err := nmGetConnectionData(cpath)
	// if err != nil {
	// 	return
	// }
	// if getSettingConnectionReadOnly(connData) {
	// 	logger.Debug("read only connection, create new")
	// 	return m.CreateConnection(generalGetConnectionType(connData), devPath)
	// }

	// session, err = NewConnectionSessionByOpen(uuid, devPath)
	// if err != nil {
	// 	logger.Error(err)
	// 	return
	// }

	// // install dbus session
	// err = dbus.InstallOnSession(session)
	// if err != nil {
	// 	logger.Error(err)
	// 	return
	// }

	return
}

// GetConnectionUuidByAccessPoint return the connection's uuid of access point, return empty if none.
func (m *Manager) GetConnectionUuidByAccessPoint(apPath dbus.ObjectPath) (uuid string, err error) {
	ap, err := nmNewAccessPoint(apPath)
	if err != nil {
		return
	}

	cpath, ok := nmGetWirelessConnectionBySsid(ap.Ssid.Get())
	if !ok {
		return
	}

	uuid = nmGetConnectionUuid(cpath)

	logger.Debugf("GetConnectionUuidByAccessPoint: apPath=%s, uuid=%s", apPath, uuid) // TODO test
	return
}