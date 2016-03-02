package application_test

import (
	"bytes"

	"github.com/pivotal-cf-experimental/bosh-bootloader/application"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Logger", func() {
	var (
		buffer *bytes.Buffer
		logger *application.Logger
	)

	BeforeEach(func() {
		buffer = bytes.NewBuffer([]byte{})
		logger = application.NewLogger(buffer)
	})

	Describe("Step", func() {
		It("prints the step message", func() {
			logger.Step("creating key pair")

			Expect(buffer.String()).To(Equal("step: creating key pair\n"))
		})
	})

	Describe("Dot", func() {
		It("prints a dot", func() {
			logger.Dot()
			logger.Dot()
			logger.Dot()

			Expect(buffer.String()).To(Equal("\u2022\u2022\u2022"))
		})
	})

	Describe("mixing steps and dots", func() {
		It("prints out a coherent set of lines", func() {
			logger.Step("creating keypair")
			logger.Step("generating cloudformation template")
			logger.Step("applying cloudformation template")
			logger.Dot()
			logger.Dot()
			logger.Dot()
			logger.Dot()
			logger.Dot()
			logger.Step("completed applying cloudformation template")

			Expect(buffer.String()).To(Equal(`step: creating keypair
step: generating cloudformation template
step: applying cloudformation template
•••••
step: completed applying cloudformation template
`))
		})
	})
})
