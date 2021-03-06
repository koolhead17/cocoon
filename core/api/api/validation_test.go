package api

import (
	"testing"

	"github.com/ellcrys/util"
	"github.com/ellcrys/cocoon/core/scheduler"
	"github.com/ellcrys/cocoon/core/types"
	. "github.com/smartystreets/goconvey/convey"
)

func TestValidation(t *testing.T) {
	Convey("Validation", t, func() {
		Convey(".ValidateCocoon", func() {
			Convey("should return expected errors", func() {

				err := ValidateCocoon(&types.Cocoon{})
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "id: id is required")

				err = ValidateCocoon(&types.Cocoon{
					ID: "some id",
				})
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "id: id is not a valid resource name")

				err = ValidateCocoon(&types.Cocoon{
					ID: util.UUID4(),
				})
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "resources.memory: memory is required")

				err = ValidateCocoon(&types.Cocoon{
					ID:     util.UUID4(),
					Memory: 512,
				})
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "resources.cpuShare: CPU share is required")

				err = ValidateCocoon(&types.Cocoon{
					ID:       util.UUID4(),
					Memory:   100,
					CPUShare: 122,
				})
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "resources: Unknown resource set")

				err = ValidateCocoon(&types.Cocoon{
					ID:             util.UUID4(),
					Memory:         512,
					CPUShare:       100,
					NumSignatories: 1,
					SigThreshold:   1,
					Signatories:    []string{"id1", "id2"},
				})
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "signatories.signatories: max signatories already added. You can't add more")
			})
		})

		Convey(".ValidateRelease", func() {

			err := ValidateRelease(&types.Release{})
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, "id is required")

			err = ValidateRelease(&types.Release{
				ID: "some id",
			})
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, "id is not a valid uuid")

			err = ValidateRelease(&types.Release{
				ID: util.UUID4(),
			})
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, "cocoon id is required")

			err = ValidateRelease(&types.Release{
				ID:       util.UUID4(),
				CocoonID: "cocoon-123",
			})
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, "url is required")

			err = ValidateRelease(&types.Release{
				ID:       util.UUID4(),
				CocoonID: "cocoon-123",
				URL:      "http://google.com",
			})
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, "url is not a valid github repo url")

			err = ValidateRelease(&types.Release{
				ID:       util.UUID4(),
				CocoonID: "cocoon-123",
				URL:      "https://github.com/ellcrys/cocoon-example-01",
			})
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, "language is required")

			err = ValidateRelease(&types.Release{
				ID:       util.UUID4(),
				CocoonID: "cocoon-123",
				URL:      "https://github.com/ellcrys/cocoon-example-01",
				Language: "abc",
			})
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldContainSubstring, "language is not supported")

			err = ValidateRelease(&types.Release{
				ID:         util.UUID4(),
				CocoonID:   "cocoon-123",
				URL:        "https://github.com/ellcrys/cocoon-example-01",
				Language:   scheduler.SupportedCocoonCodeLang[0],
				BuildParam: "non json",
			})
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldContainSubstring, "build parameter is not valid json")
		})

		Convey(".ValidateFirewallRules", func() {
			_, errs := ValidateFirewallRules("")
			So(len(errs), ShouldEqual, 1)
			So(errs[0].Error(), ShouldEqual, "empty string passed")

			_, errs = ValidateFirewallRules(123)
			So(len(errs), ShouldEqual, 1)
			So(errs[0].Error(), ShouldEqual, "invalid type. expects a json string or a slice of map")

			_, errs = ValidateFirewallRules(`abc`)
			So(len(errs), ShouldEqual, 1)
			So(errs[0].Error(), ShouldEqual, "malformed json")

			_, errs = ValidateFirewallRules(`{}`)
			So(len(errs), ShouldEqual, 1)
			So(errs[0].Error(), ShouldEqual, "malformed json")

			_, errs = ValidateFirewallRules(`[{}]`)
			So(len(errs), ShouldEqual, 2)
			So(errs[0].Error(), ShouldEqual, "rule 0: destination is required")
			So(errs[1].Error(), ShouldEqual, "rule 0: port is required")

			_, errs = ValidateFirewallRules(`[{ "destination": "0.0.0.0.0" }]`)
			So(len(errs), ShouldEqual, 2)
			So(errs[0].Error(), ShouldEqual, "rule 0: destination is not a valid IP or host")
			So(errs[1].Error(), ShouldEqual, "rule 0: port is required")

			_, errs = ValidateFirewallRules(`[{ "destination": "http://google.com" }]`)
			So(len(errs), ShouldEqual, 2)
			So(errs[0].Error(), ShouldEqual, "rule 0: destination is not a valid IP or host")
			So(errs[1].Error(), ShouldEqual, "rule 0: port is required")

			_, errs = ValidateFirewallRules(`[{ "destination": "google.com", "protocol": "icmp" }]`)
			So(len(errs), ShouldEqual, 2)
			So(errs[0].Error(), ShouldEqual, "rule 0: port is required")
			So(errs[1].Error(), ShouldEqual, "rule 0: invalid protocol")

			_, errs = ValidateFirewallRules(`[{ "destination": "0.0.0.0", "port": "3000" }]`)
			So(len(errs), ShouldEqual, 0)
		})

		Convey(".ValidateEnvVariables", func() {
			errs := ValidateEnvVariables(map[string]string{
				"VAR$A": "value",
			})
			So(len(errs), ShouldEqual, 1)

			errs = ValidateEnvVariables(map[string]string{
				"VAR-B":    "value",
				"VAR@Flag": "value",
			})
			So(len(errs), ShouldEqual, 2)

			errs = ValidateEnvVariables(map[string]string{
				"VAR B":    "value",
				"VAR-B":    "value",
				"VAR@Flag": "value",
			})
			So(len(errs), ShouldEqual, 3)
		})
	})
}
