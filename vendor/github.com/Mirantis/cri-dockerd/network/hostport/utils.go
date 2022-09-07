package hostport

import (
	"bytes"
	"fmt"

	utiliptables "k8s.io/kubernetes/pkg/util/iptables"
)

var (
	commitBytes = []byte("COMMIT")
	spaceBytes  = []byte(" ")
)

const (
	// K3S: Adding this here since v1.25.0 removed this
	// exported variable

	// kubeMarkMasqChain is the mark-for-masquerade chain
	KubeMarkMasqChain utiliptables.Chain = "KUBE-MARK-MASQ"
)

// K3S: GetChainLines was removed in v1.25.0 https://github.com/kubernetes/kubernetes/pull/110328

// GetChainLines parses a table's iptables-save data to find chains in the table.
// It returns a map of iptables.Chain to []byte where the []byte is the chain line
// from save (with counters etc.).
// Note that to avoid allocations memory is SHARED with save.
func GetChainLines(table utiliptables.Table, save []byte) map[utiliptables.Chain][]byte {
	chainsMap := make(map[utiliptables.Chain][]byte)
	tablePrefix := []byte("*" + string(table))
	readIndex := 0
	// find beginning of table
	for readIndex < len(save) {
		line, n := readLine(readIndex, save)
		readIndex = n
		if bytes.HasPrefix(line, tablePrefix) {
			break
		}
	}
	// parse table lines
	for readIndex < len(save) {
		line, n := readLine(readIndex, save)
		readIndex = n
		if len(line) == 0 {
			continue
		}
		if bytes.HasPrefix(line, commitBytes) || line[0] == '*' {
			break
		} else if line[0] == '#' {
			continue
		} else if line[0] == ':' && len(line) > 1 {
			// We assume that the <line> contains space - chain lines have 3 fields,
			// space delimited. If there is no space, this line will panic.
			spaceIndex := bytes.Index(line, spaceBytes)
			if spaceIndex == -1 {
				panic(fmt.Sprintf("Unexpected chain line in iptables-save output: %v", string(line)))
			}
			chain := utiliptables.Chain(line[1:spaceIndex])
			chainsMap[chain] = line
		}
	}
	return chainsMap
}

func readLine(readIndex int, byteArray []byte) ([]byte, int) {
	currentReadIndex := readIndex

	// consume left spaces
	for currentReadIndex < len(byteArray) {
		if byteArray[currentReadIndex] == ' ' {
			currentReadIndex++
		} else {
			break
		}
	}

	// leftTrimIndex stores the left index of the line after the line is left-trimmed
	leftTrimIndex := currentReadIndex

	// rightTrimIndex stores the right index of the line after the line is right-trimmed
	// it is set to -1 since the correct value has not yet been determined.
	rightTrimIndex := -1

	for ; currentReadIndex < len(byteArray); currentReadIndex++ {
		if byteArray[currentReadIndex] == ' ' {
			// set rightTrimIndex
			if rightTrimIndex == -1 {
				rightTrimIndex = currentReadIndex
			}
		} else if (byteArray[currentReadIndex] == '\n') || (currentReadIndex == (len(byteArray) - 1)) {
			// end of line or byte buffer is reached
			if currentReadIndex <= leftTrimIndex {
				return nil, currentReadIndex + 1
			}
			// set the rightTrimIndex
			if rightTrimIndex == -1 {
				rightTrimIndex = currentReadIndex
				if currentReadIndex == (len(byteArray)-1) && (byteArray[currentReadIndex] != '\n') {
					// ensure that the last character is part of the returned string,
					// unless the last character is '\n'
					rightTrimIndex = currentReadIndex + 1
				}
			}
			// Avoid unnecessary allocation.
			return byteArray[leftTrimIndex:rightTrimIndex], currentReadIndex + 1
		} else {
			// unset rightTrimIndex
			rightTrimIndex = -1
		}
	}
	return nil, currentReadIndex
}
