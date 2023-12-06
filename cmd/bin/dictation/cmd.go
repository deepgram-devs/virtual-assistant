// Copyright 2023 Deepgram Virtual Assistant contributors. All Rights Reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// SPDX-License-Identifier: MIT

package main

import (
	"bufio"
	"fmt"
	"os"

	assistant "github.com/dvonthenen/open-virtual-assistant/pkg/assistant"
	interfaces "github.com/dvonthenen/open-virtual-assistant/pkg/assistant/interfaces"
	initlib "github.com/dvonthenen/open-virtual-assistant/pkg/init"

	dictation "github.com/deepgram-devs/virtual-assistant/pkg/dictation"
)

func main() {
	/*
		Init
	*/
	initlib.Init(initlib.AssistantInit{
		LogLevel: initlib.LogLevelStandard, // LogLevelStandard / LogLevelFull / LogLevelTrace / LogLevelVerbose
	})

	/*
		Assistant
	*/
	myAssistant, err := dictation.New()
	if err != nil {
		fmt.Printf("dictation.New failed. Err: %v\n", err)
		os.Exit(1)
	}

	var assistImpl interfaces.AssistantImpl
	assistImpl = myAssistant

	assist, err := assistant.New(&assistImpl, &assistant.AssistantOptions{})
	if err != nil {
		fmt.Printf("assistant.New failed. Err: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("\nStarting the Dictation Assistant...\n\n")

	// blocking call
	err = assist.Start()
	if err != nil {
		fmt.Printf("assist.Start failed. Err: %v\n", err)
		os.Exit(1)
	}

	fmt.Print("Press ENTER to exit!\n\n")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()

	// clean up
	assist.Stop()
}
