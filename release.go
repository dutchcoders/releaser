package releaser

type Release struct {
	Assets []struct {
		BrowserDownloadURL string `json:"browser_download_url"`
		ContentType        string `json:"content_type"`
		CreatedAt          string `json:"created_at"`
		DownloadCount      int64  `json:"download_count"`
		ID                 int64  `json:"id"`
		Label              string `json:"label"`
		Name               string `json:"name"`
		Size               int64  `json:"size"`
		State              string `json:"state"`
		UpdatedAt          string `json:"updated_at"`
		Uploader           struct {
			AvatarURL         string `json:"avatar_url"`
			EventsURL         string `json:"events_url"`
			FollowersURL      string `json:"followers_url"`
			FollowingURL      string `json:"following_url"`
			GistsURL          string `json:"gists_url"`
			GravatarID        string `json:"gravatar_id"`
			HtmlURL           string `json:"html_url"`
			ID                int64  `json:"id"`
			Login             string `json:"login"`
			OrganizationsURL  string `json:"organizations_url"`
			ReceivedEventsURL string `json:"received_events_url"`
			ReposURL          string `json:"repos_url"`
			SiteAdmin         bool   `json:"site_admin"`
			StarredURL        string `json:"starred_url"`
			SubscriptionsURL  string `json:"subscriptions_url"`
			Type              string `json:"type"`
			URL               string `json:"url"`
		} `json:"uploader"`
		URL string `json:"url"`
	} `json:"assets"`
	AssetsURL string `json:"assets_url"`
	Author    struct {
		AvatarURL         string `json:"avatar_url"`
		EventsURL         string `json:"events_url"`
		FollowersURL      string `json:"followers_url"`
		FollowingURL      string `json:"following_url"`
		GistsURL          string `json:"gists_url"`
		GravatarID        string `json:"gravatar_id"`
		HtmlURL           string `json:"html_url"`
		ID                int64  `json:"id"`
		Login             string `json:"login"`
		OrganizationsURL  string `json:"organizations_url"`
		ReceivedEventsURL string `json:"received_events_url"`
		ReposURL          string `json:"repos_url"`
		SiteAdmin         bool   `json:"site_admin"`
		StarredURL        string `json:"starred_url"`
		SubscriptionsURL  string `json:"subscriptions_url"`
		Type              string `json:"type"`
		URL               string `json:"url"`
	} `json:"author"`
	Body            string `json:"body"`
	CreatedAt       string `json:"created_at"`
	Draft           bool   `json:"draft"`
	HtmlURL         string `json:"html_url"`
	ID              int64  `json:"id"`
	Name            string `json:"name"`
	Prerelease      bool   `json:"prerelease"`
	PublishedAt     string `json:"published_at"`
	TagName         string `json:"tag_name"`
	TarballURL      string `json:"tarball_url"`
	TargetCommitish string `json:"target_commitish"`
	UploadURL       string `json:"upload_url"`
	URL             string `json:"url"`
	ZipballURL      string `json:"zipball_url"`
}

type Commit struct {
	Author struct {
		AvatarUrl         string `json:"avatar_url"`
		EventsUrl         string `json:"events_url"`
		FollowersUrl      string `json:"followers_url"`
		FollowingUrl      string `json:"following_url"`
		GistsUrl          string `json:"gists_url"`
		GravatarId        string `json:"gravatar_id"`
		HtmlUrl           string `json:"html_url"`
		Id                int64  `json:"id"`
		Login             string `json:"login"`
		NodeId            string `json:"node_id"`
		OrganizationsUrl  string `json:"organizations_url"`
		ReceivedEventsUrl string `json:"received_events_url"`
		ReposUrl          string `json:"repos_url"`
		SiteAdmin         bool   `json:"site_admin"`
		StarredUrl        string `json:"starred_url"`
		SubscriptionsUrl  string `json:"subscriptions_url"`
		Type              string `json:"type"`
		Url               string `json:"url"`
	} `json:"author"`
	CommentsUrl string `json:"comments_url"`
	Commit      struct {
		Author struct {
			Date  string `json:"date"`
			Email string `json:"email"`
			Name  string `json:"name"`
		} `json:"author"`
		CommentCount int64 `json:"comment_count"`
		Committer    struct {
			Date  string `json:"date"`
			Email string `json:"email"`
			Name  string `json:"name"`
		} `json:"committer"`
		Message string `json:"message"`
		Tree    struct {
			Sha string `json:"sha"`
			Url string `json:"url"`
		} `json:"tree"`
		Url          string `json:"url"`
		Verification struct {
			Payload   interface{} `json:"payload"`
			Reason    string      `json:"reason"`
			Signature interface{} `json:"signature"`
			Verified  bool        `json:"verified"`
		} `json:"verification"`
	} `json:"commit"`
	Committer struct {
		AvatarUrl         string `json:"avatar_url"`
		EventsUrl         string `json:"events_url"`
		FollowersUrl      string `json:"followers_url"`
		FollowingUrl      string `json:"following_url"`
		GistsUrl          string `json:"gists_url"`
		GravatarId        string `json:"gravatar_id"`
		HtmlUrl           string `json:"html_url"`
		Id                int64  `json:"id"`
		Login             string `json:"login"`
		NodeId            string `json:"node_id"`
		OrganizationsUrl  string `json:"organizations_url"`
		ReceivedEventsUrl string `json:"received_events_url"`
		ReposUrl          string `json:"repos_url"`
		SiteAdmin         bool   `json:"site_admin"`
		StarredUrl        string `json:"starred_url"`
		SubscriptionsUrl  string `json:"subscriptions_url"`
		Type              string `json:"type"`
		Url               string `json:"url"`
	} `json:"committer"`
	HtmlUrl string `json:"html_url"`
	NodeId  string `json:"node_id"`
	Parents []struct {
		HtmlUrl string `json:"html_url"`
		Sha     string `json:"sha"`
		Url     string `json:"url"`
	} `json:"parents"`
	Sha string `json:"sha"`
	Url string `json:"url"`
}
