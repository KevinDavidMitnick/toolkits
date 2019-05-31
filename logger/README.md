logger
======
	if g.Config().Debug {
		g.InitLog("logfmt", "debug")
	} else {
		g.InitLog("logfmt", "error")
	}

	level.Info(g.Logger).Log("msg", "start gateway......")

