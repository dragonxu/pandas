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
	"time"

	"github.com/cloustone/pandas/models"
	"github.com/cloustone/pandas/models/cache"
	"github.com/cloustone/pandas/models/factory"
	modelsoptions "github.com/cloustone/pandas/models/options"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

type viewFactory struct {
	modelDB        *gorm.DB
	cache          cache.Cache
	servingOptions *modelsoptions.ServingOptions
}

func newViewFactory(servingOptions *modelsoptions.ServingOptions) factory.Factory {
	servingOptions = modelsoptions.NewServingOptions()
	modelDB, err := gorm.Open(servingOptions.StorePath, "pandas-projects.db")
	//modelDB, err := gorm.Open("sqlite3", "pandas-projects.db")
	if err != nil {
		logrus.Fatal(err)
	}
	modelDB.AutoMigrate(&models.View{})
	return &viewFactory{
		modelDB:        modelDB,
		cache:          cache.NewCache(servingOptions),
		servingOptions: servingOptions,
	}
}

func (pf *viewFactory) Save(owner factory.Owner, obj models.Model) (models.Model, error) {
	view := obj.(*models.View)
	view.ViewCreatedAt = time.Now()
	pf.modelDB.Save(view)

	if err := factory.Error(pf.modelDB); err != nil {
		return nil, err
	}
	// update cache
	pf.cache.Set(factory.NewCacheID(owner, view.ViewID), view)
	return view, nil
}

func (pf *viewFactory) List(owner factory.Owner, query *models.Query) ([]models.Model, error) {
	views := []*models.Project{}
	pf.modelDB.Where("userId = ?", owner.User()).Find(views)

	if err := factory.Error(pf.modelDB); err != nil {
		return nil, err
	}
	results := []models.Model{}
	for _, view := range views {
		results = append(results, view)
	}
	return results, nil
}

func (pf *viewFactory) Get(owner factory.Owner, viewID string) (models.Model, error) {
	view := models.View{}
	if err := pf.cache.Get(factory.NewCacheID(owner, viewID), &view); err == nil {
		return &view, nil
	}
	pf.modelDB.Where("userId = ? AND projectId = ?", owner.User(), viewID).Find(&view)
	if err := factory.Error(pf.modelDB); err != nil {
		return nil, err
	}

	return &view, nil
}

func (pf *viewFactory) Delete(owner factory.Owner, viewID string) error {
	pf.modelDB.Delete(&models.View{
		ProjectID: owner.Project(),
		ViewID:    viewID,
	})
	pf.cache.Delete(factory.NewCacheID(owner, viewID))
	return nil
}

func (pf *viewFactory) Update(factory.Owner, models.Model) error {
	return nil
}
