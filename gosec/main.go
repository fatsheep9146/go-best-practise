package main

import (
	"log"
	"os"
	"os/exec"
)

func RunCmdWithConstString() {
	cmd := exec.Command("ls")
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Waiting for command to finish...")
	err = cmd.Wait()
}

func RunCmdWithVariableWithConstString() {
	cmdstr := "ls"
	cmd := exec.Command(cmdstr)
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Waiting for command to finish...")
	err = cmd.Wait()
}

// func RunCmdWithInputArgs() {
// 	cmd := exec.Command(os.Args[0])
// 	err := cmd.Start()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	log.Printf("Waiting for command to finish...")
// 	err = cmd.Wait()
// }

func RunCmdWithShortDeclartionsWithInputArgs() {
	cmdstr := os.Args[0]
	cmd := exec.Command(cmdstr)
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Waiting for command to finish...")
	err = cmd.Wait()
}

func RunCmdWithVariableDeclartionWithInputArgs() {
	var cmdstr = os.Args[0]
	cmd := exec.Command(cmdstr)
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Waiting for command to finish...")
	err = cmd.Wait()
}

func validate(cmd string) string {
	return cmd
}

func RunCmdWithVariableValidated() {
	var cmdarg = os.Args[0]
	cmdstr := validate(cmdarg)
	cmd := exec.Command(cmdstr)
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Waiting for command to finish...")
	err = cmd.Wait()
}

func main() {
	RunCmdWithConstString()
	RunCmdWithVariableWithConstString()
	// RunCmdWithInputArgs()
	RunCmdWithShortDeclartionsWithInputArgs()
	RunCmdWithVariableDeclartionWithInputArgs()
	RunCmdWithVariableValidated()
}
