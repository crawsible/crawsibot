package config_test

import (
	"reflect"

	"github.com/crawsible/crawsibot/config"
	"github.com/crawsible/crawsibot/config/mocks"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Config", func() {
	Describe("#MakeFlags", func() {
		var (
			fakeFlagSet *mocks.FakeFlagSet
			args        []string
			cfg         *config.Config
		)

		BeforeEach(func() {
			fakeFlagSet = mocks.NewFakeFlagSet()

			args = []string{
				"-a", "localhost:3000",
				"-n", "some-nick",
				"-p", "oauth:key",
				"-c", "somechannel",
			}
			cfg = &config.Config{}
			cfg.MakeFlags(fakeFlagSet, args)
		})

		It("defines the appropriate flags", func() {
			Expect(fakeFlagSet.StringVarCalls).To(Equal(4))

			flag0 := []string{"a", "irc.twitch.tv:6667", "The IRC server address (<hostname>:<port>)"}
			flag1 := []string{"n", "", "Your IRC nickname"}
			flag2 := []string{"p", "", "Your IRC password"}
			flag3 := []string{"c", "", "The channel you're connecting to (omit '#')"}

			Expect(fakeFlagSet.DefinedFlags).To(ContainElement(flag0))
			Expect(fakeFlagSet.DefinedFlags).To(ContainElement(flag1))
			Expect(fakeFlagSet.DefinedFlags).To(ContainElement(flag2))
			Expect(fakeFlagSet.DefinedFlags).To(ContainElement(flag3))
		})

		It("binds the flags to the config receiver's fields", func() {
			addressPtr := reflect.ValueOf(&(cfg.Address)).Pointer()
			nickPtr := reflect.ValueOf(&(cfg.Nick)).Pointer()
			passwordPtr := reflect.ValueOf(&(cfg.Password)).Pointer()
			channelPtr := reflect.ValueOf(&(cfg.Channel)).Pointer()

			Expect(fakeFlagSet.BoundVarPtrs["a"]).To(Equal(addressPtr))
			Expect(fakeFlagSet.BoundVarPtrs["n"]).To(Equal(nickPtr))
			Expect(fakeFlagSet.BoundVarPtrs["p"]).To(Equal(passwordPtr))
			Expect(fakeFlagSet.BoundVarPtrs["c"]).To(Equal(channelPtr))
		})

		It("calls Parse with the provided args after defining the flags", func() {
			Expect(fakeFlagSet.ParseCalls).To(Equal(1))
			Expect(fakeFlagSet.ParseArgs).To(Equal(args))
			Expect(fakeFlagSet.StringVarCallsAtParse).To(Equal(4))
		})
	})
})
