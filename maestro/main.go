package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"complyagent.com/maestro/gen/common"
	"complyagent.com/maestro/gen/ocr"

	"google.golang.org/grpc"
)

func main() {
	env := os.Getenv("ENV")
	agentPort := os.Getenv("AGENT_OCR_PORT")
	if agentPort == "" {
		agentPort = "50051"
	}

	fmt.Println("-------------------------------------------------")
	fmt.Println(" 🚀 Maestro starting up")
	fmt.Println("-------------------------------------------------")
	fmt.Printf(" Environment : %s\n", env)
	fmt.Printf(" Agent OCR   : agent_ocr:%s\n", agentPort)
	fmt.Println("-------------------------------------------------")

	// Connect to agent_ocr
	address := fmt.Sprintf("agent_ocr:%s", agentPort)
	conn, err := grpc.Dial(
		address,
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithTimeout(5*time.Second),
	)
	if err != nil {
		log.Fatalf("❌ Could not connect to agent_ocr at %s: %v", address, err)
	}
	defer conn.Close()
	log.Printf("✅ Connected to agent_ocr at %s", address)

	// Create client
	client := ocr.NewOCRServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Example: read an image and send it
	imageBytes, err := os.ReadFile("sample.png")
	if err != nil {
		log.Fatalf("❌ Failed to read sample.png: %v", err)
	}

	req := &ocr.OCRExtractRequest{
		Image: &common.Image{
			Data:   imageBytes,
			Format: "png",
		},
	}

	resp, err := client.Extract(ctx, req)
	if err != nil {
		log.Fatalf("❌ Extract failed: %v", err)
	}

	log.Printf("📩 OCR Text: %s", resp.Text)

	// Stay alive
	log.Println("Maestro is now running... waiting for future work")
	select {}
}
