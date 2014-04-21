package main

import (
	"fmt"
)

// TODO doc

const NM_SETTING_IP6_CONFIG_SETTING_NAME = "ipv6"

const (
	NM_SETTING_IP6_CONFIG_METHOD             = "method"
	NM_SETTING_IP6_CONFIG_DNS                = "dns"
	NM_SETTING_IP6_CONFIG_DNS_SEARCH         = "dns-search"
	NM_SETTING_IP6_CONFIG_ADDRESSES          = "addresses"
	NM_SETTING_IP6_CONFIG_ROUTES             = "routes"
	NM_SETTING_IP6_CONFIG_IGNORE_AUTO_ROUTES = "ignore-auto-routes"
	NM_SETTING_IP6_CONFIG_IGNORE_AUTO_DNS    = "ignore-auto-dns"
	NM_SETTING_IP6_CONFIG_NEVER_DEFAULT      = "never-default"
	NM_SETTING_IP6_CONFIG_MAY_FAIL           = "may-fail"
	NM_SETTING_IP6_CONFIG_IP6_PRIVACY        = "ip6-privacy"
	NM_SETTING_IP6_CONFIG_DHCP_HOSTNAME      = "dhcp-hostname"
)

const (
	NM_SETTING_IP6_CONFIG_METHOD_IGNORE     = "ignore"
	NM_SETTING_IP6_CONFIG_METHOD_AUTO       = "auto"
	NM_SETTING_IP6_CONFIG_METHOD_DHCP       = "dhcp"
	NM_SETTING_IP6_CONFIG_METHOD_LINK_LOCAL = "link-local"
	NM_SETTING_IP6_CONFIG_METHOD_MANUAL     = "manual"
	NM_SETTING_IP6_CONFIG_METHOD_SHARED     = "shared"
)

// Get available keys
func getSettingIp6ConfigAvailableKeys(data _ConnectionData) (keys []string) {
	method := getSettingIp6ConfigMethod(data)
	switch method {
	default:
		Logger.Error("ip6 config method is invalid:", method)
	case NM_SETTING_IP6_CONFIG_METHOD_IGNORE:
		keys = []string{
			NM_SETTING_IP6_CONFIG_METHOD,
		}
	case NM_SETTING_IP6_CONFIG_METHOD_AUTO:
		keys = []string{
			NM_SETTING_IP6_CONFIG_METHOD,
		}
		keys = appendStrArrayUnion(keys, getRelatedAvailableVirtualKeys(fieldIPv6, NM_SETTING_IP6_CONFIG_DNS)...)
	case NM_SETTING_IP6_CONFIG_METHOD_DHCP: // ignore
		keys = []string{
			NM_SETTING_IP6_CONFIG_METHOD,
		}
		keys = appendStrArrayUnion(keys, getRelatedAvailableVirtualKeys(fieldIPv6, NM_SETTING_IP6_CONFIG_DNS)...)
	case NM_SETTING_IP6_CONFIG_METHOD_LINK_LOCAL: // ignore
	case NM_SETTING_IP6_CONFIG_METHOD_MANUAL:
		keys = []string{
			NM_SETTING_IP6_CONFIG_METHOD,
		}
		keys = appendStrArrayUnion(keys, getRelatedAvailableVirtualKeys(fieldIPv6, NM_SETTING_IP6_CONFIG_DNS)...)
		keys = appendStrArrayUnion(keys, getRelatedAvailableVirtualKeys(fieldIPv6, NM_SETTING_IP6_CONFIG_ADDRESSES)...)
	case NM_SETTING_IP6_CONFIG_METHOD_SHARED: // ignore
	}
	return
}

// Get available values
func getSettingIp6ConfigAvailableValues(key string) (values []string, customizable bool) {
	customizable = true
	switch key {
	case NM_SETTING_IP6_CONFIG_METHOD:
		values = []string{
			// NM_SETTING_IP6_CONFIG_METHOD_IGNORE, // ignore
			NM_SETTING_IP6_CONFIG_METHOD_AUTO,
			// NM_SETTING_IP6_CONFIG_METHOD_DHCP, // ignore
			// NM_SETTING_IP6_CONFIG_METHOD_LINK_LOCAL, // ignore
			NM_SETTING_IP6_CONFIG_METHOD_MANUAL,
			// NM_SETTING_IP6_CONFIG_METHOD_SHARED,// ignore
		}
		customizable = false
	}
	return
}

// Check whether the values are correct
func checkSettingIp6ConfigValues(data _ConnectionData) (errs map[string]string) {
	errs = make(map[string]string)

	// check method
	ensureSettingIp6ConfigMethodNoEmpty(data, errs)
	switch getSettingIp6ConfigMethod(data) {
	default:
		rememberError(errs, NM_SETTING_IP6_CONFIG_METHOD, NM_KEY_ERROR_INVALID_VALUE)
		return
	case NM_SETTING_IP6_CONFIG_METHOD_IGNORE: // ignore
		checkSettingIp6MethodConflict(data, errs)
	case NM_SETTING_IP6_CONFIG_METHOD_AUTO:
	case NM_SETTING_IP6_CONFIG_METHOD_DHCP: // ignore
	case NM_SETTING_IP6_CONFIG_METHOD_LINK_LOCAL: // ignore
		checkSettingIp6MethodConflict(data, errs)
	case NM_SETTING_IP6_CONFIG_METHOD_MANUAL:
		ensureSettingIp6ConfigAddressesNoEmpty(data, errs)
	case NM_SETTING_IP6_CONFIG_METHOD_SHARED: // ignore
		checkSettingIp6MethodConflict(data, errs)
	}

	// check value of dns
	checkSettingIp6ConfigDns(data, errs)

	// check value of address
	checkSettingIp6ConfigAddresses(data, errs)

	// TODO check value of route

	return
}

func checkSettingIp6MethodConflict(data _ConnectionData, errs map[string]string) {
	// check dns
	if isSettingIp6ConfigDnsExists(data) {
		rememberVkError(errs, fieldIPv6, NM_SETTING_IP6_CONFIG_DNS, fmt.Sprintf(NM_KEY_ERROR_IP6_METHOD_CONFLICT, NM_SETTING_IP6_CONFIG_DNS))
	}
	// check dns search
	if isSettingIp6ConfigDnsSearchExists(data) {
		rememberError(errs, NM_SETTING_IP6_CONFIG_DNS_SEARCH, fmt.Sprintf(NM_KEY_ERROR_IP6_METHOD_CONFLICT, NM_SETTING_IP6_CONFIG_DNS_SEARCH))
	}
	// check address
	if isSettingIp6ConfigAddressesExists(data) {
		rememberVkError(errs, fieldIPv6, NM_SETTING_IP6_CONFIG_ADDRESSES, fmt.Sprintf(NM_KEY_ERROR_IP6_METHOD_CONFLICT, NM_SETTING_IP6_CONFIG_ADDRESSES))
	}
	// check route
	if isSettingIp6ConfigRoutesExists(data) {
		rememberVkError(errs, fieldIPv6, NM_SETTING_IP6_CONFIG_ROUTES, fmt.Sprintf(NM_KEY_ERROR_IP6_METHOD_CONFLICT, NM_SETTING_IP6_CONFIG_ROUTES))
	}
}

func checkSettingIp6ConfigDns(data _ConnectionData, errs map[string]string) {
	if !isSettingIp6ConfigDnsExists(data) {
		return
	}
	dnses := getSettingIp6ConfigDns(data)
	for _, dns := range dnses {
		if !isIpv6AddressValid(dns) {
			rememberError(errs, NM_SETTING_VK_IP6_CONFIG_DNS, NM_KEY_ERROR_INVALID_VALUE)
			return
		}
		if isIpv6AddressZero(dns) {
			rememberError(errs, NM_SETTING_VK_IP6_CONFIG_DNS, NM_KEY_ERROR_INVALID_VALUE)
			return
		}
	}
}

func checkSettingIp6ConfigAddresses(data _ConnectionData, errs map[string]string) {
	if !isSettingIp6ConfigAddressesExists(data) {
		return
	}
	addresses := getSettingIp6ConfigAddresses(data)
	for _, addr := range addresses {
		// check address
		if !isIpv6AddressValid(addr.Address) || isIpv6AddressZero(addr.Address) {
			rememberError(errs, NM_SETTING_VK_IP6_CONFIG_ADDRESSES_ADDRESS, NM_KEY_ERROR_INVALID_VALUE)
			return
		}
		// check prefix
		if addr.Prefix < 1 || addr.Prefix > 128 {
			rememberError(errs, NM_SETTING_VK_IP6_CONFIG_ADDRESSES_PREFIX, NM_KEY_ERROR_INVALID_VALUE)
			return
		}
		// check gateway
		if !isIpv6AddressValid(addr.Gateway) {
			rememberError(errs, NM_SETTING_VK_IP6_CONFIG_ADDRESSES_GATEWAY, NM_KEY_ERROR_INVALID_VALUE)
			return
		}
	}
}

// Logic setter
func logicSetSettingIp6ConfigMethodJSON(data _ConnectionData, valueJSON string) {
	setSettingIp6ConfigMethodJSON(data, valueJSON)

	value := getSettingIp6ConfigMethod(data)
	logicSetSettingIp6ConfigMethod(data, value)
}
func logicSetSettingIp6ConfigMethod(data _ConnectionData, value string) {
	switch value {
	case NM_SETTING_IP6_CONFIG_METHOD_IGNORE: // ignore
		removeSettingKeyBut(data, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_METHOD)
	case NM_SETTING_IP6_CONFIG_METHOD_AUTO:
		removeSettingIp6ConfigAddresses(data)
	case NM_SETTING_IP6_CONFIG_METHOD_DHCP: // ignore
	case NM_SETTING_IP6_CONFIG_METHOD_LINK_LOCAL: // ignore
		removeSettingIp6ConfigDns(data)
		removeSettingIp6ConfigDnsSearch(data)
		removeSettingIp6ConfigAddresses(data)
		removeSettingIp6ConfigRoutes(data)
	case NM_SETTING_IP6_CONFIG_METHOD_MANUAL:
	case NM_SETTING_IP6_CONFIG_METHOD_SHARED: // ignore
		removeSettingIp6ConfigDns(data)
		removeSettingIp6ConfigDnsSearch(data)
		removeSettingIp6ConfigAddresses(data)
		removeSettingIp6ConfigRoutes(data)
	}
	setSettingIp6ConfigMethod(data, value)
}