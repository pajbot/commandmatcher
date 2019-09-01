package commandmatcher

import (
	"strings"
	"sync"
)

type CommandMatcher struct {
	commandsMutex *sync.RWMutex
	commands      map[string]interface{}

	Separator string
}

func New() *CommandMatcher {
	m := &CommandMatcher{
		commandsMutex: &sync.RWMutex{},
		commands:      make(map[string]interface{}),

		Separator: " ",
	}

	return m
}

func (m *CommandMatcher) Register(aliases []string, command interface{}) interface{} {
	m.commandsMutex.Lock()
	defer m.commandsMutex.Unlock()

	for _, alias := range aliases {
		m.commands[alias] = command
	}

	return command
}

func (m *CommandMatcher) DeregisterAliases(aliases []string) {
	m.commandsMutex.Lock()
	defer m.commandsMutex.Unlock()

	for _, alias := range aliases {
		delete(m.commands, alias)
	}
}

func (m *CommandMatcher) Deregister(command interface{}) {
	m.commandsMutex.Lock()
	defer m.commandsMutex.Unlock()

	var aliasesToRemove []string
	for alias, cmd := range m.commands {
		if cmd == command {
			aliasesToRemove = append(aliasesToRemove, alias)
		}
	}

	for _, alias := range aliasesToRemove {
		delete(m.commands, alias)
	}
}

func (m *CommandMatcher) Match(text string) (interface{}, []string) {
	parts := strings.Split(text, m.Separator)

	m.commandsMutex.RLock()
	defer m.commandsMutex.RUnlock()

	if command, ok := m.commands[parts[0]]; ok {
		return command, parts
	}

	return nil, parts
}

func (m *CommandMatcher) ForEach(cb func([]string, interface{})) {
	m.commandsMutex.Lock()
	defer m.commandsMutex.Unlock()

	uniqueCommands := map[interface{}][]string{}

	for alias, cmd := range m.commands {
		uniqueCommands[cmd] = append(uniqueCommands[cmd], alias)
	}

	for cmd, aliases := range uniqueCommands {
		cb(aliases, cmd)
	}
}

func (m *CommandMatcher) Find(alias string) interface{} {
	m.commandsMutex.Lock()
	defer m.commandsMutex.Unlock()

	cmd, ok := m.commands[alias]
	if !ok {
		return nil
	}

	return cmd
}
