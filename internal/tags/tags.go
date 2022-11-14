package tags

type Tag struct {
	Ref    string `json:"ref"`
	NodeId string `json:"node_id"`
	Url    string `json:"url"`
	Object *struct {
		Sha  string `json:"sha"`
		Type string `json:"type"`
		Url  string `json:"url"`
	} `json:"object"`
}
