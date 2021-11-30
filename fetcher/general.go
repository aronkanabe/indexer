package fetcher

const (
	RARIBLE    = "Rarible"
	CONTEXT    = "Context"
	CONVO      = "Convo"
	TWITTER    = "Twtter"
	OPENSEA    = "Opensea"
	ZORA       = "Zora"
	FOUNDATION = "Foundation"
	SHOWTIME   = "Showtime"
	SYBIL      = "Sybil"
	SUPERRARE  = "Superrare"
	INFURA     = "Infura"
)

const (
	SuperrareContractAddress  = "0x41a322b28d0ff354040e2cbc676f0320d8c8850d"
	OpenSeaContractAddress    = "0x495f947276749ce646f68ac8c248420045cb7b5e"
	RaribleContractAddress    = "0xd07dc4262bcdbf85190c01c996b4c06a461d2430"
	FoundationContractAddress = "0x3b3ee1931dc30c1957379fac9aba94d1c48a5405"
	ZoraContractAddress       = "0xabefbc9fd2f806065b4f3c237d4b59d9a97bcac7"
	ContextContractAddress    = "ctx"
)

const (
	ContextUrl          = "https://context.app/api/profile/%s"
	SuperrareUrl        = "https://superrare.com/api/v2/user?address=%s"
	RaribleFollowingUrl = "https://api-mainnet.rarible.com/marketplace/api/v4/followings?owner=%s"
	RaribleFollowerUrl  = "https://api-mainnet.rarible.com/marketplace/api/v4/followers?user=%s"
	PoapSubGraphUrl     = "https://api.thegraph.com/subgraphs/name/amxx/poap"
	PoapApiUrl          = "https://api.poap.xyz"
)

type ConnectionEntryList struct {
	Conn []ConnectionEntry
	Err  error
	msg  string
}
type ConnectionEntry struct {
	From     string
	To       string
	Platform string
}

type IdentityEntryList struct {
	OpenSea    []UserOpenSeaIdentity
	Twitter    []UserTwitterIdentity
	Superrare  []UserSuperrareIdentity
	Rarible    []UserRaribleIdentity
	Context    []UserContextIdentity
	Zora       []UserZoraIdentity
	Foundation []UserFoundationIdentity
	Showtime   []UserShowtimeIdentity
	Poap       []UserPoapIdentity
	Ens        string
}

type IdentityEntry struct {
	OpenSea    *UserOpenSeaIdentity
	Twitter    *UserTwitterIdentity
	Superrare  *UserSuperrareIdentity
	Rarible    *UserRaribleIdentity
	Context    *UserContextIdentity
	Zora       *UserZoraIdentity
	Ens        *UserEnsIdentity
	Foundation *UserFoundationIdentity
	Showtime   *UserShowtimeIdentity
	Poap       []UserPoapIdentity
	Err        error
	Msg        string
}

type UserTwitterIdentity struct {
	Handle     string
	DataSource string
}

type UserRaribleIdentity struct {
	Username        string
	Homepage        string
	ItemSold        int
	AmountSoldInEth float64
	DataSource      string
}

type UserOpenSeaIdentity struct {
	Username   string
	Homepage   string
	DataSource string
}

type UserEnsIdentity struct {
	Ens        string
	DataSource string
}

type UserContextIdentity struct {
	FollowerCount int
	Username      string
	Website       string
	DataSource    string
}

type UserSuperrareIdentity struct {
	Username       string
	Homepage       string
	Location       string
	Bio            string
	InstagramLink  string
	TwitterLink    string
	SteemitLink    string
	Website        string
	SpotifyLink    string
	SoundCloudLink string
	DataSource     string
}

type UserFoundationIdentity struct {
	Username   string
	Bio        string
	Tiktok     string
	Twitch     string
	Discord    string
	Twitter    string
	Website    string
	Youtube    string
	Facebook   string
	Snapchat   string
	Instagram  string
	DataSource string
}

type UserZoraIdentity struct {
	Username   string
	Website    string
	DataSource string
}

type UserShowtimeIdentity struct {
	Name             string
	Username         string
	Bio              string
	TwitterHandle    string
	LinkTreeHandle   string
	CryptoArtHandle  string
	FoundationHandle string
	HicetnuncHandle  string
	OpenseaHandle    string
	RaribleHandle    string
	DataSource       string
}

type UserPoapIdentity struct {
	Address         string
	EventID         int
	TokenID         int
	FancyID         string
	Name            string
	EventUrl        string
	ImageUrl        string
	Country         string
	City            string
	Description     string
	Year            int
	StartDate       string
	EndDate         string
	ExpiryDate      string
	FromAdmin       bool
	VirtualEvent    bool
	EventTemplateID int
	EventHostID     int
	PrivateEvent    bool
	Recommendations []PoapRecommendation
}

type PoapRecommendation struct {
	Address string
	EventID int
	TokenID int
}

type RaribleConnectionResp struct {
	Following struct {
		From string `json:"owner"`
		To   string `json:"user"`
	} `json:"following"`
}

type ContextAppResp struct {
	FollowerCount int               `json:"followerCount"`
	Ens           map[string]string `json:"ens"`
	Profiles      map[string]([]struct {
		Address  string `json:"address,omitempty"`
		Contract string `json:"contract,omitempty"`
		Url      string `json:"url,omitempty"`
		Website  string `json:"website,omitempty"`
		Username string `json:"username,omitempty"`
	}) `json:"profiles"`
}

type ContextConnection struct {
	Relationships []struct {
		Actor string `json:"actor"`
	} `json:"relationships"`
	Profiles map[string]([]struct {
		Address string `json:"address"`
	}) `json:"profiles"`
}

type SuperrareProfile struct {
	Result struct {
		Username       string `json:"username"`
		Location       string `json:"location"`
		Bio            string `json:"bio"`
		InstagramLink  string `json:"instagramLink"`
		TwitterLink    string `json:"twitterLink"`
		SteemitLink    string `json:"steemitLink"`
		Website        string `json:"website"`
		SpotifyLink    string `json:"spotifyLink"`
		SoundCloudLink string `json:"soundcloudLink"`
	} `json:"result"`
}

type FoundationIdentity struct {
	Data struct {
		User struct {
			Username string `json:"username"`
			Bio      string `json:"bio"`
			Links    struct {
				Tiktok struct {
					Handle string `json:"handle"`
				} `json:"tiktok"`
				Twitch struct {
					Handle string `json:"handle"`
				} `json:"twitch"`
				Discord struct {
					Handle string `json:"handle"`
				} `json:"discord"`
				Twitter struct {
					Handle string `json:"handle"`
				} `json:"twitter"`
				Website struct {
					Handle string `json:"handle"`
				} `json:"website"`
				Youtube struct {
					Handle string `json:"handle"`
				} `json:"youtube"`
				Facebook struct {
					Handle string `json:"handle"`
				} `json:"facebook"`
				Snapchat struct {
					Handle string `json:"handle"`
				} `json:"snapchat"`
				Instagram struct {
					Handle string `json:"handle"`
				} `json:"instagram"`
			} `json:"links"`
			TwitSocialVerifs []struct {
				Username string `json:"username"`
			} `json:"twitSocialVerifs"`
			InstaSocialVerifs []struct {
				Username string `json:"username"`
			} `json:"instaSocialVerifs"`
		} `json:"user"`
	} `json:"data"`
}

type PoapGraphResponse struct {
	Data struct {
		Account struct {
			ID     string `json:"id"`
			Tokens []struct {
				ID    int `json:"id,string"`
				Event struct {
					ID     int `json:"id,string"`
					Tokens []struct {
						ID    int `json:"id,string"`
						Owner struct {
							ID string `json:"id"`
						} `json:"owner"`
					} `json:"tokens"`
				} `json:"event"`
			} `json:"tokens"`
		} `json:"account"`
	} `json:"data"`
}

type PoapApiResponse struct {
	ID              int    `json:"id"`
	FancyID         string `json:"fancy_id"`
	Name            string `json:"name"`
	EventURL        string `json:"event_url"`
	ImageURL        string `json:"image_url"`
	Country         string `json:"country"`
	City            string `json:"city"`
	Description     string `json:"description"`
	Year            int    `json:"year"`
	StartDate       string `json:"start_date"`
	EndDate         string `json:"end_date"`
	ExpiryDate      string `json:"expiry_date"`
	FromAdmin       bool   `json:"from_admin"`
	VirtualEvent    bool   `json:"virtual_event"`
	EventTemplateID int    `json:"event_template_id"`
	EventHostID     int    `json:"event_host_id"`
	PrivateEvent    bool   `json:"private_event"`
}
