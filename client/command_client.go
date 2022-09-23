package main

import "main/client/agent"

func main() {
	agent := agent.NewAgent()
	agent.Run()
}
