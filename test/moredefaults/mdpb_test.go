// Code generated by protoc-gen-gogo.
// source: md.proto
// DO NOT EDIT!

/*
Package moredefaults is a generated protocol buffer package.

It is generated from these files:
	md.proto

It has these top-level messages:
	MoreDefaultsB
	MoreDefaultsA
*/
package moredefaults

import testing "testing"
import math_rand "math/rand"
import time "time"
import github_com_maditya_protobuf_proto "github.com/maditya/protobuf/proto"
import github_com_maditya_protobuf_jsonpb "github.com/maditya/protobuf/jsonpb"
import proto "github.com/maditya/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/maditya/protobuf/gogoproto"
import _ "github.com/maditya/protobuf/test/example"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func TestMoreDefaultsBProto(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	p := NewPopulatedMoreDefaultsB(popr, false)
	data, err := github_com_maditya_protobuf_proto.Marshal(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	msg := &MoreDefaultsB{}
	if err := github_com_maditya_protobuf_proto.Unmarshal(data, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	littlefuzz := make([]byte, len(data))
	copy(littlefuzz, data)
	for i := range data {
		data[i] = byte(popr.Intn(256))
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
	if len(littlefuzz) > 0 {
		fuzzamount := 100
		for i := 0; i < fuzzamount; i++ {
			littlefuzz[popr.Intn(len(littlefuzz))] = byte(popr.Intn(256))
			littlefuzz = append(littlefuzz, byte(popr.Intn(256)))
		}
		// shouldn't panic
		_ = github_com_maditya_protobuf_proto.Unmarshal(littlefuzz, msg)
	}
}

func TestMoreDefaultsAProto(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	p := NewPopulatedMoreDefaultsA(popr, false)
	data, err := github_com_maditya_protobuf_proto.Marshal(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	msg := &MoreDefaultsA{}
	if err := github_com_maditya_protobuf_proto.Unmarshal(data, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	littlefuzz := make([]byte, len(data))
	copy(littlefuzz, data)
	for i := range data {
		data[i] = byte(popr.Intn(256))
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
	if len(littlefuzz) > 0 {
		fuzzamount := 100
		for i := 0; i < fuzzamount; i++ {
			littlefuzz[popr.Intn(len(littlefuzz))] = byte(popr.Intn(256))
			littlefuzz = append(littlefuzz, byte(popr.Intn(256)))
		}
		// shouldn't panic
		_ = github_com_maditya_protobuf_proto.Unmarshal(littlefuzz, msg)
	}
}

func TestMoreDefaultsBJSON(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	p := NewPopulatedMoreDefaultsB(popr, true)
	marshaler := github_com_maditya_protobuf_jsonpb.Marshaler{}
	jsondata, err := marshaler.MarshalToString(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	msg := &MoreDefaultsB{}
	err = github_com_maditya_protobuf_jsonpb.UnmarshalString(jsondata, msg)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Json Equal %#v", seed, msg, p)
	}
}
func TestMoreDefaultsAJSON(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	p := NewPopulatedMoreDefaultsA(popr, true)
	marshaler := github_com_maditya_protobuf_jsonpb.Marshaler{}
	jsondata, err := marshaler.MarshalToString(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	msg := &MoreDefaultsA{}
	err = github_com_maditya_protobuf_jsonpb.UnmarshalString(jsondata, msg)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Json Equal %#v", seed, msg, p)
	}
}
func TestMoreDefaultsBProtoText(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	p := NewPopulatedMoreDefaultsB(popr, true)
	data := github_com_maditya_protobuf_proto.MarshalTextString(p)
	msg := &MoreDefaultsB{}
	if err := github_com_maditya_protobuf_proto.UnmarshalText(data, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
}

func TestMoreDefaultsBProtoCompactText(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	p := NewPopulatedMoreDefaultsB(popr, true)
	data := github_com_maditya_protobuf_proto.CompactTextString(p)
	msg := &MoreDefaultsB{}
	if err := github_com_maditya_protobuf_proto.UnmarshalText(data, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
}

func TestMoreDefaultsAProtoText(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	p := NewPopulatedMoreDefaultsA(popr, true)
	data := github_com_maditya_protobuf_proto.MarshalTextString(p)
	msg := &MoreDefaultsA{}
	if err := github_com_maditya_protobuf_proto.UnmarshalText(data, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
}

func TestMoreDefaultsAProtoCompactText(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	p := NewPopulatedMoreDefaultsA(popr, true)
	data := github_com_maditya_protobuf_proto.CompactTextString(p)
	msg := &MoreDefaultsA{}
	if err := github_com_maditya_protobuf_proto.UnmarshalText(data, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
}

//These tests are generated by github.com/maditya/protobuf/plugin/testgen
