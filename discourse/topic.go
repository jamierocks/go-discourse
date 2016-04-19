package discourse

import (
    "time"
)

type Topic struct {
    PostStream TopicPostStream `json:"post_stream"`

    ID int `json:"id"`
    Title string `json:"title"`
    FancyTitle string `json:"fancy_title"`
    PostsCounts int `json:"posts_count"`
    CreatedAt time.Time `json:"created_at"`
    Views int `json:"views"`
    ReplyCount int `json:"reply_count"`
    ParticipantCount int `json:"participant_count"`
    LikeCount int `json:"like_count"`
    LastPostedAt time.Time `json:"last_posted_at"`
    Visible bool `json:"visible"`
    Closed bool `json:"closed"`
    Archived bool `json:"archived"`
    HasSummary bool `json:"has_summary"`
    Archetype string `json:"archetype"`
    Slug string `json:"slug"`
    CategoryID int `json:"category_id"`
    WordCount int `json:"word_count"`
    UserID int `json:"user_id"`
    DraftKey string `json:"draft_key"`
    DraftSequence int `json:"draft_sequence"`
    Posted bool `json:"posted"`
    PinnedGlobally bool `json:"pinned_globally"`
    Pinned bool `json:"pinned"`
    Details TopicDetails `json:"details"`
    HighestPostNumber int `json:"highest_post_number"`
    LastReadPostNumber int `json:"last_read_post_number"`
    ActionsSummary []TopicActionSummary `json:"actions_summary"`
    ChunkSize int `json:"chunk_size"`
    Bookmarked bool `json:"bookmarked"`
}

type TopicDetails struct {
    AutoCloseBasedOnLastPost bool `json:"auto_close_based_on_last_post"`
    CreatedBy TopicDetailsUser `json:"created_by"`
    LastPoster TopicDetailsUser `json:"last_poster"`
    Participants []TopicDetailsParticipant `json:"participants"`
    SuggestedTopics []TopicDetailsSuggestedTopic `json:"suggested_topics"`
    Links []TopicDetailsLink `json:"links"`
    NotificationLevel int `json:"notification_level"`
    CanInviteTo bool `json:"can_invite_to"`
    CanCreatePost bool `json:"can_create_post"`
    CanReplyAsNewTopic bool `json:"can_reply_as_new_topic"`
    CanFlagTopic bool `json:"can_flag_topic"`
}

type TopicDetailsUser struct {
    ID int `json:"id"`
    Username string `json:"username"`
    AvatarTemplate string `json:"avatar_template"`
}

type TopicDetailsParticipant struct {
    ID int `json:"id"`
    Username string `json:"username"`
    AvatarTemplate string `json:"avatar_template"`
    PostCount int `json:"post_count"`
}

type TopicDetailsSuggestedTopic struct {
    ID int `json:"id"`
    Title string `json:"title"`
    FancyTitle string `json:"fancy_title"`
    Slug string `json:"slug"`
    PostsCount int `json:"posts_count"`
    ReplyCount int `json:"reply_count"`
    HighestPostNumber int `json:"highes_post_number"`
    ImageURL string `json:"image_url"`
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
    Archetype string `json:"archetype"`
    LikeCount int `json:"like_count"`
    Views int `json:"views"`
    CategoryID int `json:"category_id"`
}

type TopicDetailsLink struct {
    URL string `json:"url"`
    Title string `json:"title"`
    Internal bool `json:"internal"`
    Attachment bool `json:"attachment"`
    Reflection bool `json:"reflection"`
    Clicks int `json:"clicks"`
    UserID int `json:"user_id"`
    Domain string `json:"domain"`
}

type TopicPostStream struct {
    Posts []TopicPost `json:"posts"`
    Stream []int `json:"stream"`
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
    ReplyCount int `json:"reply_count"`
    QuoteCount int `json:"quote_count"`
    AverageTime int `json:"avg_time"`
    IncomingLinkCount int `json:"incoming_link_count"`
    Reads int `json:"reads"`
    Score float64 `json:"score"`
    Yours bool `json:"yours"`
    TopicID int `json:"topic_id"`
    TopicSlug string `json:"topic_slug"`
    DisplayUsername string `json:"display_username"`
    PrimaryGroupName string `json:"primary_group_name"`
    Version int `json:"version"`
    CanEdit bool `json:"can_edit"`
    CanDelete bool `json:"can_delete"`
    CanRecover bool `json:"can_recover"`
    CanWiki bool `json:"can_wiki"`
    LinkCounts []TopicLinkCount `json:"link_counts"`
    Read bool `json:"read"`
    UserTitle string `json:"user_title"`
    ActionsSummary []TopicActionSummary `json:"actions_summary"`
    Moderator bool `json:"moderator"`
    Admin bool `json:"admin"`
    Staff bool `json:"staff"`
    UserID int `json:"user_id"`
    Hidden bool `json:"hidden"`
    TrustLevel int `json:"trust_level"`
    UserDeleted bool `json:"user_deleted"`
    CanViewEditHistory bool `json:"can_view_edit_history"`
    Wiki bool `json:"wiki"`
    CanAcceptAnswer bool `json:"can_accept_answer"`
    CanUnacceptAnswer bool `json:"can_unaccept_answer"`
    AcceptedAnswer bool `json:"accepted_answer"`
}

type TopicLinkCount struct {
    URL string `json:"url"`
    Internal bool `json:"internal"`
    Reflection bool `json:"reflection"`
    Title string `json:"title"`
    Clicks int `json:"clicks"`
}

type TopicActionSummary struct {
    ID int `json:"id"`
    CanAct bool `json:"can_act"`
}
