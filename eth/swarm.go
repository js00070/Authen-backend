package eth

import (
	"fmt"
	"io/ioutil"
	"log"

	bzzclient "github.com/ethersphere/swarm/api/client"
)

// SwarmClient swarm client
var SwarmClient *bzzclient.Client

// Init 初始化
func Init(swarmHost string) {
	SwarmClient = bzzclient.NewClient(fmt.Sprintf("http://%s:8500", swarmHost))
}

// UploadETH 上传至以太坊swarm
func UploadETH(path string) (string, error) {
	file, err := bzzclient.Open(path)
	if err != nil {
		return "", err
	}
	manifestHash, err := SwarmClient.Upload(file, "", false, false, false)
	if err != nil {
		return "", err
	}

	fmt.Println("manifestHash is ", manifestHash)
	return manifestHash, nil
}

// DownloadETH 下载
func DownloadETH(manifestHash string) {
	manifest, isEncrypted, err := SwarmClient.DownloadManifest(manifestHash)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(isEncrypted) // false

	for _, entry := range manifest.Entries {
		fmt.Println(entry.Hash)        // manifestHash
		fmt.Println(entry.ContentType) // text/plain; charset=utf-8
		fmt.Println(entry.Size)        // 12
		fmt.Println(entry.Path)        // ""
	}

	file, err := SwarmClient.Download(manifestHash, "")
	if err != nil {
		log.Fatal(err)
	}

	content, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(content)) // hello world
}
