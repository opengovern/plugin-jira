// Implement types for each resource

package provider

type Metadata struct{}

type ProjectListResponse struct {
	IsLast  bool          `json:"isLast"`
	StartAt int           `json:"startAt"`
	Values  []ProjectJSON `json:"values"`
}

type AvatarUrlsJSON struct {
	Small16x16  string `json:"16x16"`
	Small24x24  string `json:"24x24"`
	Medium32x32 string `json:"32x32"`
	Large48x48  string `json:"48x48"`
}

type AvatarUrls struct {
	Small16x16  string
	Small24x24  string
	Medium32x32 string
	Large48x48  string
}

type InsightJSON struct {
	LastIssueUpdateTime string `json:"lastIssueUpdateTime"`
	TotalIssueCount     int    `json:"totalIssueCount"`
}

type Insight struct {
	LastIssueUpdateTime string
	TotalIssueCount     int
}

type ProjectCategoryJSON struct {
	Description string `json:"description"`
	ID          string `json:"id"`
	Name        string `json:"name"`
	Self        string `json:"self"`
}

type ProjectCategory struct {
	Description string
	ID          string
	Name        string
	Self        string
}

type ProjectJSON struct {
	AvatarUrls      AvatarUrlsJSON      `json:"avatarUrls"`
	ID              string              `json:"id"`
	Insight         InsightJSON         `json:"insight"`
	Key             string              `json:"key"`
	Name            string              `json:"name"`
	ProjectCategory ProjectCategoryJSON `json:"projectCategory"`
	Self            string              `json:"self"`
	Simplified      bool                `json:"simplified"`
	Style           string              `json:"style"`
}

type ProjectDescription struct {
	AvatarUrls      AvatarUrls
	ID              string
	Insight         Insight
	Key             string
	Name            string
	ProjectCategory ProjectCategory
	Self            string
	Simplified      bool
	Style           string
}

type IssueListResponse struct {
	Expand          string      `json:"expand"`
	Issues          []IssueJSON `json:"issues"`
	MaxResults      int         `json:"maxResults"`
	StartAt         int         `json:"startAt"`
	Total           int         `json:"total"`
	WarningMessages []string    `json:"warningMessages"`
}

type IssueJSON struct {
	ID     string     `json:"id"`
	Key    string     `json:"key"`
	Self   string     `json:"self"`
	Fields FieldsJSON `json:"fields"`
	Expand string     `json:"expand"`
}

type IssueDescription struct {
	ID     string
	Key    string
	Self   string
	Fields Fields
	Expand string
}

type FieldsJSON struct {
	Watcher      WatcherJSON        `json:"watcher"`
	Attachment   []AttachmentJSON   `json:"attachment"`
	SubTasks     []SubTaskJSON      `json:"sub-tasks"`
	Project      ProjectJSON        `json:"project"`
	Comment      []CommentJSON      `json:"comment"`
	IssueLinks   []IssueLinkJSON    `json:"issuelinks"`
	Worklog      []WorklogEntryJSON `json:"worklog"`
	TimeTracking TimeTrackingJSON   `json:"timetracking"`
	Updated      string             `json:"updated"`
}

type Fields struct {
	Watcher      Watcher
	Attachment   []Attachment
	SubTasks     []SubTask
	Project      ProjectDescription
	Comment      []Comment
	IssueLinks   []IssueLink
	Worklog      []WorklogEntry
	TimeTracking TimeTracking
	Updated      string
}

type WatcherJSON struct {
	IsWatching bool   `json:"isWatching"`
	Self       string `json:"self"`
	WatchCount int    `json:"watchCount"`
}

type Watcher struct {
	IsWatching bool
	Self       string
	WatchCount int
}

type AttachmentJSON struct {
	Author   UserJSON `json:"author"`
	Content  string   `json:"content"`
	Created  string   `json:"created"`
	Filename string   `json:"filename"`
	ID       int      `json:"id"`
	MimeType string   `json:"mimeType"`
	Self     string   `json:"self"`
	Size     int      `json:"size"`
}

type Attachment struct {
	Author   User
	Content  string
	Created  string
	Filename string
	ID       int
	MimeType string
	Self     string
	Size     int
}

type UserJSON struct {
	AccountID   string         `json:"accountId"`
	AccountType string         `json:"accountType"`
	Active      bool           `json:"active"`
	AvatarUrls  AvatarUrlsJSON `json:"avatarUrls"`
	DisplayName string         `json:"displayName"`
	Key         string         `json:"key"`
	Name        string         `json:"name"`
	Self        string         `json:"self"`
}

type User struct {
	AccountID   string
	AccountType string
	Active      bool
	AvatarUrls  AvatarUrls
	DisplayName string
	Key         string
	Name        string
	Self        string
}

type SubTaskJSON struct {
	ID           string           `json:"id"`
	OutwardIssue OutwardIssueJSON `json:"outwardIssue"`
	Type         IssueTypeJSON    `json:"type"`
}

type SubTask struct {
	ID           string
	OutwardIssue OutwardIssue
	Type         IssueType
}

type OutwardIssueJSON struct {
	Fields StatusFieldJSON `json:"fields"`
	ID     string          `json:"id"`
	Key    string          `json:"key"`
	Self   string          `json:"self"`
}

type OutwardIssue struct {
	Fields StatusField
	ID     string
	Key    string
	Self   string
}

type IssueTypeJSON struct {
	ID      string `json:"id"`
	Inward  string `json:"inward"`
	Name    string `json:"name"`
	Outward string `json:"outward"`
}

type IssueType struct {
	ID      string
	Inward  string
	Name    string
	Outward string
}

type StatusFieldJSON struct {
	Status StatusJSON `json:"status"`
}

type StatusField struct {
	Status Status
}

type StatusJSON struct {
	IconURL string `json:"iconUrl"`
	Name    string `json:"name"`
}

type Status struct {
	IconURL string
	Name    string
}

type ContentJSON struct {
	Type    string              `json:"type"`
	Version int                 `json:"version"`
	Content []ContentDetailJSON `json:"content"`
}

type Content struct {
	Type    string
	Version int
	Content []ContentDetail
}

type ContentDetailJSON struct {
	Type    string            `json:"type"`
	Content []TextContentJSON `json:"content"`
}

type ContentDetail struct {
	Type    string
	Content []TextContent
}

type TextContentJSON struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type TextContent struct {
	Type string
	Text string
}

type CommentJSON struct {
	Author       UserJSON       `json:"author"`
	Created      string         `json:"created"`
	ID           string         `json:"id"`
	Self         string         `json:"self"`
	UpdateAuthor UserJSON       `json:"updateAuthor"`
	Updated      string         `json:"updated"`
	Visibility   VisibilityJSON `json:"visibility"`
}

type Comment struct {
	Author       User
	Created      string
	ID           string
	Self         string
	UpdateAuthor User
	Updated      string
	Visibility   Visibility
}

type VisibilityJSON struct {
	Identifier string `json:"identifier"`
	Type       string `json:"type"`
	Value      string `json:"value"`
}

type Visibility struct {
	Identifier string
	Type       string
	Value      string
}

type IssueLinkJSON struct {
	ID           string            `json:"id"`
	OutwardIssue *OutwardIssueJSON `json:"outwardIssue,omitempty"`
	InwardIssue  *OutwardIssueJSON `json:"inwardIssue,omitempty"`
	Type         IssueTypeJSON     `json:"type"`
}

type IssueLink struct {
	ID           string
	OutwardIssue *OutwardIssue
	InwardIssue  *OutwardIssue
	Type         IssueType
}

type WorklogEntryJSON struct {
	Author           UserJSON       `json:"author"`
	ID               string         `json:"id"`
	IssueID          string         `json:"issueId"`
	Self             string         `json:"self"`
	Started          string         `json:"started"`
	TimeSpent        string         `json:"timeSpent"`
	TimeSpentSeconds int            `json:"timeSpentSeconds"`
	UpdateAuthor     UserJSON       `json:"updateAuthor"`
	Updated          string         `json:"updated"`
	Visibility       VisibilityJSON `json:"visibility"`
}

type WorklogEntry struct {
	Author           User
	ID               string
	IssueID          string
	Self             string
	Started          string
	TimeSpent        string
	TimeSpentSeconds int
	UpdateAuthor     User
	Updated          string
	Visibility       Visibility
}

type TimeTrackingJSON struct {
	OriginalEstimate         string `json:"originalEstimate"`
	OriginalEstimateSeconds  int    `json:"originalEstimateSeconds"`
	RemainingEstimate        string `json:"remainingEstimate"`
	RemainingEstimateSeconds int    `json:"remainingEstimateSeconds"`
	TimeSpent                string `json:"timeSpent"`
	TimeSpentSeconds         int    `json:"timeSpentSeconds"`
}

type TimeTracking struct {
	OriginalEstimate         string
	OriginalEstimateSeconds  int
	RemainingEstimate        string
	RemainingEstimateSeconds int
	TimeSpent                string
	TimeSpentSeconds         int
}

type BoardListResponse struct {
	IsLast  bool        `json:"isLast"`
	StartAt int         `json:"startAt"`
	Values  []BoardJSON `json:"values"`
}

type BoardJSON struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Self string `json:"self"`
	Type string `json:"type"`
}

type BoardDescription struct {
	ID   int
	Name string
	Self string
	Type string
}
