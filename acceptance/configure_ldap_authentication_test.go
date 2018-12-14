package acceptance

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"os/exec"

	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pivotal-cf/om/api"
)

var _ = Describe("configure-ldap-authentication command", func() {
	It("configures the admin user account on OpsManager with LDAP", func() {
		var auth struct {
			Setup api.SetupInput `json:"setup"`
		}
		var ensureAvailabilityCallCount int

		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			switch req.URL.Path {
			case "/api/v0/setup":
				err := json.NewDecoder(req.Body).Decode(&auth)
				Expect(err).NotTo(HaveOccurred())
			case "/login/ensure_availability":
				ensureAvailabilityCallCount++

				if ensureAvailabilityCallCount == 1 {
					w.Header().Set("Location", "/setup")
					w.WriteHeader(http.StatusFound)
					return
				}

				if ensureAvailabilityCallCount < 3 {
					w.WriteHeader(http.StatusOK)
					w.Write([]byte("Waiting for authentication system to start..."))
					return
				}

				w.Header().Set("Location", "/auth/cloudfoundry")
				w.WriteHeader(http.StatusFound)
			default:
				out, err := httputil.DumpRequest(req, true)
				Expect(err).NotTo(HaveOccurred())
				Fail(fmt.Sprintf("unexpected request: %s", out))
			}
		}))

		defer server.Close()

		command := exec.Command(pathToMain,
			"--target", server.URL,
			"--skip-ssl-validation",
			"configure-ldap-authentication",
			"--decryption-passphrase", "some-passphrase",
			"--server-url", "ldap://YOUR-LDAP-SERVER",
			"--ldap-username", "cn=admin,dc=opsmanager,dc=com",
			"--ldap-password", "password",
			"--user-search-base", "ou=users,dc=opsmanager,dc=com",
			"--user-search-filter", "cn={0}",
			"--group-search-base", "ou=groups,dc=opsmanager,dc=com",
			"--group-search-filter", "member={0}",
			"--ldap-rbac-admin-group-name", "cn=opsmgradmins,ou=groups,dc=opsmanager,dc=com",
			"--email-attribute", "mail",
			"--ldap-referrals", "follow",
		)

		session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
		Expect(err).NotTo(HaveOccurred())

		Eventually(session, "5s").Should(gexec.Exit(0))

		Expect(auth.Setup).To(Equal(api.SetupInput{
			IdentityProvider:                 "ldap",
			DecryptionPassphrase:             "some-passphrase",
			DecryptionPassphraseConfirmation: "some-passphrase",
			EULAAccepted:                     "true",
			LDAPSettings: &api.LDAPSettings{
				EmailAttribute:     "mail",
				GroupSearchBase:    "ou=groups,dc=opsmanager,dc=com",
				GroupSearchFilter:  "member={0}",
				LDAPPassword:       "password",
				LDAPRBACAdminGroup: "cn=opsmgradmins,ou=groups,dc=opsmanager,dc=com",
				LDAPReferral:       "follow",
				LDAPUsername:       "cn=admin,dc=opsmanager,dc=com",
				ServerURL:          "ldap://YOUR-LDAP-SERVER",
				UserSearchBase:     "ou=users,dc=opsmanager,dc=com",
				UserSearchFilter:   "cn={0}",
			},
		}))

		Expect(ensureAvailabilityCallCount).To(Equal(3))

		Expect(session.Out).To(gbytes.Say("configuring LDAP authentication..."))
		Expect(session.Out).To(gbytes.Say("waiting for configuration to complete..."))
		Expect(session.Out).To(gbytes.Say("configuration complete"))
	})
})
