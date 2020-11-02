package awsnetworking

import (
	"github.com/danielrive/awsnetworking"
)

func Createnetworking(name string, CIDR string) {
	vpcinfo := awsnetworking.Vpcinfo{
		name,
		CIDR,
	}
	awsnetworking.Createvpc(vpcinfo)
}
