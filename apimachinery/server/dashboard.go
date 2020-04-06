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
package server

import (
	"github.com/cloustone/pandas/apimachinery/restapi/operations/dashboard"
	"github.com/cloustone/pandas/apimachinery/models"

	"github.com/go-openapi/runtime/middleware"
)

func GetDashboard(params dashboard.GetDashboardParams, principal *models.Principal) middleware.Responder {
	return dashboard.NewGetDashboardOK()
}
