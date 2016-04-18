package discourse

import (
    "time"
)

type Topic struct {
    PostStream TopicPostStream `json:"post_stream"`
}

type TopicPostStream struct {
    Posts []TopicPost `json:"posts"`
}

type TopicPost struct {
    ID int `json:"id"`
    Name string `json:"name"`
    Username string `json:"username"`
    AvatarTemplate string `json:"avatar_template"`
    CreatedAt time.Time `json:"created_at"`
    Cooked string `json:"cooked"`
    PostNumber int `json:"post_number"`
    PostType int `json:"post_type"`
    UpdatedAt time.Time `json:"updated_at"`
}
