// Copyright 2023 Deepgram Virtual Assistant contributors. All Rights Reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// SPDX-License-Identifier: MIT

package dictation

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"text/template"

	interfaces "github.com/dvonthenen/open-virtual-assistant/pkg/speech/interfaces"
	gomail "gopkg.in/mail.v2"
)

// My Assistant
func New() (*MyAssistant, error) {
	// config file
	var smtpConfig string
	if v := os.Getenv("EMAIL_CONFIG_FILE"); v != "" {
		fmt.Printf("EMAIL_CONFIG_FILE found\n")
		smtpConfig = v
	} else {
		fmt.Printf("Using config.json for configuration file\n")
		smtpConfig = "config.json"
	}
	os.Setenv("ASSISTANT_TRANSCRIBER", "deepgram")

	byData, err := os.ReadFile(smtpConfig)
	if err != nil {
		return nil, err
	}
	// fmt.Printf("\n\nbyData:\n%s\n\n", string(byData))

	config := &Config{}
	err = json.Unmarshal(byData, &config)
	if err != nil {
		return nil, err
	}

	// password for env
	var stmpPassword string
	if v := os.Getenv("EMAIL_SMTP_PASSWORD"); v != "" {
		fmt.Printf("EMAIL_SMTP_PASSWORD found\n")
		stmpPassword = v
	} else {
		fmt.Printf("EMAIL_SMTP_PASSWORD not found\n")
		return nil, ErrInvalidInput
	}
	config.EmailSmtpPassword = stmpPassword

	// template
	template, err := template.ParseFiles(config.Template)
	if err != nil {
		return nil, err
	}

	return &MyAssistant{
		config:   config,
		template: template,
		list:     make([]string, 0),
	}, nil
}

func (a *MyAssistant) SetSpeech(s *interfaces.Speech) {
	// Don't need to set... we're not using it
}

func (a *MyAssistant) Response(text string) error {
	tmp := strings.TrimSpace(text)
	if len(tmp) == 0 {
		return nil
	}
	cmd := strings.ToLower(tmp)
	// fmt.Printf("text: %s\n", text)

	if strings.Contains(strings.TrimSpace(cmd), "resume") {
		fmt.Printf("Unpausing...\n")
		a.paused = false
		return nil
	} else if strings.Contains(strings.TrimSpace(cmd), "pause") {
		fmt.Printf("Pausing...\n")
		a.paused = true
		return nil
	} else if strings.Contains(strings.TrimSpace(cmd), "send email") {
		fmt.Println("Sending email....")
		errEmail := a.sendEmail()
		if errEmail != nil {
			fmt.Printf("Failed to send email. Err: %v\n", errEmail)
			return errEmail
		}

		fmt.Printf("Email sent successfully!\n")
		a.list = make([]string, 0)
		return nil
	}

	// we are paused skip!
	if a.paused {
		// fmt.Printf("We are paused, skipping...\n")
		return nil
	}

	// append to string builder
	fmt.Printf("text: %s\n", text)
	a.list = append(a.list, text)

	return nil
}

func (a *MyAssistant) sendEmail() error {
	// convert string port to int
	ismtpPort, err := strconv.Atoi(a.config.EmailSmtpPort)
	if err != nil {
		return err
	}

	// body
	var body bytes.Buffer
	a.template.Execute(&body, struct {
		Transcription []string
		To            string
		From          string
	}{
		Transcription: a.list,
		To:            a.config.EmailTo,
		From:          a.config.EmailFrom,
	})

	// setup and send email
	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", a.config.EmailFrom)

	// Set E-Mail receivers
	m.SetHeader("To", a.config.EmailTo)

	// Set E-Mail subject
	m.SetHeader("Subject", a.config.EmailSubject)

	// Set E-Mail body. You can set plain text or html with text/html
	m.SetBody("text/plain", body.String())

	// Settings for SMTP server
	d := gomail.NewDialer(a.config.EmailSmtpAddr, ismtpPort, a.config.EmailSmtpUsername, a.config.EmailSmtpPassword)

	// skip server auth
	if a.config.SkipServerAuth {
		// TODO: add verification later, pick up from ENV or FILE
		/* #nosec G402 */
		d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
