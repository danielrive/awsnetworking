package awsnetworking

import (
	awsnetworking "github.com/danielrive/awsnetworking/vpc"
)

func Createnetworking(name string, CIDR string) {
	vpcinfo := awsnetworking.Vpcinfo{
		name,
		CIDR,
	}
	awsnetworking.Createvpc(vpcinfo)
}
