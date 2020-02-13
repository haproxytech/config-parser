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

func TestTCPRequestConnectionAccept(t *testing.T) {

	parser := &tcp.Requests{}

	line := strings.TrimSpace("tcp-request connection accept")

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
		returnLine = result[0].Data
	} else {
		returnLine = fmt.Sprintf("%s # %s", result[0].Data, result[0].Comment)
	}

	if line != returnLine {
		t.Errorf(fmt.Sprintf("error: has [%s] expects [%s]", returnLine, line))
	}

}

func TestTCPRequestConnectionAcceptWithCondition(t *testing.T) {

	parser := &tcp.Requests{}

	line := strings.TrimSpace("tcp-request connection accept if !HTTP")

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
		returnLine = result[0].Data
	} else {
		returnLine = fmt.Sprintf("%s # %s", result[0].Data, result[0].Comment)
	}

	if line != returnLine {
		t.Errorf(fmt.Sprintf("error: has [%s] expects [%s]", returnLine, line))
	}

}

func TestTCPRequestConnectionReject(t *testing.T) {

	parser := &tcp.Requests{}

	line := strings.TrimSpace("tcp-request connection reject")

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
		returnLine = result[0].Data
	} else {
		returnLine = fmt.Sprintf("%s # %s", result[0].Data, result[0].Comment)
	}

	if line != returnLine {
		t.Errorf(fmt.Sprintf("error: has [%s] expects [%s]", returnLine, line))
	}

}

func TestTCPRequestConnectionRejectWithCondition(t *testing.T) {

	parser := &tcp.Requests{}

	line := strings.TrimSpace("tcp-request connection reject if !HTTP")

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
		returnLine = result[0].Data
	} else {
		returnLine = fmt.Sprintf("%s # %s", result[0].Data, result[0].Comment)
	}

	if line != returnLine {
		t.Errorf(fmt.Sprintf("error: has [%s] expects [%s]", returnLine, line))
	}

}

func TestTCPRequestConnectionExpectProxy(t *testing.T) {

	parser := &tcp.Requests{}

	line := strings.TrimSpace("tcp-request connection expect-proxy layer4 if { src -f proxies.lst }")

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

func TestTCPRequestConnectionExpectNetscalerCip(t *testing.T) {

	parser := &tcp.Requests{}

	line := strings.TrimSpace("tcp-request connection expect-netscaler-cip layer4")

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

func TestTCPRequestConnectionCapture(t *testing.T) {

	parser := &tcp.Requests{}

	line := strings.TrimSpace("tcp-request connection capture req.payload(0,6) len 6")

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

func TestTCPRequestConnectionTrackSc0(t *testing.T) {

	parser := &tcp.Requests{}

	line := strings.TrimSpace("tcp-request connection track-sc0 src")

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

func TestTCPRequestConnectionTrackSc0WithCondition(t *testing.T) {

	parser := &tcp.Requests{}

	line := strings.TrimSpace("tcp-request connection track-sc0 src if some_check")

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

func TestTCPRequestConnectionTrackSc1(t *testing.T) {

	parser := &tcp.Requests{}

	line := strings.TrimSpace("tcp-request connection track-sc1 src")

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

func TestTCPRequestConnectionTrackSc1WithCondition(t *testing.T) {

	parser := &tcp.Requests{}

	line := strings.TrimSpace("tcp-request connection track-sc1 src if some_check")

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

func TestTCPRequestConnectionTrackSc2(t *testing.T) {

	parser := &tcp.Requests{}

	line := strings.TrimSpace("tcp-request connection track-sc2 src")

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

func TestTCPRequestConnectionTrackSc2WithCondition(t *testing.T) {

	parser := &tcp.Requests{}

	line := strings.TrimSpace("tcp-request connection track-sc2 src if some_check")

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

func TestTCPRequestConnectionScIncGpc0(t *testing.T) {

	parser := &tcp.Requests{}

	line := strings.TrimSpace("tcp-request connection sc-inc-gpc0(2)")

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

func TestTCPRequestConnectionScIncGpc0WithCondition(t *testing.T) {

	parser := &tcp.Requests{}

	line := strings.TrimSpace("tcp-request connection sc-inc-gpc0(2) if is-error")

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

func TestTCPRequestConnectionScIncGpc1(t *testing.T) {

	parser := &tcp.Requests{}

	line := strings.TrimSpace("tcp-request connection sc-inc-gpc1(2)")

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

func TestTCPRequestConnectionScIncGpc1WithCondition(t *testing.T) {

	parser := &tcp.Requests{}

	line := strings.TrimSpace("tcp-request connection sc-inc-gpc1(2) if is-error")

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

func TestTCPRequestConnectionScSetGpt0(t *testing.T) {

	parser := &tcp.Requests{}

	line := strings.TrimSpace("tcp-request connection sc-set-gpt0(0) 1337")

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

func TestTCPRequestConnectionScSetGpt0WithCondition(t *testing.T) {

	parser := &tcp.Requests{}

	line := strings.TrimSpace("tcp-request connection sc-set-gpt0(0) 1337 if exceeds_limit")

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

func TestTCPRequestConnectionSetSrc(t *testing.T) {

	parser := &tcp.Requests{}

	line := strings.TrimSpace("tcp-request connection set-src src,ipmask(24)")

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

func TestTCPRequestConnectionSetSrcWithCondition(t *testing.T) {

	parser := &tcp.Requests{}

	line := strings.TrimSpace("tcp-request connection set-src src,ipmask(24) if some_check")

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

func TestTCPRequestConnectionSetSrcSecond(t *testing.T) {

	parser := &tcp.Requests{}

	line := strings.TrimSpace("tcp-request connection set-src hdr(x-forwarded-for)")

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

func TestTCPRequestConnectionSetSrcSecondWithCondition(t *testing.T) {

	parser := &tcp.Requests{}

	line := strings.TrimSpace("tcp-request connection set-src hdr(x-forwarded-for) if some_check")

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
