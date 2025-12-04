package commands

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/urfave/cli/v2"

	"github.com/thash/asana/api"
)

func Download(c *cli.Context) {
	args := c.Args()
	if args.Len() < 1 {
		fmt.Println("Usage: asana download <task_index> <attachment_index>")
		fmt.Println("       asana download <attachment_gid>")
		fmt.Println("\nTo see attachments, use: asana task <task_index>")
		return
	}

	var attachment api.Attachment_t

	// Check if first argument is an attachment GID (starts with "1" and is long)
	if len(args.First()) > 10 && args.Len() == 1 {
		// Assume it's an attachment GID
		attachment = api.Attachment(args.First())
	} else if args.Len() >= 2 {
		// Assume it's task_index and attachment_index
		taskId := api.FindTaskId(args.First(), false)
		attachments := api.Attachments(taskId)

		if len(attachments) == 0 {
			fmt.Println("No attachments found for this task")
			return
		}

		attIndex, err := strconv.Atoi(args.Get(1))
		if err != nil || attIndex < 0 || attIndex >= len(attachments) {
			fmt.Printf("Invalid attachment index. Valid range: 0-%d\n", len(attachments)-1)
			return
		}

		// Get full attachment details with download URL
		attachment = api.Attachment(attachments[attIndex].Gid)
	} else {
		fmt.Println("Usage: asana download <task_index> <attachment_index>")
		fmt.Println("       asana download <attachment_gid>")
		return
	}

	// Get download URL with full details
	if attachment.DownloadUrl == "" {
		fmt.Println("Error: No download URL available for this attachment")
		return
	}

	// Determine output filename
	outputPath := c.String("output")
	if outputPath == "" {
		outputPath = attachment.Name
	}

	fmt.Printf("Downloading: %s\n", attachment.Name)
	fmt.Printf("Saving to: %s\n", outputPath)

	// Download the file
	resp, err := http.Get(attachment.DownloadUrl)
	if err != nil {
		fmt.Printf("Error downloading file: %v\n", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Printf("Error: HTTP %d\n", resp.StatusCode)
		return
	}

	// Create output file
	os.MkdirAll(filepath.Dir(outputPath), 0755)
	out, err := os.Create(outputPath)
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return
	}
	defer out.Close()

	// Copy data
	written, err := io.Copy(out, resp.Body)
	if err != nil {
		fmt.Printf("Error writing file: %v\n", err)
		return
	}

	fmt.Printf("Downloaded successfully! (%d bytes)\n", written)
}
