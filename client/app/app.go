package app

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
	"time"
)

func Run() error {
	config, err := LoadConfig()
	if err != nil {
		return err
	}

	conn, err := net.Dial("tcp", config.ServerAddress)
	if err != nil {
		return fmt.Errorf("failed to connect to server: %w", err)
	}
	defer conn.Close()

	reader := bufio.NewReader(conn)

	challengeLine, err := reader.ReadString('\n')
	if err != nil {
		return fmt.Errorf("failed to read challenge from server: %w", err)
	}
	challengeLine = strings.TrimSpace(challengeLine)
	parts := strings.Split(challengeLine, ":")
	if len(parts) != 2 {
		return fmt.Errorf("invalid challenge format: %s", challengeLine)
	}
	challenge := parts[0]
	difficulty, err := strconv.Atoi(parts[1])
	if err != nil {
		return fmt.Errorf("invalid difficulty value: %w", err)
	}

	log.Printf("Received challenge: %s, difficulty: %d", challenge, difficulty)

	start := time.Now()
	nonce := solvePoW(challenge, difficulty)
	elapsed := time.Since(start)
	log.Printf("Found nonce: %s in %s", nonce, elapsed)

	_, err = conn.Write([]byte(nonce + "\n"))
	if err != nil {
		return fmt.Errorf("failed to send nonce to server: %w", err)
	}

	quote, err := reader.ReadString('\n')
	if err != nil {
		return fmt.Errorf("failed to read quote from server: %w", err)
	}
	quote = strings.TrimSpace(quote)
	fmt.Printf("Quote: %s\n", quote)

	return nil
}

func solvePoW(challenge string, difficulty int) string {
	prefix := strings.Repeat("0", difficulty)
	var nonce int64 = 0
	for {
		candidate := fmt.Sprintf("%d", nonce)
		data := challenge + candidate
		hashBytes := sha256.Sum256([]byte(data))
		hashStr := hex.EncodeToString(hashBytes[:])
		if strings.HasPrefix(hashStr, prefix) {
			return candidate
		}
		nonce++
	}
}
