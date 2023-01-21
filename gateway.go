package gateway

import (
	"errors"
	"net"
	"runtime"
)

var (
	errNoGateway      = errors.New("no gateway found")
	errCantParse      = errors.New("can't parse string output")
	errNotImplemented = errors.New("not implemented for OS: " + runtime.GOOS)
)

// DiscoverGateway is the OS independent function to get the default gateway
func DiscoverGateway() (ip net.IP, err error) {
	return discoverGatewayOSSpecific()
}

// DiscoverInterface is the OS independent function to call to get the default network interface IP that uses the default gateway
func DiscoverInterfaceIp() (ip net.IP, err error) {
	ipOfGateway, err1 := discoverGatewayOSSpecific()
	if err1 != nil {
		return nil, err1
	}
	allInts, _ := net.Interfaces()
	for _, netInt := range allInts {
		addrs, _ := netInt.Addrs()
		for _, addr := range addrs {

			ipA, ipnetA, err2 := net.ParseCIDR(addr.String())
			if err2 != nil {
				return nil, err1
			}
			if ipnetA.Contains(ipOfGateway) {
				return ipA, nil

			}

		}
	}
	return nil, errors.New("Could not find interface")

}

func DiscoverInterface() (netInterface *net.Interface, err error) {
	ipOfGateway, err1 := discoverGatewayOSSpecific()
	if err1 != nil {
		return nil, err1
	}
	allInts, _ := net.Interfaces()
	for _, netInt := range allInts {
		addrs, _ := netInt.Addrs()
		for _, addr := range addrs {

			_, ipnetA, err2 := net.ParseCIDR(addr.String())
			if err2 != nil {
				return nil, err1
			}
			if ipnetA.Contains(ipOfGateway) {
				return &netInt, nil

			}

		}
	}
	return nil, errors.New("Could not find interface")

}
