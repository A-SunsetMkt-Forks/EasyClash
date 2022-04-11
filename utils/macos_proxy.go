package utils

import "strconv"

func systemProxy(address string, httpPort, sockPort int) error {
	status := "off"

	if address == "" {
		address = "127.0.0.1"
	}

	if httpPort > 0 || sockPort > 0 {
		status = "on"
	}

	if httpPort == 0 && sockPort == 0 {
		address = ""
	}

	_httpPort := ""
	if httpPort > 0 {
		_httpPort = strconv.FormatInt(int64(httpPort), 10)
	}
	_sockPort := ""
	if sockPort > 0 {
		_sockPort = strconv.FormatInt(int64(sockPort), 10)
	}

	_, err := Exec("networksetup", "-setwebproxy", "Wi-Fi", address, _httpPort)
	if err != nil {
		return err
	}
	_, err = Exec("networksetup", "-setsecurewebproxy", "Wi-Fi", address, _httpPort)
	if err != nil {
		return err
	}
	_, err = Exec("networksetup", "-setsocksfirewallproxy", "Wi-Fi", address, _sockPort)
	if err != nil {
		return err
	}
	_, err = Exec("networksetup", "-setwebproxystate", "Wi-Fi", status)
	if err != nil {
		return err
	}
	_, err = Exec("networksetup", "-setsecurewebproxystate", "Wi-Fi", status)
	if err != nil {
		return err
	}
	_, err = Exec("networksetup", "-setsocksfirewallproxystate", "Wi-Fi", status)
	if err != nil {
		return err
	}

	return nil
}

func SetSystemProxy(httpPort, sockPort int) error {
	return systemProxy("", httpPort, sockPort)
}

func UnSetSystemProxy() error {
	return systemProxy("", 0, 0)
}
