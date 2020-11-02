package vpc

import (
	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type Vpcinfo struct {
	Name string
	CIDR string
}

func Createvpc(info Vpcinfo) {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ec2.NewVpc(ctx, info.Name, &ec2.VpcArgs{
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
