package respx

import (
	"fmt"
	"net/http"
)

func HttpExcelResult(w http.ResponseWriter, fileName string) {
	HttpExcelResultWithCache(w, fileName, nil)
}

func HttpExcelResultWithCache(w http.ResponseWriter, fileName string, cacheTime *int) {
	w.Header().Add("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	w.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
	if cacheTime != nil {
		if *cacheTime == 0 {
			w.Header().Add("Cache-Control", "no-cache, no-store, must-revalidate")
			w.Header().Add("Pragma", "no-cache")
			w.Header().Add("Expires", "0")
		} else {
			age := fmt.Sprintf("%d", *cacheTime)
			w.Header().Add("Cache-Control", fmt.Sprintf("max-age=%s", age))
			w.Header().Add("Pragma", "cache")
			w.Header().Add("Expires", age)
		}
	}
}
