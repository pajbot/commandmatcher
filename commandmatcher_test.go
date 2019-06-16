package commandmatcher

import (
	"testing"

	"github.com/pajbot/testhelper"
)

type test struct {
}

var (
	command = &test{}
)

func TestNonMatchingCommand(t *testing.T) {
	const testString = `!test lol`
	m := New()
	match, _ := m.Match(testString)
	testhelper.AssertNil(t, match, "No command should be matched here")
}

func TestRegisterCommand(t *testing.T) {
	const testString = `!test lol`
	m := New()
	cmd2 := m.Register([]string{"!test"}, command)
	match, _ := m.Match(testString)
	testhelper.AssertNotNil(t, match, "A command should be matched here")
	testhelper.AssertInterfacesEqual(t, command, match, "A command should be matched here")
	testhelper.AssertInterfacesEqual(t, command, cmd2, "A command should be matched here")
}

func TestDeregisterCommand(t *testing.T) {
	const testString = `!test lol`
	m := New()
	cmd2 := m.Register([]string{"!test"}, command)
	match, _ := m.Match(testString)
	testhelper.AssertNotNil(t, match, "A command should be matched here")
	testhelper.AssertInterfacesEqual(t, command, match, "A command should be matched here")
	testhelper.AssertInterfacesEqual(t, command, cmd2, "A command should be matched here")
	m.Deregister(cmd2)
	match, _ = m.Match(testString)
	testhelper.AssertNil(t, match, "No command should be matched here")
}
