// Commercial secret, LLC "RevTech". Refer to CONFIDENTIAL file in the root for details

package conformance

import (
	"testing"
	"time"

	"github.com/planetscale/vtprotobuf/testproto/aliases"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestUnmarshalVTBasic(t *testing.T) {
	msg := &TestAllTypesProto3{
		OptionalTimestamp: timestamppb.Now(),
		OptionalDuration:  durationpb.New(time.Second),
		OptionalValue:     structpb.NewStringValue("kek"),
	}
	serializedOrig, err := proto.Marshal(msg)
	require.NoError(t, err)

	got := &TestAllTypesProto3{}
	require.NoError(t, got.UnmarshalVT(serializedOrig))
	require.True(t, proto.Equal(msg, got))
	require.True(t, msg.EqualVT(got))
}

func TestUnmarshalAliasVTBasic(t *testing.T) {
	msg := &TestAllTypesProto3{
		OptionalString: "some-string",
		RepeatedString: []string{"one-string", "two-strings"},
	}
	serializedOrig, err := proto.Marshal(msg)
	require.NoError(t, err)

	got := &TestAllTypesProto3{}
	require.NoError(t, got.UnmarshalAliasVT(serializedOrig))
	require.True(t, proto.Equal(msg, got))
	require.True(t, msg.EqualVT(got))
}

func TestUnmarshalAliasVTFull(t *testing.T) {
	msg := &aliases.AliasedFields{
		Bytes:         []byte("some-bytes"),
		RepeatedBytes: [][]byte{[]byte("one-bytes"), []byte("two-bytes")},
		OptionalBytes: []byte("optional-bytes"),

		String_:        "some-string",
		RepeatedString: []string{"one-string", "two-strings"},
		OptionalString: proto.String("opt-string"),
		MapStringBytes: map[string][]byte{
			"key":         []byte("value"),
			"another-key": []byte("another-value"),
		},
		MapStringString: map[string]string{
			"more-keys": "more-values",
			"foo":       "bar",
		},
		Oneof: &aliases.AliasedFields_OneofString{
			OneofString: "some-string",
		},
		Any: &anypb.Any{
			TypeUrl: "some-url",
			Value:   []byte("more-bytes"),
		},
	}
	serializedOrig, err := proto.Marshal(msg)
	require.NoError(t, err)

	got := &aliases.AliasedFields{}
	require.NoError(t, got.UnmarshalAliasVT(serializedOrig))
	require.True(t, proto.Equal(msg, got))
}
