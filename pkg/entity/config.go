package entity

import "errors"

// PolarisClusterConf polaris 集群的控制面配置
type PolarisClusterConf struct {
	Name  string `json:"name"`  // 集群名
	Host  string `json:"host"`  // 集群控制面域名，或者"ip:port"
	Token string `json:"token"` // 集群控制面的鉴权 token
}

// PolarisCtl cli 配置
type PolarisCtlConf struct {
	DefaultCluster string               `json:"default_culster"`
	Clusters       []PolarisClusterConf `json:"clusters"`
}

// Check 检查必填配置
func (cnf PolarisCtlConf) Check() error {
	if len(cnf.DefaultCluster) == 0 {
		return errors.New("default_culster empty")
	}

	if len(cnf.Clusters) == 0 {
		return errors.New("clusters empty")
	}

	if _, err := cnf.FindCluster(cnf.DefaultCluster); err != nil {
		return err
	}

	return nil

}

// FindCluster 找到某个集群的配置
func (cnf PolarisCtlConf) FindCluster(clusterName string) (PolarisClusterConf, error) {
	for _, cluster := range cnf.Clusters {
		if cluster.Name == clusterName {
			return cluster, nil
		}
	}
	return PolarisClusterConf{}, errors.New("cannot find cluster conf:" + clusterName)
}
