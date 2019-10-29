package hdrhistogram_test

import (
	"github.com/filipecosta90/hdrhistogram"
	"reflect"
	"testing"
)

var (
	LOWEST           int64 = 1
	HIGHEST          int64 = 3600 * 1000 * 1000
	SIGNIFICANT            = 3
	TEST_VALUE_LEVEL       = 4
	INTERVAL               = 10000
)

func TestHistogram_Load(t *testing.T) {
	//HDR_LOG_NAME := "hdr.log"
	//min := int64(1)
	//max := int64(10000000)
	//sigfigs := 3
	//encoded := "HISTFAAAAEV42pNpmSzMwMCgyAABTBDKT4GBgdnNYMcCBvsPEBEJISEuATEZMQ4uASkhIR4nrxg9v2lMaxhvMekILGZkKmcCAEf2CsI="
	encoded := "HISTFAAAACl4nJNpmSzMwMBgyAABzFCaEURcm7yEwf4DROA8/4I5jNM7mJgAlWkH9g=="
	_, err := hdrhistogram.Load(encoded)
	if err != nil {
		t.Errorf("TestHistogram_Load() error = %v", err)
		return
	}

}

func Test_DecodeCompressedHeaderFormat(t *testing.T) {
	type args struct {
		decoded []byte
	}
	tests := []struct {
		name    string
		args    args
		want    int32
		want1   int32
		want2   int32
		want3   int32
		want4   int64
		want5   int64
		want6   float64
		wantErr bool
	}{
		{"Test_DecodeCompressedHeaderFormat_1",
			args{[]byte{28, 132, 147, 19, 0, 0, 0, 33, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 78, 32, 0, 0, 3, 70, 48, 184, 160, 0, 63, 240, 0, 0, 0, 0, 0, 0, 24, 18, 18, 10, 16, 22, 28, 22, 8, 10, 16, 26, 18, 18, 12, 66, 74, 92, 46, 78, 150, 2, 172, 1, 218, 2, 44, 16, 163, 1, 2, 119, 2,},},
			478450451, 33, 0, 2, 20000, 3600000000000, 1.0, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2, got3, got4, got5, got6, err := hdrhistogram.DecodeDeCompressedHeaderFormat(tt.args.decoded)
			if (err != nil) != tt.wantErr {
				t.Errorf("decodeCompressedHeaderFormat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("decodeCompressedHeaderFormat() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("decodeCompressedHeaderFormat() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("decodeCompressedHeaderFormat() got2 = %v, want %v", got2, tt.want2)
			}
			if got3 != tt.want3 {
				t.Errorf("decodeCompressedHeaderFormat() got3 = %v, want %v", got3, tt.want3)
			}
			if got4 != tt.want4 {
				t.Errorf("decodeCompressedHeaderFormat() got4 = %v, want %v", got4, tt.want4)
			}
			if got5 != tt.want5 {
				t.Errorf("decodeCompressedHeaderFormat() got5 = %v, want %v", got5, tt.want5)
			}
			if got6 != tt.want6 {
				t.Errorf("decodeCompressedHeaderFormat() got6 = %v, want %v", got6, tt.want6)
			}
		})
	}
}

func TestDecodeCompressedFormat(t *testing.T) {
	type args struct {
		decoded []byte
	}
	tests := []struct {
		name    string
		args    args
		wantRh  *hdrhistogram.Histogram
		wantErr bool
	}{
		{"TestDecodeCompressedFormat_1",
			args{[]byte{28, 132, 147, 20, 0, 0, 0, 69, 120, 218, 147, 105, 153, 44, 204, 192, 192, 160, 200, 0, 1, 76, 16,
				202, 79, 129, 129, 129, 217, 205, 96, 199, 2, 6, 251, 15, 16, 17, 9, 33, 33, 46, 1, 49, 25, 49, 14,
				46, 1, 41, 33, 33, 30, 39, 175, 24, 61, 191, 105, 76, 107, 24, 111, 49, 233, 8, 44, 102, 100, 42,
				103, 2, 0, 71, 246, 10, 194,},},
			nil,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := hdrhistogram.DecodeCompressedFormat(tt.args.decoded, 40)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecodeCompressedFormat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			//if !reflect.DeepEqual(gotRh, tt.wantRh) {
			//	t.Errorf("DecodeCompressedFormat() gotRh = %v, want %v", gotRh, tt.wantRh)
			//}
		})
	}
}

func getTestDefaultHistogram() (rh *hdrhistogram.Histogram) {
	rh = hdrhistogram.New(LOWEST, HIGHEST, SIGNIFICANT)
	// record this value with a count of 10, 000
	rh.RecordValues(1000, 10000)
	rh.RecordValue(100000000)
	return

}

func TestEncodeLEB128_64b9B_variant(t *testing.T) {
	type args struct {
		b []byte
		s int64
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{"12", args{[]byte{}, 12}, []byte{24},},
		{"-60", args{[]byte{}, -60}, []byte{119},},
		{"1", args{[]byte{24, 18, 18, 10, 16, 22, 28, 22, 8, 10, 16, 26, 18, 18, 12, 66, 74,
			92, 46, 78, 150, 2, 172, 1, 218, 2, 44, 16, 163, 1, 2, 119}, 1}, []byte{24, 18, 18, 10, 16, 22, 28, 22, 8, 10, 16, 26, 18, 18, 12, 66, 74,
			92, 46, 78, 150, 2, 172, 1, 218, 2, 44, 16, 163, 1, 2, 119, 2},},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hdrhistogram.EncodeLEB128_64b9B_variant(tt.args.b, tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EncodeLEB128_64b9B_variant() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecodeLEB128_64b9B_variant(t *testing.T) {
	type args struct {
		buffer []byte
	}
	tests := []struct {
		name  string
		args  args
		wantS int64
		wantN uint8
	}{
		{"2", args{[]byte{2}}, 1, 1,},
		{"60", args{[]byte{119}}, -60, 1,},
		{"12", args{[]byte{24, 18, 18, 10, 16, 22, 28, 22, 8, 10, 16, 26, 18, 18, 12, 66, 74,
			92, 46, 78, 150, 2, 172, 1, 218, 2, 44, 16, 163, 1, 2, 119, 2}}, 12, 1,},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotS, gotN := hdrhistogram.DecodeLEB128_64b9B_variant(tt.args.buffer)
			if gotS != tt.wantS {
				t.Errorf("DecodeLEB128_64b9B_variant() gotS = %v, want %v", gotS, tt.wantS)
			}
			if gotN != tt.wantN {
				t.Errorf("DecodeLEB128_64b9B_variant() gotN = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}
