package commands

type ScanCLI struct {
	Headers      bool `help:"Get the headers of the email"`
	Hash         bool `help:"Get the hash of the email"`
	Links        bool `help:"Get the links in the email file"`
	Digests      bool `help:"Get the digests of the eml file"`
	Attachements bool `help:"Get the attachments from the eml file"`

	Filename string `name:"filename" help:"The eml file name to scan" required:""`
}
