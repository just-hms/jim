package actions

// TODO:
// - add --synch
// - decide how to work with collision probs

// - add --checkout (change from local to global aliases)
//	- find a way to devide the dbs (maybe two file and a .gitignore on the local one)
// - find a way to list and use both local and global aliases

// - add --move
//	move aliases between local and global

// HOW TO CREATE A NEW REPO ON GITHUB
// $ git init
// $ git add .
// $ git commit -m "first commit"

// $ git remote add origin git@github.com:alexpchin/<reponame>.git
// $ git push -u origin master

var Sync = &Action{
	Value: func(args []string) {

	},
	ArgumentsCheck: func(args []string) bool {
		return len(args) == 1
	},
}
