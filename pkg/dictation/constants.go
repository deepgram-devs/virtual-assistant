// Copyright 2023 Deepgram Virtual Assistant contributors. All Rights Reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// SPDX-License-Identifier: MIT

package dictation

import "errors"

// response
const (
	ResponseWhatIsYourQuest string = "To seek the Holy Grail."
)

var (
	// ErrInvalidInput Invalid input received
	ErrInvalidInput = errors.New("Invalid input received")
)
