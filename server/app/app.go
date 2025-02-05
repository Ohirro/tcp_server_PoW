package app

import (
	"bufio"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"net"
	"strings"
)

var quotes = []string{
	"The only true wisdom is in knowing you know nothing. - Socrates",
	"Do not take life too seriously. You will never get out of it alive. - Elbert Hubbard",
	"In the middle of difficulty lies opportunity. - Albert Einstein",
	"The journey of a thousand miles begins with one step. - Lao Tzu",
	"Life is what happens when you're busy making other plans. - John Lennon",
}

func generateChallenge() (string, error) {
	bytes := make([]byte, 16)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func verifyPoW(challenge string, nonce string, difficulty int) bool {
	data := challenge + nonce
	hashBytes := sha256.Sum256([]byte(data))
	hashStr := hex.EncodeToString(hashBytes[:])
	prefix := strings.Repeat("0", difficulty)
	return strings.HasPrefix(hashStr, prefix)
}

func getRandomQuote() (string, error) {
	maxInt := big.NewInt(int64(len(quotes)))
	nBig, err := rand.Int(rand.Reader, maxInt)
	if err != nil {
		return "", err
	}
	return quotes[nBig.Int64()], nil
}

func Run() error {
	config, err := LoadConfig()
	if err != nil {
		return err
	}

	address := ":" + config.Port
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}
	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			log.Printf("failed to close listener: %v", err)
		}
	}(listener)

	log.Printf("Server is running on port %s", config.Port)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Failed to accept connection: %v", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			log.Printf("failed to close connection: %v", err)
		}
	}(conn)
	log.Printf("New client connected: %s", conn.RemoteAddr())

	difficulty := 4
	challenge, err := generateChallenge()
	if err != nil {
		log.Printf("Failed to generate challenge: %v", err)
		return
	}

	message := fmt.Sprintf("%s:%d\n", challenge, difficulty)
	_, err = conn.Write([]byte(message))
	if err != nil {
		log.Printf("Failed to send challenge: %v", err)
		return
	}

	reader := bufio.NewReader(conn)
	nonce, err := reader.ReadString('\n')
	if err != nil {
		log.Printf("Failed to read nonce from client: %v", err)
		return
	}
	nonce = strings.TrimSpace(nonce)
	log.Printf("Received nonce: %s from %s", nonce, conn.RemoteAddr())

	if verifyPoW(challenge, nonce, difficulty) {
		quote, err := getRandomQuote()
		if err != nil {
			log.Printf("Failed to retrieve quote: %v", err)
			return
		}
		_, err = conn.Write([]byte(quote + "\n"))
		if err != nil {
			log.Printf("Failed to send quote: %v", err)
		} else {
			log.Printf("Sent quote to client %s", conn.RemoteAddr())
		}
	} else {
		_, err := conn.Write([]byte("Invalid proof of work\n"))
		if err != nil {
			log.Printf("Failed to send challenge: %v", err)
			return
		}
		log.Printf("Invalid PoW solution from %s", conn.RemoteAddr())
	}
}
