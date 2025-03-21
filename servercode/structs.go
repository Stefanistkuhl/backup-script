package main

type Config struct {
	UploadDir                string `json:"uploadDir"`
	CheckIntervall           string `json:"checkIntervall"`
	GenerationCount          int    `json:"generationCount"`
	GenerationsDir           string `json:"generationsDir"`
	SshPort                  int    `json:"sshPort"`
	Atomicity                bool   `json:"atomicity"`
	GenerationsDirNamePrefix string `json:"genDirPrefix"`
}
