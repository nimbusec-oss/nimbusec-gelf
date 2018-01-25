package nimbusec

import (
	"context"
	"net/http"
)

type DomainService service

func (srv DomainService) List(ctx context.Context) ([]Domain, error) {
	domains := []Domain{}
	err := srv.client.Do(ctx, http.MethodGet, "/v3/domains", nil, &domains)
	return domains, err
}

func (srv *DomainService) Get(ctx context.Context, id DomainID) (Domain, error) {
	domain := Domain{}
	err := srv.client.Do(ctx, http.MethodGet, string(id), nil, &domain)
	return domain, err
}

func (srv *DomainService) Create(ctx context.Context, create Domain) (Domain, error) {
	domain := Domain{}
	err := srv.client.Do(ctx, http.MethodPost, "/v3/domains", create, &domain)
	return domain, err
}

func (srv *DomainService) Update(ctx context.Context, id DomainID, update Domain) (Domain, error) {
	domain := Domain{}
	err := srv.client.Do(ctx, http.MethodPut, string(id), update, &domain)
	return domain, err
}

func (srv *DomainService) Delete(ctx context.Context, id DomainID) error {
	err := srv.client.Do(ctx, http.MethodDelete, string(id), nil, nil)
	return err
}
