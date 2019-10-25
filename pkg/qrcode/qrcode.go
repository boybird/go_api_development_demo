package qrcode

import (
	"image/jpeg"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/boybird/hello/pkg/file"
	"github.com/boybird/hello/pkg/setting"
	"github.com/boybird/hello/pkg/util"
)

// QrCode struct
type QrCode struct {
	URL    string
	Width  int
	Height int
	Ext    string
	Level  qr.ErrorCorrectionLevel
	Mode   qr.Encoding
}

const (
	// EXT_JPG image jpg
	_JPG = ".jpg"
)

// NewQrCode get qrCode struct
func NewQrCode(url string, width, height int, level qr.ErrorCorrectionLevel, mode qr.Encoding) *QrCode {
	return &QrCode{
		URL:    url,
		Width:  width,
		Height: height,
		Level:  level,
		Mode:   mode,
		Ext:    _JPG,
	}
}

// GetQrCodePath get save path
func GetQrCodePath() string {
	return setting.AppSetting.QrCodeSavePath
}

// GetQrCodeFullPath get full save path
func GetQrCodeFullPath() string {
	return setting.AppSetting.RuntimeRootPath + setting.AppSetting.QrCodeSavePath
}

// GetQrCodeFullURL get the full access path
func GetQrCodeFullURL(name string) string {
	return setting.AppSetting.PrefixURL + "/" + GetQrCodePath() + name
}

// GetQrCodeFileName get qr file name
func GetQrCodeFileName(value string) string {
	return util.EncodeMD5(value)
}

// GetQrCodeExt get qr file ext
func (q *QrCode) GetQrCodeExt() string {
	return q.Ext
}

// Encode generate QR code
func (q *QrCode) Encode(path string) (string, string, error) {
	name := GetQrCodeFileName(q.URL) + q.GetQrCodeExt()
	src := path + name
	if file.CheckNotExist(src) == true {
		code, err := qr.Encode(q.URL, q.Level, q.Mode)
		if err != nil {
			return "", "", err
		}
		code, err = barcode.Scale(code, q.Width, q.Height)
		if err != nil {
			return "", "", err
		}
		f, err := file.MustOpen(name, path)
		if err != nil {
			return "", "", err
		}
		defer f.Close()
		err = jpeg.Encode(f, code, nil)
		if err != nil {
			return "", "", err
		}
	}
	return name, path, nil
}
