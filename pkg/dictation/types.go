// Copyright 2023 Deepgram Virtual Assistant contributors. All Rights Reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// SPDX-License-Identifier: MIT

package dictation

import (
	"text/template"
)

/*
	Config
*/
type Config struct {
	Template          string `json:"template,omitempty"`
	SkipServerAuth    bool   `json:"skipServerAuth,omitempty"`
	EmailTo           string `json:"emailTo,omitempty"`
	EmailFrom         string `json:"emailFrom,omitempty"`
	EmailSubject      string `json:"emailSubject,omitempty"`
	EmailSmtpAddr     string `json:"emailSmtpAddr,omitempty"`
	EmailSmtpPort     string `json:"emailSmtpPort,omitempty"`
	EmailSmtpUsername string `json:"emailSmtpUsername,omitempty"`
	EmailSmtpPassword string `json:"emailSmtpPassword,omitempty"`
}

/*
	My Assistant
*/
type MyAssistant struct {
	config *Config

	template *template.Template

	list   []string
	paused bool
}
