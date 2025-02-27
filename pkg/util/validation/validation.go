package validation

import (
	"fmt"
	"net/url"

	kubevalidation "k8s.io/apimachinery/pkg/util/validation"
)

const clusterNameMaxLength int = 48

// LabelValueMaxLength is a label's max length
const LabelValueMaxLength int = 63

// ValidateClusterName tests whether the cluster name passed is valid.
// If the cluster name is not valid, a list of error strings is returned. Otherwise an empty list (or nil) is returned.
// Rules of a valid cluster name:
// - Must be a valid label value as per RFC1123.
//   * An alphanumeric (a-z, and 0-9) string, with a maximum length of 63 characters,
//     with the '-' character allowed anywhere except the first or last character.
// - Length must be less than 48 characters.
//   * Since cluster name used to generate execution namespace by adding a prefix, so reserve 15 characters for the prefix.
func ValidateClusterName(name string) []string {
	if len(name) > clusterNameMaxLength {
		return []string{fmt.Sprintf("must be no more than %d characters", clusterNameMaxLength)}
	}

	return kubevalidation.IsDNS1123Label(name)
}

// ValidateClusterProxyURL tests whether the proxyURL is valid.
// If not valid, a list of error string is returned. Otherwise an empty list (or nil) is returned.
func ValidateClusterProxyURL(proxyURL string) []string {
	u, err := url.Parse(proxyURL)
	if err != nil {
		return []string{fmt.Sprintf("cloud not parse: %s, %v", proxyURL, err)}
	}

	switch u.Scheme {
	case "http", "https", "socks5":
	default:
		return []string{fmt.Sprintf("unsupported scheme %q, must be http, https, or socks5", u.Scheme)}
	}

	return nil
}
