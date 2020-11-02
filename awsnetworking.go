package awsnetworking

import (
	awsnetworkingvpc "github.com/danielrive/awsnetworking/vpc"
)

func Createnetworking(name string, CIDR string) {
	vpcinfo := awsnetworkingvpc.Vpcinfo{
		name,
		CIDR,
	}
	awsnetworkingvpc.Createvpc(vpcinfo)
}
