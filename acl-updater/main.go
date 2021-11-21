package main

import (
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/CzarSimon/httputil/environ"
	"github.com/CzarSimon/httputil/logger"
	"github.com/scaleway/scaleway-sdk-go/api/k8s/v1"
	"github.com/scaleway/scaleway-sdk-go/api/rdb/v1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

var log = logger.GetDefaultLogger("acl-updater/main").Sugar()

type config struct {
	region     scw.Region
	clusterID  string
	instanceID string
}

type aclUpdate struct {
	toAdd    []*rdb.ACLRuleRequest
	toRemove []string
}

func main() {
	c := getClient()
	cfg := getConfig()
	update := getAclUpdate(c, cfg)
	updateACLRules(c, cfg, update)
}

func getClient() *scw.Client {
	accessKey := environ.MustGet("SCW_ACCESS_KEY")
	secretKey := environ.MustGet("SCW_SECRET_KEY")
	c, err := scw.NewClient(scw.WithAuth(accessKey, secretKey))
	if err != nil {
		log.Fatal("failed to initalize client. Error: %w", err)
	}

	return c
}

func getAclUpdate(c *scw.Client, cfg config) aclUpdate {
	nodeCIDRs := listNodeCIDRs(c, cfg)
	ruleCIDRs := listACLRuleCIDRs(c, cfg)

	update := aclUpdate{
		toAdd:    make([]*rdb.ACLRuleRequest, 0),
		toRemove: make([]string, 0),
	}
	for cidr, n := range nodeCIDRs {
		_, ok := ruleCIDRs[cidr]
		if ok {
			continue
		}

		newACL := newACLRequest(cidr, n)
		update.toAdd = append(update.toAdd, newACL)
	}

	k8sACLSuffix := fmt.Sprintf("cluster/%s", cfg.clusterID)
	for cidr, r := range ruleCIDRs {
		_, ok := nodeCIDRs[cidr]
		if ok || !strings.Contains(r.Description, k8sACLSuffix) {
			continue
		}

		update.toRemove = append(update.toRemove, cidr)
	}

	return update
}

func newACLRequest(cidr string, node *k8s.Node) *rdb.ACLRuleRequest {
	_, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		log.Fatal("failed to parse cidr %s. Error: %w", cidr, err)
	}

	return &rdb.ACLRuleRequest{
		IP: scw.IPNet{
			IPNet: *ipnet,
		},
		Description: fmt.Sprintf("allow ingress from node/%s in cluster/%s", node.ID, node.ClusterID),
	}
}

func listNodeCIDRs(c *scw.Client, cfg config) map[string]*k8s.Node {
	req := &k8s.ListNodesRequest{
		Region:    cfg.region,
		ClusterID: cfg.clusterID,
	}

	res, err := k8s.NewAPI(c).ListNodes(req)
	if err != nil {
		log.Fatal("failed to list k8s nodes for %s. Error: %w", cfg, err)
	}

	return mapNodeCIDRs(res.Nodes)
}

func mapNodeCIDRs(nodes []*k8s.Node) map[string]*k8s.Node {
	cidrs := make(map[string]*k8s.Node)
	for _, n := range nodes {
		cidr := fmt.Sprintf("%s/32", n.PublicIPV4.String())
		cidrs[cidr] = n
	}

	return cidrs
}

func listACLRuleCIDRs(c *scw.Client, cfg config) map[string]*rdb.ACLRule {
	req := &rdb.ListInstanceACLRulesRequest{
		Region:     cfg.region,
		InstanceID: cfg.instanceID,
	}

	res, err := rdb.NewAPI(c).ListInstanceACLRules(req)
	if err != nil {
		log.Fatal("failed to list database ACLs for %s. Error: %w", cfg, err)
	}

	return mapAclCIDRs(res.Rules)
}

func mapAclCIDRs(rules []*rdb.ACLRule) map[string]*rdb.ACLRule {
	cidrs := make(map[string]*rdb.ACLRule)
	for _, r := range rules {
		cidrs[r.IP.String()] = r
	}

	return cidrs
}

func updateACLRules(c *scw.Client, cfg config, update aclUpdate) {
	api := rdb.NewAPI(c)
	addACLRules(api, cfg, update.toAdd)
	time.Sleep(time.Second)
	deleteACLRules(api, cfg, update.toRemove)
}

func deleteACLRules(api *rdb.API, cfg config, toRemove []string) {
	if len(toRemove) < 1 {
		return
	}

	deleteRequest := &rdb.DeleteInstanceACLRulesRequest{
		Region:     cfg.region,
		InstanceID: cfg.instanceID,
		ACLRuleIPs: toRemove,
	}
	_, err := api.DeleteInstanceACLRules(deleteRequest)
	if err != nil {
		log.Fatal("failed to delete ACLs for %v. Error: %w", toRemove, err)
	}
}

func addACLRules(api *rdb.API, cfg config, toAdd []*rdb.ACLRuleRequest) {
	if len(toAdd) < 1 {
		return
	}

	addRequest := &rdb.AddInstanceACLRulesRequest{
		Region:     cfg.region,
		InstanceID: cfg.instanceID,
		Rules:      toAdd,
	}
	_, err := api.AddInstanceACLRules(addRequest)
	if err != nil {
		log.Fatal("failed to add ACLs for %v. Error: %w", toAdd, err)
	}
}

func getConfig() config {
	regionStr := environ.MustGet("SCALEWAY_REGION")
	region, err := scw.ParseRegion(regionStr)
	if err != nil {
		log.Fatal("failed to parse region %s. Error: %w", regionStr, err)
	}

	return config{
		region:     region,
		clusterID:  environ.MustGet("K8S_CLUSTER_ID"),
		instanceID: environ.MustGet("RDB_INSTANCE_ID"),
	}
}

func (cfg config) String() string {
	return fmt.Sprintf(
		"config(region=%s, clusterId=%s, instanceId=%s)",
		cfg.region,
		cfg.clusterID,
		cfg.instanceID,
	)
}
