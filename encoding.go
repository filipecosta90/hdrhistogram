package hdrhistogram

import (
	"bytes"
	"compress/zlib"
	"encoding/base64"
	"encoding/binary"
	"errors"
	"fmt"
	"io/ioutil"
)

const (
	V0EncodingCookieBase           = 0x1c849308
	V0CompressedEncodingCookieBase = 0x1c849309
	V1EncodingCookieBase           = 0x1c849301
	V1CompressedEncodingCookieBase = 0x1c849302
	V2EncodingCookieBase           = 0x1c849303
	V2CompressedEncodingCookieBase = 0x1c849304

	encodingCookieBase           = V2EncodingCookieBase
	compressedEncodingCookieBase = V2CompressedEncodingCookieBase
	ENCODING_HEADER_SIZE         = 40
	V0_ENCODING_HEADER_SIZE      = 32

	// LEB128-64b9B + ZigZag require up to 9 bytes per word
	V2maxWordSizeInBytes = 9
)

// Dump returns a snapshot view of the Histogram. This can be later passed to
// Import to construct a new Histogram with the same state.
func Load(encoded string) (rh *Histogram, err error) {

	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return
	}
	Cookie, err := extractCookie(decoded)
	if err != nil {
		return
	}
	headerSize := 0
	if ((getCookieBase(Cookie) == compressedEncodingCookieBase) ||
		(getCookieBase(Cookie) == V1CompressedEncodingCookieBase)) {

		headerSize = ENCODING_HEADER_SIZE

	} else if (getCookieBase(Cookie) != V0CompressedEncodingCookieBase) {

		headerSize = V0_ENCODING_HEADER_SIZE

	} else {
		err = errors.New(fmt.Sprintf("The buffer does not contain a compressed Histogram (no valid cookie found). Got %d want either %d, %d or %d", getCookieBase(Cookie), compressedEncodingCookieBase, V1CompressedEncodingCookieBase, V0CompressedEncodingCookieBase))
		return
	}

	rh, err = DecodeCompressedFormat(decoded, headerSize)

	return
}

func extractCookie(decoded []byte) (Cookie int32, err error) {
	if len(decoded) < 4 {
		err = errors.New(fmt.Sprintf("Cookie byte slice needs to have at least 4 Bytes. Got %d", len(decoded)))
		return
	}
	Cookie = int32(binary.BigEndian.Uint32(decoded[0:4]))
	return
}

func getCookieBase(cookie int32) int32 {
	return (cookie & ^0xf0);
}

func DecodeCompressedFormat(decoded []byte, headerSize int) (rh *Histogram, err error) {
	initialTargetPosition := 0
	compressedContents := decoded[initialTargetPosition+8 : len(decoded)]
	b := bytes.NewReader(compressedContents)
	z, err := zlib.NewReader(b)
	if err != nil {
		return
	}
	defer z.Close()
	decompressedSlice, err := ioutil.ReadAll(z)
	if err != nil {
		return
	}

	decompressedSliceLen := int32(len(decompressedSlice))
	_, _, _, NumberOfSignificantValueDigits, LowestTrackableValue, HighestTrackableValue, _, err := DecodeDeCompressedHeaderFormat(decompressedSlice[0:headerSize])
	if err != nil {
		return
	}
	actualPayloadLen := decompressedSliceLen - int32(headerSize)
	//if PayloadLength != actualPayloadLen {
	//	err = errors.New(fmt.Sprintf("PayloadLength should have the same size of the actual payload. Got %d want %d", actualPayloadLen, PayloadLength))
	//	return
	//}
	rh = New(LowestTrackableValue, HighestTrackableValue, int(NumberOfSignificantValueDigits))
	payload := decompressedSlice[headerSize:]
	var i uint8 = 0
	index := 0
	for int(actualPayloadLen) > int(i) {
		s, n := DecodeLEB128_64b9B_variant(payload[i:])
		i += n
		if s > 0 {
			rh.counts[index] = s
		} else {
			s = -s
			index = int(s) + index
		}
		index++
	}
	return rh, err
}

func getEncodingCookie() int32 {
	return encodingCookieBase | 0x10; // LSBit of wordsize byte indicates TLZE Encoding
}

/**
 * Encode this histogram into a ByteBuffer
 * @param buffer The buffer to encode into
 * @return The number of bytes written to the buffer
 */
func EncodeIntoByteBuffer() []byte {
	getEncodingCookie()
	//
	//buffer.putInt(getEncodingCookie());
	//buffer.putInt(0); // Placeholder for payload length in bytes.
	//buffer.putInt(getNormalizingIndexOffset());
	//buffer.putInt(numberOfSignificantValueDigits);
	//buffer.putLong(lowestDiscernibleValue);
	//buffer.putLong(highestTrackableValue);
	//buffer.putDouble(getIntegerToDoubleValueConversionRatio());
	//
	//int payloadStartPosition = buffer.position();
	//fillBufferFromCountsArray(buffer);
	//buffer.putInt(initialPosition + 4, buffer.position() - payloadStartPosition); // Record the payload length
	//
	//return buffer.position() - initialPosition;
	return []byte{}
}

func DecodeLEB128_64b9B_variant(buffer []byte) (s int64, n uint8) {
	var u uint64 = 0
	l := uint8(len(buffer) & 0xff)
	// The longest LEB128 encoded sequence is 9 byte long (8 0xff's and 1 0x7f)
	// so make sure we won't overflow.
	if l > 9 {
		l = 9
	}
	var i uint8 = 0
	for i = 0; (i < l); i++ {
		u |= uint64(buffer[i]&0x7f) << (7 * i)
		if buffer[i]&0x80 == 0 {
			n = i + 1
			break
		}
	}
	abs := int64(u >> 1)
	if (u % 2) != 0 {
		//^   bitwise XOR
		s = ^abs
	} else {
		s = abs
	}
	return
}

// AppendUleb128 appends v to b using unsigned LEB128 64b9B variant encoding.
func EncodeLEB128_64b9B_variant(b []byte, s int64) []byte {

	var v uint64 = uint64((s << 1) ^ (s >> 63))
	for {
		c := uint8(v & 0x7f)
		v >>= 7

		if v != 0 {
			c |= 0x80
		}
		b = append(b, c)

		if c&0x80 == 0 {
			break
		}
	}
	return b
}

// AppendUleb128 appends v to b using unsigned LEB128 64b9B variant encoding.
func EncodeLEB128_64b9B_variantInt(b []byte, s int32) []byte {

	var v uint32 = uint32((s << 1) ^ (s >> 31))
	for {
		c := uint8(v & 0x7f)
		v >>= 7

		if v != 0 {
			c |= 0x80
		}
		b = append(b, c)

		if c&0x80 == 0 {
			break
		}
	}
	return b
}

func DecodeDeCompressedHeaderFormat(decoded []byte) (Cookie int32, PayloadLength int32, NormalizingIndexOffSet int32, NumberOfSignificantValueDigits int32, LowestTrackableValue int64, HighestTrackableValue int64, IntegerToDoubleConversionRatio float64, err error) {
	Cookie = int32(binary.BigEndian.Uint32(decoded[0:4]))
	PayloadLength = int32(binary.BigEndian.Uint32(decoded[4:8]))
	NormalizingIndexOffSet = int32(binary.BigEndian.Uint32(decoded[8:12]))
	NumberOfSignificantValueDigits = int32(binary.BigEndian.Uint32(decoded[12:16]))

	rbuf := bytes.NewBuffer(decoded[16:32])
	r64 := make([]int64, 2)
	err = binary.Read(rbuf, binary.BigEndian, &r64)
	if err != nil {
		return
	}

	LowestTrackableValue = r64[0]
	HighestTrackableValue = r64[1]

	buf := bytes.NewReader(decoded[32:40])
	err = binary.Read(buf, binary.BigEndian, &IntegerToDoubleConversionRatio)

	return
}

// Dump returns a snapshot view of the Histogram. This can be later passed to
// Import to construct a new Histogram with the same state.
func (h *Histogram) Dump() []byte {

	Cookie := uint32(V2EncodingCookieBase)
	CompressedHeader := []byte{}
	LengthOfCompressedContents := uint32(0)

	buf := make([]byte, 8, 8)
	binary.BigEndian.PutUint32(buf, Cookie)
	binary.BigEndian.PutUint32(buf, LengthOfCompressedContents)
	buf = append(buf, CompressedHeader...)
	return buf
}
