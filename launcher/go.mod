module netease.com/launcher

replace netease.com/reqs => ../reqs

replace netease.com/core => ../core

replace netease.com/controller => ../controller

go 1.16

require (
	github.com/gin-gonic/gin v1.7.1
	netease.com/controller v0.0.0-00010101000000-000000000000
	netease.com/core v0.0.0-00010101000000-000000000000
	netease.com/reqs v0.0.0-00010101000000-000000000000 // indirect

)
