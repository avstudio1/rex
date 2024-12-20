package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

const xAPIURL = "https://api.twitter.com/2/users/me"

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the CLI and test connection to X API",
	Long: `This command initializes the rex CLI by loading the API key from
the .env file and testing the connection to the X API.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := loadEnv()
		if err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}

		apiKey := os.Getenv("API_KEY")
		if apiKey == "" {
			log.Fatal("API_KEY is not set in .env")
		}

		err = testXAPIConnection(apiKey)
		if err != nil {
			log.Fatalf("Connection test failed: %v", err)
		}

		fmt.Println("Initialization successful! Connection to X API verified.")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

// loadEnv loads the .env file
func loadEnv() error {
	return godotenv.Load()
}

// testXAPIConnection tests the connection to the X API using the provided API key
func testXAPIConnection(apiKey string) error {
	client := &http.Client{}
	req, err := http.NewRequest("GET", xAPIURL, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+apiKey)
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to perform request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}
