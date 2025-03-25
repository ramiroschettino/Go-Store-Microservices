package repository

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"github.com/ramiroschettino/Go-Store-Microservices/checkout-service/internal/domain"
)

type CartRepositoryImpl struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) domain.CartRepository {
	return &CartRepositoryImpl{db: db}
}

func (r *CartRepositoryImpl) GetByID(ctx context.Context, id uint) (*domain.Cart, error) {
	var cart domain.Cart
	err := r.db.WithContext(ctx).Preload("Items").First(&cart, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &cart, err
}

func (r *CartRepositoryImpl) GetByUserID(ctx context.Context, userID uint) (*domain.Cart, error) {
	var cart domain.Cart
	err := r.db.WithContext(ctx).Preload("Items").Where("user_id = ?", userID).First(&cart).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &cart, err
}

func (r *CartRepositoryImpl) Create(ctx context.Context, cart *domain.Cart) error {
	return r.db.WithContext(ctx).Create(cart).Error
}

func (r *CartRepositoryImpl) Update(ctx context.Context, cart *domain.Cart) error {
	return r.db.WithContext(ctx).Save(cart).Error
}

func (r *CartRepositoryImpl) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&domain.CartItem{}, "cart_id = ?", id).Error; err != nil {
			return err
		}
		return tx.Delete(&domain.Cart{}, id).Error
	})
}

func (r *CartRepositoryImpl) AddItem(ctx context.Context, cartID uint, item *domain.CartItem) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		item.CartID = cartID
		if err := tx.Create(item).Error; err != nil {
			return err
		}
		return tx.Model(&domain.Cart{}).
			Where("id = ?", cartID).
			Update("amount", gorm.Expr("amount + ?", item.Price*float64(item.Quantity))).
			Error
	})
}

func (r *CartRepositoryImpl) RemoveItem(ctx context.Context, cartID uint, itemID uint) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var item domain.CartItem
		if err := tx.First(&item, itemID).Error; err != nil {
			return err
		}
		
		if err := tx.Delete(&item).Error; err != nil {
			return err
		}
		
		return tx.Model(&domain.Cart{}).
			Where("id = ?", cartID).
			Update("amount", gorm.Expr("amount - ?", item.Price*float64(item.Quantity))).
			Error
	})
}

func (r *CartRepositoryImpl) ClearItems(ctx context.Context, cartID uint) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&domain.CartItem{}, "cart_id = ?", cartID).Error; err != nil {
			return err
		}
		return tx.Model(&domain.Cart{}).
			Where("id = ?", cartID).
			Update("amount", 0).
			Error
	})
}