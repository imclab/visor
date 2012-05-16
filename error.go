// Copyright (c) 2012, SoundCloud Ltd., Alexis Sellier, Alexander Simmerl, Daniel Bornkessel, Tomás Senart
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// Source code and contact info at http://github.com/soundcloud/visor

package visor

import (
	"errors"
)

var (
	ErrKeyConflict   = errors.New("key is already set")
	ErrKeyNotFound   = errors.New("key not found")
	ErrTicketClaimed = errors.New("ticket already claimed")
	ErrUnauthorized  = errors.New("operation is not permitted")
	ErrInvalidState  = errors.New("invalid state")
)

type Error struct{}
