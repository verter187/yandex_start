module example.com/m

go 1.20

replace person => ../person

replace company => ../company

replace robot => ../robot

require (
	company v0.0.0-00010101000000-000000000000
	person v0.0.0-00010101000000-000000000000
	robot v0.0.0-00010101000000-000000000000
)
