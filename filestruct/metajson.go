package filestruct

type MetaData struct {
	CreatedAt int64 `json:"createdAt"`
	Files     []struct {
		Path       string `json:"path"`
		PathC      string `json:"pathc"`
		PathH      string `json:"pathh"`
		SHA1       string `json:"sha1"`
		Size       int64  `json:"size"`
		HashedPath string `json:"hashed_path"`
		HashedSHA1 string `json:"hashed_sha1"`
	} `json:"files"`
}
