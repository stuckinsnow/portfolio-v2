package services

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gofiber/fiber/v2"
)

type S3Service struct {
	client *s3.Client
	region string
}

func NewS3Service() *S3Service {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
	if err != nil {
		log.Printf("Failed to load AWS config: %v", err)
		return nil
	}

	return &S3Service{
		client: s3.NewFromConfig(cfg),
		region: "us-east-1",
	}
}

func (s *S3Service) HandleListObjects(c *fiber.Ctx) error {
	log.Println("S3 ListObjects endpoint called")

	bucketName := c.Query("bucket")
	if bucketName == "" {
		bucketName = os.Getenv("S3_BUCKET")
	}

	log.Printf("Listing objects in bucket: %s", bucketName)

	result, err := s.client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: &bucketName,
	})
	if err != nil {
		log.Printf("S3 ListObjects error: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Unable to list objects",
		})
	}

	log.Printf("Found %d objects", len(result.Contents))

	var objects []string
	for _, obj := range result.Contents {
		objects = append(objects, *obj.Key)
		log.Printf("Object: %s", *obj.Key)
	}

	return c.JSON(fiber.Map{
		"bucket":  bucketName,
		"objects": objects,
	})
}

func (s *S3Service) HandleGetObject(c *fiber.Ctx) error {
	bucketName := c.Query("bucket")
	if bucketName == "" {
		bucketName = os.Getenv("S3_BUCKET")
	}

	objectKey := c.Query("key")
	if objectKey == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Object key is required",
		})
	}

	result, err := s.client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: &bucketName,
		Key:    &objectKey,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Unable to get object",
		})
	}
	defer result.Body.Close()

	c.Set("Cache-Control", "public, max-age=3600")
	c.Set("Content-Type", "image/jpeg")

	body, err := io.ReadAll(result.Body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to read object",
		})
	}

	return c.Send(body)
}
