/* IBM Confidential
*  Object Code Only Source Materials
*  5747-SM3
*  (c) Copyright IBM Corp. 2017,2021
*
*  The source code for this program is not published or otherwise divested
*  of its trade secrets, irrespective of what has been deposited with the
*  U.S. Copyright Office. */

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccIBMISVirtualEndpointGatewayIP_Basic(t *testing.T) {
	t.Skip()
	var endpointGateway string
	name := "ibm_is_virtual_endpoint_gateway.virtual_endpoint_gateway"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		// No requirement for CheckDestory of this resource as by reaching this point it must have already been deleted from CIS.
		Steps: []resource.TestStep{
			{
				Config: testAccCheckisVirtualEndpointGatewayIPConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckisVirtualEndpointGatewayIPExists(name, &endpointGateway),
					resource.TestCheckResourceAttr(name, "reserved_ip_id", subnetID),
				),
			},
		},
	})
}

func TestAccIBMISVirtualEndpointGatewayIP_import(t *testing.T) {
	name := "ibm_is_virtual_endpoint_gateway.virtual_endpoint_gateway"
	t.Skip()
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckisVirtualEndpointGatewayIPDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckisVirtualEndpointGatewayIPConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "ip_id", subnetID),
				),
			},
			{
				ResourceName:      name,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIBMISVirtualEndpointGatewayIP_CreateAfterManualDestroy(t *testing.T) {
	var monitorOne, monitorTwo string
	name := "ibm_is_virtual_endpoint_gateway.virtual_endpoint_gateway"
	t.Skip()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckisVirtualEndpointGatewayIPDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckisVirtualEndpointGatewayIPConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckisVirtualEndpointGatewayIPExists(name, &monitorOne),
					testAccisVirtualEndpointGatewayIPManuallyDelete(&monitorOne),
				),
			},
			{
				Config: testAccCheckisVirtualEndpointGatewayIPConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckisVirtualEndpointGatewayIPExists(name, &monitorTwo),
					func(state *terraform.State) error {
						if monitorOne == monitorTwo {
							return fmt.Errorf("load balancer monitor id is unchanged even after we thought we deleted it ( %s )",
								monitorTwo)
						}
						return nil
					},
				),
			},
		},
	})
}

func testAccisVirtualEndpointGatewayIPManuallyDelete(tfEndpointGwIPID *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		sess, err := testAccProvider.Meta().(ClientSession).VpcV1API()
		if err != nil {
			return err
		}
		tfEndpointGwIP := *tfEndpointGwIPID
		parts, err := idParts(tfEndpointGwIP)
		if err != nil {
			return err
		}
		gatewayID := parts[0]
		ipID := parts[1]
		opt := sess.NewRemoveEndpointGatewayIPOptions(gatewayID, ipID)
		response, err := sess.RemoveEndpointGatewayIP(opt)
		if err != nil {
			return fmt.Errorf("Delete Endpoint Gateway IP failed: %v", response)
		}
		return nil
	}
}

func testAccCheckisVirtualEndpointGatewayIPDestroy(s *terraform.State) error {
	sess, err := testAccProvider.Meta().(ClientSession).VpcV1API()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_virtual_endpoint_gateway" {
			continue
		}
		parts, err := idParts(rs.Primary.ID)
		if err != nil {
			return err
		}
		gatewayID := parts[0]
		ipID := parts[1]
		opt := sess.NewGetEndpointGatewayIPOptions(gatewayID, ipID)
		_, response, err := sess.GetEndpointGatewayIP(opt)
		if err == nil {
			return fmt.Errorf("Endpoint Gateway IP still exists: %v", response)
		}
	}

	return nil
}

func testAccCheckisVirtualEndpointGatewayIPExists(n string, tfEndpointGwIPID *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No endpoint gateway ID is set")
		}

		sess, err := testAccProvider.Meta().(ClientSession).VpcV1API()
		if err != nil {
			return err
		}

		parts, err := idParts(rs.Primary.ID)
		if err != nil {
			return err
		}
		gatewayID := parts[0]
		ipID := parts[1]
		opt := sess.NewGetEndpointGatewayIPOptions(gatewayID, ipID)
		_, response, err := sess.GetEndpointGatewayIP(opt)
		if err != nil {
			return fmt.Errorf("Endpoint Gateway IP does not exist: %s", response)
		}
		*tfEndpointGwIPID = fmt.Sprintf("%s/%s", gatewayID, ipID)
		return nil
	}
}

func testAccCheckisVirtualEndpointGatewayIPConfigBasic() string {
	vpcname1 := fmt.Sprintf("tfvpngw-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname1 := fmt.Sprintf("tfvpngw-subnet-%d", acctest.RandIntRange(10, 100))
	name1 := fmt.Sprintf("tfvpngw-createname-%d", acctest.RandIntRange(10, 100))
	return testAccCheckisVirtualEndpointGatewayConfigBasic(vpcname1, subnetname1, name1) + fmt.Sprintf(`
	resource "ibm_is_virtual_endpoint_gateway_ip" "virtual_endpoint_gateway_ip" {
		gateway = ibm_is_virtual_endpoint_gateway.endpoint_gateway.id
		reserved_ip = "%[1]s"
	  }
	`, subnetID)
}
