//  Licensed under the Apache License, Version 2.0 (the "License"); you may
//  not use p file except in compliance with the License. You may obtain
//  a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
//  WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
//  License for the specific language governing permissions and limitations
//  under the License.
package rulechain

import (
	"fmt"
	"strconv"

	"github.com/cloustone/pandas/rulechain/manifest"
	"github.com/cloustone/pandas/rulechain/nodes"
	"github.com/sirupsen/logrus"
)

// ruleChainInstance is rulechain's runtime instance that manage all nodes in this chain,
// validate and apply datanly one input node exist in chain as precondition,
// and with many output nodes, Relations within nodes is maintained by link object
type ruleChainInstance struct {
	name            string
	firstRuleNodeId string
	root            bool
	debugMode       bool
	channel         string
	subTopic        string
	configuration   map[string]interface{}
	nodes           map[string]nodes.Node
}

func newRuleChainInstance(Channel string, SubTopic string, data []byte) (*ruleChainInstance, []error) {
	errors := []error{}

	manifest, err := manifest.New(data)
	if err != nil {
		errors = append(errors, err)
		logrus.WithError(err).Errorf("invalidi manifest file")
		return nil, errors
	}
	return newInstanceWithManifest(Channel, SubTopic, manifest)
}

// newWithManifest create rule chain by user's manifest file
func newInstanceWithManifest(Channel string, SubTopic string, m *manifest.Manifest) (*ruleChainInstance, []error) {
	errs := []error{}

	r := &ruleChainInstance{
		name:            m.RuleChain.Name,
		firstRuleNodeId: m.RuleChain.FirstRuleNodeId,
		root:            m.RuleChain.Root,
		debugMode:       m.RuleChain.DebugMode,
		channel:         Channel,
		subTopic:        SubTopic,
		configuration:   m.RuleChain.Configuration,
		nodes:           make(map[string]nodes.Node),
	}
	// Create All nodes
	for _, n := range m.Metadata.Nodes {
		metadata := nodes.NewMetadataWithValues(n.Configuration).With("debugMode", r.debugMode)
		node, err := nodes.NewNode(n.Type, n.Name, metadata)
		if err != nil {
			errs = append(errs, err)
			continue
		}
		if _, found := r.nodes[n.Name]; found {
			err := fmt.Errorf("node '%s' already exist in rulechain '%s'", n.Name, m.RuleChain.Name)
			errs = append(errs, err)
			continue
		}
		r.nodes[n.Name] = node
	}

	// Create All node connections
	for _, conn := range m.Metadata.Connections {
		originalNode, found := r.nodes[strconv.Itoa(conn.FromIndex)]
		if !found {
			err := fmt.Errorf("original node '%s' no exist in rulechain '%s'", originalNode.Name(), m.RuleChain.Name)
			errs = append(errs, err)
			continue
		}
		targetNode, found := r.nodes[strconv.Itoa(conn.ToIndex)]
		if !found {
			err := fmt.Errorf("target node '%s' no exist in rulechain '%s'", targetNode.Name(), m.RuleChain.Name)
			errs = append(errs, err)
			continue
		}
		originalNode.AddLinkedNode(conn.Type, targetNode)
	}
	// some labels must be satisified
	for name, node := range r.nodes {
		targetNodes := node.GetLinkedNodes()
		mustLabels := node.MustLabels()
		for _, label := range mustLabels {
			if _, found := targetNodes[label]; !found {
				err := fmt.Errorf("the label '%s' in node '%s' no exist'", label, name)
				errs = append(errs, err)
				continue
			}
		}
	}

	return r, errs
}
