/*
Copyright 2019 HAProxy Technologies

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package tests

import (
	"fmt"
	"strings"
	"testing"

	"github.com/haproxytech/config-parser/v2/parsers/tcp"
)

func TestTCPRequestContentAccept(t *testing.T) {

	parser := &tcp.Requests{}

	line := strings.TrimSpace("tcp-request content accept")

	err := ProcessLine(line, parser)

	if err != nil {
		t.Errorf(err.Error())
	}

	result, err := parser.Result()

	if err != nil {
		t.Errorf(err.Error())
	}

	var returnLine string

	if result[0].Comment == "" {
		returnLine = fmt.Sprintf("%s", result[0].Data)
	} else {
		returnLine = fmt.Sprintf("%s # %s", result[0].Data, result[0].Comment)
	}

	if line != returnLine {
		t.Errorf(fmt.Sprintf("error: has [%s] expects [%s]", returnLine, line))
	}

}

func TestTCPRequestContentAcceptWithCondition(t *testing.T) {

	parser := &tcp.Requests{}

	line := strings.TrimSpace("tcp-request content accept if !HTTP")

	err := ProcessLine(line, parser)

	if err != nil {
		t.Errorf(err.Error())
	}

	result, err := parser.Result()

	if err != nil {
		t.Errorf(err.Error())
	}

	var returnLine string

	if result[0].Comment == "" {
		returnLine = fmt.Sprintf("%s", result[0].Data)
	} else {
		returnLine = fmt.Sprintf("%s # %s", result[0].Data, result[0].Comment)
	}

	if line != returnLine {
		t.Errorf(fmt.Sprintf("error: has [%s] expects [%s]", returnLine, line))
	}

}

func TestTCPRequestContentReject(t *testing.T) {

	parser := &tcp.Requests{}

	line := strings.TrimSpace("tcp-request content reject")

	err := ProcessLine(line, parser)

	if err != nil {
		t.Errorf(err.Error())
	}

	result, err := parser.Result()

	if err != nil {
		t.Errorf(err.Error())
	}

	var returnLine string

	if result[0].Comment == "" {
		returnLine = fmt.Sprintf("%s", result[0].Data)
	} else {
		returnLine = fmt.Sprintf("%s # %s", result[0].Data, result[0].Comment)
	}

	if line != returnLine {
		t.Errorf(fmt.Sprintf("error: has [%s] expects [%s]", returnLine, line))
	}

}
func TestTCPRequestContentRejectWithCondition(t *testing.T) {

	parser := &tcp.Requests{}

	line := strings.TrimSpace("tcp-request content reject if !HTTP")

	err := ProcessLine(line, parser)

	if err != nil {
		t.Errorf(err.Error())
	}

	result, err := parser.Result()

	if err != nil {
		t.Errorf(err.Error())
	}

	var returnLine string

	if result[0].Comment == "" {
		returnLine = fmt.Sprintf("%s", result[0].Data)
	} else {
		returnLine = fmt.Sprintf("%s # %s", result[0].Data, result[0].Comment)
	}

	if line != returnLine {
		t.Errorf(fmt.Sprintf("error: has [%s] expects [%s]", returnLine, line))
	}

}

func TestTCPRequestContentCapture(t *testing.T) {

	parser := &tcp.Requests{}

	line := strings.TrimSpace("tcp-request content capture req.payload(0,6) len 6")

	err := ProcessLine(line, parser)

	if err != nil {
		t.Errorf(err.Error())
	}

	result, err := parser.Result()

	if err != nil {
		t.Errorf(err.Error())
	}

	var returnLine string

	if result[0].Comment == "" {
		returnLine = fmt.Sprintf("%s", result[0].Data)
	} else {
		returnLine = fmt.Sprintf("%s # %s", result[0].Data, result[0].Comment)
	}

	if line != returnLine {
		t.Errorf(fmt.Sprintf("error: has [%s] expects [%s]", returnLine, line))
	}

}

func TestTCPRequestContentCaptureWithCondition(t *testing.T) {

	parser := &tcp.Requests{}

	line := strings.TrimSpace("tcp-request content capture req.payload(0,6) len 6 if !HTTP")

	err := ProcessLine(line, parser)

	if err != nil {
		t.Errorf(err.Error())
	}

	result, err := parser.Result()

	if err != nil {
		t.Errorf(err.Error())
	}

	var returnLine string

	if result[0].Comment == "" {
		returnLine = fmt.Sprintf("%s", result[0].Data)
	} else {
		returnLine = fmt.Sprintf("%s # %s", result[0].Data, result[0].Comment)
	}

	if line != returnLine {
		t.Errorf(fmt.Sprintf("error: has [%s] expects [%s]", returnLine, line))
	}

}

func TestTCPRequestContentSetPriorityClass(t *testing.T) {

	parser := &tcp.Requests{}

	line := strings.TrimSpace("tcp-request content set-priority-class int(1)")

	err := ProcessLine(line, parser)

	if err != nil {
		t.Errorf(err.Error())
	}

	result, err := parser.Result()

	if err != nil {
		t.Errorf(err.Error())
	}

	var returnLine string

	if result[0].Comment == "" {
		returnLine = fmt.Sprintf("%s", result[0].Data)
	} else {
		returnLine = fmt.Sprintf("%s # %s", result[0].Data, result[0].Comment)
	}

	if line != returnLine {
		t.Errorf(fmt.Sprintf("error: has [%s] expects [%s]", returnLine, line))
	}

}

func TestTCPRequestContentSetPriorityClassWithConditions(t *testing.T) {

	parser := &tcp.Requests{}

	line := strings.TrimSpace("tcp-request content set-priority-class int(1) if some_check")

	err := ProcessLine(line, parser)

	if err != nil {
		t.Errorf(err.Error())
	}

	result, err := parser.Result()

	if err != nil {
		t.Errorf(err.Error())
	}

	var returnLine string

	if result[0].Comment == "" {
		returnLine = fmt.Sprintf("%s", result[0].Data)
	} else {
		returnLine = fmt.Sprintf("%s # %s", result[0].Data, result[0].Comment)
	}

	if line != returnLine {
		t.Errorf(fmt.Sprintf("error: has [%s] expects [%s]", returnLine, line))
	}

}

func TestTCPRequestContentSetPriorityOffset(t *testing.T) {

	parser := &tcp.Requests{}

	line := strings.TrimSpace("tcp-request content set-priority-offset int(10)")

	err := ProcessLine(line, parser)

	if err != nil {
		t.Errorf(err.Error())
	}

	result, err := parser.Result()

	if err != nil {
		t.Errorf(err.Error())
	}

	var returnLine string

	if result[0].Comment == "" {
		returnLine = fmt.Sprintf("%s", result[0].Data)
	} else {
		returnLine = fmt.Sprintf("%s # %s", result[0].Data, result[0].Comment)
	}

	if line != returnLine {
		t.Errorf(fmt.Sprintf("error: has [%s] expects [%s]", returnLine, line))
	}

}

func TestTCPRequestContentSetPriorityOffsetWithConditions(t *testing.T) {

	parser := &tcp.Requests{}

	line := strings.TrimSpace("tcp-request content set-priority-offset int(10) if some_check")

	err := ProcessLine(line, parser)

	if err != nil {
		t.Errorf(err.Error())
	}

	result, err := parser.Result()

	if err != nil {
		t.Errorf(err.Error())
	}

	var returnLine string

	if result[0].Comment == "" {
		returnLine = fmt.Sprintf("%s", result[0].Data)
	} else {
		returnLine = fmt.Sprintf("%s # %s", result[0].Data, result[0].Comment)
	}

	if line != returnLine {
		t.Errorf(fmt.Sprintf("error: has [%s] expects [%s]", returnLine, line))
	}

}

func TestTCPRequestContentTrackSc0(t *testing.T) {

	parser := &tcp.Requests{}

	line := strings.TrimSpace("tcp-request content track-sc0 src")

	err := ProcessLine(line, parser)

	if err != nil {
		t.Errorf(err.Error())
	}

	result, err := parser.Result()

	if err != nil {
		t.Errorf(err.Error())
	}

	var returnLine string

	if result[0].Comment == "" {
		returnLine = fmt.Sprintf("%s", result[0].Data)
	} else {
		returnLine = fmt.Sprintf("%s # %s", result[0].Data, result[0].Comment)
	}

	if line != returnLine {
		t.Errorf(fmt.Sprintf("error: has [%s] expects [%s]", returnLine, line))
	}

}

func TestTCPRequestContentTrackSc0WithCondition(t *testing.T) {

	parser := &tcp.Requests{}

	line := strings.TrimSpace("tcp-request content track-sc0 src if some_check")

	err := ProcessLine(line, parser)

	if err != nil {
		t.Errorf(err.Error())
	}

	result, err := parser.Result()

	if err != nil {
		t.Errorf(err.Error())
	}

	var returnLine string

	if result[0].Comment == "" {
		returnLine = fmt.Sprintf("%s", result[0].Data)
	} else {
		returnLine = fmt.Sprintf("%s # %s", result[0].Data, result[0].Comment)
	}

	if line != returnLine {
		t.Errorf(fmt.Sprintf("error: has [%s] expects [%s]", returnLine, line))
	}

}

func TestTCPRequestContentTrackSc1(t *testing.T) {

	parser := &tcp.Requests{}

	line := strings.TrimSpace("tcp-request content track-sc1 src")

	err := ProcessLine(line, parser)

	if err != nil {
		t.Errorf(err.Error())
	}

	result, err := parser.Result()

	if err != nil {
		t.Errorf(err.Error())
	}

	var returnLine string

	if result[0].Comment == "" {
		returnLine = fmt.Sprintf("%s", result[0].Data)
	} else {
		returnLine = fmt.Sprintf("%s # %s", result[0].Data, result[0].Comment)
	}

	if line != returnLine {
		t.Errorf(fmt.Sprintf("error: has [%s] expects [%s]", returnLine, line))
	}

}

func TestTCPRequestContentTrackSc1WithCondition(t *testing.T) {

	parser := &tcp.Requests{}

	line := strings.TrimSpace("tcp-request content track-sc1 src if some_check")

	err := ProcessLine(line, parser)

	if err != nil {
		t.Errorf(err.Error())
	}

	result, err := parser.Result()

	if err != nil {
		t.Errorf(err.Error())
	}

	var returnLine string

	if result[0].Comment == "" {
		returnLine = fmt.Sprintf("%s", result[0].Data)
	} else {
		returnLine = fmt.Sprintf("%s # %s", result[0].Data, result[0].Comment)
	}

	if line != returnLine {
		t.Errorf(fmt.Sprintf("error: has [%s] expects [%s]", returnLine, line))
	}

}

func TestTCPRequestContentTrackSc2(t *testing.T) {

	parser := &tcp.Requests{}

	line := strings.TrimSpace("tcp-request content track-sc2 src")

	err := ProcessLine(line, parser)

	if err != nil {
		t.Errorf(err.Error())
	}

	result, err := parser.Result()

	if err != nil {
		t.Errorf(err.Error())
	}

	var returnLine string

	if result[0].Comment == "" {
		returnLine = fmt.Sprintf("%s", result[0].Data)
	} else {
		returnLine = fmt.Sprintf("%s # %s", result[0].Data, result[0].Comment)
	}

	if line != returnLine {
		t.Errorf(fmt.Sprintf("error: has [%s] expects [%s]", returnLine, line))
	}

}

func TestTCPRequestContentTrackSc2WithCondition(t *testing.T) {

	parser := &tcp.Requests{}

	line := strings.TrimSpace("tcp-request content track-sc2 src if some_check")

	err := ProcessLine(line, parser)

	if err != nil {
		t.Errorf(err.Error())
	}

	result, err := parser.Result()

	if err != nil {
		t.Errorf(err.Error())
	}

	var returnLine string

	if result[0].Comment == "" {
		returnLine = fmt.Sprintf("%s", result[0].Data)
	} else {
		returnLine = fmt.Sprintf("%s # %s", result[0].Data, result[0].Comment)
	}

	if line != returnLine {
		t.Errorf(fmt.Sprintf("error: has [%s] expects [%s]", returnLine, line))
	}

}

func TestTCPRequestContentSetDst(t *testing.T) {

	parser := &tcp.Requests{}

	line := strings.TrimSpace("tcp-request content set-dst ipv4(10.0.0.1)")

	err := ProcessLine(line, parser)

	if err != nil {
		t.Errorf(err.Error())
	}

	result, err := parser.Result()

	if err != nil {
		t.Errorf(err.Error())
	}

	var returnLine string

	if result[0].Comment == "" {
		returnLine = fmt.Sprintf("%s", result[0].Data)
	} else {
		returnLine = fmt.Sprintf("%s # %s", result[0].Data, result[0].Comment)
	}

	if line != returnLine {
		t.Errorf(fmt.Sprintf("error: has [%s] expects [%s]", returnLine, line))
	}

}

func TestTCPRequestContentSilentDrop(t *testing.T) {

	parser := &tcp.Requests{}

	line := strings.TrimSpace("tcp-request content silent-drop")

	err := ProcessLine(line, parser)

	if err != nil {
		t.Errorf(err.Error())
	}

	result, err := parser.Result()

	if err != nil {
		t.Errorf(err.Error())
	}

	var returnLine string

	if result[0].Comment == "" {
		returnLine = fmt.Sprintf("%s", result[0].Data)
	} else {
		returnLine = fmt.Sprintf("%s # %s", result[0].Data, result[0].Comment)
	}

	if line != returnLine {
		t.Errorf(fmt.Sprintf("error: has [%s] expects [%s]", returnLine, line))
	}

}

func TestTCPRequestContentSilentDropWithCondition(t *testing.T) {

	parser := &tcp.Requests{}

	line := strings.TrimSpace("tcp-request content silent-drop if !HTTP")

	err := ProcessLine(line, parser)

	if err != nil {
		t.Errorf(err.Error())
	}

	result, err := parser.Result()

	if err != nil {
		t.Errorf(err.Error())
	}

	var returnLine string

	if result[0].Comment == "" {
		returnLine = fmt.Sprintf("%s", result[0].Data)
	} else {
		returnLine = fmt.Sprintf("%s # %s", result[0].Data, result[0].Comment)
	}

	if line != returnLine {
		t.Errorf(fmt.Sprintf("error: has [%s] expects [%s]", returnLine, line))
	}

}

func TestTCPRequestContentSendSpoeGroup(t *testing.T) {

	parser := &tcp.Requests{}

	line := strings.TrimSpace("tcp-request content send-spoe-group engine group")

	err := ProcessLine(line, parser)

	if err != nil {
		t.Errorf(err.Error())
	}

	result, err := parser.Result()

	if err != nil {
		t.Errorf(err.Error())
	}

	var returnLine string

	if result[0].Comment == "" {
		returnLine = fmt.Sprintf("%s", result[0].Data)
	} else {
		returnLine = fmt.Sprintf("%s # %s", result[0].Data, result[0].Comment)
	}

	if line != returnLine {
		t.Errorf(fmt.Sprintf("error: has [%s] expects [%s]", returnLine, line))
	}

}

func TestTCPRequestContentUseService(t *testing.T) {

	// TODO: tcp-request content use-service lua.deny { src -f /etc/haproxy/blacklist.lst }

	parser := &tcp.Requests{}

	line := strings.TrimSpace("tcp-request content use-service lua.deny")

	err := ProcessLine(line, parser)

	if err != nil {
		t.Errorf(err.Error())
	}

	result, err := parser.Result()

	if err != nil {
		t.Errorf(err.Error())
	}

	var returnLine string

	if result[0].Comment == "" {
		returnLine = fmt.Sprintf("%s", result[0].Data)
	} else {
		returnLine = fmt.Sprintf("%s # %s", result[0].Data, result[0].Comment)
	}

	if line != returnLine {
		t.Errorf(fmt.Sprintf("error: has [%s] expects [%s]", returnLine, line))
	}

}

func TestTCPRequestContentUseServiceWithCondition(t *testing.T) {

	// TODO: tcp-request content use-service lua.deny { src -f /etc/haproxy/blacklist.lst }

	parser := &tcp.Requests{}

	line := strings.TrimSpace("tcp-request content use-service lua.deny if !HTTP")

	err := ProcessLine(line, parser)

	if err != nil {
		t.Errorf(err.Error())
	}

	result, err := parser.Result()

	if err != nil {
		t.Errorf(err.Error())
	}

	var returnLine string

	if result[0].Comment == "" {
		returnLine = fmt.Sprintf("%s", result[0].Data)
	} else {
		returnLine = fmt.Sprintf("%s # %s", result[0].Data, result[0].Comment)
	}

	if line != returnLine {
		t.Errorf(fmt.Sprintf("error: has [%s] expects [%s]", returnLine, line))
	}

}
