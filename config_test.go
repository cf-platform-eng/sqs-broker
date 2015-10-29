package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/cf-platform-eng/sqs-broker"

	"github.com/cf-platform-eng/sqs-broker/sqsbroker"
)

var _ = Describe("Config", func() {
	var (
		config Config

		validConfig = Config{
			LogLevel: "DEBUG",
			Username: "broker-username",
			Password: "broker-password",
			SQSConfig: sqsbroker.Config{
				Region:      "sqs-region",
				QueuePrefix: "cf",
			},
		}
	)

	Describe("Validate", func() {
		BeforeEach(func() {
			config = validConfig
		})

		It("does not return error if all sections are valid", func() {
			err := config.Validate()
			Expect(err).ToNot(HaveOccurred())
		})

		It("returns error if LogLevel is not valid", func() {
			config.LogLevel = ""

			err := config.Validate()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("Must provide a non-empty LogLevel"))
		})

		It("returns error if Username is not valid", func() {
			config.Username = ""

			err := config.Validate()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("Must provide a non-empty Username"))
		})

		It("returns error if Password is not valid", func() {
			config.Password = ""

			err := config.Validate()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("Must provide a non-empty Password"))
		})

		It("returns error if SQS configuration is not valid", func() {
			config.SQSConfig = sqsbroker.Config{}

			err := config.Validate()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("Validating SQS configuration"))
		})
	})
})
