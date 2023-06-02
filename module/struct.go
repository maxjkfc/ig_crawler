package module

type Response struct {
	Data Data `json:"data"`
}

type Data struct {
	User User `json:"user"`
}

type User struct {
	Media EdgeOwnerToTimelineMedia `json:"edge_owner_to_timeline_media"`
}

type EdgeOwnerToTimelineMedia struct {
	Count    int64    `json:"count"`
	PageInfo PageInfo `json:"page_info"`
	Edges    []Edges  `json:"edges"`
}

type PageInfo struct {
	HasNextPage bool   `json:"has_next_page"`
	EndCursor   string `json:"end_cursor"`
}

type Edges struct {
	Node EdgesNode `json:"node"`
}

type EdgesNode struct {
	ID                   string              `json:"id"`
	Shortcode            string              `json:"shortcode"`
	EdgeMediaPreviewLike EdgeMediaPrivewLike `json:"edge_media_preview_like"`
	EdgeMediaToCaption   EdgeMediaToCaption  `json:"edge_media_to_caption"`
	Owner                Owner               `json:"owner"`
}

type EdgeMediaPrivewLike struct {
	Count int64 `json:"count"`
}

type Owner struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

type EdgeMediaToCaption struct {
	Edges []EdgeBody `json:"edges"`
}

type EdgeBody struct {
	Node EdgeBodyNode `json:"node"`
}

type EdgeBodyNode struct {
	Text string `json:"text"`
}
