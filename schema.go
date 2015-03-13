/*
 * schema.go
 * 
 * Copyright 2015 Christian Diener <ch.diener@gmail.com>
 * 
 * MIT license. See LICENSE for more information.
 */

package main

type Place struct {
	Address string `json:"address"`
	PostalCode string `json:"postalCode"`
	City string `json:"city"`
	Country string `json:"countryCode"`
	Region string `json:"region"`
}

type Profile struct {
	Network string `json:"network"`
	User string `json:"username"`
	Url string `json:"url"`
}

type BasicInfo struct {
	First string `json:"firstname"`
	Last string `json:"lastname"`
	Label string `json:"label"`
	Picture string `json:"picture"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Website string `json:"website"`
	Summary string `json:"summary"`
	Location Place `json:"location"`
	Profiles []Profile `json:"profiles"`
}

type Workplace struct {
	Company string `json:"company"`
	Position string `json:"position"`
	Website string `json:"website"`
	Begin string `json:"startDate"`
	End string `json:"endDate"`
	Summary string `json:"summary"`
	Highlights []string `json:"highlights"`
}

type VolunteerInfo struct {
	Organization string `json:"organization"`
	Position string `json:"position"`
	Website string `json:"website"`
	Begin string `json:"startDate"`
	End string `json:"endDate"`
	Summary string `json:"summary"`
	Highlights []string `json:"highlights"`
}

type EducationInfo struct {
	Institution string `json:"institution"`
	Area string `json:"area"`
	Title string `json:"studyType"`
	Begin string `json:"startDate"`
	End string `json:"endDate"`
	Grade string `json:"gpa"`
	Courses []string `json:"courses"`
}

type Award struct {
	Title string `json:"title"`
	Date string `json:"date"`
	Institution string `json:"awarder"`
	Summary string `json:"summary"`
}

type Publication struct {
	Title string `json:"name"`
	Publisher string `json:"publisher"`
	Date string `json:"releaseDate"`
	Url string `json:"website"`
	Summary string `json:"summary"`
}

type Skill struct {
	Name string `json:"name"`
	Level string `json:"level"`
	Keywords []string `json:"keywords"`
}

type Language struct {
	Language string `json:"language"`
	Fluency string `json:"fluency"`
}

type Interest struct {
	Name string `json:"name"`
	Keywords []string `json:"keywords"`
}

type Reference struct {
	Name string `json:"name"`
	Quote string `json:"reference"`
}

type Resume struct {
	Basics BasicInfo `json:"basics"`
	Work []Workplace `json:"work"` 
	Volunteer []VolunteerInfo `json:"volunteer"`
	Education []EducationInfo `json:"education"`
	Awards []Award `json:"awards"`
	Publications []Publication `json:"publications"`
	Skills []Skill `json:"skills"`
	Languages []Language `json:"languages"`
	Interests []Interest `json:"interests"`
	References []Reference `json:"References"`
}


