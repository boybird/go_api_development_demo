package export

import "github.com/boybird/hello/pkg/setting"

// EXT Excel export as excel 2007
const EXT = ".xlsx"

// GetExcelFullURL get the full access path of the Excel file
func GetExcelFullURL(name string) string {
	return setting.AppSetting.PrefixURL + "/" + GetExcelPath() + name
}

// GetExcelPath get the relative save path of the Excel file
func GetExcelPath() string {
	return setting.AppSetting.ExportSavePath
}

// GetExcelFullPath Get the full save path of the Excel file
func GetExcelFullPath() string {
	return setting.AppSetting.RuntimeRootPath + GetExcelPath()
}
