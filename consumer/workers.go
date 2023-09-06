package consumer

import (
	"log"
	"message_queue_service/utils"
	"message_queue_service/wire"
	"strconv"
	"sync"
)

func worker(jobs chan string, workerID int) {
	for job := range jobs {
		log.Printf("worker : %d started Job : %s\n", workerID, job)
		productID, err := strconv.ParseInt(job, 10, 64)
		if err != nil {
			return
		}

		pID := int(productID)

		urls, dbErr := wire.ProductRepo.GetProductURLs(pID)
		if dbErr != nil {
			return
		}

		var wg sync.WaitGroup

		for offset, url := range urls {
			imageURL := url
			wg.Add(1)
			go func(pID, offset int, wg *sync.WaitGroup) {
				defer wg.Done()

				fileErr := wire.FileService.DownloadAndCompress(imageURL, utils.GetFileName(pID, offset))
				if fileErr != nil {
					return
				}

				path, pathErr := wire.FileService.GetFilePath(utils.GetFileName(pID, offset))
				if pathErr != nil {
					return
				}

				saveErr := wire.ProductRepo.SaveLocalPath(pID, path)
				if saveErr != nil {
					return
				}
			}(pID, offset, &wg)
		}

		wg.Wait()
		log.Printf("worker : %d completed Job : %s\n", workerID, job)
	}
}
