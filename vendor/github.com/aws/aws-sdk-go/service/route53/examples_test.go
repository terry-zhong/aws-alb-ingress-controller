// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package route53_test

import (
	"fmt"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/route53"
)

var _ time.Duration
var _ strings.Reader
var _ aws.Config

func parseTime(layout, value string) *time.Time {
	t, err := time.Parse(layout, value)
	if err != nil {
		panic(err)
	}
	return &t
}

// To associate a VPC with a hosted zone
//
// The following example associates the VPC with ID vpc-1a2b3c4d with the hosted zone
// with ID Z3M3LMPEXAMPLE.
func ExampleRoute53_AssociateVPCWithHostedZone_shared00() {
	svc := route53.New(session.New())
	input := &route53.AssociateVPCWithHostedZoneInput{
		Comment:      aws.String(""),
		HostedZoneId: aws.String("Z3M3LMPEXAMPLE"),
		VPC: &route53.VPC{
			VPCId:     aws.String("vpc-1a2b3c4d"),
			VPCRegion: aws.String("us-east-2"),
		},
	}

	result, err := svc.AssociateVPCWithHostedZone(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case route53.ErrCodeNoSuchHostedZone:
				fmt.Println(route53.ErrCodeNoSuchHostedZone, aerr.Error())
			case route53.ErrCodeNotAuthorizedException:
				fmt.Println(route53.ErrCodeNotAuthorizedException, aerr.Error())
			case route53.ErrCodeInvalidVPCId:
				fmt.Println(route53.ErrCodeInvalidVPCId, aerr.Error())
			case route53.ErrCodeInvalidInput:
				fmt.Println(route53.ErrCodeInvalidInput, aerr.Error())
			case route53.ErrCodePublicZoneVPCAssociation:
				fmt.Println(route53.ErrCodePublicZoneVPCAssociation, aerr.Error())
			case route53.ErrCodeConflictingDomainExists:
				fmt.Println(route53.ErrCodeConflictingDomainExists, aerr.Error())
			case route53.ErrCodeLimitsExceeded:
				fmt.Println(route53.ErrCodeLimitsExceeded, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To create a basic resource record set
//
// The following example creates a resource record set that routes Internet traffic
// to a resource with an IP address of 192.0.2.44.
func ExampleRoute53_ChangeResourceRecordSets_shared00() {
	svc := route53.New(session.New())
	input := &route53.ChangeResourceRecordSetsInput{
		ChangeBatch: &route53.ChangeBatch{
			Changes: []*route53.Change{
				{
					Action: aws.String("CREATE"),
				},
			},
			Comment: aws.String("Web server for example.com"),
		},
		HostedZoneId: aws.String("Z3M3LMPEXAMPLE"),
	}

	result, err := svc.ChangeResourceRecordSets(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case route53.ErrCodeNoSuchHostedZone:
				fmt.Println(route53.ErrCodeNoSuchHostedZone, aerr.Error())
			case route53.ErrCodeNoSuchHealthCheck:
				fmt.Println(route53.ErrCodeNoSuchHealthCheck, aerr.Error())
			case route53.ErrCodeInvalidChangeBatch:
				fmt.Println(route53.ErrCodeInvalidChangeBatch, aerr.Error())
			case route53.ErrCodeInvalidInput:
				fmt.Println(route53.ErrCodeInvalidInput, aerr.Error())
			case route53.ErrCodePriorRequestNotComplete:
				fmt.Println(route53.ErrCodePriorRequestNotComplete, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To create weighted resource record sets
//
// The following example creates two weighted resource record sets. The resource with
// a Weight of 100 will get 1/3rd of traffic (100/100+200), and the other resource will
// get the rest of the traffic for example.com.
func ExampleRoute53_ChangeResourceRecordSets_shared01() {
	svc := route53.New(session.New())
	input := &route53.ChangeResourceRecordSetsInput{
		ChangeBatch: &route53.ChangeBatch{
			Changes: []*route53.Change{
				{
					Action: aws.String("CREATE"),
				},
				{
					Action: aws.String("CREATE"),
				},
			},
			Comment: aws.String("Web servers for example.com"),
		},
		HostedZoneId: aws.String("Z3M3LMPEXAMPLE"),
	}

	result, err := svc.ChangeResourceRecordSets(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case route53.ErrCodeNoSuchHostedZone:
				fmt.Println(route53.ErrCodeNoSuchHostedZone, aerr.Error())
			case route53.ErrCodeNoSuchHealthCheck:
				fmt.Println(route53.ErrCodeNoSuchHealthCheck, aerr.Error())
			case route53.ErrCodeInvalidChangeBatch:
				fmt.Println(route53.ErrCodeInvalidChangeBatch, aerr.Error())
			case route53.ErrCodeInvalidInput:
				fmt.Println(route53.ErrCodeInvalidInput, aerr.Error())
			case route53.ErrCodePriorRequestNotComplete:
				fmt.Println(route53.ErrCodePriorRequestNotComplete, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To create an alias resource record set
//
// The following example creates an alias resource record set that routes traffic to
// a CloudFront distribution.
func ExampleRoute53_ChangeResourceRecordSets_shared02() {
	svc := route53.New(session.New())
	input := &route53.ChangeResourceRecordSetsInput{
		ChangeBatch: &route53.ChangeBatch{
			Changes: []*route53.Change{
				{
					Action: aws.String("CREATE"),
				},
			},
			Comment: aws.String("CloudFront distribution for example.com"),
		},
		HostedZoneId: aws.String("Z3M3LMPEXAMPLE"),
	}

	result, err := svc.ChangeResourceRecordSets(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case route53.ErrCodeNoSuchHostedZone:
				fmt.Println(route53.ErrCodeNoSuchHostedZone, aerr.Error())
			case route53.ErrCodeNoSuchHealthCheck:
				fmt.Println(route53.ErrCodeNoSuchHealthCheck, aerr.Error())
			case route53.ErrCodeInvalidChangeBatch:
				fmt.Println(route53.ErrCodeInvalidChangeBatch, aerr.Error())
			case route53.ErrCodeInvalidInput:
				fmt.Println(route53.ErrCodeInvalidInput, aerr.Error())
			case route53.ErrCodePriorRequestNotComplete:
				fmt.Println(route53.ErrCodePriorRequestNotComplete, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To create weighted alias resource record sets
//
// The following example creates two weighted alias resource record sets that route
// traffic to ELB load balancers. The resource with a Weight of 100 will get 1/3rd of
// traffic (100/100+200), and the other resource will get the rest of the traffic for
// example.com.
func ExampleRoute53_ChangeResourceRecordSets_shared03() {
	svc := route53.New(session.New())
	input := &route53.ChangeResourceRecordSetsInput{
		ChangeBatch: &route53.ChangeBatch{
			Changes: []*route53.Change{
				{
					Action: aws.String("CREATE"),
				},
				{
					Action: aws.String("CREATE"),
				},
			},
			Comment: aws.String("ELB load balancers for example.com"),
		},
		HostedZoneId: aws.String("Z3M3LMPEXAMPLE"),
	}

	result, err := svc.ChangeResourceRecordSets(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case route53.ErrCodeNoSuchHostedZone:
				fmt.Println(route53.ErrCodeNoSuchHostedZone, aerr.Error())
			case route53.ErrCodeNoSuchHealthCheck:
				fmt.Println(route53.ErrCodeNoSuchHealthCheck, aerr.Error())
			case route53.ErrCodeInvalidChangeBatch:
				fmt.Println(route53.ErrCodeInvalidChangeBatch, aerr.Error())
			case route53.ErrCodeInvalidInput:
				fmt.Println(route53.ErrCodeInvalidInput, aerr.Error())
			case route53.ErrCodePriorRequestNotComplete:
				fmt.Println(route53.ErrCodePriorRequestNotComplete, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To create latency resource record sets
//
// The following example creates two latency resource record sets that route traffic
// to EC2 instances. Traffic for example.com is routed either to the Ohio region or
// the Oregon region, depending on the latency between the user and those regions.
func ExampleRoute53_ChangeResourceRecordSets_shared04() {
	svc := route53.New(session.New())
	input := &route53.ChangeResourceRecordSetsInput{
		ChangeBatch: &route53.ChangeBatch{
			Changes: []*route53.Change{
				{
					Action: aws.String("CREATE"),
				},
				{
					Action: aws.String("CREATE"),
				},
			},
			Comment: aws.String("EC2 instances for example.com"),
		},
		HostedZoneId: aws.String("Z3M3LMPEXAMPLE"),
	}

	result, err := svc.ChangeResourceRecordSets(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case route53.ErrCodeNoSuchHostedZone:
				fmt.Println(route53.ErrCodeNoSuchHostedZone, aerr.Error())
			case route53.ErrCodeNoSuchHealthCheck:
				fmt.Println(route53.ErrCodeNoSuchHealthCheck, aerr.Error())
			case route53.ErrCodeInvalidChangeBatch:
				fmt.Println(route53.ErrCodeInvalidChangeBatch, aerr.Error())
			case route53.ErrCodeInvalidInput:
				fmt.Println(route53.ErrCodeInvalidInput, aerr.Error())
			case route53.ErrCodePriorRequestNotComplete:
				fmt.Println(route53.ErrCodePriorRequestNotComplete, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To create latency alias resource record sets
//
// The following example creates two latency alias resource record sets that route traffic
// for example.com to ELB load balancers. Requests are routed either to the Ohio region
// or the Oregon region, depending on the latency between the user and those regions.
func ExampleRoute53_ChangeResourceRecordSets_shared05() {
	svc := route53.New(session.New())
	input := &route53.ChangeResourceRecordSetsInput{
		ChangeBatch: &route53.ChangeBatch{
			Changes: []*route53.Change{
				{
					Action: aws.String("CREATE"),
				},
				{
					Action: aws.String("CREATE"),
				},
			},
			Comment: aws.String("ELB load balancers for example.com"),
		},
		HostedZoneId: aws.String("Z3M3LMPEXAMPLE"),
	}

	result, err := svc.ChangeResourceRecordSets(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case route53.ErrCodeNoSuchHostedZone:
				fmt.Println(route53.ErrCodeNoSuchHostedZone, aerr.Error())
			case route53.ErrCodeNoSuchHealthCheck:
				fmt.Println(route53.ErrCodeNoSuchHealthCheck, aerr.Error())
			case route53.ErrCodeInvalidChangeBatch:
				fmt.Println(route53.ErrCodeInvalidChangeBatch, aerr.Error())
			case route53.ErrCodeInvalidInput:
				fmt.Println(route53.ErrCodeInvalidInput, aerr.Error())
			case route53.ErrCodePriorRequestNotComplete:
				fmt.Println(route53.ErrCodePriorRequestNotComplete, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To create failover resource record sets
//
// The following example creates primary and secondary failover resource record sets
// that route traffic to EC2 instances. Traffic is generally routed to the primary resource,
// in the Ohio region. If that resource is unavailable, traffic is routed to the secondary
// resource, in the Oregon region.
func ExampleRoute53_ChangeResourceRecordSets_shared06() {
	svc := route53.New(session.New())
	input := &route53.ChangeResourceRecordSetsInput{
		ChangeBatch: &route53.ChangeBatch{
			Changes: []*route53.Change{
				{
					Action: aws.String("CREATE"),
				},
				{
					Action: aws.String("CREATE"),
				},
			},
			Comment: aws.String("Failover configuration for example.com"),
		},
		HostedZoneId: aws.String("Z3M3LMPEXAMPLE"),
	}

	result, err := svc.ChangeResourceRecordSets(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case route53.ErrCodeNoSuchHostedZone:
				fmt.Println(route53.ErrCodeNoSuchHostedZone, aerr.Error())
			case route53.ErrCodeNoSuchHealthCheck:
				fmt.Println(route53.ErrCodeNoSuchHealthCheck, aerr.Error())
			case route53.ErrCodeInvalidChangeBatch:
				fmt.Println(route53.ErrCodeInvalidChangeBatch, aerr.Error())
			case route53.ErrCodeInvalidInput:
				fmt.Println(route53.ErrCodeInvalidInput, aerr.Error())
			case route53.ErrCodePriorRequestNotComplete:
				fmt.Println(route53.ErrCodePriorRequestNotComplete, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To create failover alias resource record sets
//
// The following example creates primary and secondary failover alias resource record
// sets that route traffic to ELB load balancers. Traffic is generally routed to the
// primary resource, in the Ohio region. If that resource is unavailable, traffic is
// routed to the secondary resource, in the Oregon region.
func ExampleRoute53_ChangeResourceRecordSets_shared07() {
	svc := route53.New(session.New())
	input := &route53.ChangeResourceRecordSetsInput{
		ChangeBatch: &route53.ChangeBatch{
			Changes: []*route53.Change{
				{
					Action: aws.String("CREATE"),
				},
				{
					Action: aws.String("CREATE"),
				},
			},
			Comment: aws.String("Failover alias configuration for example.com"),
		},
		HostedZoneId: aws.String("Z3M3LMPEXAMPLE"),
	}

	result, err := svc.ChangeResourceRecordSets(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case route53.ErrCodeNoSuchHostedZone:
				fmt.Println(route53.ErrCodeNoSuchHostedZone, aerr.Error())
			case route53.ErrCodeNoSuchHealthCheck:
				fmt.Println(route53.ErrCodeNoSuchHealthCheck, aerr.Error())
			case route53.ErrCodeInvalidChangeBatch:
				fmt.Println(route53.ErrCodeInvalidChangeBatch, aerr.Error())
			case route53.ErrCodeInvalidInput:
				fmt.Println(route53.ErrCodeInvalidInput, aerr.Error())
			case route53.ErrCodePriorRequestNotComplete:
				fmt.Println(route53.ErrCodePriorRequestNotComplete, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To create geolocation resource record sets
//
// The following example creates four geolocation resource record sets that use IPv4
// addresses to route traffic to resources such as web servers running on EC2 instances.
// Traffic is routed to one of four IP addresses, for North America (NA), for South
// America (SA), for Europe (EU), and for all other locations (*).
func ExampleRoute53_ChangeResourceRecordSets_shared08() {
	svc := route53.New(session.New())
	input := &route53.ChangeResourceRecordSetsInput{
		ChangeBatch: &route53.ChangeBatch{
			Changes: []*route53.Change{
				{
					Action: aws.String("CREATE"),
				},
				{
					Action: aws.String("CREATE"),
				},
				{
					Action: aws.String("CREATE"),
				},
				{
					Action: aws.String("CREATE"),
				},
			},
			Comment: aws.String("Geolocation configuration for example.com"),
		},
		HostedZoneId: aws.String("Z3M3LMPEXAMPLE"),
	}

	result, err := svc.ChangeResourceRecordSets(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case route53.ErrCodeNoSuchHostedZone:
				fmt.Println(route53.ErrCodeNoSuchHostedZone, aerr.Error())
			case route53.ErrCodeNoSuchHealthCheck:
				fmt.Println(route53.ErrCodeNoSuchHealthCheck, aerr.Error())
			case route53.ErrCodeInvalidChangeBatch:
				fmt.Println(route53.ErrCodeInvalidChangeBatch, aerr.Error())
			case route53.ErrCodeInvalidInput:
				fmt.Println(route53.ErrCodeInvalidInput, aerr.Error())
			case route53.ErrCodePriorRequestNotComplete:
				fmt.Println(route53.ErrCodePriorRequestNotComplete, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To create geolocation alias resource record sets
//
// The following example creates four geolocation alias resource record sets that route
// traffic to ELB load balancers. Traffic is routed to one of four IP addresses, for
// North America (NA), for South America (SA), for Europe (EU), and for all other locations
// (*).
func ExampleRoute53_ChangeResourceRecordSets_shared09() {
	svc := route53.New(session.New())
	input := &route53.ChangeResourceRecordSetsInput{
		ChangeBatch: &route53.ChangeBatch{
			Changes: []*route53.Change{
				{
					Action: aws.String("CREATE"),
				},
				{
					Action: aws.String("CREATE"),
				},
				{
					Action: aws.String("CREATE"),
				},
				{
					Action: aws.String("CREATE"),
				},
			},
			Comment: aws.String("Geolocation alias configuration for example.com"),
		},
		HostedZoneId: aws.String("Z3M3LMPEXAMPLE"),
	}

	result, err := svc.ChangeResourceRecordSets(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case route53.ErrCodeNoSuchHostedZone:
				fmt.Println(route53.ErrCodeNoSuchHostedZone, aerr.Error())
			case route53.ErrCodeNoSuchHealthCheck:
				fmt.Println(route53.ErrCodeNoSuchHealthCheck, aerr.Error())
			case route53.ErrCodeInvalidChangeBatch:
				fmt.Println(route53.ErrCodeInvalidChangeBatch, aerr.Error())
			case route53.ErrCodeInvalidInput:
				fmt.Println(route53.ErrCodeInvalidInput, aerr.Error())
			case route53.ErrCodePriorRequestNotComplete:
				fmt.Println(route53.ErrCodePriorRequestNotComplete, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To add or remove tags from a hosted zone or health check
//
// The following example adds two tags and removes one tag from the hosted zone with
// ID Z3M3LMPEXAMPLE.
func ExampleRoute53_ChangeTagsForResource_shared00() {
	svc := route53.New(session.New())
	input := &route53.ChangeTagsForResourceInput{
		AddTags: []*route53.Tag{
			{
				Key:   aws.String("apex"),
				Value: aws.String("3874"),
			},
			{
				Key:   aws.String("acme"),
				Value: aws.String("4938"),
			},
		},
		RemoveTagKeys: []*string{
			aws.String("Nadir"),
		},
		ResourceId:   aws.String("Z3M3LMPEXAMPLE"),
		ResourceType: aws.String("hostedzone"),
	}

	result, err := svc.ChangeTagsForResource(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case route53.ErrCodeInvalidInput:
				fmt.Println(route53.ErrCodeInvalidInput, aerr.Error())
			case route53.ErrCodeNoSuchHealthCheck:
				fmt.Println(route53.ErrCodeNoSuchHealthCheck, aerr.Error())
			case route53.ErrCodeNoSuchHostedZone:
				fmt.Println(route53.ErrCodeNoSuchHostedZone, aerr.Error())
			case route53.ErrCodePriorRequestNotComplete:
				fmt.Println(route53.ErrCodePriorRequestNotComplete, aerr.Error())
			case route53.ErrCodeThrottlingException:
				fmt.Println(route53.ErrCodeThrottlingException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To get information about a hosted zone
//
// The following example gets information about the Z3M3LMPEXAMPLE hosted zone.
func ExampleRoute53_GetHostedZone_shared00() {
	svc := route53.New(session.New())
	input := &route53.GetHostedZoneInput{
		Id: aws.String("Z3M3LMPEXAMPLE"),
	}

	result, err := svc.GetHostedZone(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case route53.ErrCodeNoSuchHostedZone:
				fmt.Println(route53.ErrCodeNoSuchHostedZone, aerr.Error())
			case route53.ErrCodeInvalidInput:
				fmt.Println(route53.ErrCodeInvalidInput, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}