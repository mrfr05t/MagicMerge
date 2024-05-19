package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
	"time"
)

func startNgrok() (string, error) {
	cmd := exec.Command("ngrok", "http", "8080")
	if err := cmd.Start(); err != nil {
		return "", err
	}

	time.Sleep(2 * time.Second)

	resp, err := http.Get("http://127.0.0.1:4040/api/tunnels")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var data struct {
		Tunnels []struct {
			PublicURL string `json:"public_url"`
		} `json:"tunnels"`
	}

	if err := json.Unmarshal(body, &data); err != nil {
		return "", err
	}

	if len(data.Tunnels) > 0 {
		return data.Tunnels[0].PublicURL, nil
	}
	return "", fmt.Errorf("no tunnels found")
}
