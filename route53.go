package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path"
	"strings"

	"golang.org/x/net/publicsuffix"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/route53"
	"github.com/aws/aws-sdk-go/service/route53/route53iface"
)

var logpath = os.Getenv("HOME") + "/.cache/lacrosse.log"
var comment = "Updated by lacrosse: github.com/lowply/lacrosse"

type Route53 struct {
	Client route53iface.Route53API
	Id     string
	Req    *Request
}

func NewRoute53(args []string) (*Route53, error) {
	req, err := NewRequest(args[1:])
	if err != nil {
		return nil, err
	}

	config := &aws.Config{
		Credentials: credentials.NewSharedCredentials("", req.Profile),
	}

	sess, err := session.NewSession(config)
	if err != nil {
		return nil, err
	}

	r := &Route53{
		Client: route53.New(sess),
		Id:     "",
		Req:    req,
	}

	return r, nil
}

func (r *Route53) GetHostedZoneId(fqdn string) (string, error) {
	suffix, _ := publicsuffix.PublicSuffix(fqdn)
	name := strings.Split(strings.Replace(fqdn, "."+suffix, "", -1), ".")
	domain := name[len(name)-1] + "." + suffix

	params := &route53.ListHostedZonesInput{}
	resp, err := r.Client.ListHostedZones(params)
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

func (r *Route53) CheckStatus(resp *route53.ChangeResourceRecordSetsOutput) error {
	id := &route53.GetChangeInput{Id: resp.ChangeInfo.Id}
	fmt.Println("Waiting for request " + *id.Id + " to be reflected...")
	err := r.Client.WaitUntilResourceRecordSetsChanged(id)
	if err != nil {
		return err
	}
	fmt.Println("Done.")
	return nil
}

func (r *Route53) CreateNewParams() *route53.ChangeResourceRecordSetsInput {
	params := &route53.ChangeResourceRecordSetsInput{
		ChangeBatch: &route53.ChangeBatch{
			Changes: []*route53.Change{
				{
					Action: aws.String(r.Req.Action),
					ResourceRecordSet: &route53.ResourceRecordSet{
						Name: aws.String(r.Req.Domain),
						Type: aws.String(r.Req.Type),
						ResourceRecords: []*route53.ResourceRecord{
							{
								Value: aws.String(r.Req.Value),
							},
						},
						TTL: aws.Int64(r.Req.TTL),
					},
				},
			},
			Comment: aws.String(comment),
		},
		HostedZoneId: aws.String(r.Id),
	}
	return params
}

func (r *Route53) Logger() error {
	_ = os.Mkdir(path.Dir(logpath), 0777)

	logfile, err := os.OpenFile(logpath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}

	b, err := json.Marshal(r.Req)
	if err != nil {
		return err
	}

	logfile.Write(b)
	logfile.WriteString("\n")
	return nil
}

func (r *Route53) RequestChange() error {
	id, err := r.GetHostedZoneId(r.Req.Domain)
	if err != nil {
		return err
	}
	r.Id = id

	params := r.CreateNewParams()

	resp, err := r.Client.ChangeResourceRecordSets(params)
	if err != nil {
		return err
	}

	err = r.CheckStatus(resp)
	if err != nil {
		return err
	}

	err = r.Logger()
	if err != nil {
		return err
	}
	return nil
}
