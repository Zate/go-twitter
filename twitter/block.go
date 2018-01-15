package twitter

import (
	"net/http"

	"github.com/dghubble/sling"
)

// User represents a Twitter User.
// https://dev.twitter.com/overview/api/users
// type User struct {
// 	ContributorsEnabled            bool          `json:"contributors_enabled"`
// 	CreatedAt                      string        `json:"created_at"`
// 	DefaultProfile                 bool          `json:"default_profile"`
// 	DefaultProfileImage            bool          `json:"default_profile_image"`
// 	Description                    string        `json:"description"`
// 	Email                          string        `json:"email"`
// 	Entities                       *UserEntities `json:"entities"`
// 	FavouritesCount                int           `json:"favourites_count"`
// 	FollowRequestSent              bool          `json:"follow_request_sent"`
// 	Following                      bool          `json:"following"`
// 	FollowersCount                 int           `json:"followers_count"`
// 	FriendsCount                   int           `json:"friends_count"`
// 	GeoEnabled                     bool          `json:"geo_enabled"`
// 	ID                             int64         `json:"id"`
// 	IDStr                          string        `json:"id_str"`
// 	IsTranslator                   bool          `json:"is_translator"`
// 	Lang                           string        `json:"lang"`
// 	ListedCount                    int           `json:"listed_count"`
// 	Location                       string        `json:"location"`
// 	Name                           string        `json:"name"`
// 	Notifications                  bool          `json:"notifications"`
// 	ProfileBackgroundColor         string        `json:"profile_background_color"`
// 	ProfileBackgroundImageURL      string        `json:"profile_background_image_url"`
// 	ProfileBackgroundImageURLHttps string        `json:"profile_background_image_url_https"`
// 	ProfileBackgroundTile          bool          `json:"profile_background_tile"`
// 	ProfileBannerURL               string        `json:"profile_banner_url"`
// 	ProfileImageURL                string        `json:"profile_image_url"`
// 	ProfileImageURLHttps           string        `json:"profile_image_url_https"`
// 	ProfileLinkColor               string        `json:"profile_link_color"`
// 	ProfileSidebarBorderColor      string        `json:"profile_sidebar_border_color"`
// 	ProfileSidebarFillColor        string        `json:"profile_sidebar_fill_color"`
// 	ProfileTextColor               string        `json:"profile_text_color"`
// 	ProfileUseBackgroundImage      bool          `json:"profile_use_background_image"`
// 	Protected                      bool          `json:"protected"`
// 	ScreenName                     string        `json:"screen_name"`
// 	ShowAllInlineMedia             bool          `json:"show_all_inline_media"`
// 	Status                         *Tweet        `json:"status"`
// 	StatusesCount                  int           `json:"statuses_count"`
// 	Timezone                       string        `json:"time_zone"`
// 	URL                            string        `json:"url"`
// 	UtcOffset                      int           `json:"utc_offset"`
// 	Verified                       bool          `json:"verified"`
// 	WithheldInCountries            []string      `json:"withheld_in_countries"`
// 	WithholdScope                  string        `json:"withheld_scope"`
// }

// BlockService provides methods for accessing Twitter user API endpoints for blocking other users.
type BlockService struct {
	sling *sling.Sling
}

// NewBlockService returns a new BlockService.
func NewBlockService(sling *sling.Sling) *BlockService {
	return &BlockService{
		sling: sling.Path("blocks/"),
	}
}

// BlockUserParams is a struct for the params supplied to both Create and Destroy
type BlockUserParams struct {
	UserID          int64  `url:"user_id,omitempty"`          //optional - The screen name of the potentially blocked user. Helpful for disambiguating when a valid screen name is also a user ID.
	ScreenName      string `url:"screen_name,omitempty"`      // optional - The ID of the potentially blocked user. Helpful for disambiguating when a valid user ID is also a valid screen name.
	IncludeEntities *bool  `url:"include_entities,omitempty"` // optional - The entities node will not be included when set to false.  Default: false
	SkipStatus      *bool  `url:"skip_status,omitempty"`      // optional - When set to either true , t or 1 statuses will not be included in the returned user objects.  Default: true
}

// Create blocks a user
func (s *BlockService) Create(params *BlockUserParams) (*User, *http.Response, error) {
	user := new(User)
	apiError := new(APIError)
	resp, err := s.sling.New().Post("create.json").QueryStruct(params).Receive(user, apiError)
	return user, resp, relevantError(err, *apiError)
}

// Destroy removes a block on a user
func (s *BlockService) Destroy(params *BlockUserParams) (*User, *http.Response, error) {
	user := new(User)
	apiError := new(APIError)
	resp, err := s.sling.New().Post("destroy.json").QueryStruct(params).Receive(user, apiError)
	return user, resp, relevantError(err, *apiError)
}
