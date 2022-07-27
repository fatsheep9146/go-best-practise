package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
)

type alertCondition struct {
	DataType     string `json:"data_type"`
	DataSourceID string `json:"datasource_id"`
	Expr         string `json:"expr"`
	Duration     string `json:"duration"`
	Reducer      string `json:"reducer"`
	Evaluator    string `json:"evaluator"`
	Threshold    string `json:"threshold"`
}

// Manifest defines the conditions of an alert rule
type Manifest []alertCondition

type metadata struct {
	Labels    map[string]string `yaml:"labels"`
	Name      string            `yaml:"name"`
	Namespace string            `yaml:"namespace"`
}

type alertRule struct {
	Alert       string            `yaml:"alert"`
	Annotations map[string]string `yaml:"annotations"`
	Expr        string            `yaml:"expr"`
	For         string            `yaml:"for"`
	Labels      map[string]string `yaml:"labels"`
}

type group struct {
	Name  string       `yaml:"name"`
	Rules []*alertRule `yaml:"rules"`
}

type spec struct {
	Groups []group `yaml:"groups"`
}

type prometheusRule struct {
	APIVersion string   `yaml:"apiVersion"`
	Kind       string   `yaml:"kind"`
	Metadata   metadata `yaml:"metadata"`
	Spec       spec     `yaml:"spec"`
}

func toPrometheusRule(rule *store.WorkloadAlertRule) (raw []byte, namespace string, err error) {
	// parse all alert rule
	alerts, err := toAlertRules(rule.Manifest)
	if err != nil {
		return
	}

	rulelabels, err := util.ParseLabelsStr(rule.Labels, "#")
	if err != nil {
		return
	}

	for _, alert := range alerts {
		alert.Alert = prometheusRuleName(rule)
		alert.Labels = rulelabels
		alert.Annotations = map[string]string{
			"dashboard":         rule.Dashboards,
			"template":          rule.Detail,
			"owner":             rule.Owners,
			"playbook":          rule.Playbooks,
			"description":       rule.Description,
			"alert_template_id": rule.RuleGroupID,
			"alert_rule_id":     rule.UUID,
		}
	}

	if _, exist := rulelabels["prometheusrule_ignored"]; exist {
		return
	}

	p := prometheusRule{
		APIVersion: "monitoring.coreos.com/v1",
		Kind:       "PrometheusRule",
		Metadata: metadata{
			Name:      strings.ToLower(prometheusRuleName(rule)),
			Namespace: reduceNamespace(rulelabels["cluster_type"]),
			Labels:    reducePrometheusRuleLabels(rulelabels["cluster_type"]),
		},
		Spec: spec{
			Groups: []group{
				{
					Name:  prometheusRuleName(rule),
					Rules: alerts,
				},
			},
		},
	}

	if raw, err = yaml.Marshal(p); err != nil {
		err = fmt.Errorf("marshal prometheusRule failed, err: %v", err)
		return
	}

	return
}

func prometheusRuleName(rule *store.WorkloadAlertRule) string {
	return rule.Name
}

func toAlertRules(manifest string) (rules []*alertRule, err error) {
	var (
		conditions = make([]alertCondition, 0)
		operator   = ""
	)

	if err = json.Unmarshal([]byte(manifest), &conditions); err != nil {
		return nil, fmt.Errorf("unmarshal manifest failed: %v", err)
	}

	switch conditions[0].Evaluator {
	case "gt":
		operator = ">"
	case "lt":
		operator = "<"
	}

	rules = make([]*alertRule, 0)
	rule := &alertRule{
		Expr: fmt.Sprintf("(%s) %s %s", conditions[0].Expr, operator, conditions[0].Threshold),
		For:  conditions[0].Duration,
	}
	rules = append(rules, rule)

	return
}

func reduceNamespace(clusterType string) string {
	switch clusterType {
	case "primary":
		return "monitoring"
	case "guest":
	case "guest_default":
		return "kube-system"
	}

	return ""
}

func reducePrometheusRuleLabels(clusterType string) (prometheusRuleLabels map[string]string) {
	switch clusterType {
	case "primary":
		return map[string]string{
			"app":    "prometheus",
			"source": "deploy",
			"type":   "alerting",
		}
	case "guest":
	case "guest_default":
		return map[string]string{
			"app":    "ruler",
			"cato":   "single",
			"source": "deploy",
			"type":   "alerting",
		}
	}

	return
}

func main() {
	
}