run:
	nodemon --watch './**/*.go' --signal SIGTERM --exec 'go build -o safex main.go"' --verbose

reflect:
	reflex -r '\.go' -- sh -c './build.sh'
