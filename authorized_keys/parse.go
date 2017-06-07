package authorized_keys

import (
	"fmt"
	"strings"
)

var (
	BLOCK_START_MESSAGE = "\n##### DO NOT MODIFY BELOW EDITED BY OPSCOIN BLOCKSSH AGENT #####\n\n"
	BLOCK_END_MESSAGE   = "\n##### FINISHED DO NOT MODIFY BLOCK BY OPSCOIN BLOCKSSH AGENT #####\n"
)

// AddKeysToFile parses a .ssh/authorized_key file and adds in the keys you want
func AddKeysToFile(filedata string, keys []string) (string, error) {
	//TODO change this to be a streaming function, still rough
	output := filedata

	//Delete any old blocks
	count := 0
	for {
		count++
		startMsgIdx := strings.Index(output, BLOCK_START_MESSAGE)
		lastMsgIdx := strings.Index(output, BLOCK_END_MESSAGE)

		//Only delete if we find a start and end block
		if startMsgIdx == -1 || lastMsgIdx == -1 || count >= 3 {
			break
		}

		end := lastMsgIdx + len(BLOCK_END_MESSAGE)
		tmp := output[0:startMsgIdx]
		output = tmp + output[end:]
	}

	output += BLOCK_START_MESSAGE
	for _, key := range keys {
		//TODO support different key formats in future
		//right now its only EDSA
		output += fmt.Sprintf("ecdsa-sha2-nistp256 %s misc\n", key)
	}
	output += BLOCK_END_MESSAGE

	return output, nil
}
