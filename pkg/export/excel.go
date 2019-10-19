package export

import "github.com/boybird/go_api_development_demo/pkg/setting"

// EXT excel export extenstion
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
