package e2e

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/openshift/rosa/tests/ci/labels"
	"github.com/openshift/rosa/tests/utils/config"
	"github.com/openshift/rosa/tests/utils/exec/rosacli"
)

var _ = Describe("Edit default ingress",
	labels.Day2,
	labels.FeatureIngress,
	func() {
		defer GinkgoRecover()

		var (
			clusterID      string
			rosaClient     *rosacli.Client
			ingressService rosacli.IngressService
			isHosted       bool
		)

		BeforeEach(func() {
			By("Get the cluster")
			clusterID = config.GetClusterID()
			Expect(clusterID).ToNot(Equal(""), "ClusterID is required. Please export CLUSTER_ID")

			By("Init the client")
			rosaClient = rosacli.NewClient()
			ingressService = rosaClient.Ingress

			By("Check cluster is hosted")
			var err error
			isHosted, err = rosaClient.Cluster.IsHostedCPCluster(clusterID)
			Expect(err).ToNot(HaveOccurred())

		})

		It("can update on rosa HCP cluster - [id:63323]",
			labels.Critical,
			func() {
				By("Retrieve cluster and get default ingress id")
				if !isHosted {
					Skip("This case is for HCP cluster")
				}
				output, err := ingressService.ListIngress(clusterID)
				Expect(err).ToNot(HaveOccurred())

				ingressList, err := ingressService.ReflectIngressList(output)
				Expect(err).ToNot(HaveOccurred())
				var defaultID, originalValue string
				for _, v := range ingressList.Ingresses {
					if v.Default == "yes" {
						defaultID = v.ID
						originalValue = v.Private
					}
				}

				By("Edit the default ingress on rosa HCP cluster to different value")
				updatedValue := "no"
				if originalValue == "no" {
					updatedValue = "yes"
				}
				testvalue := map[string]string{
					"yes": "true",
					"no":  "false",
				}
				cmdFlag := fmt.Sprintf("--private=%s", testvalue[updatedValue])
				output, err = ingressService.EditIngress(clusterID, defaultID,
					cmdFlag)
				Expect(err).ToNot(HaveOccurred())
				textData := rosaClient.Parser.TextData.Input(output).Parse().Tip()
				Expect(textData).Should(ContainSubstring("INFO: Updated ingress '%s' on cluster '%s'", defaultID, clusterID))

				defer func() {
					_, err = ingressService.EditIngress(clusterID, defaultID,
						fmt.Sprintf("--private=%s", testvalue[originalValue]))
					Expect(err).ToNot(HaveOccurred())

					output, err = ingressService.ListIngress(clusterID)
					Expect(err).ToNot(HaveOccurred())

					ingressList, err = ingressService.ReflectIngressList(output)
					Expect(err).ToNot(HaveOccurred())

					in := ingressList.Ingress(defaultID)
					Expect(in.Private).To(Equal(originalValue))
				}()

				output, err = ingressService.ListIngress(clusterID)
				Expect(err).ToNot(HaveOccurred())

				ingressList, err = ingressService.ReflectIngressList(output)
				Expect(err).ToNot(HaveOccurred())

				in := ingressList.Ingress(defaultID)
				Expect(in.Private).To(Equal(updatedValue))

				By("Edit the default ingress on rosa HCP cluster with current value")
				output, err = ingressService.EditIngress(clusterID, defaultID, cmdFlag)
				Expect(err).ToNot(HaveOccurred())
				textData = rosaClient.Parser.TextData.Input(output).Parse().Tip()
				Expect(textData).Should(ContainSubstring("WARN: No need to update ingress as there are no changes"))

				By("Edit the default ingress only with --private")
				output, err = ingressService.EditIngress(clusterID, defaultID, "--private")
				Expect(err).ToNot(HaveOccurred())
				textData = rosaClient.Parser.TextData.Input(output).Parse().Tip()
				if updatedValue == "yes" {
					Expect(textData).Should(ContainSubstring("WARN: No need to update ingress as there are no changes"))
				} else {
					Expect(textData).Should(ContainSubstring("Updated ingress '%s' on cluster '%s'", defaultID, clusterID))
				}

				By("Run command to edit a default ingress with --label-match")
				output, err = ingressService.EditIngress(clusterID, defaultID,
					"--label-match", "aaa=bbb,ccc=ddd")
				Expect(err).To(HaveOccurred())
				textData = rosaClient.Parser.TextData.Input(output).Parse().Tip()
				Expect(textData).Should(ContainSubstring("ERR: Updating route selectors is not supported for Hosted Control Plane clusters"))
			})

		It("change load balancer type - [id:64767]",
			labels.Critical,
			labels.Day2,
			labels.NonHCPCluster,
			func() {
				if isHosted {
					Skip("This case is for standard ROSA clusters only")
				}

				output, err := ingressService.ListIngress(clusterID)
				Expect(err).ToNot(HaveOccurred())
				ingressList, err := ingressService.ReflectIngressList(output)
				Expect(err).ToNot(HaveOccurred())

				defaultIngress := func(ingressList rosacli.IngressList) (*rosacli.Ingress, bool) {
					for _, ingress := range ingressList.Ingresses {
						if ingress.Default == "yes" {
							return &ingress, true
						}
					}
					return nil, false
				}
				ingress, exists := defaultIngress(*ingressList)
				Expect(exists).To(BeTrue())
				defaultID := ingress.ID
				Expect(defaultID).ToNot(BeNil())
				output, err = ingressService.EditIngress(clusterID, defaultID, "--lb-type", "nlb")
				Expect(err).ToNot(HaveOccurred())
				textData := rosaClient.Parser.TextData.Input(output).Parse().Tip()
				Expect(textData).Should(ContainSubstring("Updated ingress '%s'", defaultID))

				output, err = ingressService.ListIngress(clusterID)
				Expect(err).ToNot(HaveOccurred())
				ingressList, err = ingressService.ReflectIngressList(output)
				Expect(err).ToNot(HaveOccurred())
				ingress, _ = defaultIngress(*ingressList)
				Expect(ingress.LBType).Should(ContainSubstring("nlb"))

				output, err = ingressService.EditIngress(clusterID, defaultID, "--lb-type", "classic")
				Expect(err).ToNot(HaveOccurred())
				textData = rosaClient.Parser.TextData.Input(output).Parse().Tip()
				Expect(textData).Should(ContainSubstring("Updated ingress '%s'", defaultID))

				output, err = ingressService.ListIngress(clusterID)
				Expect(err).ToNot(HaveOccurred())
				ingressList, err = ingressService.ReflectIngressList(output)
				Expect(err).ToNot(HaveOccurred())
				ingress, _ = defaultIngress(*ingressList)
				Expect(ingress.LBType).Should(ContainSubstring("classic"))
			})
		It("Update Ingress Controller attributes [id:65799]",
			labels.Critical,
			labels.Day2,
			labels.NonHCPCluster,
			func() {
				if isHosted {
					Skip("This case is for standard ROSA clusters only")
				}
				output, err := ingressService.ListIngress(clusterID)
				Expect(err).ToNot(HaveOccurred())

				ingressList, err := ingressService.ReflectIngressList(output)
				Expect(err).ToNot(HaveOccurred())
				defaultIngress := func(ingressList rosacli.IngressList) (*rosacli.Ingress, bool) {
					for _, ingress := range ingressList.Ingresses {
						if ingress.Default == "yes" {
							return &ingress, true
						}
					}
					return nil, false
				}

				ingress, exists := defaultIngress(*ingressList)
				Expect(exists).To(BeTrue())
				defaultID := ingress.ID
				output, err = ingressService.EditIngress(clusterID, defaultID, "--excluded-namespaces", "test-ns1,test-ns2", "--route-selector",
					"app1=test1,app2=test2", "--namespace-ownership-policy", "Strict", "--wildcard-policy", "WildcardsDisallowed")
				Expect(err).ToNot(HaveOccurred())
				textData := rosaClient.Parser.TextData.Input(output).Parse().Tip()
				Expect(textData).Should(ContainSubstring("Updated ingress '%s'", defaultID))

				output, err = ingressService.ListIngress(clusterID)
				Expect(err).ToNot(HaveOccurred())

				ingressList, err = ingressService.ReflectIngressList(output)
				Expect(err).ToNot(HaveOccurred())

				ingress, _ = defaultIngress(*ingressList)
				Expect(ingress.ExcludeNamespace).Should(ContainSubstring("test-ns1"))
				Expect(ingress.ExcludeNamespace).Should(ContainSubstring("test-ns2"))
				Expect(ingress.RouteSelectors).Should(ContainSubstring("app1=test1"))
				Expect(ingress.RouteSelectors).Should(ContainSubstring("app2=test2"))
				Expect(ingress.NamespaceOwnershipPolicy).Should(ContainSubstring("Strict"))
				Expect(ingress.WildcardPolicy).Should(ContainSubstring("WildcardsDisallowed"))
			})
	})
