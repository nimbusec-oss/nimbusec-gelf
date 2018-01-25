package nimbusec

import (
	"context"
	"net/http"
)

type BundleService service

func (srv *BundleService) List(ctx context.Context) ([]Bundle, error) {
	bundles := []Bundle{}
	err := srv.client.Do(ctx, http.MethodGet, "/v3/bundles", nil, &bundles)
	return bundles, err
}

func (srv *BundleService) Get(ctx context.Context, id BundleID) (Bundle, error) {
	bundle := Bundle{}
	err := srv.client.Do(ctx, http.MethodGet, string(id), nil, &bundle)
	return bundle, err
}
