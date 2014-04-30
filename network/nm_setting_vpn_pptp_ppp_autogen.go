// This file is automatically generated, please don't edit manully.
package main

import (
	"fmt"
)

// Get key type
func getSettingVpnPptpPppKeyType(key string) (t ktype) {
	switch key {
	default:
		t = ktypeUnknown
	case NM_SETTING_VPN_PPTP_KEY_REFUSE_EAP:
		t = ktypeBoolean
	case NM_SETTING_VPN_PPTP_KEY_REFUSE_PAP:
		t = ktypeBoolean
	case NM_SETTING_VPN_PPTP_KEY_REFUSE_CHAP:
		t = ktypeBoolean
	case NM_SETTING_VPN_PPTP_KEY_REFUSE_MSCHAP:
		t = ktypeBoolean
	case NM_SETTING_VPN_PPTP_KEY_REFUSE_MSCHAPV2:
		t = ktypeBoolean
	case NM_SETTING_VPN_PPTP_KEY_REQUIRE_MPPE:
		t = ktypeBoolean
	case NM_SETTING_VPN_PPTP_KEY_REQUIRE_MPPE_40:
		t = ktypeBoolean
	case NM_SETTING_VPN_PPTP_KEY_REQUIRE_MPPE_128:
		t = ktypeBoolean
	case NM_SETTING_VPN_PPTP_KEY_MPPE_STATEFUL:
		t = ktypeBoolean
	case NM_SETTING_VPN_PPTP_KEY_NOBSDCOMP:
		t = ktypeBoolean
	case NM_SETTING_VPN_PPTP_KEY_NODEFLATE:
		t = ktypeBoolean
	case NM_SETTING_VPN_PPTP_KEY_NO_VJ_COMP:
		t = ktypeBoolean
	case NM_SETTING_VPN_PPTP_KEY_LCP_ECHO_FAILURE:
		t = ktypeUint32
	case NM_SETTING_VPN_PPTP_KEY_LCP_ECHO_INTERVAL:
		t = ktypeUint32
	}
	return
}

// Check is key in current setting field
func isKeyInSettingVpnPptpPpp(key string) bool {
	switch key {
	case NM_SETTING_VPN_PPTP_KEY_REFUSE_EAP:
		return true
	case NM_SETTING_VPN_PPTP_KEY_REFUSE_PAP:
		return true
	case NM_SETTING_VPN_PPTP_KEY_REFUSE_CHAP:
		return true
	case NM_SETTING_VPN_PPTP_KEY_REFUSE_MSCHAP:
		return true
	case NM_SETTING_VPN_PPTP_KEY_REFUSE_MSCHAPV2:
		return true
	case NM_SETTING_VPN_PPTP_KEY_REQUIRE_MPPE:
		return true
	case NM_SETTING_VPN_PPTP_KEY_REQUIRE_MPPE_40:
		return true
	case NM_SETTING_VPN_PPTP_KEY_REQUIRE_MPPE_128:
		return true
	case NM_SETTING_VPN_PPTP_KEY_MPPE_STATEFUL:
		return true
	case NM_SETTING_VPN_PPTP_KEY_NOBSDCOMP:
		return true
	case NM_SETTING_VPN_PPTP_KEY_NODEFLATE:
		return true
	case NM_SETTING_VPN_PPTP_KEY_NO_VJ_COMP:
		return true
	case NM_SETTING_VPN_PPTP_KEY_LCP_ECHO_FAILURE:
		return true
	case NM_SETTING_VPN_PPTP_KEY_LCP_ECHO_INTERVAL:
		return true
	}
	return false
}

// Get key's default value
func getSettingVpnPptpPppKeyDefaultValueJSON(key string) (valueJSON string) {
	switch key {
	default:
		logger.Error("invalid key:", key)
	case NM_SETTING_VPN_PPTP_KEY_REFUSE_EAP:
		valueJSON = `false`
	case NM_SETTING_VPN_PPTP_KEY_REFUSE_PAP:
		valueJSON = `false`
	case NM_SETTING_VPN_PPTP_KEY_REFUSE_CHAP:
		valueJSON = `false`
	case NM_SETTING_VPN_PPTP_KEY_REFUSE_MSCHAP:
		valueJSON = `false`
	case NM_SETTING_VPN_PPTP_KEY_REFUSE_MSCHAPV2:
		valueJSON = `false`
	case NM_SETTING_VPN_PPTP_KEY_REQUIRE_MPPE:
		valueJSON = `false`
	case NM_SETTING_VPN_PPTP_KEY_REQUIRE_MPPE_40:
		valueJSON = `false`
	case NM_SETTING_VPN_PPTP_KEY_REQUIRE_MPPE_128:
		valueJSON = `false`
	case NM_SETTING_VPN_PPTP_KEY_MPPE_STATEFUL:
		valueJSON = `false`
	case NM_SETTING_VPN_PPTP_KEY_NOBSDCOMP:
		valueJSON = `false`
	case NM_SETTING_VPN_PPTP_KEY_NODEFLATE:
		valueJSON = `false`
	case NM_SETTING_VPN_PPTP_KEY_NO_VJ_COMP:
		valueJSON = `false`
	case NM_SETTING_VPN_PPTP_KEY_LCP_ECHO_FAILURE:
		valueJSON = `0`
	case NM_SETTING_VPN_PPTP_KEY_LCP_ECHO_INTERVAL:
		valueJSON = `0`
	}
	return
}

// Get JSON value generally
func generalGetSettingVpnPptpPppKeyJSON(data connectionData, key string) (value string) {
	switch key {
	default:
		logger.Error("generalGetSettingVpnPptpPppKeyJSON: invalide key", key)
	case NM_SETTING_VPN_PPTP_KEY_REFUSE_EAP:
		value = getSettingVpnPptpKeyRefuseEapJSON(data)
	case NM_SETTING_VPN_PPTP_KEY_REFUSE_PAP:
		value = getSettingVpnPptpKeyRefusePapJSON(data)
	case NM_SETTING_VPN_PPTP_KEY_REFUSE_CHAP:
		value = getSettingVpnPptpKeyRefuseChapJSON(data)
	case NM_SETTING_VPN_PPTP_KEY_REFUSE_MSCHAP:
		value = getSettingVpnPptpKeyRefuseMschapJSON(data)
	case NM_SETTING_VPN_PPTP_KEY_REFUSE_MSCHAPV2:
		value = getSettingVpnPptpKeyRefuseMschapv2JSON(data)
	case NM_SETTING_VPN_PPTP_KEY_REQUIRE_MPPE:
		value = getSettingVpnPptpKeyRequireMppeJSON(data)
	case NM_SETTING_VPN_PPTP_KEY_REQUIRE_MPPE_40:
		value = getSettingVpnPptpKeyRequireMppe40JSON(data)
	case NM_SETTING_VPN_PPTP_KEY_REQUIRE_MPPE_128:
		value = getSettingVpnPptpKeyRequireMppe128JSON(data)
	case NM_SETTING_VPN_PPTP_KEY_MPPE_STATEFUL:
		value = getSettingVpnPptpKeyMppeStatefulJSON(data)
	case NM_SETTING_VPN_PPTP_KEY_NOBSDCOMP:
		value = getSettingVpnPptpKeyNobsdcompJSON(data)
	case NM_SETTING_VPN_PPTP_KEY_NODEFLATE:
		value = getSettingVpnPptpKeyNodeflateJSON(data)
	case NM_SETTING_VPN_PPTP_KEY_NO_VJ_COMP:
		value = getSettingVpnPptpKeyNoVjCompJSON(data)
	case NM_SETTING_VPN_PPTP_KEY_LCP_ECHO_FAILURE:
		value = getSettingVpnPptpKeyLcpEchoFailureJSON(data)
	case NM_SETTING_VPN_PPTP_KEY_LCP_ECHO_INTERVAL:
		value = getSettingVpnPptpKeyLcpEchoIntervalJSON(data)
	}
	return
}

// Set JSON value generally
func generalSetSettingVpnPptpPppKeyJSON(data connectionData, key, valueJSON string) (err error) {
	switch key {
	default:
		logger.Error("generalSetSettingVpnPptpPppKeyJSON: invalide key", key)
	case NM_SETTING_VPN_PPTP_KEY_REFUSE_EAP:
		err = setSettingVpnPptpKeyRefuseEapJSON(data, valueJSON)
	case NM_SETTING_VPN_PPTP_KEY_REFUSE_PAP:
		err = setSettingVpnPptpKeyRefusePapJSON(data, valueJSON)
	case NM_SETTING_VPN_PPTP_KEY_REFUSE_CHAP:
		err = setSettingVpnPptpKeyRefuseChapJSON(data, valueJSON)
	case NM_SETTING_VPN_PPTP_KEY_REFUSE_MSCHAP:
		err = setSettingVpnPptpKeyRefuseMschapJSON(data, valueJSON)
	case NM_SETTING_VPN_PPTP_KEY_REFUSE_MSCHAPV2:
		err = setSettingVpnPptpKeyRefuseMschapv2JSON(data, valueJSON)
	case NM_SETTING_VPN_PPTP_KEY_REQUIRE_MPPE:
		err = logicSetSettingVpnPptpKeyRequireMppeJSON(data, valueJSON)
	case NM_SETTING_VPN_PPTP_KEY_REQUIRE_MPPE_40:
		err = setSettingVpnPptpKeyRequireMppe40JSON(data, valueJSON)
	case NM_SETTING_VPN_PPTP_KEY_REQUIRE_MPPE_128:
		err = setSettingVpnPptpKeyRequireMppe128JSON(data, valueJSON)
	case NM_SETTING_VPN_PPTP_KEY_MPPE_STATEFUL:
		err = setSettingVpnPptpKeyMppeStatefulJSON(data, valueJSON)
	case NM_SETTING_VPN_PPTP_KEY_NOBSDCOMP:
		err = setSettingVpnPptpKeyNobsdcompJSON(data, valueJSON)
	case NM_SETTING_VPN_PPTP_KEY_NODEFLATE:
		err = setSettingVpnPptpKeyNodeflateJSON(data, valueJSON)
	case NM_SETTING_VPN_PPTP_KEY_NO_VJ_COMP:
		err = setSettingVpnPptpKeyNoVjCompJSON(data, valueJSON)
	case NM_SETTING_VPN_PPTP_KEY_LCP_ECHO_FAILURE:
		err = setSettingVpnPptpKeyLcpEchoFailureJSON(data, valueJSON)
	case NM_SETTING_VPN_PPTP_KEY_LCP_ECHO_INTERVAL:
		err = setSettingVpnPptpKeyLcpEchoIntervalJSON(data, valueJSON)
	}
	return
}

// Check if key exists
func isSettingVpnPptpKeyRefuseEapExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_REFUSE_EAP)
}
func isSettingVpnPptpKeyRefusePapExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_REFUSE_PAP)
}
func isSettingVpnPptpKeyRefuseChapExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_REFUSE_CHAP)
}
func isSettingVpnPptpKeyRefuseMschapExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_REFUSE_MSCHAP)
}
func isSettingVpnPptpKeyRefuseMschapv2Exists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_REFUSE_MSCHAPV2)
}
func isSettingVpnPptpKeyRequireMppeExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_REQUIRE_MPPE)
}
func isSettingVpnPptpKeyRequireMppe40Exists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_REQUIRE_MPPE_40)
}
func isSettingVpnPptpKeyRequireMppe128Exists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_REQUIRE_MPPE_128)
}
func isSettingVpnPptpKeyMppeStatefulExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_MPPE_STATEFUL)
}
func isSettingVpnPptpKeyNobsdcompExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_NOBSDCOMP)
}
func isSettingVpnPptpKeyNodeflateExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_NODEFLATE)
}
func isSettingVpnPptpKeyNoVjCompExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_NO_VJ_COMP)
}
func isSettingVpnPptpKeyLcpEchoFailureExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_LCP_ECHO_FAILURE)
}
func isSettingVpnPptpKeyLcpEchoIntervalExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_LCP_ECHO_INTERVAL)
}

// Ensure field and key exists and not empty
func ensureFieldSettingVpnPptpPppExists(data connectionData, errs fieldErrors, relatedKey string) {
	if !isSettingFieldExists(data, NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME) {
		rememberError(errs, relatedKey, NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME, fmt.Sprintf(NM_KEY_ERROR_MISSING_SECTION, NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME))
	}
	fieldData, _ := data[NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME]
	if len(fieldData) == 0 {
		rememberError(errs, relatedKey, NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME, fmt.Sprintf(NM_KEY_ERROR_EMPTY_SECTION, NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME))
	}
}
func ensureSettingVpnPptpKeyRefuseEapNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingVpnPptpKeyRefuseEapExists(data) {
		rememberError(errs, NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_REFUSE_EAP, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingVpnPptpKeyRefusePapNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingVpnPptpKeyRefusePapExists(data) {
		rememberError(errs, NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_REFUSE_PAP, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingVpnPptpKeyRefuseChapNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingVpnPptpKeyRefuseChapExists(data) {
		rememberError(errs, NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_REFUSE_CHAP, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingVpnPptpKeyRefuseMschapNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingVpnPptpKeyRefuseMschapExists(data) {
		rememberError(errs, NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_REFUSE_MSCHAP, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingVpnPptpKeyRefuseMschapv2NoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingVpnPptpKeyRefuseMschapv2Exists(data) {
		rememberError(errs, NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_REFUSE_MSCHAPV2, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingVpnPptpKeyRequireMppeNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingVpnPptpKeyRequireMppeExists(data) {
		rememberError(errs, NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_REQUIRE_MPPE, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingVpnPptpKeyRequireMppe40NoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingVpnPptpKeyRequireMppe40Exists(data) {
		rememberError(errs, NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_REQUIRE_MPPE_40, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingVpnPptpKeyRequireMppe128NoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingVpnPptpKeyRequireMppe128Exists(data) {
		rememberError(errs, NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_REQUIRE_MPPE_128, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingVpnPptpKeyMppeStatefulNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingVpnPptpKeyMppeStatefulExists(data) {
		rememberError(errs, NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_MPPE_STATEFUL, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingVpnPptpKeyNobsdcompNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingVpnPptpKeyNobsdcompExists(data) {
		rememberError(errs, NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_NOBSDCOMP, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingVpnPptpKeyNodeflateNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingVpnPptpKeyNodeflateExists(data) {
		rememberError(errs, NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_NODEFLATE, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingVpnPptpKeyNoVjCompNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingVpnPptpKeyNoVjCompExists(data) {
		rememberError(errs, NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_NO_VJ_COMP, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingVpnPptpKeyLcpEchoFailureNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingVpnPptpKeyLcpEchoFailureExists(data) {
		rememberError(errs, NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_LCP_ECHO_FAILURE, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingVpnPptpKeyLcpEchoIntervalNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingVpnPptpKeyLcpEchoIntervalExists(data) {
		rememberError(errs, NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_LCP_ECHO_INTERVAL, NM_KEY_ERROR_MISSING_VALUE)
	}
}

// Getter
func getSettingVpnPptpKeyRefuseEap(data connectionData) (value bool) {
	value, _ = getSettingKey(data, NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_REFUSE_EAP).(bool)
	return
}
func getSettingVpnPptpKeyRefusePap(data connectionData) (value bool) {
	value, _ = getSettingKey(data, NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_REFUSE_PAP).(bool)
	return
}
func getSettingVpnPptpKeyRefuseChap(data connectionData) (value bool) {
	value, _ = getSettingKey(data, NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_REFUSE_CHAP).(bool)
	return
}
func getSettingVpnPptpKeyRefuseMschap(data connectionData) (value bool) {
	value, _ = getSettingKey(data, NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_REFUSE_MSCHAP).(bool)
	return
}
func getSettingVpnPptpKeyRefuseMschapv2(data connectionData) (value bool) {
	value, _ = getSettingKey(data, NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_REFUSE_MSCHAPV2).(bool)
	return
}
func getSettingVpnPptpKeyRequireMppe(data connectionData) (value bool) {
	value, _ = getSettingKey(data, NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_REQUIRE_MPPE).(bool)
	return
}
func getSettingVpnPptpKeyRequireMppe40(data connectionData) (value bool) {
	value, _ = getSettingKey(data, NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_REQUIRE_MPPE_40).(bool)
	return
}
func getSettingVpnPptpKeyRequireMppe128(data connectionData) (value bool) {
	value, _ = getSettingKey(data, NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_REQUIRE_MPPE_128).(bool)
	return
}
func getSettingVpnPptpKeyMppeStateful(data connectionData) (value bool) {
	value, _ = getSettingKey(data, NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_MPPE_STATEFUL).(bool)
	return
}
func getSettingVpnPptpKeyNobsdcomp(data connectionData) (value bool) {
	value, _ = getSettingKey(data, NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_NOBSDCOMP).(bool)
	return
}
func getSettingVpnPptpKeyNodeflate(data connectionData) (value bool) {
	value, _ = getSettingKey(data, NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_NODEFLATE).(bool)
	return
}
func getSettingVpnPptpKeyNoVjComp(data connectionData) (value bool) {
	value, _ = getSettingKey(data, NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_NO_VJ_COMP).(bool)
	return
}
func getSettingVpnPptpKeyLcpEchoFailure(data connectionData) (value uint32) {
	value, _ = getSettingKey(data, NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_LCP_ECHO_FAILURE).(uint32)
	return
}
func getSettingVpnPptpKeyLcpEchoInterval(data connectionData) (value uint32) {
	value, _ = getSettingKey(data, NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_LCP_ECHO_INTERVAL).(uint32)
	return
}

// Setter
func setSettingVpnPptpKeyRefuseEap(data connectionData, value bool) {
	setSettingKey(data, NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_REFUSE_EAP, value)
}
func setSettingVpnPptpKeyRefusePap(data connectionData, value bool) {
	setSettingKey(data, NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_REFUSE_PAP, value)
}
func setSettingVpnPptpKeyRefuseChap(data connectionData, value bool) {
	setSettingKey(data, NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_REFUSE_CHAP, value)
}
func setSettingVpnPptpKeyRefuseMschap(data connectionData, value bool) {
	setSettingKey(data, NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_REFUSE_MSCHAP, value)
}
func setSettingVpnPptpKeyRefuseMschapv2(data connectionData, value bool) {
	setSettingKey(data, NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_REFUSE_MSCHAPV2, value)
}
func setSettingVpnPptpKeyRequireMppe(data connectionData, value bool) {
	setSettingKey(data, NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_REQUIRE_MPPE, value)
}
func setSettingVpnPptpKeyRequireMppe40(data connectionData, value bool) {
	setSettingKey(data, NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_REQUIRE_MPPE_40, value)
}
func setSettingVpnPptpKeyRequireMppe128(data connectionData, value bool) {
	setSettingKey(data, NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_REQUIRE_MPPE_128, value)
}
func setSettingVpnPptpKeyMppeStateful(data connectionData, value bool) {
	setSettingKey(data, NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_MPPE_STATEFUL, value)
}
func setSettingVpnPptpKeyNobsdcomp(data connectionData, value bool) {
	setSettingKey(data, NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_NOBSDCOMP, value)
}
func setSettingVpnPptpKeyNodeflate(data connectionData, value bool) {
	setSettingKey(data, NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_NODEFLATE, value)
}
func setSettingVpnPptpKeyNoVjComp(data connectionData, value bool) {
	setSettingKey(data, NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_NO_VJ_COMP, value)
}
func setSettingVpnPptpKeyLcpEchoFailure(data connectionData, value uint32) {
	setSettingKey(data, NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_LCP_ECHO_FAILURE, value)
}
func setSettingVpnPptpKeyLcpEchoInterval(data connectionData, value uint32) {
	setSettingKey(data, NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_LCP_ECHO_INTERVAL, value)
}

// JSON Getter
func getSettingVpnPptpKeyRefuseEapJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_REFUSE_EAP, getSettingVpnPptpPppKeyType(NM_SETTING_VPN_PPTP_KEY_REFUSE_EAP))
	return
}
func getSettingVpnPptpKeyRefusePapJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_REFUSE_PAP, getSettingVpnPptpPppKeyType(NM_SETTING_VPN_PPTP_KEY_REFUSE_PAP))
	return
}
func getSettingVpnPptpKeyRefuseChapJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_REFUSE_CHAP, getSettingVpnPptpPppKeyType(NM_SETTING_VPN_PPTP_KEY_REFUSE_CHAP))
	return
}
func getSettingVpnPptpKeyRefuseMschapJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_REFUSE_MSCHAP, getSettingVpnPptpPppKeyType(NM_SETTING_VPN_PPTP_KEY_REFUSE_MSCHAP))
	return
}
func getSettingVpnPptpKeyRefuseMschapv2JSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_REFUSE_MSCHAPV2, getSettingVpnPptpPppKeyType(NM_SETTING_VPN_PPTP_KEY_REFUSE_MSCHAPV2))
	return
}
func getSettingVpnPptpKeyRequireMppeJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_REQUIRE_MPPE, getSettingVpnPptpPppKeyType(NM_SETTING_VPN_PPTP_KEY_REQUIRE_MPPE))
	return
}
func getSettingVpnPptpKeyRequireMppe40JSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_REQUIRE_MPPE_40, getSettingVpnPptpPppKeyType(NM_SETTING_VPN_PPTP_KEY_REQUIRE_MPPE_40))
	return
}
func getSettingVpnPptpKeyRequireMppe128JSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_REQUIRE_MPPE_128, getSettingVpnPptpPppKeyType(NM_SETTING_VPN_PPTP_KEY_REQUIRE_MPPE_128))
	return
}
func getSettingVpnPptpKeyMppeStatefulJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_MPPE_STATEFUL, getSettingVpnPptpPppKeyType(NM_SETTING_VPN_PPTP_KEY_MPPE_STATEFUL))
	return
}
func getSettingVpnPptpKeyNobsdcompJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_NOBSDCOMP, getSettingVpnPptpPppKeyType(NM_SETTING_VPN_PPTP_KEY_NOBSDCOMP))
	return
}
func getSettingVpnPptpKeyNodeflateJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_NODEFLATE, getSettingVpnPptpPppKeyType(NM_SETTING_VPN_PPTP_KEY_NODEFLATE))
	return
}
func getSettingVpnPptpKeyNoVjCompJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_NO_VJ_COMP, getSettingVpnPptpPppKeyType(NM_SETTING_VPN_PPTP_KEY_NO_VJ_COMP))
	return
}
func getSettingVpnPptpKeyLcpEchoFailureJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_LCP_ECHO_FAILURE, getSettingVpnPptpPppKeyType(NM_SETTING_VPN_PPTP_KEY_LCP_ECHO_FAILURE))
	return
}
func getSettingVpnPptpKeyLcpEchoIntervalJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_LCP_ECHO_INTERVAL, getSettingVpnPptpPppKeyType(NM_SETTING_VPN_PPTP_KEY_LCP_ECHO_INTERVAL))
	return
}

// JSON Setter
func setSettingVpnPptpKeyRefuseEapJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_REFUSE_EAP, valueJSON, getSettingVpnPptpPppKeyType(NM_SETTING_VPN_PPTP_KEY_REFUSE_EAP))
}
func setSettingVpnPptpKeyRefusePapJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_REFUSE_PAP, valueJSON, getSettingVpnPptpPppKeyType(NM_SETTING_VPN_PPTP_KEY_REFUSE_PAP))
}
func setSettingVpnPptpKeyRefuseChapJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_REFUSE_CHAP, valueJSON, getSettingVpnPptpPppKeyType(NM_SETTING_VPN_PPTP_KEY_REFUSE_CHAP))
}
func setSettingVpnPptpKeyRefuseMschapJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_REFUSE_MSCHAP, valueJSON, getSettingVpnPptpPppKeyType(NM_SETTING_VPN_PPTP_KEY_REFUSE_MSCHAP))
}
func setSettingVpnPptpKeyRefuseMschapv2JSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_REFUSE_MSCHAPV2, valueJSON, getSettingVpnPptpPppKeyType(NM_SETTING_VPN_PPTP_KEY_REFUSE_MSCHAPV2))
}
func setSettingVpnPptpKeyRequireMppeJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_REQUIRE_MPPE, valueJSON, getSettingVpnPptpPppKeyType(NM_SETTING_VPN_PPTP_KEY_REQUIRE_MPPE))
}
func setSettingVpnPptpKeyRequireMppe40JSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_REQUIRE_MPPE_40, valueJSON, getSettingVpnPptpPppKeyType(NM_SETTING_VPN_PPTP_KEY_REQUIRE_MPPE_40))
}
func setSettingVpnPptpKeyRequireMppe128JSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_REQUIRE_MPPE_128, valueJSON, getSettingVpnPptpPppKeyType(NM_SETTING_VPN_PPTP_KEY_REQUIRE_MPPE_128))
}
func setSettingVpnPptpKeyMppeStatefulJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_MPPE_STATEFUL, valueJSON, getSettingVpnPptpPppKeyType(NM_SETTING_VPN_PPTP_KEY_MPPE_STATEFUL))
}
func setSettingVpnPptpKeyNobsdcompJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_NOBSDCOMP, valueJSON, getSettingVpnPptpPppKeyType(NM_SETTING_VPN_PPTP_KEY_NOBSDCOMP))
}
func setSettingVpnPptpKeyNodeflateJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_NODEFLATE, valueJSON, getSettingVpnPptpPppKeyType(NM_SETTING_VPN_PPTP_KEY_NODEFLATE))
}
func setSettingVpnPptpKeyNoVjCompJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_NO_VJ_COMP, valueJSON, getSettingVpnPptpPppKeyType(NM_SETTING_VPN_PPTP_KEY_NO_VJ_COMP))
}
func setSettingVpnPptpKeyLcpEchoFailureJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_LCP_ECHO_FAILURE, valueJSON, getSettingVpnPptpPppKeyType(NM_SETTING_VPN_PPTP_KEY_LCP_ECHO_FAILURE))
}
func setSettingVpnPptpKeyLcpEchoIntervalJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_LCP_ECHO_INTERVAL, valueJSON, getSettingVpnPptpPppKeyType(NM_SETTING_VPN_PPTP_KEY_LCP_ECHO_INTERVAL))
}

// Logic JSON Setter
func logicSetSettingVpnPptpKeyRequireMppeJSON(data connectionData, valueJSON string) (err error) {
	err = setSettingVpnPptpKeyRequireMppeJSON(data, valueJSON)
	if err != nil {
		return
	}
	if isSettingVpnPptpKeyRequireMppeExists(data) {
		value := getSettingVpnPptpKeyRequireMppe(data)
		err = logicSetSettingVpnPptpKeyRequireMppe(data, value)
	}
	return
}

// Remover
func removeSettingVpnPptpKeyRefuseEap(data connectionData) {
	removeSettingKey(data, NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_REFUSE_EAP)
}
func removeSettingVpnPptpKeyRefusePap(data connectionData) {
	removeSettingKey(data, NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_REFUSE_PAP)
}
func removeSettingVpnPptpKeyRefuseChap(data connectionData) {
	removeSettingKey(data, NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_REFUSE_CHAP)
}
func removeSettingVpnPptpKeyRefuseMschap(data connectionData) {
	removeSettingKey(data, NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_REFUSE_MSCHAP)
}
func removeSettingVpnPptpKeyRefuseMschapv2(data connectionData) {
	removeSettingKey(data, NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_REFUSE_MSCHAPV2)
}
func removeSettingVpnPptpKeyRequireMppe(data connectionData) {
	removeSettingKey(data, NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_REQUIRE_MPPE)
}
func removeSettingVpnPptpKeyRequireMppe40(data connectionData) {
	removeSettingKey(data, NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_REQUIRE_MPPE_40)
}
func removeSettingVpnPptpKeyRequireMppe128(data connectionData) {
	removeSettingKey(data, NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_REQUIRE_MPPE_128)
}
func removeSettingVpnPptpKeyMppeStateful(data connectionData) {
	removeSettingKey(data, NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_MPPE_STATEFUL)
}
func removeSettingVpnPptpKeyNobsdcomp(data connectionData) {
	removeSettingKey(data, NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_NOBSDCOMP)
}
func removeSettingVpnPptpKeyNodeflate(data connectionData) {
	removeSettingKey(data, NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_NODEFLATE)
}
func removeSettingVpnPptpKeyNoVjComp(data connectionData) {
	removeSettingKey(data, NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_NO_VJ_COMP)
}
func removeSettingVpnPptpKeyLcpEchoFailure(data connectionData) {
	removeSettingKey(data, NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_LCP_ECHO_FAILURE)
}
func removeSettingVpnPptpKeyLcpEchoInterval(data connectionData) {
	removeSettingKey(data, NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_LCP_ECHO_INTERVAL)
}