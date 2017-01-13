package main

import (
	"strings"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/service/route53"
	"github.com/aws/aws-sdk-go/service/route53/route53iface"
)

var query = "lacrosse example.com TXT 93.184.216.34 86400 private"

func ptr(s string) *string {
	return &s
}

type mockRoute53Client struct {
	route53iface.Route53API
}

func NewMockRoute53() *Route53 {
	args := strings.Split(query, " ")
	req, _ := NewRequest(args[1:])

	r := &Route53{
		Client: &mockRoute53Client{},
		Id:     "",
		Req:    req,
	}

	return r
}

func (m *mockRoute53Client) ChangeResourceRecordSets(*route53.ChangeResourceRecordSetsInput) (*route53.ChangeResourceRecordSetsOutput, error) {
	time := time.Now().UTC()
	r := &route53.ChangeResourceRecordSetsOutput{
		ChangeInfo: &route53.ChangeInfo{
			Comment:     ptr("This is a test"),
			Id:          ptr("XXXXXXXXXX"),
			Status:      ptr("INSYNC"),
			SubmittedAt: &time,
		},
	}
	return r, nil
}

func (m *mockRoute53Client) ListHostedZones(*route53.ListHostedZonesInput) (*route53.ListHostedZonesOutput, error) {
	z := []*route53.HostedZone{
		{
			CallerReference: ptr("CC54D0F3-E81B-0DF4-A688-E8E340B7A0DA"),
			Id:              ptr("/hostedzone/Z3EVG1FTEI7PMH"),
			Name:            ptr("example.com."),
		},
	}
	istruncated := false
	l := &route53.ListHostedZonesOutput{
		HostedZones: z,
		IsTruncated: &istruncated,
		Marker:      ptr(""),
		MaxItems:    ptr("100"),
	}
	return l, nil
}

func (m *mockRoute53Client) WaitUntilResourceRecordSetsChanged(*route53.GetChangeInput) error {
	return nil
}

func TestRoute53_GetHostedZoneId(t *testing.T) {
	r := NewMockRoute53()
	id, err := r.GetHostedZoneId("example.com")
	if err != nil {
		t.Errorf("Error: ", err)
	}
	expected := "Z3EVG1FTEI7PMH"
	if id != expected {
		t.Errorf("got %v\nwant %v", id, expected)
	}
}

func TestRoute53_RequestChange(t *testing.T) {
	r := NewMockRoute53()
	err := r.RequestChange()
	if err != nil {
		t.Errorf("Error: ", err)
	}
}

func TestRoute53_CheckStatus(t *testing.T) {
	r := NewMockRoute53()
	time := time.Now().UTC()
	resp := &route53.ChangeResourceRecordSetsOutput{
		ChangeInfo: &route53.ChangeInfo{
			Comment:     ptr("This is a test"),
			Id:          ptr("XXXXXXXXXX"),
			Status:      ptr("INSYNC"),
			SubmittedAt: &time,
		},
	}
	err := r.CheckStatus(resp)
	if err != nil {
		t.Errorf("Error: ", err)
	}
}
