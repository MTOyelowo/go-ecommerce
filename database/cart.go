package database

import "errors"

var (
	ErrCantFindProduct    = errors.New("Cannot find the Product")
	ErrCantDecodeProducts = errors.New("Cannot find the Product")
	ErrUserIdIsNotValid   = errors.New("This user is not valid")
	ErrCantUpdateUser     = errors.New("Cannot update the User")
	ErrCantRemoveItemCart = errors.New("Cannot remove this item from the cart")
	ErrCantGetItem        = errors.New("Was unable to get the item from the cart")
	ErrCantBuyCartItem    = errors.New("Cannot update the purchase")
)

func AddProductToCart() {

}
func RemoveCartItem() {

}
func BuyItemFromCart() {

}
func InstantBuyer() {

}
