package restaurantmodel

type Filter struct {
	OwnerId int `json:"owner_id,omitempty" form:"owner_id"`
	// "-" khong cho phep nguoi dung truyen len
	Status []int `json:"-"`
}
