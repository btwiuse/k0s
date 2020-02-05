// Copyright (c) 2016-present Cloud <cloud@txthinking.com>
//
// This program is free software; you can redistribute it and/or
// modify it under the terms of version 3 of the GNU General Public
// License as published by the Free Software Foundation.
//
// This program is distributed in the hope that it will be useful, but
// WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU
// General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>.

package sysproxy

// GetNetworkInterface returns default interface dev name.
func GetNetworkInterface() (string, error) {
	return "", nil
}

// GetDefaultGateway returns default gateway.
func GetDefaultGateway() (string, error) {
	return "", nil
}

// GetDNSServers used to get DNS servers.
func GetDNSServers() ([]string, error) {
	return nil, nil
}

// SetDNSServers used to set system DNS servers.
func SetDNSServers(servers []string) error {
	return nil
}

func TurnOnSystemProxy(pac string) error {
	return nil
}

func TurnOffSystemProxy() error {
	return nil
}
