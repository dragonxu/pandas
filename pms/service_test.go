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
package pms

import (
	"context"
	"testing"

	modelsoptions "github.com/cloustone/pandas/models/options"
	pb "github.com/cloustone/pandas/pms/grpc_pms_v1"
	. "github.com/smartystreets/goconvey/convey"
)

func TestInitialize(t *testing.T) {
	Convey("TestInitialize should return ok", t, func() {
		servingOptions := modelsoptions.NewServingOptions()
		//factory.Initialize(servingOptions)
		//fmt.Println(servingOptions.CacheConnectedUrl)
		pm := NewProjectManagementService(servingOptions)
		So(pm, ShouldNotBeNil)
	})
}

func TestCreateProject(t *testing.T) {
	Convey("TestCreateProject should return ok when project metainfo is nil", t, func() {
		servingOptions := modelsoptions.NewServingOptions()
		pm := NewProjectManagementService(servingOptions)
		_, err := pm.CreateProject(context.TODO(), &pb.CreateProjectRequest{})
		So(err, ShouldNotBeNil)
	})
	Convey("TestCreateProject should return ok when project not exist", t, func() {
		servingOptions := modelsoptions.NewServingOptions()
		pm := NewProjectManagementService(servingOptions)
		id := "12345678"
		req := pb.CreateProjectRequest{
			UserID: "hello",
			Project: &pb.Project{
				ID:          id,
				Name:        "sample",
				UserID:      "hello",
				Description: "sample project",
			},
		}

		project, err := pm.CreateProject(context.TODO(), &req)
		So(err, ShouldBeNil)
		So(project, ShouldNotBeNil)
		//So(StringSliceEqual(project.Id, id), ShouldBeTrue)
	})

	Convey("TestCreateProject should return error if project already exist", t, func() {
		servingOptions := modelsoptions.NewServingOptions()
		pm := NewProjectManagementService(servingOptions)
		id := "12345678"
		req := pb.CreateProjectRequest{
			UserID: "hello",
			Project: &pb.Project{
				ID:          id,
				Name:        "sample",
				UserID:      "hello",
				Description: "sample project",
			},
		}

		project, err := pm.CreateProject(context.TODO(), &req)
		So(err, ShouldNotBeNil)
		So(project, ShouldBeNil)
	})

}
