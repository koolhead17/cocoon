package api

import (
	"fmt"

	"github.com/asaskevich/govalidator"
	"github.com/ellcrys/util"
	cocoon_util "github.com/ncodes/cocoon-util"
	"github.com/ncodes/cocoon/core/types"
)

var (
	// supportedLanguages list the languages supported
	supportedLanguages = []string{"go"}

	// supportedMemory list the memory values supported
	supportedMemory = []string{"512m", "1g", "2g"}

	// supportedCPUShares list the cpu share values supported
	supportedCPUShares = []string{"1x", "2x"}
)

// ValidateCocoon validates a cocoon to be created
func ValidateCocoon(c *types.Cocoon) error {

	if len(c.ID) == 0 {
		return fmt.Errorf("id is required")
	}
	if !govalidator.IsUUIDv4(c.ID) {
		return fmt.Errorf("id is not a valid uuid")
	}
	if len(c.URL) == 0 {
		return fmt.Errorf("url is required")
	}
	if !cocoon_util.IsGithubRepoURL(c.URL) {
		return fmt.Errorf("url is not a valid github repo url")
	}
	if len(c.Language) == 0 {
		return fmt.Errorf("language is required")
	}
	if !util.InStringSlice(supportedLanguages, c.Language) {
		return fmt.Errorf("language is not supported. Expects one of these values %s", supportedLanguages)
	}
	if len(c.BuildParam) > 0 {
		var _c map[string]interface{}
		if util.FromJSON([]byte(c.BuildParam), &_c) != nil {
			return fmt.Errorf("build parameter is not valid json")
		}
	}
	if len(c.Memory) == 0 {
		return fmt.Errorf("memory is required")
	}
	if !util.InStringSlice(supportedMemory, c.Memory) {
		return fmt.Errorf("Memory value is not supported. Expects one of these values %s", supportedMemory)
	}
	if len(c.CPUShares) == 0 {
		return fmt.Errorf("CPU share is required")
	}
	if !util.InStringSlice(supportedCPUShares, c.CPUShares) {
		return fmt.Errorf("CPU share value is not supported. Expects one of these values %s", supportedCPUShares)
	}
	if c.NumSignatories <= 0 {
		return fmt.Errorf("number of signatories cannot be less than 1")
	}
	if c.SigThreshold <= 0 {
		return fmt.Errorf("signatory threshold cannot be less than 1")
	}
	if c.NumSignatories < int32(len(c.Signatories)) {
		return fmt.Errorf("max signatories already added. You can't add more")
	}

	return nil
}

// ValidateRelease checks whether a release field values
// are valid.
func ValidateRelease(r *types.Release) error {
	if len(r.ID) == 0 {
		return fmt.Errorf("id is required")
	}
	if !govalidator.IsUUIDv4(r.ID) {
		return fmt.Errorf("id is not a valid uuid")
	}
	if len(r.CocoonID) == 0 {
		return fmt.Errorf("cocoon id is required")
	}
	if len(r.URL) == 0 {
		return fmt.Errorf("url is required")
	}
	if !cocoon_util.IsGithubRepoURL(r.URL) {
		return fmt.Errorf("url is not a valid github repo url")
	}
	if len(r.Language) == 0 {
		return fmt.Errorf("language is required")
	}
	if !util.InStringSlice(supportedLanguages, r.Language) {
		return fmt.Errorf("language is not supported. Expects one of these values %s", supportedLanguages)
	}
	if len(r.BuildParam) > 0 {
		var _r map[string]interface{}
		if util.FromJSON([]byte(r.BuildParam), &_r) != nil {
			return fmt.Errorf("build parameter is not valid json")
		}
	}
	return nil
}

// ValidateIdentity checks whether an identity field values
// are valid.
func ValidateIdentity(i *types.Identity) error {
	if len(i.Email) == 0 {
		return fmt.Errorf("email is required")
	} else if len(i.Password) == 0 {
		return fmt.Errorf("password is required")
	} else if len(i.Password) < 8 {
		return fmt.Errorf("password is too short. Minimum of 8 characters required")
	}
	return nil
}
