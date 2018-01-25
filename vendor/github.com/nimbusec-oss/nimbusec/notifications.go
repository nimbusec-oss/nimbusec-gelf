package nimbusec

import (
	"context"
	"net/http"
)

type NotificationService service

func (srv *NotificationService) List(ctx context.Context) ([]Notification, error) {
	notifications := []Notification{}
	err := srv.client.Do(ctx, http.MethodGet, "/v3/notifications", nil, &notifications)
	return notifications, err
}

func (srv *NotificationService) ListByNotification(ctx context.Context, id NotificationID) ([]Notification, error) {
	notifications := []Notification{}
	err := srv.client.Do(ctx, http.MethodGet, string(id)+"/notifications", nil, &notifications)
	return notifications, err
}

func (srv *NotificationService) ListByUser(ctx context.Context, id UserID) ([]Notification, error) {
	notifications := []Notification{}
	err := srv.client.Do(ctx, http.MethodGet, string(id)+"/notifications", nil, &notifications)
	return notifications, err
}

func (srv *NotificationService) Get(ctx context.Context, id NotificationID) (Notification, error) {
	notification := Notification{}
	err := srv.client.Do(ctx, http.MethodGet, string(id), nil, &notification)
	return notification, err
}

func (srv *NotificationService) Create(ctx context.Context, create Notification) (Notification, error) {
	notification := Notification{}
	err := srv.client.Do(ctx, http.MethodPost, "/v3/notifications", create, &notification)
	return notification, err
}

func (srv *NotificationService) Update(ctx context.Context, id NotificationID, update NotificationUpdate) (Notification, error) {
	notification := Notification{}
	err := srv.client.Do(ctx, http.MethodPut, string(id), update, &notification)
	return notification, err
}

func (srv *NotificationService) Delete(ctx context.Context, id NotificationID) error {
	err := srv.client.Do(ctx, http.MethodDelete, string(id), nil, nil)
	return err
}
