package red

import (
	"bytes"
	"reflect"
	"testing"
)

func TestServerLinkMessage_UnmarshalBinary(t *testing.T) {
	tests := []struct {
		name string
		b    []byte
		slm  ServerLinkMessage
		err  error
	}{
		{
			name: "empty",
			err:  errInvalidPacket,
		},
		{
			name: "short",
			b:    fromHex("00 00 00 00 30 81 9f 30 0d 06 09 2a 86 48 86 f7 0d 01 01 01 05 00 03 81 8d 00 30 81 89 02 81 81 00 bb 49 20 f1 9e 70 a3 07 32 ca a1 63 ce 8d 05 26 82 73 3a 74 59 9d cc c3 83 9c c8 59 60 7e 15 5b 62 8d 53 02 aa f4 81 bf e6 b5 bc 17 88 10 4c d6 dc 6c 83 b9 c2 05 4e ed 89 99 a7 a3 fd 2d 05 d3 0d 60 b3 de 6d 16 3c 9e c8 8c 33 38 b8 3d 39 c1 23 d7 c3 ae e0 59 b6 1a b1 87 d5 b5 30 dc 2b 04 c7 92 6d 92 c4 be bf 21 ae 8a 69 ff 53 1c 41 ff a7 1d 32 8d bb 86 aa c2 50 c4 da 53 f9 24 b0 99 02 03 01 00 01 01 00 00 00 01 00 00 00 b2 00 00"),
			err:  errInvalidPacket,
		},
		{
			name: "uneven caps",
			slm:  ServerLinkMessage{},
			b:    fromHex("00 00 00 00 30 81 9f 30 0d 06 09 2a 86 48 86 f7 0d 01 01 01 05 00 03 81 8d 00 30 81 89 02 81 81 00 bb 49 20 f1 9e 70 a3 07 32 ca a1 63 ce 8d 05 26 82 73 3a 74 59 9d cc c3 83 9c c8 59 60 7e 15 5b 62 8d 53 02 aa f4 81 bf e6 b5 bc 17 88 10 4c d6 dc 6c 83 b9 c2 05 4e ed 89 99 a7 a3 fd 2d 05 d3 0d 60 b3 de 6d 16 3c 9e c8 8c 33 38 b8 3d 39 c1 23 d7 c3 ae e0 59 b6 1a b1 87 d5 b5 30 dc 2b 04 c7 92 6d 92 c4 be bf 21 ae 8a 69 ff 53 1c 41 ff a7 1d 32 8d bb 86 aa c2 50 c4 da 53 f9 24 b0 99 02 03 01 00 01 01 00 00 00 01 00 00 00 b2 00 00 00 0b 00 00"),
			err:  errInvalidPacket,
		},

		{
			name: "no caps",
			slm: ServerLinkMessage{
				Error:               0x0,
				PubKey:              [162]uint8{0x30, 0x81, 0x9f, 0x30, 0xd, 0x6, 0x9, 0x2a, 0x86, 0x48, 0x86, 0xf7, 0xd, 0x1, 0x1, 0x1, 0x5, 0x0, 0x3, 0x81, 0x8d, 0x0, 0x30, 0x81, 0x89, 0x2, 0x81, 0x81, 0x0, 0xbb, 0x49, 0x20, 0xf1, 0x9e, 0x70, 0xa3, 0x7, 0x32, 0xca, 0xa1, 0x63, 0xce, 0x8d, 0x5, 0x26, 0x82, 0x73, 0x3a, 0x74, 0x59, 0x9d, 0xcc, 0xc3, 0x83, 0x9c, 0xc8, 0x59, 0x60, 0x7e, 0x15, 0x5b, 0x62, 0x8d, 0x53, 0x2, 0xaa, 0xf4, 0x81, 0xbf, 0xe6, 0xb5, 0xbc, 0x17, 0x88, 0x10, 0x4c, 0xd6, 0xdc, 0x6c, 0x83, 0xb9, 0xc2, 0x5, 0x4e, 0xed, 0x89, 0x99, 0xa7, 0xa3, 0xfd, 0x2d, 0x5, 0xd3, 0xd, 0x60, 0xb3, 0xde, 0x6d, 0x16, 0x3c, 0x9e, 0xc8, 0x8c, 0x33, 0x38, 0xb8, 0x3d, 0x39, 0xc1, 0x23, 0xd7, 0xc3, 0xae, 0xe0, 0x59, 0xb6, 0x1a, 0xb1, 0x87, 0xd5, 0xb5, 0x30, 0xdc, 0x2b, 0x4, 0xc7, 0x92, 0x6d, 0x92, 0xc4, 0xbe, 0xbf, 0x21, 0xae, 0x8a, 0x69, 0xff, 0x53, 0x1c, 0x41, 0xff, 0xa7, 0x1d, 0x32, 0x8d, 0xbb, 0x86, 0xaa, 0xc2, 0x50, 0xc4, 0xda, 0x53, 0xf9, 0x24, 0xb0, 0x99, 0x2, 0x3, 0x1, 0x0, 0x1},
				CommonCaps:          0x0,
				ChannelCaps:         0x0,
				CapsOffset:          0xb2,
				CommonCapabilities:  nil,
				ChannelCapabilities: nil,
			},
			b: fromHex("00 00 00 00 30 81 9f 30 0d 06 09 2a 86 48 86 f7 0d 01 01 01 05 00 03 81 8d 00 30 81 89 02 81 81 00 bb 49 20 f1 9e 70 a3 07 32 ca a1 63 ce 8d 05 26 82 73 3a 74 59 9d cc c3 83 9c c8 59 60 7e 15 5b 62 8d 53 02 aa f4 81 bf e6 b5 bc 17 88 10 4c d6 dc 6c 83 b9 c2 05 4e ed 89 99 a7 a3 fd 2d 05 d3 0d 60 b3 de 6d 16 3c 9e c8 8c 33 38 b8 3d 39 c1 23 d7 c3 ae e0 59 b6 1a b1 87 d5 b5 30 dc 2b 04 c7 92 6d 92 c4 be bf 21 ae 8a 69 ff 53 1c 41 ff a7 1d 32 8d bb 86 aa c2 50 c4 da 53 f9 24 b0 99 02 03 01 00 01 00 00 00 00 00 00 00 00 b2 00 00 00"),
		},
		{
			name: "wrong caps",
			b:    fromHex("00 00 00 00 30 81 9f 30 0d 06 09 2a 86 48 86 f7 0d 01 01 01 05 00 03 81 8d 00 30 81 89 02 81 81 00 bb 49 20 f1 9e 70 a3 07 32 ca a1 63 ce 8d 05 26 82 73 3a 74 59 9d cc c3 83 9c c8 59 60 7e 15 5b 62 8d 53 02 aa f4 81 bf e6 b5 bc 17 88 10 4c d6 dc 6c 83 b9 c2 05 4e ed 89 99 a7 a3 fd 2d 05 d3 0d 60 b3 de 6d 16 3c 9e c8 8c 33 38 b8 3d 39 c1 23 d7 c3 ae e0 59 b6 1a b1 87 d5 b5 30 dc 2b 04 c7 92 6d 92 c4 be bf 21 ae 8a 69 ff 53 1c 41 ff a7 1d 32 8d bb 86 aa c2 50 c4 da 53 f9 24 b0 99 02 03 01 00 01 01 00 00 00 01 00 00 00 b2 00 00 00"),
			err:  errInvalidPacket,
		},
		{
			name: "ok",
			slm: ServerLinkMessage{
				Error:               0x0,
				PubKey:              [162]uint8{0x30, 0x81, 0x9f, 0x30, 0xd, 0x6, 0x9, 0x2a, 0x86, 0x48, 0x86, 0xf7, 0xd, 0x1, 0x1, 0x1, 0x5, 0x0, 0x3, 0x81, 0x8d, 0x0, 0x30, 0x81, 0x89, 0x2, 0x81, 0x81, 0x0, 0xbb, 0x49, 0x20, 0xf1, 0x9e, 0x70, 0xa3, 0x7, 0x32, 0xca, 0xa1, 0x63, 0xce, 0x8d, 0x5, 0x26, 0x82, 0x73, 0x3a, 0x74, 0x59, 0x9d, 0xcc, 0xc3, 0x83, 0x9c, 0xc8, 0x59, 0x60, 0x7e, 0x15, 0x5b, 0x62, 0x8d, 0x53, 0x2, 0xaa, 0xf4, 0x81, 0xbf, 0xe6, 0xb5, 0xbc, 0x17, 0x88, 0x10, 0x4c, 0xd6, 0xdc, 0x6c, 0x83, 0xb9, 0xc2, 0x5, 0x4e, 0xed, 0x89, 0x99, 0xa7, 0xa3, 0xfd, 0x2d, 0x5, 0xd3, 0xd, 0x60, 0xb3, 0xde, 0x6d, 0x16, 0x3c, 0x9e, 0xc8, 0x8c, 0x33, 0x38, 0xb8, 0x3d, 0x39, 0xc1, 0x23, 0xd7, 0xc3, 0xae, 0xe0, 0x59, 0xb6, 0x1a, 0xb1, 0x87, 0xd5, 0xb5, 0x30, 0xdc, 0x2b, 0x4, 0xc7, 0x92, 0x6d, 0x92, 0xc4, 0xbe, 0xbf, 0x21, 0xae, 0x8a, 0x69, 0xff, 0x53, 0x1c, 0x41, 0xff, 0xa7, 0x1d, 0x32, 0x8d, 0xbb, 0x86, 0xaa, 0xc2, 0x50, 0xc4, 0xda, 0x53, 0xf9, 0x24, 0xb0, 0x99, 0x2, 0x3, 0x1, 0x0, 0x1},
				CommonCaps:          0x1,
				ChannelCaps:         0x1,
				CapsOffset:          0xb2,
				CommonCapabilities:  []uint32{0xb},
				ChannelCapabilities: []uint32{0x9},
			},
			b: fromHex("00 00 00 00 30 81 9f 30 0d 06 09 2a 86 48 86 f7 0d 01 01 01 05 00 03 81 8d 00 30 81 89 02 81 81 00 bb 49 20 f1 9e 70 a3 07 32 ca a1 63 ce 8d 05 26 82 73 3a 74 59 9d cc c3 83 9c c8 59 60 7e 15 5b 62 8d 53 02 aa f4 81 bf e6 b5 bc 17 88 10 4c d6 dc 6c 83 b9 c2 05 4e ed 89 99 a7 a3 fd 2d 05 d3 0d 60 b3 de 6d 16 3c 9e c8 8c 33 38 b8 3d 39 c1 23 d7 c3 ae e0 59 b6 1a b1 87 d5 b5 30 dc 2b 04 c7 92 6d 92 c4 be bf 21 ae 8a 69 ff 53 1c 41 ff a7 1d 32 8d bb 86 aa c2 50 c4 da 53 f9 24 b0 99 02 03 01 00 01 01 00 00 00 01 00 00 00 b2 00 00 00 0b 00 00 00 09 00 00 00"),
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			var slm ServerLinkMessage
			err := (&slm).UnmarshalBinary(testCase.b)

			if want, got := testCase.err, err; want != got {
				t.Fatalf("unexpected error:\n- want: %v\n-  got: %v", want, got)
			}
			if err != nil {
				return
			}

			if want, got := testCase.slm, slm; !reflect.DeepEqual(want, got) {
				t.Fatalf("unexpected Message:\n- want: %#v\n-  got: %#v", want, got)
			}
		})
	}
}

func TestServerLinkMessage_MarshalBinary(t *testing.T) {
	tests := []struct {
		name string
		b    []byte
		slm  ServerLinkMessage
		err  error
	}{
		{
			name: "ok",
			slm: ServerLinkMessage{
				Error:               0x0,
				PubKey:              [162]uint8{0x30, 0x81, 0x9f, 0x30, 0xd, 0x6, 0x9, 0x2a, 0x86, 0x48, 0x86, 0xf7, 0xd, 0x1, 0x1, 0x1, 0x5, 0x0, 0x3, 0x81, 0x8d, 0x0, 0x30, 0x81, 0x89, 0x2, 0x81, 0x81, 0x0, 0xbb, 0x49, 0x20, 0xf1, 0x9e, 0x70, 0xa3, 0x7, 0x32, 0xca, 0xa1, 0x63, 0xce, 0x8d, 0x5, 0x26, 0x82, 0x73, 0x3a, 0x74, 0x59, 0x9d, 0xcc, 0xc3, 0x83, 0x9c, 0xc8, 0x59, 0x60, 0x7e, 0x15, 0x5b, 0x62, 0x8d, 0x53, 0x2, 0xaa, 0xf4, 0x81, 0xbf, 0xe6, 0xb5, 0xbc, 0x17, 0x88, 0x10, 0x4c, 0xd6, 0xdc, 0x6c, 0x83, 0xb9, 0xc2, 0x5, 0x4e, 0xed, 0x89, 0x99, 0xa7, 0xa3, 0xfd, 0x2d, 0x5, 0xd3, 0xd, 0x60, 0xb3, 0xde, 0x6d, 0x16, 0x3c, 0x9e, 0xc8, 0x8c, 0x33, 0x38, 0xb8, 0x3d, 0x39, 0xc1, 0x23, 0xd7, 0xc3, 0xae, 0xe0, 0x59, 0xb6, 0x1a, 0xb1, 0x87, 0xd5, 0xb5, 0x30, 0xdc, 0x2b, 0x4, 0xc7, 0x92, 0x6d, 0x92, 0xc4, 0xbe, 0xbf, 0x21, 0xae, 0x8a, 0x69, 0xff, 0x53, 0x1c, 0x41, 0xff, 0xa7, 0x1d, 0x32, 0x8d, 0xbb, 0x86, 0xaa, 0xc2, 0x50, 0xc4, 0xda, 0x53, 0xf9, 0x24, 0xb0, 0x99, 0x2, 0x3, 0x1, 0x0, 0x1},
				CommonCaps:          0x1,
				ChannelCaps:         0x1,
				CapsOffset:          0xb2,
				CommonCapabilities:  []uint32{0xb},
				ChannelCapabilities: []uint32{0x9},
			},
			b: fromHex("00 00 00 00 30 81 9f 30 0d 06 09 2a 86 48 86 f7 0d 01 01 01 05 00 03 81 8d 00 30 81 89 02 81 81 00 bb 49 20 f1 9e 70 a3 07 32 ca a1 63 ce 8d 05 26 82 73 3a 74 59 9d cc c3 83 9c c8 59 60 7e 15 5b 62 8d 53 02 aa f4 81 bf e6 b5 bc 17 88 10 4c d6 dc 6c 83 b9 c2 05 4e ed 89 99 a7 a3 fd 2d 05 d3 0d 60 b3 de 6d 16 3c 9e c8 8c 33 38 b8 3d 39 c1 23 d7 c3 ae e0 59 b6 1a b1 87 d5 b5 30 dc 2b 04 c7 92 6d 92 c4 be bf 21 ae 8a 69 ff 53 1c 41 ff a7 1d 32 8d bb 86 aa c2 50 c4 da 53 f9 24 b0 99 02 03 01 00 01 01 00 00 00 01 00 00 00 b2 00 00 00 0b 00 00 00 09 00 00 00"),
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			b, err := testCase.slm.MarshalBinary()

			if want, got := testCase.err, err; want != got {
				t.Fatalf("unexpected error:\n- want: %v\n-  got: %v", want, got)
			}
			if err != nil {
				return
			}

			if want, got := testCase.b, b; !bytes.Equal(want, got) {
				t.Fatalf("unexpected Message bytes:\n- want: [%# x]\n-  got: [%# x]", want, got)
			}
		})
	}
}