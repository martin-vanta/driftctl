package middlewares

import (
	"reflect"
	"testing"

	awssdk "github.com/aws/aws-sdk-go/aws"
	"github.com/cloudskiff/driftctl/pkg/resource"
	"github.com/cloudskiff/driftctl/pkg/resource/aws"
)

func TestAwsInstanceEIP_Execute(t *testing.T) {
	type args struct {
		remoteResources    *[]resource.Resource
		resourcesFromState *[]resource.Resource
	}
	tests := []struct {
		name     string
		args     args
		expected args
	}{
		{
			name: "test that public ip and dns are nilled with eip",
			args: args{
				remoteResources: &[]resource.Resource{
					&resource.AbstractResource{
						Id:   "instance1",
						Type: "aws_instance",
						Attrs: &resource.Attributes{
							"public_ip":  "1.2.3.4",
							"public_dns": "dns-of-eip.com",
						},
					},
					&resource.AbstractResource{
						Id:   "instance2",
						Type: "aws_instance",
						Attrs: &resource.Attributes{
							"public_ip":  "1.2.3.4",
							"public_dns": "dns-of-eip.com",
						},
					},
				},
				resourcesFromState: &[]resource.Resource{
					&resource.AbstractResource{
						Id:   "instance1",
						Type: "aws_instance",
						Attrs: &resource.Attributes{
							"public_ip":  "5.6.7.8",
							"public_dns": "example.com",
						},
					},
					&aws.AwsEip{
						Instance: awssdk.String("instance1"),
					},
				},
			},
			expected: args{
				remoteResources: &[]resource.Resource{
					&resource.AbstractResource{
						Id:    "instance1",
						Type:  "aws_instance",
						Attrs: &resource.Attributes{},
					},
					&resource.AbstractResource{
						Id:   "instance2",
						Type: "aws_instance",
						Attrs: &resource.Attributes{
							"public_ip":  "1.2.3.4",
							"public_dns": "dns-of-eip.com",
						},
					},
				},
				resourcesFromState: &[]resource.Resource{
					&resource.AbstractResource{
						Id:    "instance1",
						Type:  "aws_instance",
						Attrs: &resource.Attributes{},
					},
					&aws.AwsEip{
						Instance: awssdk.String("instance1"),
					},
				},
			},
		},
		{
			name: "test that public ip and dns are nilled when eip association",
			args: args{
				remoteResources: &[]resource.Resource{
					&resource.AbstractResource{
						Id:   "instance1",
						Type: "aws_instance",
						Attrs: &resource.Attributes{
							"public_ip":  "1.2.3.4",
							"public_dns": "dns-of-eip.com",
						},
					},
					&resource.AbstractResource{
						Id:   "instance2",
						Type: "aws_instance",
						Attrs: &resource.Attributes{
							"public_ip":  "1.2.3.4",
							"public_dns": "dns-of-eip.com",
						},
					},
				},
				resourcesFromState: &[]resource.Resource{
					&resource.AbstractResource{
						Id:   "instance1",
						Type: "aws_instance",
						Attrs: &resource.Attributes{
							"public_ip":  "5.6.7.8",
							"public_dns": "example.com",
						},
					},
					&aws.AwsEipAssociation{
						InstanceId: awssdk.String("instance1"),
					},
				},
			},
			expected: args{
				remoteResources: &[]resource.Resource{
					&resource.AbstractResource{
						Id:    "instance1",
						Type:  "aws_instance",
						Attrs: &resource.Attributes{},
					},
					&resource.AbstractResource{
						Id:   "instance2",
						Type: "aws_instance",
						Attrs: &resource.Attributes{
							"public_ip":  "1.2.3.4",
							"public_dns": "dns-of-eip.com",
						},
					},
				},
				resourcesFromState: &[]resource.Resource{
					&resource.AbstractResource{
						Id:    "instance1",
						Type:  "aws_instance",
						Attrs: &resource.Attributes{},
					},
					&aws.AwsEipAssociation{
						InstanceId: awssdk.String("instance1"),
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := AwsInstanceEIP{}
			if err := a.Execute(tt.args.remoteResources, tt.args.resourcesFromState); err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(tt.args, tt.expected) {
				t.Fatalf("Expected results mismatch")
			}
		})
	}
}
