package vpc

import (
	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Basic information to create AWS VPC
type Vpcinfo struct {
	Name      string
	CIDR      string
	EnableNAT bool
}

// Create vpc with route public and private route tables, internet gateway and NAT GW (optional)
func Createvpc(info Vpcinfo) {
	pulumi.Run(func(ctx *pulumi.Context) error {
		vpc, errvpc := ec2.NewVpc(ctx, info.Name, &ec2.VpcArgs{
			CidrBlock:          pulumi.String(info.CIDR),
			EnableDnsHostnames: pulumi.BoolPtr(true),
			EnableDnsSupport:   pulumi.BoolPtr(true),
			Tags: pulumi.StringMap{
				"Name":        pulumi.String(info.Name),
				"CreatedBu4y": pulumi.String("Pulumi"),
			},
		})

		if errvpc != nil {
			return errvpc
		}

		//creating Internet GW
		IGW, errigw := ec2.NewInternetGateway(ctx, "IGW"+info.Name, &ec2.InternetGatewayArgs{
			VpcId: vpc.ID(),
			Tags: pulumi.StringMap{
				"Name":        pulumi.String("IGW" + info.Name),
				"CreatedBu4y": pulumi.String("Pulumi"),
			},
		})
		if errigw != nil {
			return errigw
		}

		// Creating Route tables, public and private

		_, errrt := ec2.NewRouteTable(ctx, "PublicRT"+info.Name, &ec2.RouteTableArgs{
			VpcId: vpc.ID(),
			Routes: ec2.RouteTableRouteArray{
				&ec2.RouteTableRouteArgs{
					CidrBlock: pulumi.String("0.0.0.0/0"),
					GatewayId: IGW.ID(),
				},
			},
			Tags: pulumi.StringMap{
				"Name":        pulumi.String("PublicRT" + info.Name),
				"CreatedBu4y": pulumi.String("Pulumi"),
			},
		})
		if errrt != nil {
			return errrt
		}
		return nil
	})
}
