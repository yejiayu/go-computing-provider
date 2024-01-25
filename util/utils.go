package util

import (
	"fmt"
	"github.com/valyala/gozstd"
	"io"
	"net/http"
	"os"
)

func SaveMcsFileByUrlToFile(fileName, mcsUrl string) error {
	resp, err := http.Get(mcsUrl)
	if err != nil {
		return fmt.Errorf("error making request to Space API: %+v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("space API response not OK. Status Code: %d", resp.StatusCode)
	}

	outBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read body failed, error:", err)
		return err
	}

	out, err := gozstd.Decompress(nil, outBytes)
	if err != nil {
		fmt.Println("decompress response bytes failed, error:", err)
		return err
	}
	return os.WriteFile(fileName, out, 0655)
}
