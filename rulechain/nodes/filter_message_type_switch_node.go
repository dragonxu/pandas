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
package nodes

import (
	"fmt"

	"github.com/cloustone/pandas/rulechain/message"
	"github.com/sirupsen/logrus"
)

type messageTypeSwitchNode struct {
	bareNode
}
type messageTypeSwitchNodeFactory struct{}

func (f messageTypeSwitchNodeFactory) Name() string     { return "MessageTypeSwitchNode" }
func (f messageTypeSwitchNodeFactory) Category() string { return NODE_CATEGORY_FILTER }

func (f messageTypeSwitchNodeFactory) Create(id string, meta Metadata) (Node, error) {
	labels := []string{"True", "False"}
	node := &messageTypeSwitchNode{
		bareNode: newBareNode(f.Name(), id, meta, labels),
	}
	return decodePath(meta, node)
}

func (n *messageTypeSwitchNode) Handle(msg message.Message) error {
	logrus.Infof("%s handle message '%s'", n.Name(), msg.GetType())

	nodes := n.GetLinkedNodes()
	messageType := msg.GetType()
	messageTypeKey, _ := n.Metadata().Value(NODE_CONFIG_MESSAGE_TYPE_KEY)
	userMessageType := msg.GetMetadata().GetKeyValue(messageTypeKey.(string))

	for label, node := range nodes {
		if messageType == label || userMessageType == label {
			return node.Handle(msg)
		}
	}
	// if other label exist
	if node := n.GetLinkedNode("True"); node != nil {
		return node.Handle(msg)
	}

	// not found
	return fmt.Errorf("%s no label to handle message", n.Name())
}
