package main

import (
	"encoding/binary"
	"errors"
	"fmt"
	"math"
	"time"

	"github.com/gofrs/uuid"
)

// ByteReader, byte dizilerini okumak için bir yapı sağlar.
type ByteReader struct {
	Data   []byte // Okunacak byte dizisi
	Offset uint   // Mevcut okuma pozisyonu
}

// NewByteReader, yeni bir ByteReader örneği oluşturur.
func NewByteReader(data []byte) *ByteReader {
	return &ByteReader{
		Data:   data,
		Offset: 0,
	}
}

// u8 fonksiyonu, ByteReader'dan bir unsigned 8-bit (byte) değeri okur.
// Okuma başarılı olursa byte değerini ve nil (hata yok) döndürür.
// Eğer offset verinin uzunluğundan büyük veya eşitse, 0 ve bir hata döndürür.
func (b *ByteReader) u8() (byte, error) {
	if b.Offset >= uint(len(b.Data)) {
		return 0, errors.New("u8: offset verinin uzunluğundan büyük veya eşit")
	}

	val := b.Data[b.Offset]
	b.Offset += 1

	return val, nil
}

// u16 fonksiyonu, ByteReader'dan bir unsigned 16-bit (uint16) değeri okur.
// Okuma başarılı olursa uint16 değerini ve nil döndürür.
// Yeterli byte yoksa hata döndürür.
func (b *ByteReader) u16() (uint16, error) {
	if b.Offset+2 > uint(len(b.Data)) {
		return 0, errors.New("u16: okumak için yeterli byte yok")
	}

	val := binary.LittleEndian.Uint16(b.Data[b.Offset : b.Offset+2])
	b.Offset += 2

	return val, nil
}

// u32 fonksiyonu, ByteReader'dan bir unsigned 32-bit (uint32) değeri okur.
// Okuma başarılı olursa uint32 değerini ve nil döndürür.
// Yeterli byte yoksa hata döndürür.
func (b *ByteReader) u32() (uint32, error) {
	if b.Offset+4 > uint(len(b.Data)) {
		return 0, errors.New("u32: okumak için yeterli byte yok")
	}

	val := binary.LittleEndian.Uint32(b.Data[b.Offset : b.Offset+4])
	b.Offset += 4

	return val, nil
}

// u64 fonksiyonu, ByteReader'dan bir unsigned 64-bit (uint64) değeri okur.
// Okuma başarılı olursa uint64 değerini ve nil döndürür.
// Yeterli byte yoksa hata döndürür.
func (b *ByteReader) u64() (uint64, error) {
	if b.Offset+8 > uint(len(b.Data)) {
		return 0, errors.New("u64: okumak için yeterli byte yok")
	}

	val := binary.LittleEndian.Uint64(b.Data[b.Offset : b.Offset+8])
	b.Offset += 8

	return val, nil
}

// i8 fonksiyonu, ByteReader'dan bir signed 8-bit (int8) değeri okur.
// Okuma başarılı olursa int8 değerini ve nil döndürür.
// Yeterli byte yoksa hata döndürür.
func (b *ByteReader) i8() (int8, error) {
	if b.Offset >= uint(len(b.Data)) {
		return 0, errors.New("i8: okumak için yeterli byte yok")
	}

	val := int8(b.Data[b.Offset])
	b.Offset += 1

	return val, nil
}

// i16 fonksiyonu, ByteReader'dan bir signed 16-bit (int16) değeri okur.
// Okuma başarılı olursa int16 değerini ve nil döndürür.
// Yeterli byte yoksa hata döndürür.
func (b *ByteReader) i16() (int16, error) {
	if b.Offset+2 > uint(len(b.Data)) {
		return 0, errors.New("i16: okumak için yeterli byte yok")
	}

	val := int16(binary.LittleEndian.Uint16(b.Data[b.Offset : b.Offset+2]))
	b.Offset += 2

	return val, nil
}

// i32 fonksiyonu, ByteReader'dan bir signed 32-bit (int32) değeri okur.
// Okuma başarılı olursa int32 değerini ve nil döndürür.
// Yeterli byte yoksa hata döndürür.
func (b *ByteReader) i32() (int32, error) {
	if b.Offset+4 > uint(len(b.Data)) {
		return 0, errors.New("i32: okumak için yeterli byte yok")
	}

	val := int32(binary.LittleEndian.Uint32(b.Data[b.Offset : b.Offset+4]))
	b.Offset += 4

	return val, nil
}

// i64 fonksiyonu, ByteReader'dan bir signed 64-bit (int64) değeri okur.
// Okuma başarılı olursa int64 değerini ve nil döndürür.
// Yeterli byte yoksa hata döndürür.
func (b *ByteReader) i64() (int64, error) {
	if b.Offset+8 > uint(len(b.Data)) {
		return 0, errors.New("i64: okumak için yeterli byte yok")
	}

	val := int64(binary.LittleEndian.Uint64(b.Data[b.Offset : b.Offset+8]))
	b.Offset += 8

	return val, nil
}

// Rbool fonksiyonu, ByteReader'dan bir boolean değeri okur.
// Bir byte okur ve 0 değilse true, 0 ise false döndürür.
// Byte okunamıyorsa hata döndürür.
func (b *ByteReader) Rbool() (bool, error) {
	val, err := b.u8()
	if err != nil {
		return false, fmt.Errorf("Rbool: %w", err)
	}
	return val != 0, nil
}

// Rbits fonksiyonu, ByteReader'dan bir byte okur ve 8 boolean dilimi olarak döndürür.
// Byte okunamıyorsa hata döndürür.
func (b *ByteReader) Rbits() ([]bool, error) {
	val, err := b.u8()
	if err != nil {
		return nil, fmt.Errorf("Rbits: %w", err)
	}
	result := make([]bool, 8)
	for i := 0; i < 8; i++ { // Range yerine sabit 8 kullanıldı
		result[i] = (val&(1<<uint(i)) != 0) // i'yi uint'e dönüştür
	}
	return result, nil
}

// Rbytes fonksiyonu, belirtilen sayıda byte okur ve bir dilim olarak döndürür.
// Yeterli byte yoksa hata döndürür.
func (b *ByteReader) Rbytes(count uint) ([]byte, error) {
	if b.Offset+count > uint(len(b.Data)) {
		return nil, errors.New("Rbytes: okumak için yeterli byte yok")
	}

	// Dilimi kopyalayarak dışarıya sızmasını engellemek daha güvenli olabilir
	slice := make([]byte, count)
	copy(slice, b.Data[b.Offset:b.Offset+count])
	b.Offset += count

	return slice, nil
}

// PeekBytes fonksiyonu, belirtilen sayıda byte'ı okuyucunun offset'ini değiştirmeden döndürür.
// Yeterli byte yoksa hata döndürür.
func (b *ByteReader) PeekBytes(count uint) ([]byte, error) {
	if b.Offset+count > uint(len(b.Data)) {
		return nil, errors.New("PeekBytes: okumak için yeterli byte yok")
	}

	// Dilimi kopyalayarak dışarıya sızmasını engellemek daha güvenli olabilir
	slice := make([]byte, count)
	copy(slice, b.Data[b.Offset:b.Offset+count])
	return slice, nil
}

// ReadUntil fonksiyonu, belirtilen adrese kadar olan byte'ları okur ve bir dilim olarak döndürür.
// Offset'i okunan son adrese ayarlar.
// Geçersiz adres veya okunacak bir şey yoksa boş dilim ve nil hata döndürür.
func (b *ByteReader) ReadUntil(address uint) ([]byte, error) {
	end := min(address, uint(len(b.Data)))
	if b.Offset > end { // Offset zaten bitiş adresinden büyükse
		return nil, errors.New("ReadUntil: başlangıç offset'i bitiş adresinden büyük")
	}
	if b.Offset == end { // Okunacak bir şey yoksa
		return make([]byte, 0), nil
	}

	slice := make([]byte, end-b.Offset)
	copy(slice, b.Data[b.Offset:end])
	b.Offset = end

	return slice, nil
}

// offset fonksiyonu, mevcut okuma offset'ini döndürür.
func (b *ByteReader) offset() uint {
	return b.Offset
}
func (b *ByteReader) setOffset(offset uint) {
	b.Offset = offset
}

// seek fonksiyonu, okuyucunun offset'ini belirtilen değere ayarlar.
// Belirtilen offset verinin uzunluğundan büyükse hata döndürür.
func (b *ByteReader) seek(offset uint) error {
	if offset > uint(len(b.Data)) {
		return errors.New("seek: belirtilen offset verinin uzunluğundan büyük")
	}
	b.Offset = offset
	return nil
}

// uleb128 fonksiyonu, ByteReader'dan bir ULEB128 kodlu uint64 değeri okur.
// Okuma sırasında hata oluşursa 0 ve hata döndürür.
func (b *ByteReader) uleb128() (uint64, error) {
	val := uint64(0)
	shift := 0

	for {
		_byte, err := b.u8()
		if err != nil {
			return 0, fmt.Errorf("uleb128: %w", err)
		}

		val |= (uint64(_byte&0x7F) << shift)

		if (_byte & 0x80) == 0 {
			break
		}

		shift += 7
		if shift >= 64 { // Sonsuz döngüyü önlemek ve taşmayı yakalamak için
			return 0, errors.New("uleb128: çok uzun ULEB128 değeri (taşma olasılığı)")
		}
	}

	return val, nil
}

// string fonksiyonu, ByteReader'dan bir string okur.
// Eğer size nil ise, string boyutu ULEB128 olarak okunur.
// String byte'ları okunamıyorsa hata döndürür.
func (b *ByteReader) string(size *uint) (string, error) {
	var _size uint
	var err error

	if size == nil {
		ulebSize, ulebErr := b.uleb128()
		if ulebErr != nil {
			return "", fmt.Errorf("string: boyut okunamadı: %w", ulebErr)
		}
		_size = uint(ulebSize)
	} else {
		_size = *size
	}

	_bytes, err := b.Rbytes(_size)
	if err != nil {
		return "", fmt.Errorf("string: byte'lar okunamadı: %w", err)
	}

	return string(_bytes), nil
}

// uuid fonksiyonu, ByteReader'dan 16 byte okur ve bir UUID string'i olarak döndürür.
// Yeterli byte yoksa veya UUID formatı geçersizse hata döndürür.
func (b *ByteReader) uuid() (string, error) {
	_bytes, err := b.Rbytes(16)
	if err != nil {
		return "", fmt.Errorf("uuid: byte'lar okunamadı: %w", err)
	}

	myUUID, err := uuid.FromBytes(_bytes)
	if err != nil {
		return "", fmt.Errorf("uuid: UUID dönüştürme hatası: %w", err)
	}

	return myUUID.String(), nil
}

// datetime fonksiyonu, ByteReader'dan bir 64-bit zaman değeri okur ve formatlanmış string olarak döndürür.
// Zaman değeri okunamıyorsa hata döndürür.
func (b *ByteReader) datetime() (time.Time, error) {
	raw, err := b.u64()
	if err != nil {
		return time.Time{}, fmt.Errorf("datetime: zaman değeri okunamadı: %w", err)
	}

	ticks := raw & 0x3FFFFFFFFFFFFFFF

	unixOffset := uint64(621355968000000000)
	if ticks < unixOffset {
		return time.Time{}, nil // Bu özel bir durum, hata değil
	}

	return time.Now(), nil
}

// f32 fonksiyonu, ByteReader'dan bir 32-bit kayan nokta (float32) değeri okur.
// Yeterli byte yoksa hata döndürür.
func (b *ByteReader) f32() (float32, error) {
	if b.Offset+4 > uint(len(b.Data)) {
		return 0.0, errors.New("f32: okumak için yeterli byte yok")
	}

	_bytes, err := b.Rbytes(4)
	if err != nil { // Rbytes zaten kontrol ettiği için bu kontrol teknik olarak gereksiz ama yine de ekledim
		return 0.0, fmt.Errorf("f32: byte'lar okunamadı: %w", err)
	}
	value := binary.LittleEndian.Uint32(_bytes)

	return math.Float32frombits(value), nil
}

// f64 fonksiyonu, ByteReader'dan bir 64-bit kayan nokta (float64) değeri okur.
// Yeterli byte yoksa hata döndürür.
func (b *ByteReader) f64() (float64, error) {
	if b.Offset+8 > uint(len(b.Data)) {
		return 0.0, errors.New("f64: okumak için yeterli byte yok")
	}

	_bytes, err := b.Rbytes(8)
	if err != nil { // Rbytes zaten kontrol ettiği için bu kontrol teknik olarak gereksiz ama yine de ekledim
		return 0.0, fmt.Errorf("f64: byte'lar okunamadı: %w", err)
	}

	return math.Float64frombits(binary.LittleEndian.Uint64(_bytes)), nil
}

// Slice_bytes fonksiyonu, belirtilen başlangıç ve bitiş offset'leri arasındaki byte dilimini döndürür.
// Geçersiz dilim sınırları varsa hata döndürür.
func (b *ByteReader) Slice_bytes(start uint, end uint) ([]byte, error) {
	if start > end || end > uint(len(b.Data)) {
		return nil, errors.New("Slice_bytes: geçersiz dilim sınırları (start > end veya end > veri uzunluğu)")
	}
	if start > uint(len(b.Data)) { // Başlangıç da veriden büyük olamaz
		return nil, errors.New("Slice_bytes: başlangıç offset'i verinin uzunluğundan büyük")
	}

	// Dilimi kopyalayarak dışarıya sızmasını engellemek daha güvenli olabilir
	slice := make([]byte, end-start)
	copy(slice, b.Data[start:end])

	return slice, nil
}

// min, iki uint değerinden küçüğünü döndürür.
func min(a, b uint) uint {
	if a < b {
		return a
	}
	return b
}
