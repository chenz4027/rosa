package network

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	opts "github.com/openshift/rosa/pkg/options/network"
	"github.com/openshift/rosa/pkg/reporter"
)

var _ = Describe("CreateNetworkOptions", func() {
	var (
		networkOptions *Options
		userOptions    *opts.NetworkUserOptions
	)

	BeforeEach(func() {
		networkOptions = NewNetworkOptions()
		userOptions = NewNetworkUserOptions()
	})

	Context("NewNetworkUserOptions", func() {
		It("should create default user options", func() {
			Expect(userOptions.Params).To(Equal([]string{}))
		})
	})

	Context("NewNetworkOptions", func() {
		It("should create default network options", func() {
			Expect(networkOptions.reporter).To(BeAssignableToTypeOf(&reporter.Object{}))
			Expect(networkOptions.args).To(BeAssignableToTypeOf(&opts.NetworkUserOptions{}))
		})
	})

	Context("Network", func() {
		It("should return the args field", func() {
			networkOptions.args = userOptions
			Expect(networkOptions.Network()).To(Equal(userOptions))
		})
	})
})
