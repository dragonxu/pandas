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
	"sync"
	"time"

	"github.com/cloustone/pandas/rulechain/message"
	"github.com/sirupsen/logrus"
)

const DelayNodeName = "DelayNode"

type delayNode struct {
	bareNode
	PeriodTs           int               `json:"periodTs" yaml:"periodTs" jpath:"periodTs"`
	MaxPendingMessages int               `json:"maxPendingMessages" yaml:"maxPendingMessages" jpath:"maxPendingMessages"`
	messageQueue       []message.Message `jpath:"-"`
	delayTimer         *time.Timer       `jpath:"-"`
	lock               sync.Mutex        `jpath:"-"`
}

type delayNodeFactory struct{}

func (f delayNodeFactory) Name() string     { return DelayNodeName }
func (f delayNodeFactory) Category() string { return NODE_CATEGORY_ACTION }

func (f delayNodeFactory) Create(id string, meta Metadata) (Node, error) {
	labels := []string{"Success", "Failure"}
	node := &delayNode{
		bareNode: newBareNode(f.Name(), id, meta, labels),
		lock:     sync.Mutex{},
	}
	_, err := decodePath(meta, node)
	node.messageQueue = make([]message.Message, node.MaxPendingMessages)
	return node, err
}

func (n *delayNode) Handle(msg message.Message) error {
	logrus.Infof("%s handle message '%s'", n.Name(), msg.GetType())

	successLabelNode := n.GetLinkedNode("Success")
	failureLabelNode := n.GetLinkedNode("Failure")
	if successLabelNode == nil || failureLabelNode == nil {
		return fmt.Errorf("no valid label linked node in %s", n.Name())
	}

	// check wethere the time had already been started, queue message if started
	if n.delayTimer == nil {
		n.messageQueue = append(n.messageQueue, msg)
		n.delayTimer = time.NewTimer(time.Duration(n.PeriodTs) * time.Second)
		// start timecallback goroutine
		go func(n *delayNode) error {
			defer n.delayTimer.Stop()
			for {
				<-n.delayTimer.C
				n.lock.Lock()
				defer n.lock.Unlock()
				if len(n.messageQueue) > 0 {
					msg := n.messageQueue[0]
					n.messageQueue = n.messageQueue[0:]
					return successLabelNode.Handle(msg)
				}
			}
		}(n)
		return nil
	}
	// the delay timer had already been created, just queue message
	n.lock.Lock()
	defer n.lock.Unlock()
	if len(n.messageQueue) == n.MaxPendingMessages {
		return failureLabelNode.Handle(msg)
	}
	n.messageQueue = append(n.messageQueue, msg)
	return nil
}
