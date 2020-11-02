package awsnetworking

import (
	"github.com/danielrive/awsnetworking/vpc"
)

func Createnetworking(name string, CIDR string) {
	vpcinfo := vpc.Vpcinfo{
		name,
		CIDR,
	}
	vpc.Createvpc(vpcinfo)
}
