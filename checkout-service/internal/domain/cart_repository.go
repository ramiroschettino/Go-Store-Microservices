package domain

import "context"

type CartRepository interface {
    GetByID(ctx context.Context, id uint) (*Cart, error)
    GetByUserID(ctx context.Context, userID uint) (*Cart, error)
    Create(ctx context.Context, cart *Cart) error
    Update(ctx context.Context, cart *Cart) error
    Delete(ctx context.Context, id uint) error
    AddItem(ctx context.Context, cartID uint, item *CartItem) error
    RemoveItem(ctx context.Context, cartID uint, itemID uint) error
    ClearItems(ctx context.Context, cartID uint) error
}