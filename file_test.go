// Copyright (c) 2012, SoundCloud Ltd., Alexis Sellier, Alexander Simmerl, Daniel Bornkessel, Tomás Senart
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// Source code and contact info at http://github.com/soundcloud/visor

package visor

import (
	"testing"
)

func fileSetup(path string, value interface{}) *File {
	s, err := Dial(DEFAULT_ADDR, "/file-test")
	if err != nil {
		panic(err)
	}

	r, _ := s.conn.Rev()
	err = s.conn.Del("/", r)

	file := &File{Path: path, Value: value, Codec: new(ByteCodec), Snapshot: s.FastForward(-1)}

	return file
}

func TestUpdate(t *testing.T) {
	path := "update-path"
	value := "update-val"

	f := fileSetup(path, value)

	rev, _ := f.conn.Set(path, f.Rev, []byte(value))
	f = f.FastForward(rev)

	f, err := f.Update([]byte(value + "!"))
	if err != nil {
		t.Error(err)
		return
	}

	if string(f.Value.([]byte)) != value+"!" {
		t.Errorf("expected (*File).Value to be update, got %s", string(f.Value.([]byte)))
	}

	val, _, err := f.conn.Get(path, &f.Rev)
	if err != nil {
		t.Error(err)
	}

	if string(val) != value+"!" {
		t.Errorf("expected %s got %s", value+"!", val)
	}
}

func TestFastForward(t *testing.T) {
	path := "ff-path"
	value := "ff-val"

	f := fileSetup(path, value)

	newRev, err := f.conn.Set(path, f.Rev, []byte(value))
	if err != nil {
		t.Error(err)
		return
	}

	_, err = f.conn.Set(path+"-1", newRev, []byte(value))
	if err != nil {
		t.Error(err)
		return
	}

	f = f.FastForward(-1)
	if f.Rev != newRev {
		t.Errorf("expected %d got %d", newRev, f.Rev)
	}
}

func TestUpdateConflict(t *testing.T) {
	path := "conflict-path"
	value := "conflict-val"

	f := fileSetup(path, value)

	rev, _ := f.conn.Set(path, f.Rev, []byte(value))
	f = f.FastForward(rev)

	_, err := f.Update([]byte(value + "!"))
	if err != nil {
		t.Error(err)
		return
	}

	_, err = f.Update([]byte("!"))
	if err == nil {
		t.Error("expected update with old revision to fail")
		return
	}
}

func TestDel(t *testing.T) {
	path := "del-path"
	value := "del-val"

	f := fileSetup(path, value)

	_, err := f.conn.Set(path, f.Rev, []byte{})
	if err != nil {
		t.Error(err)
	}
	exists, _, err := f.conn.Exists(path, nil)
	if !exists {
		t.Error("path wasn't set properly")
		return
	}

	f = f.FastForward(-1)

	err = f.Del()
	if err != nil {
		t.Error(err)
	}

	exists, _, err = f.conn.Exists(path, nil)
	if err != nil {
		t.Error(err)
	}

	if exists {
		t.Error("path wasn't deleted")
	}
}
