package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/route53"
)

var svc *route53.Route53
var comment = "Updated by lacrosse: github.com/lowply/lacrosse"

func create_new_service(profile string) *route53.Route53 {
	sess, err := session.NewSession(&aws.Config{
		Credentials: credentials.NewSharedCredentials("", profile),
	})
	if err != nil {
		return nil
	}
	return route53.New(sess)
}

func get_hosted_zone_id(domain string) (string, error) {
	params := &route53.ListHostedZonesInput{}
	resp, err := svc.ListHostedZones(params)
	if err != nil {
		return "", err
	}
	for _, v := range resp.HostedZones {
		if *v.Name == domain+"." {
			return strings.Split(*v.Id, "/")[2], nil
		}
	}
	return "", errors.New("Domain " + domain + " was not found in the hosted zone list.")
}

func check_status(change *route53.ChangeInfo) error {
	id := &route53.GetChangeInput{Id: change.Id}
	fmt.Println("Waiting for request " + *id.Id + " to be reflected...")
	err := svc.WaitUntilResourceRecordSetsChanged(id)
	if err != nil {
		return err
	}
	fmt.Println("Done.")
	return nil
}

func create_new_params(r *Request) (*route53.ChangeResourceRecordSetsInput, error) {
	if r.Type == "TXT" {
		r.Value = "\"" + r.Value + "\""
	}
	params := &route53.ChangeResourceRecordSetsInput{
		ChangeBatch: &route53.ChangeBatch{
			Changes: []*route53.Change{
				{
					Action: aws.String(r.Action),
					ResourceRecordSet: &route53.ResourceRecordSet{
						Name: aws.String(r.Domain),
						Type: aws.String(r.Type),
						ResourceRecords: []*route53.ResourceRecord{
							{
								Value: aws.String(r.Value),
							},
						},
						TTL: aws.Int64(r.TTL),
					},
				},
			},
			Comment: aws.String(comment),
		},
		HostedZoneId: aws.String(r.Id),
	}
	return params, nil
}
