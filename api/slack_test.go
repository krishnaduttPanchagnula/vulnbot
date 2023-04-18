package api

import (
	"fmt"
	"testing"

	"github.com/slack-go/slack"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockSlackClient struct {
	mock.Mock
}

func (m *MockSlackClient) PostMessage(channelID string, options ...slack.MsgOption) (string, string, error) {
	args := m.Called(channelID, options)
	return args.String(0), args.String(1), args.Error(2)
}

func TestSendSlackMessagesSuccess(t *testing.T) {
	// Create a mock Slack client
	mockClient := new(MockSlackClient)

	// Set up test messages
	messages := map[string]string{
		"channel": "message",
	}

	// Test case: Successful send
	mockClient.On("PostMessage", "channel", mock.Anything, mock.Anything).Return("", "", nil).Once()

	// Run test
	SendSlackMessages(messages, mockClient)

	mockClient.AssertExpectations(t)
}

func TestSendSlackMessagesError(t *testing.T) {
	// Create a mock Slack client
	mockClient := new(MockSlackClient)

	// Set up test messages
	messages := map[string]string{
		"channel": "message",
	}

	// Test case: Error sending Slack message
	mockClient.On("PostMessage", "channel", mock.Anything, mock.Anything).Return("", "", fmt.Errorf("Failed to send Slack message")).Once()

	// Run test
	SendSlackMessages(messages, mockClient)

	mockClient.AssertExpectations(t)
}

func TestIsSlackTokenMissing(t *testing.T) {
	_, err := NewSlackClient("")
	assert.Error(t, err)
}

func TestSlackTokenIsNotMissing(t *testing.T) {
	_, err := NewSlackClient("slackToken")
	assert.NoError(t, err)
}
