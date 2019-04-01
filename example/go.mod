module go.mod/example

require (
	go.mod/a v0.0.0
	go.mod/b v0.0.0

)

replace (
	go.mod/a => ../a
	go.mod/b => ../b
)
