package handler

// User 用于表示用户信息的结构体
type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

// GetAllUserResponse 用于获取所有用户的响应结构体
type GetAllUserResponse struct {
	Users []User `json:"users"`
}

// LoginRequest 用于登录的请求结构体
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// RegisterRequest 用于注册的请求结构体
type RegisterRequest struct {
	Username        string `json:"username"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
	Email           string `json:"email"`
}

// RegisterResponse 用于注册的响应结构体
type RegisterResponse struct {
	Message string `json:"message"`
}

// RefreshAccessTokenRequest 用于刷新访问令牌的请求结构体
type RefreshAccessTokenRequest struct {
	RefreshToken string `json:"refreshToken"`
}

type CreatePostRequest struct {
	Title      string   `json:"title"`
	Content    string   `json:"content"`
	Categories []string `json:"categories"`
	Tags       []string `json:"tags"`
}

type CreateCategoryRequest struct {
	Category string `json:"category"`
}
