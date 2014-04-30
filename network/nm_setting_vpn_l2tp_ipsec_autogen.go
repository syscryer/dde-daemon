// This file is automatically generated, please don't edit manully.
package main

import (
	"fmt"
)

// Get key type
func getSettingVpnL2tpIpsecKeyType(key string) (t ktype) {
	switch key {
	default:
		t = ktypeUnknown
	case NM_SETTING_VPN_L2TP_KEY_IPSEC_ENABLE:
		t = ktypeBoolean
	case NM_SETTING_VPN_L2TP_KEY_IPSEC_GROUP_NAME:
		t = ktypeString
	case NM_SETTING_VPN_L2TP_KEY_IPSEC_GATEWAY_ID:
		t = ktypeString
	case NM_SETTING_VPN_L2TP_KEY_IPSEC_PSK:
		t = ktypeString
	}
	return
}

// Check is key in current setting field
func isKeyInSettingVpnL2tpIpsec(key string) bool {
	switch key {
	case NM_SETTING_VPN_L2TP_KEY_IPSEC_ENABLE:
		return true
	case NM_SETTING_VPN_L2TP_KEY_IPSEC_GROUP_NAME:
		return true
	case NM_SETTING_VPN_L2TP_KEY_IPSEC_GATEWAY_ID:
		return true
	case NM_SETTING_VPN_L2TP_KEY_IPSEC_PSK:
		return true
	}
	return false
}

// Get key's default value
func getSettingVpnL2tpIpsecKeyDefaultValueJSON(key string) (valueJSON string) {
	switch key {
	default:
		logger.Error("invalid key:", key)
	case NM_SETTING_VPN_L2TP_KEY_IPSEC_ENABLE:
		valueJSON = `false`
	case NM_SETTING_VPN_L2TP_KEY_IPSEC_GROUP_NAME:
		valueJSON = `""`
	case NM_SETTING_VPN_L2TP_KEY_IPSEC_GATEWAY_ID:
		valueJSON = `""`
	case NM_SETTING_VPN_L2TP_KEY_IPSEC_PSK:
		valueJSON = `""`
	}
	return
}

// Get JSON value generally
func generalGetSettingVpnL2tpIpsecKeyJSON(data connectionData, key string) (value string) {
	switch key {
	default:
		logger.Error("generalGetSettingVpnL2tpIpsecKeyJSON: invalide key", key)
	case NM_SETTING_VPN_L2TP_KEY_IPSEC_ENABLE:
		value = getSettingVpnL2tpKeyIpsecEnableJSON(data)
	case NM_SETTING_VPN_L2TP_KEY_IPSEC_GROUP_NAME:
		value = getSettingVpnL2tpKeyIpsecGroupNameJSON(data)
	case NM_SETTING_VPN_L2TP_KEY_IPSEC_GATEWAY_ID:
		value = getSettingVpnL2tpKeyIpsecGatewayIdJSON(data)
	case NM_SETTING_VPN_L2TP_KEY_IPSEC_PSK:
		value = getSettingVpnL2tpKeyIpsecPskJSON(data)
	}
	return
}

// Set JSON value generally
func generalSetSettingVpnL2tpIpsecKeyJSON(data connectionData, key, valueJSON string) (err error) {
	switch key {
	default:
		logger.Error("generalSetSettingVpnL2tpIpsecKeyJSON: invalide key", key)
	case NM_SETTING_VPN_L2TP_KEY_IPSEC_ENABLE:
		err = logicSetSettingVpnL2tpKeyIpsecEnableJSON(data, valueJSON)
	case NM_SETTING_VPN_L2TP_KEY_IPSEC_GROUP_NAME:
		err = setSettingVpnL2tpKeyIpsecGroupNameJSON(data, valueJSON)
	case NM_SETTING_VPN_L2TP_KEY_IPSEC_GATEWAY_ID:
		err = setSettingVpnL2tpKeyIpsecGatewayIdJSON(data, valueJSON)
	case NM_SETTING_VPN_L2TP_KEY_IPSEC_PSK:
		err = setSettingVpnL2tpKeyIpsecPskJSON(data, valueJSON)
	}
	return
}

// Check if key exists
func isSettingVpnL2tpKeyIpsecEnableExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_VF_VPN_L2TP_IPSEC_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_IPSEC_ENABLE)
}
func isSettingVpnL2tpKeyIpsecGroupNameExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_VF_VPN_L2TP_IPSEC_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_IPSEC_GROUP_NAME)
}
func isSettingVpnL2tpKeyIpsecGatewayIdExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_VF_VPN_L2TP_IPSEC_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_IPSEC_GATEWAY_ID)
}
func isSettingVpnL2tpKeyIpsecPskExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_VF_VPN_L2TP_IPSEC_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_IPSEC_PSK)
}

// Ensure field and key exists and not empty
func ensureFieldSettingVpnL2tpIpsecExists(data connectionData, errs fieldErrors, relatedKey string) {
	if !isSettingFieldExists(data, NM_SETTING_VF_VPN_L2TP_IPSEC_SETTING_NAME) {
		rememberError(errs, relatedKey, NM_SETTING_VF_VPN_L2TP_IPSEC_SETTING_NAME, fmt.Sprintf(NM_KEY_ERROR_MISSING_SECTION, NM_SETTING_VF_VPN_L2TP_IPSEC_SETTING_NAME))
	}
	fieldData, _ := data[NM_SETTING_VF_VPN_L2TP_IPSEC_SETTING_NAME]
	if len(fieldData) == 0 {
		rememberError(errs, relatedKey, NM_SETTING_VF_VPN_L2TP_IPSEC_SETTING_NAME, fmt.Sprintf(NM_KEY_ERROR_EMPTY_SECTION, NM_SETTING_VF_VPN_L2TP_IPSEC_SETTING_NAME))
	}
}
func ensureSettingVpnL2tpKeyIpsecEnableNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingVpnL2tpKeyIpsecEnableExists(data) {
		rememberError(errs, NM_SETTING_VF_VPN_L2TP_IPSEC_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_IPSEC_ENABLE, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingVpnL2tpKeyIpsecGroupNameNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingVpnL2tpKeyIpsecGroupNameExists(data) {
		rememberError(errs, NM_SETTING_VF_VPN_L2TP_IPSEC_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_IPSEC_GROUP_NAME, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingVpnL2tpKeyIpsecGroupName(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_VF_VPN_L2TP_IPSEC_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_IPSEC_GROUP_NAME, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingVpnL2tpKeyIpsecGatewayIdNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingVpnL2tpKeyIpsecGatewayIdExists(data) {
		rememberError(errs, NM_SETTING_VF_VPN_L2TP_IPSEC_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_IPSEC_GATEWAY_ID, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingVpnL2tpKeyIpsecGatewayId(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_VF_VPN_L2TP_IPSEC_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_IPSEC_GATEWAY_ID, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingVpnL2tpKeyIpsecPskNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingVpnL2tpKeyIpsecPskExists(data) {
		rememberError(errs, NM_SETTING_VF_VPN_L2TP_IPSEC_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_IPSEC_PSK, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingVpnL2tpKeyIpsecPsk(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_VF_VPN_L2TP_IPSEC_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_IPSEC_PSK, NM_KEY_ERROR_EMPTY_VALUE)
	}
}

// Getter
func getSettingVpnL2tpKeyIpsecEnable(data connectionData) (value bool) {
	value, _ = getSettingKey(data, NM_SETTING_VF_VPN_L2TP_IPSEC_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_IPSEC_ENABLE).(bool)
	return
}
func getSettingVpnL2tpKeyIpsecGroupName(data connectionData) (value string) {
	value, _ = getSettingKey(data, NM_SETTING_VF_VPN_L2TP_IPSEC_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_IPSEC_GROUP_NAME).(string)
	return
}
func getSettingVpnL2tpKeyIpsecGatewayId(data connectionData) (value string) {
	value, _ = getSettingKey(data, NM_SETTING_VF_VPN_L2TP_IPSEC_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_IPSEC_GATEWAY_ID).(string)
	return
}
func getSettingVpnL2tpKeyIpsecPsk(data connectionData) (value string) {
	value, _ = getSettingKey(data, NM_SETTING_VF_VPN_L2TP_IPSEC_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_IPSEC_PSK).(string)
	return
}

// Setter
func setSettingVpnL2tpKeyIpsecEnable(data connectionData, value bool) {
	setSettingKey(data, NM_SETTING_VF_VPN_L2TP_IPSEC_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_IPSEC_ENABLE, value)
}
func setSettingVpnL2tpKeyIpsecGroupName(data connectionData, value string) {
	setSettingKey(data, NM_SETTING_VF_VPN_L2TP_IPSEC_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_IPSEC_GROUP_NAME, value)
}
func setSettingVpnL2tpKeyIpsecGatewayId(data connectionData, value string) {
	setSettingKey(data, NM_SETTING_VF_VPN_L2TP_IPSEC_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_IPSEC_GATEWAY_ID, value)
}
func setSettingVpnL2tpKeyIpsecPsk(data connectionData, value string) {
	setSettingKey(data, NM_SETTING_VF_VPN_L2TP_IPSEC_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_IPSEC_PSK, value)
}

// JSON Getter
func getSettingVpnL2tpKeyIpsecEnableJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_VF_VPN_L2TP_IPSEC_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_IPSEC_ENABLE, getSettingVpnL2tpIpsecKeyType(NM_SETTING_VPN_L2TP_KEY_IPSEC_ENABLE))
	return
}
func getSettingVpnL2tpKeyIpsecGroupNameJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_VF_VPN_L2TP_IPSEC_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_IPSEC_GROUP_NAME, getSettingVpnL2tpIpsecKeyType(NM_SETTING_VPN_L2TP_KEY_IPSEC_GROUP_NAME))
	return
}
func getSettingVpnL2tpKeyIpsecGatewayIdJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_VF_VPN_L2TP_IPSEC_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_IPSEC_GATEWAY_ID, getSettingVpnL2tpIpsecKeyType(NM_SETTING_VPN_L2TP_KEY_IPSEC_GATEWAY_ID))
	return
}
func getSettingVpnL2tpKeyIpsecPskJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_VF_VPN_L2TP_IPSEC_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_IPSEC_PSK, getSettingVpnL2tpIpsecKeyType(NM_SETTING_VPN_L2TP_KEY_IPSEC_PSK))
	return
}

// JSON Setter
func setSettingVpnL2tpKeyIpsecEnableJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_VF_VPN_L2TP_IPSEC_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_IPSEC_ENABLE, valueJSON, getSettingVpnL2tpIpsecKeyType(NM_SETTING_VPN_L2TP_KEY_IPSEC_ENABLE))
}
func setSettingVpnL2tpKeyIpsecGroupNameJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_VF_VPN_L2TP_IPSEC_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_IPSEC_GROUP_NAME, valueJSON, getSettingVpnL2tpIpsecKeyType(NM_SETTING_VPN_L2TP_KEY_IPSEC_GROUP_NAME))
}
func setSettingVpnL2tpKeyIpsecGatewayIdJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_VF_VPN_L2TP_IPSEC_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_IPSEC_GATEWAY_ID, valueJSON, getSettingVpnL2tpIpsecKeyType(NM_SETTING_VPN_L2TP_KEY_IPSEC_GATEWAY_ID))
}
func setSettingVpnL2tpKeyIpsecPskJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_VF_VPN_L2TP_IPSEC_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_IPSEC_PSK, valueJSON, getSettingVpnL2tpIpsecKeyType(NM_SETTING_VPN_L2TP_KEY_IPSEC_PSK))
}

// Logic JSON Setter
func logicSetSettingVpnL2tpKeyIpsecEnableJSON(data connectionData, valueJSON string) (err error) {
	err = setSettingVpnL2tpKeyIpsecEnableJSON(data, valueJSON)
	if err != nil {
		return
	}
	if isSettingVpnL2tpKeyIpsecEnableExists(data) {
		value := getSettingVpnL2tpKeyIpsecEnable(data)
		err = logicSetSettingVpnL2tpKeyIpsecEnable(data, value)
	}
	return
}

// Remover
func removeSettingVpnL2tpKeyIpsecEnable(data connectionData) {
	removeSettingKey(data, NM_SETTING_VF_VPN_L2TP_IPSEC_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_IPSEC_ENABLE)
}
func removeSettingVpnL2tpKeyIpsecGroupName(data connectionData) {
	removeSettingKey(data, NM_SETTING_VF_VPN_L2TP_IPSEC_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_IPSEC_GROUP_NAME)
}
func removeSettingVpnL2tpKeyIpsecGatewayId(data connectionData) {
	removeSettingKey(data, NM_SETTING_VF_VPN_L2TP_IPSEC_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_IPSEC_GATEWAY_ID)
}
func removeSettingVpnL2tpKeyIpsecPsk(data connectionData) {
	removeSettingKey(data, NM_SETTING_VF_VPN_L2TP_IPSEC_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_IPSEC_PSK)
}