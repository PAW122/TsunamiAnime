package types

type Submit_episode_request struct {
	Title               string `json:"title"`
	EpisodeNumber       int    `json:"episodeNumber"`
	EpisodeTitle        string `json:"episodeTitle"`
	Episode_relise_date string `json:"episode_relise_date"`
	Hosting             string `json:"hosting"`
	Quality             string `json:"quality"`
	Audio_Language      string `json:"audio_Language"`
	Text_Language       string `json:"text_Language"`
	Ai_translation      bool   `json:"Ai_translation"`
	Link_embed          string `json:"link_embed"`
	Translation         string `json:"translation"`
}

type SubmitEpisodeRequest struct {
	Title              string `json:"title"`
	EpisodeNumber      int    `json:"episodeNumber"`
	EpisodeTitle       string `json:"episodeTitle"`
	EpisodeReleaseDate string `json:"episodeReleaseDate"`
	Hosting            string `json:"hosting"`
	Quality            string `json:"quality"`
	AudioLanguage      string `json:"audioLanguage"`
	TextLanguage       string `json:"textLanguage"`
	AiTranslation      bool   `json:"aiTranslation"`
	LinkEmbed          string `json:"linkEmbed"`
	Translation        string `json:"translation"`
}

// type EpisodeRequest struct {
// 	Title              string `json:"title"`
// 	EpisodeNumber      int    `json:"episodeNumber"`
// 	EpisodeTitle       string `json:"episodeTitle"`
// 	EpisodeReleaseDate string `json:"episodeReleaseDate"`
// 	Hosting            string `json:"hosting"`
// 	Quality            string `json:"quality"`
// 	AudioLanguage      string `json:"audioLanguage"`
// 	TextLanguage       string `json:"textLanguage"`
// 	AiTranslation      bool   `json:"aiTranslation"`
// 	LinkEmbed          string `json:"linkEmbed"`
// 	Translation        string `json:"translation"`
// }
