package cmd

import (
	"../framework"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const todo_file_name = "todo.json"

type jsonStructure struct {
	Entries []string
}

func (s *jsonStructure) add(entry string) {
	s.Entries = append(s.Entries, entry)
}

func readTodo() *jsonStructure {
	bBody, err := ioutil.ReadFile(todo_file_name)
	if err != nil {
		fmt.Println("err reading todo file,", err)
		return nil
	}
	s := &jsonStructure{make([]string, 0)}
	err = json.Unmarshal(bBody, s)
	if err != nil {
		fmt.Println("err unmarshalling todo file,", err)
		return nil
	}
	return s
}

func writeTodo(s jsonStructure) {
	body, err := json.Marshal(s)
	if err != nil {
		fmt.Println("err marshalling todo file,", err)
		return
	}
	err = ioutil.WriteFile(todo_file_name, body, os.ModeAppend)
	if err != nil {
		fmt.Println("err writing todo file,", err)
		return
	}
}

func AdminCommand(ctx framework.Context) {
	if ctx.User.ID != ctx.Conf.OwnerId {
		return
	}
	if len(ctx.Args) == 0 {
		ctx.Reply("Usage: music admin <subcommand>\nSubcommands: stop")
		return
	}
	switch strings.ToLower(ctx.Args[0]) {
	case "stop":
		ctx.Reply("Bye :wave:")
		ctx.Discord.Close()
		os.Exit(-1)
		break
	case "todo":
		{
			str := readTodo()
			buffer := bytes.NewBufferString("todo list entries:")
			for i, s := range str.Entries {
				buffer.WriteString(fmt.Sprintf("\n`%02d` %s", i+1, s))
			}
			ctx.Reply(buffer.String())
			break
		}
	case "addtodo":
		{
			entry := strings.Join(ctx.Args[1:], " ")
			str := readTodo()
			if str == nil {
				str = &jsonStructure{make([]string, 0)}
			}
			str.add(entry)
			writeTodo(*str)
			ctx.Reply("wrote todo list")
			break
		}
	default:
		ctx.Reply("Invalid subcommand!")
	}
}
