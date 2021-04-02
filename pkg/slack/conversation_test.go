package slack

import (
	"testing"
)

func TestNewConversationFromID(t *testing.T) {
	tc := map[string][]string{
		"Channel":        []string{"CH00001", "#test"},
		"Direct Message": []string{"DM00001", "DM"},
		"Group Message":  []string{"GR00001", "Group messaging with: @some @users"},
	}

	client := NewMockClient()
	for testID, data := range tc {
		conversation, err := NewConversationFromID(data[0], client)
		if err != nil {
			t.Logf("NewConversationFromID errored for %v: %v", testID, err)
			t.Fail()
		}
		if conversation.Name != data[1] {
			t.Logf("Conversation name was %v, instead of expected %v", conversation.Name, data[1])
			t.Fail()
		}
	}
}
