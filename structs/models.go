package structs

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name" binding:"required"`
}

type Book struct {
	ID          int     `json:"id"`
	Title       string  `json:"title" binding:"required"`
	Description string  `json:"description"`
	ImageURL    string  `json:"image_url"`
	ReleaseYear int     `json:"release_year" binding:"required"`
	Price       float64 `json:"price"`
	TotalPage   int     `json:"total_page" binding:"required"`
	Thickness   string  `json:"thickness,omitempty"`
	CategoryID  *int    `json:"category_id"`
}
