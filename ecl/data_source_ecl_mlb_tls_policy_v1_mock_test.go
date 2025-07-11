package ecl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"

	"github.com/nttcom/terraform-provider-ecl/ecl/testhelper/mock"
)

func TestMockedAccMLBV1TLSPolicyDataSource(t *testing.T) {
	mc := mock.NewMockController()
	defer mc.TerminateMockControllerSafety()

	postKeystone := fmt.Sprintf(fakeKeystonePostTmpl, mc.Endpoint(), OS_REGION_NAME)

	mc.Register(t, "keystone", "/v3/auth/tokens", postKeystone)
	mc.Register(t, "tls_policies", "/v1.0/tls_policies", testMockMLBV1TLSPoliciesListNameQuery)
	mc.Register(t, "tls_policies", "/v1.0/tls_policies", testMockMLBV1TLSPoliciesListDescriptionQuery)
	mc.Register(t, "tls_policies", "/v1.0/tls_policies", testMockMLBV1TLSPoliciesListDefaultQuery)

	mc.StartServer(t)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccMLBV1TLSPolicyDataSourceQueryName,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ecl_mlb_tls_policy_v1.tls_policy_1", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("data.ecl_mlb_tls_policy_v1.tls_policy_1", "name", "TLSv1.2_202210_01"),
					resource.TestCheckResourceAttr("data.ecl_mlb_tls_policy_v1.tls_policy_1", "description", "description"),
					resource.TestCheckResourceAttr("data.ecl_mlb_tls_policy_v1.tls_policy_1", "default", "true"),
					resource.TestCheckResourceAttr("data.ecl_mlb_tls_policy_v1.tls_policy_1", "tls_protocols.#", "2"),
					resource.TestCheckResourceAttr("data.ecl_mlb_tls_policy_v1.tls_policy_1", "tls_protocols.0", "TLSv1.2"),
					resource.TestCheckResourceAttr("data.ecl_mlb_tls_policy_v1.tls_policy_1", "tls_protocols.1", "TLSv1.3"),
					resource.TestCheckResourceAttr("data.ecl_mlb_tls_policy_v1.tls_policy_1", "tls_ciphers.#", "9"),
					resource.TestCheckResourceAttr("data.ecl_mlb_tls_policy_v1.tls_policy_1", "tls_ciphers.0", "TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256"),
					resource.TestCheckResourceAttr("data.ecl_mlb_tls_policy_v1.tls_policy_1", "tls_ciphers.1", "TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384"),
					resource.TestCheckResourceAttr("data.ecl_mlb_tls_policy_v1.tls_policy_1", "tls_ciphers.2", "TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256"),
					resource.TestCheckResourceAttr("data.ecl_mlb_tls_policy_v1.tls_policy_1", "tls_ciphers.3", "TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384"),
					resource.TestCheckResourceAttr("data.ecl_mlb_tls_policy_v1.tls_policy_1", "tls_ciphers.4", "TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256"),
					resource.TestCheckResourceAttr("data.ecl_mlb_tls_policy_v1.tls_policy_1", "tls_ciphers.5", "TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256"),
					resource.TestCheckResourceAttr("data.ecl_mlb_tls_policy_v1.tls_policy_1", "tls_ciphers.6", "TLS_AES_256_GCM_SHA384"),
					resource.TestCheckResourceAttr("data.ecl_mlb_tls_policy_v1.tls_policy_1", "tls_ciphers.7", "TLS_CHACHA20_POLY1305_SHA256"),
					resource.TestCheckResourceAttr("data.ecl_mlb_tls_policy_v1.tls_policy_1", "tls_ciphers.8", "TLS_AES_128_GCM_SHA256"),
				),
			},
			{
				Config: testAccMLBV1TLSPolicyDataSourceQueryDescription,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ecl_mlb_tls_policy_v1.tls_policy_1", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("data.ecl_mlb_tls_policy_v1.tls_policy_1", "name", "TLSv1.2_202210_01"),
					resource.TestCheckResourceAttr("data.ecl_mlb_tls_policy_v1.tls_policy_1", "description", "description"),
					resource.TestCheckResourceAttr("data.ecl_mlb_tls_policy_v1.tls_policy_1", "default", "true"),
					resource.TestCheckResourceAttr("data.ecl_mlb_tls_policy_v1.tls_policy_1", "tls_protocols.#", "2"),
					resource.TestCheckResourceAttr("data.ecl_mlb_tls_policy_v1.tls_policy_1", "tls_protocols.0", "TLSv1.2"),
					resource.TestCheckResourceAttr("data.ecl_mlb_tls_policy_v1.tls_policy_1", "tls_protocols.1", "TLSv1.3"),
					resource.TestCheckResourceAttr("data.ecl_mlb_tls_policy_v1.tls_policy_1", "tls_ciphers.#", "9"),
					resource.TestCheckResourceAttr("data.ecl_mlb_tls_policy_v1.tls_policy_1", "tls_ciphers.0", "TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256"),
					resource.TestCheckResourceAttr("data.ecl_mlb_tls_policy_v1.tls_policy_1", "tls_ciphers.1", "TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384"),
					resource.TestCheckResourceAttr("data.ecl_mlb_tls_policy_v1.tls_policy_1", "tls_ciphers.2", "TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256"),
					resource.TestCheckResourceAttr("data.ecl_mlb_tls_policy_v1.tls_policy_1", "tls_ciphers.3", "TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384"),
					resource.TestCheckResourceAttr("data.ecl_mlb_tls_policy_v1.tls_policy_1", "tls_ciphers.4", "TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256"),
					resource.TestCheckResourceAttr("data.ecl_mlb_tls_policy_v1.tls_policy_1", "tls_ciphers.5", "TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256"),
					resource.TestCheckResourceAttr("data.ecl_mlb_tls_policy_v1.tls_policy_1", "tls_ciphers.6", "TLS_AES_256_GCM_SHA384"),
					resource.TestCheckResourceAttr("data.ecl_mlb_tls_policy_v1.tls_policy_1", "tls_ciphers.7", "TLS_CHACHA20_POLY1305_SHA256"),
					resource.TestCheckResourceAttr("data.ecl_mlb_tls_policy_v1.tls_policy_1", "tls_ciphers.8", "TLS_AES_128_GCM_SHA256"),
				),
			},
			{
				Config: testAccMLBV1TLSPolicyDataSourceQueryDefault,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ecl_mlb_tls_policy_v1.tls_policy_1", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("data.ecl_mlb_tls_policy_v1.tls_policy_1", "name", "TLSv1.2_202210_01"),
					resource.TestCheckResourceAttr("data.ecl_mlb_tls_policy_v1.tls_policy_1", "description", "description"),
					resource.TestCheckResourceAttr("data.ecl_mlb_tls_policy_v1.tls_policy_1", "default", "true"),
					resource.TestCheckResourceAttr("data.ecl_mlb_tls_policy_v1.tls_policy_1", "tls_protocols.#", "2"),
					resource.TestCheckResourceAttr("data.ecl_mlb_tls_policy_v1.tls_policy_1", "tls_protocols.0", "TLSv1.2"),
					resource.TestCheckResourceAttr("data.ecl_mlb_tls_policy_v1.tls_policy_1", "tls_protocols.1", "TLSv1.3"),
					resource.TestCheckResourceAttr("data.ecl_mlb_tls_policy_v1.tls_policy_1", "tls_ciphers.#", "9"),
					resource.TestCheckResourceAttr("data.ecl_mlb_tls_policy_v1.tls_policy_1", "tls_ciphers.0", "TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256"),
					resource.TestCheckResourceAttr("data.ecl_mlb_tls_policy_v1.tls_policy_1", "tls_ciphers.1", "TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384"),
					resource.TestCheckResourceAttr("data.ecl_mlb_tls_policy_v1.tls_policy_1", "tls_ciphers.2", "TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256"),
					resource.TestCheckResourceAttr("data.ecl_mlb_tls_policy_v1.tls_policy_1", "tls_ciphers.3", "TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384"),
					resource.TestCheckResourceAttr("data.ecl_mlb_tls_policy_v1.tls_policy_1", "tls_ciphers.4", "TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256"),
					resource.TestCheckResourceAttr("data.ecl_mlb_tls_policy_v1.tls_policy_1", "tls_ciphers.5", "TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256"),
					resource.TestCheckResourceAttr("data.ecl_mlb_tls_policy_v1.tls_policy_1", "tls_ciphers.6", "TLS_AES_256_GCM_SHA384"),
					resource.TestCheckResourceAttr("data.ecl_mlb_tls_policy_v1.tls_policy_1", "tls_ciphers.7", "TLS_CHACHA20_POLY1305_SHA256"),
					resource.TestCheckResourceAttr("data.ecl_mlb_tls_policy_v1.tls_policy_1", "tls_ciphers.8", "TLS_AES_128_GCM_SHA256"),
				),
			},
		},
	})
}

var testAccMLBV1TLSPolicyDataSourceQueryName = fmt.Sprintf(`
data "ecl_mlb_tls_policy_v1" "tls_policy_1" {
  name = "TLSv1.2_202210_01"
}
`)

var testMockMLBV1TLSPoliciesListNameQuery = fmt.Sprintf(`
request:
  method: GET
  query:
    name:
      - TLSv1.2_202210_01
response:
  code: 200
  body: >
    {
      "tls_policies": [
        {
          "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
          "name": "TLSv1.2_202210_01",
          "description": "description",
          "default": true,
          "tls_protocols": [
            "TLSv1.2",
            "TLSv1.3"
          ],
          "tls_ciphers": [
            "TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256",
            "TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384",
            "TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256",
            "TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384",
            "TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256",
            "TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256",
            "TLS_AES_256_GCM_SHA384",
            "TLS_CHACHA20_POLY1305_SHA256",
            "TLS_AES_128_GCM_SHA256"
          ]
        }
      ]
    }
`)

var testAccMLBV1TLSPolicyDataSourceQueryDescription = fmt.Sprintf(`
data "ecl_mlb_tls_policy_v1" "tls_policy_1" {
  description = "description"
}
`)

var testMockMLBV1TLSPoliciesListDescriptionQuery = fmt.Sprintf(`
request:
  method: GET
  query:
    description:
      - description
response:
  code: 200
  body: >
    {
      "tls_policies": [
        {
          "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
          "name": "TLSv1.2_202210_01",
          "description": "description",
          "default": true,
          "tls_protocols": [
            "TLSv1.2",
            "TLSv1.3"
          ],
          "tls_ciphers": [
            "TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256",
            "TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384",
            "TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256",
            "TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384",
            "TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256",
            "TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256",
            "TLS_AES_256_GCM_SHA384",
            "TLS_CHACHA20_POLY1305_SHA256",
            "TLS_AES_128_GCM_SHA256"
          ]
        }
      ]
    }
`)

var testAccMLBV1TLSPolicyDataSourceQueryDefault = fmt.Sprintf(`
data "ecl_mlb_tls_policy_v1" "tls_policy_1" {
  default = "true"
}
`)

var testMockMLBV1TLSPoliciesListDefaultQuery = fmt.Sprintf(`
request:
  method: GET
  query:
    default:
      - true
response:
  code: 200
  body: >
    {
      "tls_policies": [
        {
          "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
          "name": "TLSv1.2_202210_01",
          "description": "description",
          "default": true,
          "tls_protocols": [
            "TLSv1.2",
            "TLSv1.3"
          ],
          "tls_ciphers": [
            "TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256",
            "TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384",
            "TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256",
            "TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384",
            "TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256",
            "TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256",
            "TLS_AES_256_GCM_SHA384",
            "TLS_CHACHA20_POLY1305_SHA256",
            "TLS_AES_128_GCM_SHA256"
          ]
        }
      ]
    }
`)
