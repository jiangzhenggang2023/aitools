package nas

import (
	"aitools/internal/models"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sync"
)

func processData(id int, jobs <-chan string, wg *sync.WaitGroup, kbID, accessToken string, client *http.Client) {
	defer wg.Done()

	var requst models.UpdateData
	var req *http.Request
	requst.KBID = kbID
	requst.AccessToken = accessToken

	for item := range jobs {
		// 处理每个文件的上传逻辑
		requst.Files = []string{item}
		jsonData, err := json.Marshal(requst)
		if err != nil {
			fmt.Printf("Worker %d: failed to marshal request for file %s: %v\n", id, item, err)
			continue
		}

		if id == 1 {
			fmt.Println("Worker 1 is uploading to port 5001")
			req, _ = http.NewRequest("POST", "http://10.67.112.152:5001/v1/document/upload", bytes.NewReader(jsonData))
		} else {
			fmt.Println("Worker 2 is uploading to port 6001")
			req, _ = http.NewRequest("POST", "http://10.67.112.152:6001/v1/document/upload", bytes.NewReader(jsonData))
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+accessToken)
		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("Worker %d: failed to upload file %s: %v\n", id, item, err)
			continue
		}
		resp.Body.Close()
	}
}

// GetFilesInDirectory returns a slice of file paths in the given directory.
func GetFilesInDirectory(dir string) ([]string, error) {
	files := []string{}
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	for _, entry := range entries {
		if !entry.IsDir() {
			files = append(files, dir+"/"+entry.Name())
		}
	}
	return files, nil
}

// UploadFiles uploads files to the NAS server using the provided URL and data.
func UploadFiles(data models.UpdateData) (bool, error) {
	// 如果data.Files里面的值是目录，则获取该目录中所有文件路径
	filesList := make([]string, 0)
	for _, file := range data.Files {
		if file == "" {
			continue
		}
		fileInfo, err := os.Stat(file)
		if err != nil {
			return false, fmt.Errorf("failed to stat file %s: %w", file, err)
		}
		if fileInfo.IsDir() {
			// 获取file目录下的所有文件
			files, err := GetFilesInDirectory(file)
			if err != nil {
				return false, fmt.Errorf("failed to get files in directory %s: %w", file, err)
			}
			filesList = append(filesList, files...)
		} else {
			// 如果是文件，则直接添加到filesList
			filesList = append(filesList, file)
		}
	}
	fmt.Println("-------------------")
	fmt.Println(filesList)
	fmt.Println("-------------------")
	jobChan := make(chan string, len(filesList))
	var wg sync.WaitGroup

	client := &http.Client{}

	// 启动两个协程
	for w := 1; w <= 2; w++ {
		wg.Add(1)
		go processData(w, jobChan, &wg, data.KBID, data.AccessToken, client)
	}

	// 分发任务到通道

	for _, item := range filesList {
		jobChan <- item
	}
	close(jobChan)
	// 等待所有协程完成
	wg.Wait()

	fmt.Println("All files uploaded successfully.")
	return true, nil
}
