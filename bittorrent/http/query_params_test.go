// Copyright 2016 Jimmy Zelinskie
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package http

import (
	"net/url"
	"testing"
)

var (
	baseAddr     = "https://www.subdomain.tracker.com:80/"
	testInfoHash = "01234567890123456789"
	testPeerID   = "-TEST01-6wfG2wk6wWLc"

	ValidAnnounceArguments = []url.Values{
		{"peer_id": {testPeerID}, "port": {"6881"}, "downloaded": {"1234"}, "left": {"4321"}},
		{"peer_id": {testPeerID}, "ip": {"192.168.0.1"}, "port": {"6881"}, "downloaded": {"1234"}, "left": {"4321"}},
		{"peer_id": {testPeerID}, "ip": {"192.168.0.1"}, "port": {"6881"}, "downloaded": {"1234"}, "left": {"4321"}, "numwant": {"28"}},
		{"peer_id": {testPeerID}, "ip": {"192.168.0.1"}, "port": {"6881"}, "downloaded": {"1234"}, "left": {"4321"}, "event": {"stopped"}},
		{"peer_id": {testPeerID}, "ip": {"192.168.0.1"}, "port": {"6881"}, "downloaded": {"1234"}, "left": {"4321"}, "event": {"started"}, "numwant": {"13"}},
		{"peer_id": {testPeerID}, "port": {"6881"}, "downloaded": {"1234"}, "left": {"4321"}, "no_peer_id": {"1"}},
		{"peer_id": {testPeerID}, "port": {"6881"}, "downloaded": {"1234"}, "left": {"4321"}, "compact": {"0"}, "no_peer_id": {"1"}},
		{"peer_id": {testPeerID}, "port": {"6881"}, "downloaded": {"1234"}, "left": {"4321"}, "compact": {"0"}, "no_peer_id": {"1"}, "key": {"peerKey"}},
		{"peer_id": {testPeerID}, "port": {"6881"}, "downloaded": {"1234"}, "left": {"4321"}, "compact": {"0"}, "no_peer_id": {"1"}, "key": {"peerKey"}, "trackerid": {"trackerId"}},
		{"peer_id": {"%3Ckey%3A+0x90%3E"}, "port": {"6881"}, "downloaded": {"1234"}, "left": {"4321"}, "compact": {"0"}, "no_peer_id": {"1"}, "key": {"peerKey"}, "trackerid": {"trackerId"}},
		{"peer_id": {"%3Ckey%3A+0x90%3E"}, "compact": {"1"}},
		{"peer_id": {""}, "compact": {""}},
	}

	InvalidQueries = []string{
		baseAddr + "announce/?" + "info_hash=%0%a",
	}
)

func mapArrayEqual(boxed map[string][]string, unboxed map[string]string) bool {
	if len(boxed) != len(unboxed) {
		return false
	}

	for mapKey, mapVal := range boxed {
		// Always expect box to hold only one element
		if len(mapVal) != 1 || mapVal[0] != unboxed[mapKey] {
			return false
		}
	}

	return true
}

func TestValidQueries(t *testing.T) {
	for parseIndex, parseVal := range ValidAnnounceArguments {
		parsedQueryObj, err := NewQueryParams(baseAddr + "announce/?" + parseVal.Encode())
		if err != nil {
			t.Error(err)
		}

		if !mapArrayEqual(parseVal, parsedQueryObj.params) {
			t.Errorf("Incorrect parse at item %d.\n Expected=%v\n Recieved=%v\n", parseIndex, parseVal, parsedQueryObj.params)
		}
	}
}

func TestInvalidQueries(t *testing.T) {
	for parseIndex, parseStr := range InvalidQueries {
		parsedQueryObj, err := NewQueryParams(parseStr)
		if err == nil {
			t.Error("Should have produced error", parseIndex)
		}

		if parsedQueryObj != nil {
			t.Error("Should be nil after error", parsedQueryObj, parseIndex)
		}
	}
}

func BenchmarkParseQuery(b *testing.B) {
	for bCount := 0; bCount < b.N; bCount++ {
		for parseIndex, parseStr := range ValidAnnounceArguments {
			parsedQueryObj, err := NewQueryParams(baseAddr + "announce/?" + parseStr.Encode())
			if err != nil {
				b.Error(err, parseIndex)
				b.Log(parsedQueryObj)
			}
		}
	}
}

func BenchmarkURLParseQuery(b *testing.B) {
	for bCount := 0; bCount < b.N; bCount++ {
		for parseIndex, parseStr := range ValidAnnounceArguments {
			parsedQueryObj, err := url.ParseQuery(baseAddr + "announce/?" + parseStr.Encode())
			if err != nil {
				b.Error(err, parseIndex)
				b.Log(parsedQueryObj)
			}
		}
	}
}
