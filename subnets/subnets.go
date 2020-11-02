package subnets

import (
	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type subnetinfo struct {
	Name  string
	CIDR  string
	VpcId string
}

func Createsubnet(info subnetinfo) {

	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ec2.NewSubnet(ctx, info.Name, &ec2.SubnetArgs{
			VpcId:     pulumi.String(info.VpcId),
			CidrBlock: pulumi.String(info.CIDR),
			Tags: pulumi.StringMap{
				"Name":        pulumi.String(info.Name),
				"CreatedBu4y": pulumi.String("Pulumi"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
