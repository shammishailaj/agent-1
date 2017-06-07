package authorized_keys

import (
	"fmt"
	"strings"
	"testing"
)

func TestAddInHelpBlock(t *testing.T) {
	filedata := "ABC\nDEF\n99999\n"
	var keys []string = []string{"key1", "key2", "key3"}

	result, err := AddKeysToFile(filedata, keys)
	if err != nil {
		t.Fatal(err)
	}
	corevalidate(result, t)
}

func TestWithExistingpBlock(t *testing.T) {
	filedata := "ABC\nDEF\n99999\n" + BLOCK_START_MESSAGE + "ASDF AFSD \n" + BLOCK_END_MESSAGE
	var keys []string = []string{"key1", "key2", "key3"}

	result, err := AddKeysToFile(filedata, keys)
	if err != nil {
		t.Fatal(err)
	}
	corevalidate(result, t)
}

func corevalidate(result string, t *testing.T) {
	blockHasStartMsgIdx := strings.Index(result, BLOCK_START_MESSAGE)
	blockHasStartMsgLastIdx := strings.LastIndex(result, BLOCK_START_MESSAGE)

	if blockHasStartMsgIdx < 0 || blockHasStartMsgIdx != blockHasStartMsgLastIdx {
		t.Errorf("Expected %d start message block got idx %d", blockHasStartMsgIdx, blockHasStartMsgLastIdx)
	}

	blockHasLastMsgIdx := strings.Index(result, BLOCK_END_MESSAGE)
	blockHasLastMsgLastIdx := strings.LastIndex(result, BLOCK_END_MESSAGE)

	if blockHasLastMsgIdx < 0 || blockHasLastMsgIdx != blockHasLastMsgLastIdx {
		t.Errorf("Expected %d start message block got idx %d", blockHasLastMsgIdx, blockHasLastMsgLastIdx)
	}

	//ensure the old keys didnt get deleted
	oldKey1 := strings.Index(result, "ABC")
	if oldKey1 != 0 {
		t.Errorf("Expected %d index for old key but got idx %d", 0, oldKey1)
	}

	fmt.Printf("data \n " + result + "\n")
	//See if the new key got added
	newKey1 := strings.Index(result, "key1")
	if newKey1 < 0 {
		t.Errorf("Expected %d index for new key but got idx %d", 0, newKey1)
	}
}
