package starter

import (
	"fmt"
	"os"
	"time"

	"github.com/couchbase/gocb/v2"
	"github.com/joho/godotenv"
)

func Run() error {
	// Loads the env variable in the project
	godotenv.Load()
	// Uncomment following line to enable logging
	// gocb.SetLogger(gocb.VerboseStdioLogger())

	// Update this to your cluster details

	connectionString := os.Getenv("GO_COUCHBASE_CONNECTION_STRING")
	bucketName := os.Getenv("COUCHBASE_DEFAULT_BUCKET")
	username := os.Getenv("GO_COUCHBASE_USERNAME")
	password := os.Getenv("GO_COUCHBASE_PASSWORD")
	fmt.Println(bucketName)

	options := gocb.ClusterOptions{
		Authenticator: gocb.PasswordAuthenticator{
			Username: username,
			Password: password,
		},
	}

	// Sets a pre-configured profile called "wan-development" to help avoid latency issues
	// when accessing Capella from a different Wide Area Network
	// or Availability Zone (e.g. your laptop).
	if err := options.ApplyProfile(gocb.ClusterConfigProfileWanDevelopment); err != nil {
		return err
	}

	// Initialize the Connection
	cluster, err := gocb.Connect(connectionString, options)
	if err != nil {
		return err
	}

	bucket := cluster.Bucket(bucketName)

	err = bucket.WaitUntilReady(10*time.Second, nil)
	if err != nil {
		return err
	}

	// Get a reference to the default collection, required for older Couchbase server versions
	//Uncomment the next line to use the default collection
	// defaultcol := bucket.DefaultCollection()

	col := bucket.Scope("tenant_agent_00").Collection("users")

	// Create and store a Document
	type User struct {
		Name      string   `json:"name"`
		Email     string   `json:"email"`
		Interests []string `json:"interests"`
	}
	// Upsert 
	_, err = col.Upsert("u:john",
		User{
			Name:      "John",
			Email:     "john@test-email.com",
			Interests: []string{"Basketball", "Rugby"},
		}, nil)
	if err != nil {
		return err
	}

	// Get the document back
	getResult, err := col.Get("u:john", nil)
	if err != nil {
		return err
	}

	var inUser User
	err = getResult.Content(&inUser)
	if err != nil {
		return err
	}
	fmt.Printf("User: %v\n", inUser)

	// Perform a N1QL Query
	inventoryScope := bucket.Scope("inventory")
	queryResult, err := inventoryScope.Query(
		fmt.Sprintf("SELECT * FROM airline WHERE id=10"),
		&gocb.QueryOptions{},
	)
	if err != nil {
		return err
	}

	// Print each found Row
	for queryResult.Next() {
		var result interface{}
		err := queryResult.Row(&result)
		if err != nil {
			return err
		}
		fmt.Println(result)
	}

	if err := queryResult.Err(); err != nil {
		return err
	}
	return nil
}
