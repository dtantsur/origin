package oauth

import (
	g "github.com/onsi/ginkgo"
	o "github.com/onsi/gomega"

	admissionapi "k8s.io/pod-security-admission/api"

	exutil "github.com/openshift/origin/test/extended/util"
)

var _ = g.Describe("[sig-auth][Feature:LDAP] LDAP", func() {
	defer g.GinkgoRecover()
	var (
		oc = exutil.NewCLIWithPodSecurityLevel("oauth-ldap", admissionapi.LevelPrivileged)
	)

	g.It("should start an OpenLDAP test server", func() {
		_, _, _, _, err := exutil.CreateLDAPTestServer(oc)
		o.Expect(err).NotTo(o.HaveOccurred())
	})
})
