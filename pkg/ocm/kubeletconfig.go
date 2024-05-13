package ocm

import (
	"context"
	"fmt"
	"net/http"

	cmv1 "github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1"
)

type KubeletConfigArgs struct {
	PodPidsLimit int
	Name         string
}

func (c *Client) GetClusterKubeletConfig(clusterID string) (*cmv1.KubeletConfig, bool, error) {
	response, err := c.ocm.ClustersMgmt().V1().Clusters().Cluster(clusterID).KubeletConfig().Get().Send()

	if response.Status() == http.StatusNotFound {
		return nil, false, nil
	}

	if err != nil {
		return nil, false, handleErr(response.Error(), err)
	}

	return response.Body(), true, nil
}

func (c *Client) DeleteKubeletConfig(clusterID string) error {
	response, err := c.ocm.ClustersMgmt().V1().Clusters().Cluster(clusterID).KubeletConfig().Delete().Send()
	if err != nil {
		return handleErr(response.Error(), err)
	}
	return nil
}

func toOCMKubeletConfig(args KubeletConfigArgs) (*cmv1.KubeletConfig, error) {
	builder := &cmv1.KubeletConfigBuilder{}
	kubeletConfig, err := builder.PodPidsLimit(args.PodPidsLimit).Name(args.Name).Build()
	if err != nil {
		return nil, err
	}

	return kubeletConfig, nil
}

func (c *Client) CreateKubeletConfig(clusterID string, args KubeletConfigArgs) (*cmv1.KubeletConfig, error) {

	kubeletConfig, err := toOCMKubeletConfig(args)
	if err != nil {
		return nil, err
	}

	response, err := c.ocm.ClustersMgmt().V1().Clusters().Cluster(clusterID).
		KubeletConfig().Post().Body(kubeletConfig).Send()

	if err != nil {
		return nil, handleErr(response.Error(), err)
	}

	return response.Body(), nil
}

func (c *Client) UpdateKubeletConfig(ctx context.Context,
	clusterID string, kubeletConfigId string, args KubeletConfigArgs) (*cmv1.KubeletConfig, error) {
	kubeletConfig, err := toOCMKubeletConfig(args)
	if err != nil {
		return nil, err
	}

	response, err := c.ocm.ClustersMgmt().V1().Clusters().Cluster(clusterID).
		KubeletConfigs().KubeletConfig(kubeletConfigId).Update().Body(kubeletConfig).SendContext(ctx)

	if err != nil {
		return nil, handleErr(response.Error(), err)
	}

	return response.Body(), nil
}

func (c *Client) FindKubeletConfigByName(
	ctx context.Context, clusterId string, name string) (*cmv1.KubeletConfig, bool, error) {

	if name == "" {
		return nil, false, fmt.Errorf("'name' for the KubeletConfig cannot be empty")
	}

	/*
		In-memory searching by name as endpoint does not currently support searching. We expect
		the number of KubeletConfigs on a cluster to be small e.g less-than-10, so this call is hopefully
		not too expensive. We will return to this if it becomes a problem.
	*/

	list, err := c.ListKubeletConfigs(ctx, clusterId)
	if err != nil {
		return nil, false, err
	}

	for _, k := range list {
		if k.Name() == name {
			return k, true, nil
		}
	}

	return nil, false, nil
}

func (c *Client) ListKubeletConfigs(ctx context.Context, clusterId string) ([]*cmv1.KubeletConfig, error) {
	response, err := c.ocm.ClustersMgmt().V1().Clusters().Cluster(clusterId).KubeletConfigs().List().SendContext(ctx)
	if err != nil {
		return []*cmv1.KubeletConfig{}, handleErr(response.Error(), err)
	}

	return response.Items().Slice(), nil
}
