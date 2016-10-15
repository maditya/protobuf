// Protocol Buffers for Go with Gadgets
//
// Copyright (c) 2016, The GoGo Authors. All rights reserved.
// http://github.com/maditya/protobuf
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
//     * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//     * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package types

import (
	math_rand "math/rand"
	"testing"
	"time"

	"github.com/maditya/protobuf/jsonpb"
	"github.com/maditya/protobuf/proto"
)

func TestFullCircleProtoToStd(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	protoMsg := NewPopulatedProtoTypes(popr, true)
	protoData, err := proto.Marshal(protoMsg)
	if err != nil {
		t.Fatal(err)
	}
	stdMsg := &StdTypes{}
	if err := proto.Unmarshal(protoData, stdMsg); err != nil {
		t.Fatal(err)
	}
	stdData, err := proto.Marshal(stdMsg)
	if err != nil {
		t.Fatal(err)
	}
	protoMsgOut := &ProtoTypes{}
	if err := proto.Unmarshal(stdData, protoMsgOut); err != nil {
		t.Fatal(err)
	}
	if !protoMsg.Equal(protoMsgOut) {
		t.Fatalf("want %#v got %#v", protoMsg, protoMsgOut)
	}
}

func TestJsonFullCircleProtoToStd(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	protoMsg := NewPopulatedProtoTypes(popr, true)
	j := &jsonpb.Marshaler{}
	protoData, err := j.MarshalToString(protoMsg)
	if err != nil {
		t.Fatal(err)
	}
	stdMsg := &StdTypes{}
	if err := jsonpb.UnmarshalString(protoData, stdMsg); err != nil {
		t.Fatal(err)
	}
	stdData, err := j.MarshalToString(stdMsg)
	if err != nil {
		t.Fatal(err)
	}
	protoMsgOut := &ProtoTypes{}
	if err := jsonpb.UnmarshalString(stdData, protoMsgOut); err != nil {
		t.Fatal(err)
	}
	if !protoMsg.Equal(protoMsgOut) {
		t.Fatalf("want %#v got %#v", protoMsg, protoMsgOut)
	}
}

func TestFullCircleRepProtoToStd(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	protoMsg := NewPopulatedRepProtoTypes(popr, true)
	protoData, err := proto.Marshal(protoMsg)
	if err != nil {
		t.Fatal(err)
	}
	stdMsg := &RepStdTypes{}
	if err := proto.Unmarshal(protoData, stdMsg); err != nil {
		t.Fatal(err)
	}
	stdData, err := proto.Marshal(stdMsg)
	if err != nil {
		t.Fatal(err)
	}
	protoMsgOut := &RepProtoTypes{}
	if err := proto.Unmarshal(stdData, protoMsgOut); err != nil {
		t.Fatal(err)
	}
	if !protoMsg.Equal(protoMsgOut) {
		t.Fatalf("want %#v got %#v", protoMsg, protoMsgOut)
	}
}

func TestJsonFullCircleRepProtoToStd(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	protoMsg := NewPopulatedRepProtoTypes(popr, true)
	j := &jsonpb.Marshaler{}
	protoData, err := j.MarshalToString(protoMsg)
	if err != nil {
		t.Fatal(err)
	}
	stdMsg := &RepStdTypes{}
	if err := jsonpb.UnmarshalString(protoData, stdMsg); err != nil {
		t.Fatal(err)
	}
	stdData, err := j.MarshalToString(stdMsg)
	if err != nil {
		t.Fatal(err)
	}
	protoMsgOut := &RepProtoTypes{}
	if err := jsonpb.UnmarshalString(stdData, protoMsgOut); err != nil {
		t.Fatal(err)
	}
	if !protoMsg.Equal(protoMsgOut) {
		t.Fatalf("want %#v got %#v", protoMsg, protoMsgOut)
	}
}

func TestFullCircleMapProtoToStd(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	protoMsg := NewPopulatedMapProtoTypes(popr, true)
	protoData, err := proto.Marshal(protoMsg)
	if err != nil {
		t.Fatal(err)
	}
	stdMsg := &MapStdTypes{}
	if err := proto.Unmarshal(protoData, stdMsg); err != nil {
		t.Fatal(err)
	}
	stdData, err := proto.Marshal(stdMsg)
	if err != nil {
		t.Fatal(err)
	}
	protoMsgOut := &MapProtoTypes{}
	if err := proto.Unmarshal(stdData, protoMsgOut); err != nil {
		t.Fatal(err)
	}
	if !protoMsg.Equal(protoMsgOut) {
		t.Fatalf("want %#v got %#v", protoMsg, protoMsgOut)
	}
}

func TestJsonFullCircleMapProtoToStd(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	protoMsg := NewPopulatedMapProtoTypes(popr, true)
	j := &jsonpb.Marshaler{}
	protoData, err := j.MarshalToString(protoMsg)
	if err != nil {
		t.Fatal(err)
	}
	stdMsg := &MapStdTypes{}
	if err := jsonpb.UnmarshalString(protoData, stdMsg); err != nil {
		t.Fatal(err)
	}
	stdData, err := j.MarshalToString(stdMsg)
	if err != nil {
		t.Fatal(err)
	}
	protoMsgOut := &MapProtoTypes{}
	if err := jsonpb.UnmarshalString(stdData, protoMsgOut); err != nil {
		t.Fatal(err)
	}
	if !protoMsg.Equal(protoMsgOut) {
		t.Fatalf("want %#v got %#v", protoMsg, protoMsgOut)
	}
}

func TestFullCircleOneofProtoToStd(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	protoMsg := NewPopulatedOneofProtoTypes(popr, true)
	protoData, err := proto.Marshal(protoMsg)
	if err != nil {
		t.Fatal(err)
	}
	stdMsg := &OneofStdTypes{}
	if err := proto.Unmarshal(protoData, stdMsg); err != nil {
		t.Fatal(err)
	}
	stdData, err := proto.Marshal(stdMsg)
	if err != nil {
		t.Fatal(err)
	}
	protoMsgOut := &OneofProtoTypes{}
	if err := proto.Unmarshal(stdData, protoMsgOut); err != nil {
		t.Fatal(err)
	}
	if !protoMsg.Equal(protoMsgOut) {
		t.Fatalf("want %#v got %#v", protoMsg, protoMsgOut)
	}
}

func TestJsonFullCircleOneofProtoToStd(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	protoMsg := NewPopulatedOneofProtoTypes(popr, true)
	j := &jsonpb.Marshaler{}
	protoData, err := j.MarshalToString(protoMsg)
	if err != nil {
		t.Fatal(err)
	}
	stdMsg := &OneofStdTypes{}
	if err := jsonpb.UnmarshalString(protoData, stdMsg); err != nil {
		t.Fatal(err)
	}
	stdData, err := j.MarshalToString(stdMsg)
	if err != nil {
		t.Fatal(err)
	}
	protoMsgOut := &OneofProtoTypes{}
	if err := jsonpb.UnmarshalString(stdData, protoMsgOut); err != nil {
		t.Fatal(err)
	}
	if !protoMsg.Equal(protoMsgOut) {
		t.Fatalf("want %#v got %#v", protoMsg, protoMsgOut)
	}
}
