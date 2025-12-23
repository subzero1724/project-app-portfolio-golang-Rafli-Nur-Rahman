package model

type Profile struct {
	Name  string
	Role  string
	About string
}

// type Project struct {
// 	Title       string
// 	Description string
// }

type PageData struct {
	Title    string
	Profile  Profile
	Projects []Project
}
