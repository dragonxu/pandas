// Copyright (c) Mainflux
// SPDX-License-Identifier: Apache-2.0

package tracing

import (
	"context"

	"github.com/cloustone/pandas/vms"
	opentracing "github.com/opentracing/opentracing-go"
)

const (
	saveViewOp               = "save_view"
	saveViewsOp              = "save_view"
	updateViewOp             = "update_view"
	updateViewKeyOp          = "update_view_by_key"
	retrieveViewByIDOp       = "retrieve_view_by_id"
	retrieveViewByKeyOp      = "retrieve_view_by_key"
	retrieveAllViewsOp       = "retrieve_all_views"
	retrieveViewsByChannelOp = "retrieve_vms_by_chan"
	removeViewOp             = "remove_view"
	retrieveViewIDByKeyOp    = "retrieve_id_by_key"
)

var (
	_ vms.ViewRepository = (*viewRepositoryMiddleware)(nil)
	_ vms.ViewCache      = (*viewCacheMiddleware)(nil)
)

type viewRepositoryMiddleware struct {
	tracer opentracing.Tracer
	repo   vms.ViewRepository
}

// ViewRepositoryMiddleware tracks request and their latency, and adds spans
// to context.
func ViewRepositoryMiddleware(tracer opentracing.Tracer, repo vms.ViewRepository) vms.ViewRepository {
	return viewRepositoryMiddleware{
		tracer: tracer,
		repo:   repo,
	}
}

func (trm viewRepositoryMiddleware) Save(ctx context.Context, ths ...vms.View) ([]vms.View, error) {
	span := createSpan(ctx, trm.tracer, saveViewsOp)
	defer span.Finish()
	ctx = opentracing.ContextWithSpan(ctx, span)

	return trm.repo.Save(ctx, ths...)
}

func (trm viewRepositoryMiddleware) Update(ctx context.Context, th vms.View) error {
	span := createSpan(ctx, trm.tracer, updateViewOp)
	defer span.Finish()
	ctx = opentracing.ContextWithSpan(ctx, span)

	return trm.repo.Update(ctx, th)
}

func (trm viewRepositoryMiddleware) RetrieveByID(ctx context.Context, owner, id string) (vms.View, error) {
	span := createSpan(ctx, trm.tracer, retrieveViewByIDOp)
	defer span.Finish()
	ctx = opentracing.ContextWithSpan(ctx, span)

	return trm.repo.RetrieveByID(ctx, owner, id)
}

func (trm viewRepositoryMiddleware) RetrieveAll(ctx context.Context, owner string, offset, limit uint64, name string, metadata vms.Metadata) (vms.ViewsPage, error) {
	span := createSpan(ctx, trm.tracer, retrieveAllViewsOp)
	defer span.Finish()
	ctx = opentracing.ContextWithSpan(ctx, span)

	return trm.repo.RetrieveAll(ctx, owner, offset, limit, name, metadata)
}

func (trm viewRepositoryMiddleware) Remove(ctx context.Context, owner, id string) error {
	span := createSpan(ctx, trm.tracer, removeViewOp)
	defer span.Finish()
	ctx = opentracing.ContextWithSpan(ctx, span)

	return trm.repo.Remove(ctx, owner, id)
}

type viewCacheMiddleware struct {
	tracer opentracing.Tracer
	cache  vms.ViewCache
}

// ViewCacheMiddleware tracks request and their latency, and adds spans
// to context.
func ViewCacheMiddleware(tracer opentracing.Tracer, cache vms.ViewCache) vms.ViewCache {
	return viewCacheMiddleware{
		tracer: tracer,
		cache:  cache,
	}
}

func (tcm viewCacheMiddleware) Save(ctx context.Context, viewKey string, viewID string) error {
	span := createSpan(ctx, tcm.tracer, saveViewOp)
	defer span.Finish()
	ctx = opentracing.ContextWithSpan(ctx, span)

	return tcm.cache.Save(ctx, viewKey, viewID)
}

func (tcm viewCacheMiddleware) ID(ctx context.Context, viewKey string) (string, error) {
	span := createSpan(ctx, tcm.tracer, retrieveViewIDByKeyOp)
	defer span.Finish()
	ctx = opentracing.ContextWithSpan(ctx, span)

	return tcm.cache.ID(ctx, viewKey)
}

func (tcm viewCacheMiddleware) Remove(ctx context.Context, viewID string) error {
	span := createSpan(ctx, tcm.tracer, removeViewOp)
	defer span.Finish()
	ctx = opentracing.ContextWithSpan(ctx, span)

	return tcm.cache.Remove(ctx, viewID)
}
