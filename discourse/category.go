package discourse

import (
    "time"
)

type Category struct {
    Users []CategoryUser `json:"users"`
    TopicList CategoryTopicList `json:"topic_list"`
}

type CategoryUser struct {
    ID int `json:"id"`
    Username string `json:"username"`
    AvatarTemplate string `json:"avatar_template"`
}

type CategoryTopicList struct {
    CanCreateTopic bool `json:"can_create_topic"`
    MoreTopicUrl string `json:"more_topics_url"`
    DraftKey string `json:"draft_key"`
    DraftSequence string `json:"draft_sequence"`
    PerPage int `json:"per_page"`
    Topics []CategoryTopic `json:"topics"`
}

type CategoryTopic struct {
    ID int `json:"id"`
    Title string `json:"title"`
    FancyTitle string `json:"fancy_title"`
    Slug string `json:"slug"`
    PostsCount int `json:"posts_count"`
    ReplyCount int `json:"reply_count"`
    HighestPostNumber int `json:"highest_post_number"`
    CreatedAt time.Time `json:"created_at"`
    LastPostedAt time.Time `json:"last_posted_at"`
    Bumped bool `json:"bumped"`
    BumpedAt time.Time `json:"bumped_at"`
    Unseen bool `json:"unseen"`
    LastReadPostNumber int `json:"last_read_post_number"`
    Unread int `json:"unread"`
    NewPosts int `json:"new_posts"`
    Pinned bool `json:"pinned"`
    Visible bool `json:"visible"`
    Closed bool `json:"closed"`
    Archived bool `json:"archived"`
    NotificationLevel int `json:"notification_level"`
    Bookmarked bool `json:"bookmarked"`
    Liked bool `json:"liked"`
    Views int `json:"views"`
    LikeCount int `json:"like_count"`
    HasSummary string `json:"has_summary"`
    Archetype string `json:"archetype"`
    LastPosterUsername string `json:"last_poster_username"`
    CategoryID int `json:"category_id"`
    PinnedGlobally bool `json:"pinned_globally"`
    Posters []CategoryTopicPoster `json:"posters"`
}

type CategoryTopicPoster struct {
    Description string `json:"description"`
    UserID int `json:"user_id"`
}
