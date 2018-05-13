package clair

import (
	"context"

	"github.com/coreos/clair/api/v3/clairpb"
	"google.golang.org/grpc"
)

type Clair struct {
	client clairpb.AncestryServiceClient
}

func NewClient(url string) (*Clair, error) {
	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return &Clair{client: clairpb.NewAncestryServiceClient(conn)}, nil
}

func (c *Clair) Analyze(image string) ([]*clairpb.Vulnerability, error) {
	request := &clairpb.GetAncestryRequest{
		AncestryName:        image,
		WithFeatures:        true,
		WithVulnerabilities: true,
	}

	response, err := c.client.GetAncestry(context.Background(), request)
	if err != nil {
		return nil, err
	}

	var vulnerabilities []*clairpb.Vulnerability
	for _, f := range response.Ancestry.Features {
		for _, v := range f.Vulnerabilities {
			vulnerabilities = append(vulnerabilities, v)
		}
	}
	return vulnerabilities, nil
}
