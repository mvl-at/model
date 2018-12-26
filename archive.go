package model

//defines the archive which stores all scores
type Archive struct {
	Id        int64  `json:"id"`
	Title     string `json:"title" roles:"archive"`
	Style     string `json:"style" roles:"archive"`
	Level     string `json:"level" roles:"archive"`
	Composer  string `json:"composer" roles:"archive"`
	Arranger  string `json:"arranger" roles:"archive"`
	Publisher string `json:"publisher" roles:"archive"`
	Subtitles string `json:"subtitles" roles:"archive"`
	Score     bool   `json:"score" roles:"archive"`
	Location  string `json:"location" roles:"archive"`
	Note      string `json:"note" roles:"archive"`
}
