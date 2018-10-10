// Copyright 2018 gopcua authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package services

import (
	"testing"
	"time"

	"github.com/wmnsk/gopcua/datatypes"

	"github.com/google/go-cmp/cmp"
)

var createSessionResponseCases = []struct {
	description string
	structured  *CreateSessionResponse
	serialized  []byte
}{
	{
		"normal",
		NewCreateSessionResponse(
			NewResponseHeader(
				time.Date(2018, time.August, 10, 23, 0, 0, 0, time.UTC),
				1, 0, NewNullDiagnosticInfo(), []string{}, NewNullAdditionalHeader(), nil,
			),
			datatypes.NewNumericNodeID(0, 1),
			datatypes.NewOpaqueNodeID(0, []byte{
				0x08, 0x22, 0x87, 0x62, 0xba, 0x81, 0xe1, 0x11,
				0xa6, 0x43, 0xf8, 0x77, 0x7b, 0xc6, 0x2f, 0xc8,
			}),
			6000000, nil, nil,
			NewSignatureData(
				"http://www.w3.org/2000/09/xmldsig#rsa-sha1",
				nil,
			),
			65534,
			NewEndpointDescription(
				"ep-url",
				NewApplicationDescription(
					"app-uri", "prod-uri", "app-name", AppTypeServer,
					"gw-uri", "prof-uri", []string{"discov-uri-1", "discov-uri-2"},
				),
				[]byte{},
				SecModeNone,
				"sec-uri",
				NewUserTokenPolicyArray(
					[]*UserTokenPolicy{
						NewUserTokenPolicy(
							"1", UserTokenAnonymous,
							"issued-token", "issuer-uri", "sec-uri",
						),
						NewUserTokenPolicy(
							"1", UserTokenAnonymous,
							"issued-token", "issuer-uri", "sec-uri",
						),
					},
				),
				"trans-uri",
				0,
			),
			NewEndpointDescription(
				"ep-url",
				NewApplicationDescription(
					"app-uri", "prod-uri", "app-name", AppTypeServer,
					"gw-uri", "prof-uri", []string{"discov-uri-1", "discov-uri-2"},
				),
				[]byte{},
				SecModeNone,
				"sec-uri",
				NewUserTokenPolicyArray(
					[]*UserTokenPolicy{
						NewUserTokenPolicy(
							"1", UserTokenAnonymous,
							"issued-token", "issuer-uri", "sec-uri",
						),
						NewUserTokenPolicy(
							"1", UserTokenAnonymous,
							"issued-token", "issuer-uri", "sec-uri",
						),
					},
				),
				"trans-uri",
				0,
			),
		),
		[]byte{
			// TypeID
			0x01, 0x00, 0xd0, 0x01,
			// Timestamp
			0x00, 0x98, 0x67, 0xdd, 0xfd, 0x30, 0xd4, 0x01,
			// RequestHandle
			0x01, 0x00, 0x00, 0x00,
			// ServiceResult
			0x00, 0x00, 0x00, 0x00,
			// ServiceDiagnostics
			0x00,
			// StringTable
			0x00, 0x00, 0x00, 0x00,
			// AdditionalHeader
			0x00, 0x00, 0x00,
			// SessionID
			0x02, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00,
			// AuthenticationToken
			0x05, 0x00, 0x00, 0x10, 0x00, 0x00, 0x00, 0x08,
			0x22, 0x87, 0x62, 0xba, 0x81, 0xe1, 0x11, 0xa6,
			0x43, 0xf8, 0x77, 0x7b, 0xc6, 0x2f, 0xc8,
			// RevisedSessionTimeout
			0x80, 0x8d, 0x5b, 0x00, 0x00, 0x00, 0x00, 0x00,
			// ServerNonce
			0xff, 0xff, 0xff, 0xff,
			// ServerCertificate
			0xff, 0xff, 0xff, 0xff,
			// ServerEndpoints
			// ArraySize
			0x02, 0x00, 0x00, 0x00,
			// EndpointURL
			0x06, 0x00, 0x00, 0x00, 0x65, 0x70, 0x2d, 0x75, 0x72, 0x6c,
			// Server
			// ApplicationURI
			0x07, 0x00, 0x00, 0x00, 0x61, 0x70, 0x70, 0x2d, 0x75, 0x72, 0x69,
			// ProductURI
			0x08, 0x00, 0x00, 0x00, 0x70, 0x72, 0x6f, 0x64, 0x2d, 0x75, 0x72, 0x69,
			// ApplicationName
			0x02, 0x08, 0x00, 0x00, 0x00, 0x61, 0x70, 0x70, 0x2d, 0x6e, 0x61, 0x6d, 0x65,
			// ApplicationType
			0x00, 0x00, 0x00, 0x00,
			// GatewayServerURI
			0x06, 0x00, 0x00, 0x00, 0x67, 0x77, 0x2d, 0x75, 0x72, 0x69,
			// DiscoveryProfileURI
			0x08, 0x00, 0x00, 0x00, 0x70, 0x72, 0x6f, 0x66, 0x2d, 0x75, 0x72, 0x69,
			// DiscoveryURIs
			0x02, 0x00, 0x00, 0x00,
			0x0c, 0x00, 0x00, 0x00, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x2d, 0x75, 0x72, 0x69, 0x2d, 0x31,
			0x0c, 0x00, 0x00, 0x00, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x2d, 0x75, 0x72, 0x69, 0x2d, 0x32,
			// ServerCertificate
			0xff, 0xff, 0xff, 0xff,
			// MessageSecurityMode
			0x01, 0x00, 0x00, 0x00,
			// SecurityPolicyURI
			0x07, 0x00, 0x00, 0x00, 0x73, 0x65, 0x63, 0x2d, 0x75, 0x72, 0x69,
			// UserIdentityTokens
			0x02, 0x00, 0x00, 0x00,
			0x01, 0x00, 0x00, 0x00, 0x31, 0x00, 0x00, 0x00, 0x00,
			0x0c, 0x00, 0x00, 0x00, 0x69, 0x73, 0x73, 0x75, 0x65, 0x64, 0x2d, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
			0x0a, 0x00, 0x00, 0x00, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x2d, 0x75, 0x72, 0x69,
			0x07, 0x00, 0x00, 0x00, 0x73, 0x65, 0x63, 0x2d, 0x75, 0x72, 0x69,
			0x01, 0x00, 0x00, 0x00, 0x31, 0x00, 0x00, 0x00, 0x00,
			0x0c, 0x00, 0x00, 0x00, 0x69, 0x73, 0x73, 0x75, 0x65, 0x64, 0x2d, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
			0x0a, 0x00, 0x00, 0x00, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x2d, 0x75, 0x72, 0x69,
			0x07, 0x00, 0x00, 0x00, 0x73, 0x65, 0x63, 0x2d, 0x75, 0x72, 0x69,
			// TransportProfileURI
			0x09, 0x00, 0x00, 0x00, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x2d, 0x75, 0x72, 0x69,
			// SecurityLevel
			0x00,
			// EndpointURL
			0x06, 0x00, 0x00, 0x00, 0x65, 0x70, 0x2d, 0x75, 0x72, 0x6c,
			// Server
			// ApplicationURI
			0x07, 0x00, 0x00, 0x00, 0x61, 0x70, 0x70, 0x2d, 0x75, 0x72, 0x69,
			// ProductURI
			0x08, 0x00, 0x00, 0x00, 0x70, 0x72, 0x6f, 0x64, 0x2d, 0x75, 0x72, 0x69,
			// ApplicationName
			0x02, 0x08, 0x00, 0x00, 0x00, 0x61, 0x70, 0x70, 0x2d, 0x6e, 0x61, 0x6d, 0x65,
			// ApplicationType
			0x00, 0x00, 0x00, 0x00,
			// GatewayServerURI
			0x06, 0x00, 0x00, 0x00, 0x67, 0x77, 0x2d, 0x75, 0x72, 0x69,
			// DiscoveryProfileURI
			0x08, 0x00, 0x00, 0x00, 0x70, 0x72, 0x6f, 0x66, 0x2d, 0x75, 0x72, 0x69,
			// DiscoveryURLs
			0x02, 0x00, 0x00, 0x00,
			0x0c, 0x00, 0x00, 0x00, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x2d, 0x75, 0x72, 0x69, 0x2d, 0x31,
			0x0c, 0x00, 0x00, 0x00, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x2d, 0x75, 0x72, 0x69, 0x2d, 0x32,
			// ServerCertificate
			0xff, 0xff, 0xff, 0xff,
			// MessageSecurityMode
			0x01, 0x00, 0x00, 0x00,
			// SecurityPolicyURI
			0x07, 0x00, 0x00, 0x00, 0x73, 0x65, 0x63, 0x2d, 0x75, 0x72, 0x69,
			// UserIdentityTokens
			0x02, 0x00, 0x00, 0x00,
			0x01, 0x00, 0x00, 0x00, 0x31, 0x00, 0x00, 0x00, 0x00,
			0x0c, 0x00, 0x00, 0x00, 0x69, 0x73, 0x73, 0x75, 0x65, 0x64, 0x2d, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
			0x0a, 0x00, 0x00, 0x00, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x2d, 0x75, 0x72, 0x69,
			0x07, 0x00, 0x00, 0x00, 0x73, 0x65, 0x63, 0x2d, 0x75, 0x72, 0x69,
			0x01, 0x00, 0x00, 0x00, 0x31, 0x00, 0x00, 0x00, 0x00,
			0x0c, 0x00, 0x00, 0x00, 0x69, 0x73, 0x73, 0x75, 0x65, 0x64, 0x2d, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
			0x0a, 0x00, 0x00, 0x00, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x2d, 0x75, 0x72, 0x69,
			0x07, 0x00, 0x00, 0x00, 0x73, 0x65, 0x63, 0x2d, 0x75, 0x72, 0x69,
			// TransportProfileURI
			0x09, 0x00, 0x00, 0x00, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x2d, 0x75, 0x72, 0x69,
			// SecurityLevel
			0x00,
			// ServerSoftwareCertificates
			0x00, 0x00, 0x00, 0x00,
			// ServerSignature
			0x2a, 0x00, 0x00, 0x00, 0x68, 0x74, 0x74, 0x70, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x77,
			0x33, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x32, 0x30, 0x30, 0x30, 0x2f, 0x30, 0x39, 0x2f, 0x78, 0x6d,
			0x6c, 0x64, 0x73, 0x69, 0x67, 0x23, 0x72, 0x73, 0x61, 0x2d, 0x73, 0x68, 0x61, 0x31,
			0xff, 0xff, 0xff, 0xff,
			// MaxRequestMessageSize
			0xfe, 0xff, 0x00, 0x00,
		},
	},
}

func TestDecodeCreateSessionResponse(t *testing.T) {
	for _, c := range createSessionResponseCases {
		got, err := DecodeCreateSessionResponse(c.serialized)
		if err != nil {
			t.Fatal(err)
		}

		// need to clear Payload here.
		got.Payload = nil

		if diff := cmp.Diff(got, c.structured, decodeCmpOpt); diff != "" {
			t.Errorf("%s failed\n%s", c.description, diff)
		}
	}
}

func TestSerializeCreateSessionResponse(t *testing.T) {
	for _, c := range createSessionResponseCases {
		got, err := c.structured.Serialize()
		if err != nil {
			t.Fatal(err)
		}

		if diff := cmp.Diff(got, c.serialized); diff != "" {
			t.Errorf("%s failed\n%s", c.description, diff)
		}
	}
}

func TestCreateSessionResponseLen(t *testing.T) {
	for _, c := range createSessionResponseCases {
		got := c.structured.Len()

		if diff := cmp.Diff(got, len(c.serialized)); diff != "" {
			t.Errorf("%s failed\n%s", c.description, diff)
		}
	}
}

func TestCreateSessionResponseServiceType(t *testing.T) {
	for _, c := range createSessionResponseCases {
		if c.structured.ServiceType() != ServiceTypeCreateSessionResponse {
			t.Errorf(
				"ServiceType doesn't match. Want: %d, Got: %d",
				ServiceTypeCreateSessionResponse,
				c.structured.ServiceType(),
			)
		}
	}
}
