package main

type commandManager struct {
    commands map[string]func(context)
}

func newCommandManager() *commandManager {
    return &commandManager{make(map[string]func(context))}
}

func (manager commandManager) isCommand(name string) bool {
    if _, ok := manager.commands[name]; ok {
        return true
    }
    return false
}

func (manager *commandManager) register(name string, cmd func(ctx context)) {
    manager.commands[name] = cmd
}