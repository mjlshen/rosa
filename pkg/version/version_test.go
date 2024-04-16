package version

import (
	"fmt"

	"go.uber.org/mock/gomock"

	goVer "github.com/hashicorp/go-version"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/sirupsen/logrus"
)

var _ = Describe("IsLatest", func() {
	var (
		mockRetriever *MockRetriever
		mockCtrl      *gomock.Controller
		v             RosaVersion
		logger        *logrus.Logger
		currentVer    string
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockRetriever = NewMockRetriever(mockCtrl)
		logger = logrus.New()
		v = &rosaVersion{
			logger:    logger,
			retriever: mockRetriever,
		}
		currentVer = "1.0.0"
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	When("the current version is less than the latest version", func() {
		It("should return the latest version and false", func() {
			latestVer := "2.0.0"
			mockRetriever.EXPECT().RetrieveLatestVersionFromMirror().Return(goVer.Must(goVer.NewVersion(latestVer)), nil)

			latestVersion, isLatest, err := v.IsLatest(currentVer)

			Expect(err).To(BeNil())
			Expect(isLatest).To(BeFalse())
			Expect(latestVersion.String()).To(Equal(latestVer))
		})
	})

	When("the current version is equal to the latest version", func() {
		It("should return nil and true", func() {
			latestVer := "1.0.0"
			mockRetriever.EXPECT().RetrieveLatestVersionFromMirror().Return(goVer.Must(goVer.NewVersion(latestVer)), nil)

			latestVersion, isLatest, err := v.IsLatest(currentVer)

			Expect(err).To(BeNil())
			Expect(isLatest).To(BeTrue())
			Expect(latestVersion).To(BeNil())
		})
	})

	When("there is an error retrieving the latest version", func() {
		It("should return an error", func() {
			mockRetriever.EXPECT().RetrieveLatestVersionFromMirror().Return(nil, fmt.Errorf("error"))

			latestVersion, isLatest, err := v.IsLatest(currentVer)

			Expect(err).To(HaveOccurred())
			Expect(isLatest).To(BeFalse())
			Expect(latestVersion).To(BeNil())
		})
	})
})

var _ = Describe("NewRosaVersion", func() {
	When("all dependencies are created successfully", func() {
		It("should return a new RosaVersion instance", func() {
			rosaVersion, err := NewRosaVersion()
			Expect(err).To(BeNil())
			Expect(rosaVersion).ToNot(BeNil())
		})
	})
})
